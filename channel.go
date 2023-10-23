package scraper

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/ayes-web/rjson"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Channel struct {
	Subscribers      string
	IsVerified       bool
	IsVerifiedArtist bool
	ChannelID        string
	NewChannelID     string
	Username         string
	Description      string
	VideosAmount     string // e.g "15", "1.5K"
}

type Video struct {
	VideoID string `json:"VideoID"`
	Title   string `json:"Title"`

	// Will be empty if its livestream
	// example value 7:03
	Length string `json:"Length,omitempty"`

	/*
		Will be empty if its livestream
		Examples
			- 100 views
			- 6,396,043 views
	*/
	Views string `json:"Views,omitempty"`

	/*
		Will be empty if its livestream

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
	Date string `json:"Date,omitempty"`

	Username     string `json:"Username"`
	ChannelID    string `json:"ChannelID"`
	NewChannelID string `json:"NewChannelID"` // @username

	Viewers                string `json:"Viewers,omitempty"` // Empty if it's not a livestream
	IsLive                 bool   `json:"IsLive"`
	WasLive                bool   `json:"WasLive"`
	AuthorIsVerified       bool   `json:"AuthorIsVerified"`
	AuthorIsVerifiedArtist bool   `json:"AuthorIsVerifiedArtist"`
}

type ChannelScraper struct {
	//baseChannelUrl *url.URL
	streamsUrl string
	videosUrl  string
	channel    Channel

	videosInitialComplete   bool
	videosContinueInputJson []byte

	streamsInitialComplete   bool
	streamsContinueInputJson []byte
}

// NewChannelScraper accepts normal id or @username
func NewChannelScraper(id string) (c ChannelScraper, err error) {
	rawUrl, err := url.Parse("https://www.youtube.com/")
	if err != nil {
		return
	}

	if strings.HasPrefix(id, "@") {
		rawUrl = rawUrl.JoinPath(id)
	} else {
		rawUrl = rawUrl.JoinPath("channel", id)
	}

	rawVideosUrl := rawUrl.JoinPath("videos")
	q := rawVideosUrl.Query()
	q.Set("hl", "en")
	rawVideosUrl.RawQuery = q.Encode()
	c.videosUrl = rawVideosUrl.String()

	rawStreamsUrl := rawUrl.JoinPath("streams")
	q = rawStreamsUrl.Query()
	q.Set("hl", "en")
	rawStreamsUrl.RawQuery = q.Encode()
	c.streamsUrl = rawStreamsUrl.String()

	return
}

// GetChannelInfo will output the internal channel struct which will become available after the first call to NextVideosPage() or NextStreamsPage()
func (c *ChannelScraper) GetChannelInfo() (available bool, channel Channel) {
	if c.videosInitialComplete || c.streamsInitialComplete {
		channel = c.channel
		available = true
	} else {
		available = false
	}

	return
}

type channelInitialAccount struct {
	Subscribers     string   `rjson:"header.c4TabbedHeaderRenderer.subscriberCountText.simpleText"`
	ChannelID       string   `rjson:"metadata.channelMetadataRenderer.externalId"`
	NewChannelID    string   `rjson:"header.c4TabbedHeaderRenderer.channelHandleText.runs[0].text"`
	Username        string   `rjson:"metadata.channelMetadataRenderer.title"`
	Description     string   `rjson:"metadata.channelMetadataRenderer.description"`
	RawVideosAmount string   `rjson:"header.c4TabbedHeaderRenderer.videosCountText.runs[0].text"`
	Badges          []string `rjson:"header.c4TabbedHeaderRenderer.badges[].metadataBadgeRenderer.tooltip"`
}

// videoRenderer json type
type videoRenderer struct {
	VideoID string `rjson:"videoId"`
	Title   string `rjson:"title.runs[0].text"`
	Length  string `rjson:"lengthText.simpleText"`
	Views   string `rjson:"viewCountText.simpleText"`
	Viewers string `rjson:"viewCountText.runs[0].text"`
	Date    string `rjson:"publishedTimeText.simpleText"`
}

func (video videoRenderer) ToVideo(channel *Channel) Video {
	date, wasLive := strings.CutPrefix(video.Date, "Streamed ")
	return Video{
		VideoID:                video.VideoID,
		Title:                  video.Title,
		Length:                 video.Length,
		Views:                  video.Views,
		Viewers:                video.Viewers,
		Date:                   date,
		ChannelID:              channel.ChannelID,
		NewChannelID:           channel.NewChannelID,
		WasLive:                wasLive,
		IsLive:                 len(video.Viewers) > 0,
		AuthorIsVerified:       channel.IsVerified,
		AuthorIsVerifiedArtist: channel.IsVerifiedArtist,
	}
}

type channelVideosInitialOutput struct {
	Channel                 channelInitialAccount `rjson:"."`
	Videos                  []videoRenderer       `rjson:"contents.twoColumnBrowseResultsRenderer.tabs[1].tabRenderer.content.richGridRenderer.contents[].richItemRenderer.content.videoRenderer"`
	VideosContinuationToken string                `rjson:"contents.twoColumnBrowseResultsRenderer.tabs[1].tabRenderer.content.richGridRenderer.contents[-].continuationItemRenderer.continuationEndpoint.continuationCommand.token"`
}
type channelStreamsInitialOutput struct {
	Channel channelInitialAccount `rjson:"."`
	Tabs    []struct {
		Title  string          `rjson:"title"`
		Token  string          `rjson:"content.richGridRenderer.contents[-].continuationItemRenderer.continuationEndpoint.continuationCommand.token"`
		Videos []videoRenderer `rjson:"content.richGridRenderer.contents[].richItemRenderer.content.videoRenderer"`
	} `rjson:"contents.twoColumnBrowseResultsRenderer.tabs[].tabRenderer"`
}

type channelContinueOutput struct {
	Videos        []videoRenderer `rjson:"onResponseReceivedActions[0].appendContinuationItemsAction.continuationItems[].richItemRenderer.content.videoRenderer"`
	ContinueToken string          `rjson:"onResponseReceivedActions[0].appendContinuationItemsAction.continuationItems[-]continuationItemRenderer.continuationEndpoint.continuationCommand.token"`
}

func genericChannelInitial(initialComplete *bool, url string, channel *Channel, continueInputJson *[]byte, outputGeneric func(rawJson []byte) (rawChannel channelInitialAccount, rawVideos []videoRenderer, rawToken string, err error)) (videos []Video, err error) {
	var rawJson string
	rawJson, err = extractInitialData(url)
	if err != nil {
		return
	}

	rawChannel, rawVideos, rawToken, err := outputGeneric([]byte(rawJson))
	if err != nil {
		return
	}

	*channel = Channel{
		Subscribers:  rawChannel.Subscribers,
		ChannelID:    rawChannel.ChannelID,
		NewChannelID: rawChannel.NewChannelID,
		Username:     rawChannel.Username,
		Description:  rawChannel.Description,
	}

	for _, badge := range rawChannel.Badges {
		switch badge {
		case "Verified":
			channel.IsVerified = true
		case "Official Artist Channel":
			channel.IsVerifiedArtist = true
		}
	}

	channel.VideosAmount = rawChannel.RawVideosAmount
	if rawChannel.RawVideosAmount == "No videos" {
		channel.VideosAmount = "0"
	}

	*continueInputJson, err = continueInput{Continuation: rawToken}.FillGenericInfo().Construct()
	if err != nil {
		return
	}

	for _, video := range rawVideos {
		if video.VideoID != "" {
			videos = append(videos, video.ToVideo(channel))
		}
	}

	*initialComplete = true
	return
}

func genericChannelPage(channel *Channel, continueInputJson *[]byte, outputGeneric func(rawJson []byte) (rawToken string, rawVideos []videoRenderer, err error)) (videos []Video, err error) {
	var resp *http.Response
	resp, err = http.Post("https://www.youtube.com/youtubei/v1/browse", "application/json", bytes.NewReader(*continueInputJson))
	if err != nil {
		return
	}
	*continueInputJson = []byte{}

	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	rawToken, rawVideos, err := outputGeneric(body)
	if err != nil {
		return
	}

	*continueInputJson, err = continueInput{Continuation: rawToken}.FillGenericInfo().Construct()
	if err != nil {
		return
	}

	for _, video := range rawVideos {
		if video.VideoID != "" {
			videos = append(videos, video.ToVideo(channel))
		}
	}

	return
}

// NextVideosPage scrapes pages of the `/videos` endpoint on channel page
func (c *ChannelScraper) NextVideosPage() (videos []Video, err error) {
	if !c.videosInitialComplete {
		return genericChannelInitial(&c.videosInitialComplete, c.videosUrl, &c.channel, &c.videosContinueInputJson, func(rawJson []byte) (rawChannel channelInitialAccount, rawVideos []videoRenderer, rawToken string, err error) {
			debugFileOutput(rawJson, "channel_videos_initial.json")

			var output channelVideosInitialOutput
			if err = rjson.Unmarshal(rawJson, &output); err != nil {
				if errors.Is(err, rjson.ErrCantFindField) {
					if Debug {
						log.Println("WARNING:", err)
					}
					err = nil
				}
				return
			}

			rawChannel = output.Channel
			rawVideos = output.Videos
			rawToken = output.VideosContinuationToken

			return
		})
	} else {
		return genericChannelPage(&c.channel, &c.videosContinueInputJson, func(rawJson []byte) (rawToken string, rawVideos []videoRenderer, err error) {
			debugFileOutput(rawJson, "channel_videos.json")

			var output channelContinueOutput
			if err = rjson.Unmarshal(rawJson, &output); err != nil {
				if errors.Is(err, rjson.ErrCantFindField) {
					err = nil
				}
				return
			}

			rawToken = output.ContinueToken
			rawVideos = output.Videos
			return
		})
	}
}

// NextStreamsPage scrapes pages of the `/streams` endpoint on channel page
func (c *ChannelScraper) NextStreamsPage() (videos []Video, err error) {
	if !c.streamsInitialComplete {
		videos, err = genericChannelInitial(&c.streamsInitialComplete, c.streamsUrl, &c.channel, &c.streamsContinueInputJson, func(rawJson []byte) (rawChannel channelInitialAccount, rawVideos []videoRenderer, rawToken string, err error) {
			debugFileOutput(rawJson, "channel_streams_initial.json")

			var output channelStreamsInitialOutput
			if err = rjson.Unmarshal(rawJson, &output); err != nil {
				if errors.Is(err, rjson.ErrCantFindField) {
					if Debug {
						log.Println("WARNING:", err)
					}
					err = nil
				}
				return
			}

			rawChannel = output.Channel

			for _, tab := range output.Tabs {
				if tab.Title == "Live" {
					rawVideos = tab.Videos
					rawToken = tab.Token
				}
			}

			return
		})
		if err != nil {
			return
		}

		// fix for pagination api sometimes not working
		if len(videos) == 0 {
			return c.NextStreamsPage()
		} else {
			return
		}
	} else {
		// fix for pagination api sometimes not working
		for i := 0; i < 3; i++ {
			videos, err = genericChannelPage(&c.channel, &c.streamsContinueInputJson, func(rawJson []byte) (rawToken string, rawVideos []videoRenderer, err error) {
				debugFileOutput(rawJson, "channel_streams.json")

				var output channelContinueOutput
				if err = rjson.Unmarshal(rawJson, &output); err != nil {
					if errors.Is(errors.Unwrap(err), rjson.ErrCantFindField) {
						err = nil
					}
					return
				}

				fmt.Println(rawToken)
				rawToken = output.ContinueToken
				rawVideos = output.Videos

				return
			})

			if err != nil {
				return
			}

			if len(videos) > 0 {
				break
			}
		}

		return
	}
}
