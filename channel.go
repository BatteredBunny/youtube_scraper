package youtube_scraper

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

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

	ChannelID    string `json:"ChannelID"`
	NewChannelID string `json:"NewChannelID"` // @username
}

type ChannelVideosScraper struct {
	url string

	ChannelID    string
	NewChannelID string // @username

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

func constructContinue(token string) (ContinueInputJson []byte, err error) {
	continueInput := accountScrapeContinueInput{}
	continueInput.Context.Client.Hl = "en"
	continueInput.Context.Client.ClientName = "WEB"
	continueInput.Context.Client.ClientVersion = "2.20230706.00.00"
	continueInput.Continuation = token

	ContinueInputJson, err = json.Marshal(continueInput)
	return
}

func (c *ChannelVideosScraper) NextPage() (videos []Video, err error) {
	var resp *http.Response

	if !c.InitialComplete {
		resp, err = http.Get(c.url)
		if err != nil {
			return
		}

		var doc *goquery.Document
		doc, err = goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			return
		}

		var a accountScrapeInitial
		doc.Find("script").Each(func(i int, s *goquery.Selection) {
			if cut, valid := strings.CutPrefix(s.Text(), "var ytInitialData = "); valid {
				j, _ := strings.CutSuffix(cut, ";")
				json.Unmarshal([]byte(j), &a)
			}
		})

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
						if c.NewChannelID == "" {
							c.NewChannelID = a.Header.C4TabbedHeaderRenderer.ChannelHandleText.Runs[0].Text
						}
						if c.ChannelID == "" {
							c.ChannelID = a.Contents.TwoColumnBrowseResultsRenderer.Tabs[0].TabRenderer.Endpoint.BrowseEndpoint.BrowseId
						}

						videos = append(videos, Video{
							VideoID:      r.VideoId,
							Title:        r.Title.Runs[0].Text,
							Length:       r.LengthText.SimpleText,
							Views:        r.ViewCountText.SimpleText,
							Date:         r.PublishedTimeText.SimpleText,
							ChannelID:    c.ChannelID,
							NewChannelID: c.NewChannelID,
						})
					}
				}
			}
		}

		c.InitialComplete = true
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

		var output accountScrapeContinueOutput
		if err = json.Unmarshal(body, &output); err != nil {
			return
		}

		if len(output.OnResponseReceivedActions) > 0 {
			for _, rawVideo := range output.OnResponseReceivedActions[0].AppendContinuationItemsAction.ContinuationItems {
				if rawVideo.RichItemRenderer.TrackingParams == "" {
					c.ContinueInputJson, err = constructContinue(rawVideo.ContinuationItemRenderer.ContinuationEndpoint.ContinuationCommand.Token)
					if err != nil {
						return
					}
				} else {
					r := rawVideo.RichItemRenderer.Content.VideoRenderer

					videos = append(videos, Video{
						VideoID:      r.VideoId,
						Title:        r.Title.Runs[0].Text,
						Length:       r.LengthText.SimpleText,
						Views:        r.ViewCountText.SimpleText,
						Date:         r.PublishedTimeText.SimpleText,
						ChannelID:    c.ChannelID,
						NewChannelID: c.NewChannelID,
					})
				}
			}
		}
	}

	return
}
