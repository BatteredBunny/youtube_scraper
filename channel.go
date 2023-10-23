package youtube_scraper

import (
	"bytes"
	"encoding/json"
	"io"
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
	url string

	channel Channel
	// ChannelID    string
	// NewChannelID string // @username

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

func constructContinue(token string) (ContinueInputJson []byte, err error) {
	continueInput := continueInput{}
	continueInput.Context.Client.Hl = "en"
	continueInput.Context.Client.ClientName = "WEB"
	continueInput.Context.Client.ClientVersion = "2.20230706.00.00"
	continueInput.Continuation = token

	ContinueInputJson, err = json.Marshal(continueInput)
	return
}

func (c *ChannelVideosScraper) runInitial() (videos []Video, err error) {
	var (
		a       initialData
		rawjson string
	)
	rawjson, err = ExtractInitialData(c.url)
	if err != nil {
		return
	}

	if DEBUG {
		os.WriteFile("channel_initial.json", []byte(rawjson), 0777)
	}

	json.Unmarshal([]byte(rawjson), &a)

	// channel
	c.channel = Channel{
		Subscribers:  a.Header.C4TabbedHeaderRenderer.SubscriberCountText.SimpleText,
		ChannelID:    a.Metadata.ChannelMetadataRenderer.ExternalId,
		NewChannelID: a.Header.C4TabbedHeaderRenderer.ChannelHandleText.Runs[0].Text,
		Username:     a.Metadata.ChannelMetadataRenderer.Title,
		Description:  a.Metadata.ChannelMetadataRenderer.Description,
	}

	for _, badge := range a.Header.C4TabbedHeaderRenderer.Badges {
		if badge.MetadataBadgeRenderer.Tooltip == "Verified" {
			c.channel.IsVerified = true
		}
	}

	rawVideosAmount := a.Header.C4TabbedHeaderRenderer.VideosCountText.Runs[0].Text
	if rawVideosAmount == "No videos" {
		c.channel.VideosAmount = 0
	} else {
		c.channel.VideosAmount, err = strconv.Atoi(rawVideosAmount)
		if err != nil {
			return
		}
	}

	// videos
	for _, tab := range a.Contents.TwoColumnBrowseResultsRenderer.Tabs {
		if strings.HasSuffix(tab.TabRenderer.Endpoint.CommandMetadata.WebCommandMetadata.Url, "/videos") {
			for _, video := range tab.TabRenderer.Content.RichGridRenderer.Contents {
				r := video.RichItemRenderer.Content.VideoRenderer
				if r.VideoId == "" {
					c.ContinueInputJson, err = constructContinue(video.ContinuationItemRenderer.ContinuationEndpoint.ContinuationCommand.Token)
					if err != nil {
						return
					}
				} else {
					videos = append(videos, Video{
						VideoID:      r.VideoId,
						Title:        r.Title.Runs[0].Text,
						Length:       r.LengthText.SimpleText,
						Views:        r.ViewCountText.SimpleText,
						Date:         r.PublishedTimeText.SimpleText,
						ChannelID:    c.channel.ChannelID,
						NewChannelID: c.channel.NewChannelID,
					})
				}
			}
		}
	}

	c.InitialComplete = true
	return
}

func (c *ChannelVideosScraper) NextPage() (videos []Video, err error) {
	var resp *http.Response

	if !c.InitialComplete {
		videos, err = c.runInitial()
		if err != nil {
			return
		}
	} else {
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

		var output continueOutput
		if err = json.Unmarshal(body, &output); err != nil {
			return
		}

		if len(output.OnResponseReceivedActions) > 0 {
			for _, rawVideo := range output.OnResponseReceivedActions[0].AppendContinuationItemsAction.ContinuationItems {
				r := rawVideo.RichItemRenderer.Content.VideoRenderer
				if r.VideoId == "" {
					c.ContinueInputJson, err = constructContinue(rawVideo.ContinuationItemRenderer.ContinuationEndpoint.ContinuationCommand.Token)
					if err != nil {
						return
					}
				} else {
					videos = append(videos, Video{
						VideoID:      r.VideoId,
						Title:        r.Title.Runs[0].Text,
						Length:       r.LengthText.SimpleText,
						Views:        r.ViewCountText.SimpleText,
						Date:         r.PublishedTimeText.SimpleText,
						ChannelID:    c.channel.ChannelID,
						NewChannelID: c.channel.NewChannelID,
					})
				}
			}
		}
	}

	return
}
