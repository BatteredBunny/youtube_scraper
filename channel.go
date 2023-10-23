package scraper

import (
	"bytes"
	"errors"
	"github.com/ayes-web/rjson"
	"io"
	"log"
	"net/http"
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

type ChannelScraper struct {
	baseChannelUrl string
	channel        Channel

	videosInitialComplete   bool
	videosContinueInputJson []byte
}

// NewChannelScraper accepts normal id or @username
func NewChannelScraper(id string) (c ChannelScraper) {
	c.baseChannelUrl = "https://www.youtube.com/"

	if strings.HasPrefix(id, "@") {
		c.baseChannelUrl += id
	} else {
		c.baseChannelUrl += "channel/" + id
	}

	return
}

// GetChannelInfo will output the internal channel struct after its become available with first NextVideosPage() call
func (c *ChannelScraper) GetChannelInfo() (available bool, channel Channel) {
	if c.videosInitialComplete {
		channel = c.channel
		available = true
	} else {
		available = false
	}

	return
}

type channelVideosInitialOutput struct {
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

func (c *ChannelScraper) runVideosInitial() (videos []Video, err error) {
	var rawJson string
	rawJson, err = extractInitialData(c.baseChannelUrl + "/videos?hl=en")
	if err != nil {
		return
	}

	debugFileOutput([]byte(rawJson), "channel_initial.json")

	var output channelVideosInitialOutput
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

	c.videosContinueInputJson, err = continueInput{Continuation: output.VideosContinuationToken}.FillGenericInfo().Construct()
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

	c.videosInitialComplete = true
	return
}

type channelVideosContinueOutput struct {
	Videos []struct {
		VideoID string `rjson:"richItemRenderer.content.videoRenderer.videoId"`
		Title   string `rjson:"richItemRenderer.content.videoRenderer.title.runs[0].text"`
		Length  string `rjson:"richItemRenderer.content.videoRenderer.lengthText.simpleText"`
		Views   string `rjson:"richItemRenderer.content.videoRenderer.viewCountText.simpleText"`
		Date    string `rjson:"richItemRenderer.content.videoRenderer.publishedTimeText.simpleText"`
	} `rjson:"onResponseReceivedActions[0].appendContinuationItemsAction.continuationItems"`

	ContinueToken string `rjson:"onResponseReceivedActions[0].appendContinuationItemsAction.continuationItems[-]continuationItemRenderer.continuationEndpoint.continuationCommand.token"`
}

// NextVideosPage scrapes pages of the `/videos` endpoint on channel page
func (c *ChannelScraper) NextVideosPage() (videos []Video, err error) {
	if !c.videosInitialComplete {
		return c.runVideosInitial()
	} else {
		var resp *http.Response
		resp, err = http.Post("https://www.youtube.com/youtubei/v1/browse", "application/json", bytes.NewReader(c.videosContinueInputJson))
		if err != nil {
			return
		}
		c.videosContinueInputJson = []byte{}

		var body []byte
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return
		}

		debugFileOutput(body, "channel_videos.json")

		var output channelVideosContinueOutput
		if err = rjson.Unmarshal(body, &output); err != nil {
			if errors.Unwrap(err) == rjson.ErrCantFindField {
				err = nil
			}
			return
		}

		c.videosContinueInputJson, err = continueInput{Continuation: output.ContinueToken}.FillGenericInfo().Construct()
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
