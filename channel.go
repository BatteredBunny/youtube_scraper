package scraper

import (
	"bytes"
	"errors"
	"github.com/ayes-web/rjson"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Channel struct {
	Subscribers  string
	IsVerified   bool
	ChannelID    string
	NewChannelID string
	Username     string
	Description  string
	VideosAmount int
}

type Video struct {
	VideoID string `json:"VideoID"`
	Title   string `json:"Title"`

	// 7:03
	Length string `json:"Length"`

	/*
		Examples
			- 100 views
			- 6,396,043 views
	*/
	Views string `json:"Views"`

	/*
		Years
			- 2-11 years ago
			- 1 year ago

		Months
			- 2-11 months ago
			- 1 month ago

		Weeks
			- 2-4 weeks ago

		Days
			- 2-13 days ago
			- 1 day ago

		Hours
			- 2-23 hours ago
			- 1 hour ago

		Minutes
			- 2-59 minutes ago
			- 1 minute ago

		Seconds
			- 2-59 seconds ago
			- 1 second ago
	*/
	Date string `json:"Date"`

	Username     string `json:"Username"`
	ChannelID    string `json:"ChannelID"`
	NewChannelID string `json:"NewChannelID"` // @username
}

type ChannelVideosScraper struct {
	url     string
	channel Channel

	InitialComplete   bool
	ContinueInputJson []byte
}

// NewChannelVideosScraper accepts normal id or @username
func NewChannelVideosScraper(id string) (c ChannelVideosScraper) {
	c.url = "https://www.youtube.com/"

	if strings.HasPrefix(id, "@") {
		c.url += id
	} else {
		c.url += "channel/" + id
	}

	c.url += "/videos?hl=en"
	return
}

// GetChannelInfo will output the internal channel struct after its become available with first NextPage() call
func (c *ChannelVideosScraper) GetChannelInfo() (available bool, channel Channel) {
	if c.InitialComplete {
		channel = c.channel
		available = true
	} else {
		available = false
	}

	return
}

type channelInitialOutput struct {
	Subscribers     string `rjson:"header.c4TabbedHeaderRenderer.subscriberCountText.simpleText"`
	ChannelID       string `rjson:"metadata.channelMetadataRenderer.externalId"`
	NewChannelID    string `rjson:"header.c4TabbedHeaderRenderer.channelHandleText.runs[0].text"`
	Username        string `rjson:"metadata.channelMetadataRenderer.title"`
	Description     string `rjson:"metadata.channelMetadataRenderer.description"`
	RawVideosAmount string `rjson:"header.c4TabbedHeaderRenderer.videosCountText.runs[0].text"`

	Badges []string `rjson:"header.c4TabbedHeaderRenderer.badges[].metadataBadgeRenderer.tooltip"`

	Videos []struct {
		VideoID string `rjson:"richItemRenderer.content.videoRenderer.videoId"`
		Title   string `rjson:"richItemRenderer.content.videoRenderer.title.runs[0].text"`
		Length  string `rjson:"richItemRenderer.content.videoRenderer.lengthText.simpleText"`
		Views   string `rjson:"richItemRenderer.content.videoRenderer.viewCountText.simpleText"`
		Date    string `rjson:"richItemRenderer.content.videoRenderer.publishedTimeText.simpleText"`
	} `rjson:"contents.twoColumnBrowseResultsRenderer.tabs[1].tabRenderer.content.richGridRenderer.contents"`
	VideosContinuationToken string `rjson:"contents.twoColumnBrowseResultsRenderer.tabs[1].tabRenderer.content.richGridRenderer.contents[-].continuationItemRenderer.continuationEndpoint.continuationCommand.token"`
}

func (c *ChannelVideosScraper) runInitial() (videos []Video, err error) {
	var rawJson string
	rawJson, err = extractInitialData(c.url)
	if err != nil {
		return
	}

	if Debug {
		os.WriteFile("channel_initial.json", []byte(rawJson), 0777)
	}

	var output channelInitialOutput
	if err = rjson.Unmarshal([]byte(rawJson), &output); err != nil {
		if errors.Unwrap(err) == rjson.ErrCantFindField {
			if Debug {
				log.Println("WARNING:", err)
			}
			err = nil
		}
		return
	}

	// channel
	c.channel = Channel{
		Subscribers:  output.Subscribers,
		ChannelID:    output.ChannelID,
		NewChannelID: output.NewChannelID,
		Username:     output.Username,
		Description:  output.Description,
	}

	for _, badge := range output.Badges {
		if badge == "Verified" {
			c.channel.IsVerified = true
		}
	}

	if output.RawVideosAmount == "No videos" {
		c.channel.VideosAmount = 0
	} else {
		c.channel.VideosAmount, err = strconv.Atoi(output.RawVideosAmount)
		if err != nil {
			return
		}
	}

	c.ContinueInputJson, err = continueInput{Continuation: output.VideosContinuationToken}.FillGenericInfo().Construct()
	if err != nil {
		return
	}

	for _, video := range output.Videos {
		if video.VideoID == "" {
			continue
		}

		videos = append(videos, Video{
			VideoID:      video.VideoID,
			Title:        video.Title,
			Length:       video.Length,
			Views:        video.Views,
			Date:         video.Date,
			ChannelID:    c.channel.ChannelID,
			NewChannelID: c.channel.NewChannelID,
		})
	}

	c.InitialComplete = true
	return
}

type channelContinueOutput struct {
	Videos []struct {
		VideoID string `rjson:"richItemRenderer.content.videoRenderer.videoId"`
		Title   string `rjson:"richItemRenderer.content.videoRenderer.title.runs[0].text"`
		Length  string `rjson:"richItemRenderer.content.videoRenderer.lengthText.simpleText"`
		Views   string `rjson:"richItemRenderer.content.videoRenderer.viewCountText.simpleText"`
		Date    string `rjson:"richItemRenderer.content.videoRenderer.publishedTimeText.simpleText"`
	} `rjson:"onResponseReceivedActions[0].appendContinuationItemsAction.continuationItems"`

	ContinueToken string `rjson:"onResponseReceivedActions[0].appendContinuationItemsAction.continuationItems[-]continuationItemRenderer.continuationEndpoint.continuationCommand.token"`
}

func (c *ChannelVideosScraper) NextPage() (videos []Video, err error) {
	if !c.InitialComplete {
		return c.runInitial()
	} else {
		var resp *http.Response
		resp, err = http.Post("https://www.youtube.com/youtubei/v1/browse", "application/json", bytes.NewReader(c.ContinueInputJson))
		if err != nil {
			return
		}
		c.ContinueInputJson = []byte{}

		var body []byte
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return
		}

		if Debug {
			os.WriteFile("channel_videos.json", body, 0777)
		}

		var output channelContinueOutput
		if err = rjson.Unmarshal(body, &output); err != nil {
			if errors.Unwrap(err) == rjson.ErrCantFindField {
				err = nil
			}
			return
		}

		c.ContinueInputJson, err = continueInput{Continuation: output.ContinueToken}.FillGenericInfo().Construct()
		if err != nil {
			return
		}

		for _, video := range output.Videos {
			if video.VideoID == "" {
				continue
			}

			videos = append(videos, Video{
				VideoID:      video.VideoID,
				Title:        video.Title,
				Length:       video.Length,
				Views:        video.Views,
				Date:         video.Date,
				ChannelID:    c.channel.ChannelID,
				NewChannelID: c.channel.NewChannelID,
			})
		}
	}

	return
}
