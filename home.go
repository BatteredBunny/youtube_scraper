package youtube_scraper

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
)

type HomeVideosScraper struct {
	url string

	ChannelID    string
	NewChannelID string // @username

	InitialComplete   bool
	ContinueInput     continueInput
	ContinueInputJson []byte
}

const CollectShorts = false

func NewHomeVideosScraper() (h HomeVideosScraper) {
	h.url = "https://www.youtube.com/?hl=en"

	return
}

func (h *HomeVideosScraper) NextPage() (videos []Video, err error) {
	if !h.InitialComplete {
		var rawjson string
		rawjson, err = ExtractInitialData(h.url)
		if err != nil {
			return
		}

		if DEBUG {
			os.WriteFile("home_initial.json", []byte(rawjson), 0777)
		}

		var a initialData
		if err = json.Unmarshal([]byte(rawjson), &a); err != nil {
			return
		}

		h.ContinueInput.Context.Client.Hl = "en"
		h.ContinueInput.Context.Client.ClientName = "WEB"
		h.ContinueInput.Context.Client.ClientVersion = "2.20230706.00.00"
		h.ContinueInput.BrowseId = "FEwhat_to_watch"
		h.ContinueInput.InlineSettingStatus = "INLINE_SETTING_STATUS_ON"
		h.ContinueInput.Context.Client.VisitorData = a.ResponseContext.WebResponseContextExtensionData.YtConfigData.VisitorData

		for _, video := range a.Contents.TwoColumnBrowseResultsRenderer.Tabs[0].TabRenderer.Content.RichGridRenderer.Contents {
			// If it finds continuation token, fill it
			if video.ContinuationItemRenderer.ContinuationEndpoint.ContinuationCommand.Token != "" {
				h.ContinueInput.Continuation = video.ContinuationItemRenderer.ContinuationEndpoint.ContinuationCommand.Token

				h.ContinueInputJson, err = json.Marshal(h.ContinueInput)
				if err != nil {
					return
				}
			}

			if video.RichSectionRenderer.TrackingParams != "" {
				for _, shelf := range video.RichSectionRenderer.Content.RichShelfRenderer.Contents {
					if shelf.RichItemRenderer.Content.ReelItemRenderer.VideoId != "" { // shorts
						if CollectShorts {
							r := shelf.RichItemRenderer.Content.ReelItemRenderer
							if r.VideoId != "" {
								videos = append(videos, Video{
									VideoID: r.VideoId,
									Title:   r.Headline.SimpleText,
									Views:   r.ViewCountText.SimpleText,

									// Possibly get length from here, example: Daily dose of cute animals for you ❤️v29 Chill Lofi - 1 minute - play video
									//Length:  r.Accessibility.AccessibilityData.Label,
								})
							}
						}
					} else {
						r := shelf.RichItemRenderer.Content.VideoRenderer
						if r.VideoId != "" {
							videos = append(videos, Video{
								VideoID:      r.VideoId,
								Title:        r.Title.Runs[0].Text,
								Length:       r.LengthText.SimpleText,
								Views:        r.ViewCountText.SimpleText,
								Date:         r.PublishedTimeText.SimpleText,
								Username:     r.LongBylineText.Runs[0].Text,
								ChannelID:    r.LongBylineText.Runs[0].NavigationEndpoint.BrowseEndpoint.BrowseId,
								NewChannelID: strings.TrimPrefix(r.LongBylineText.Runs[0].NavigationEndpoint.BrowseEndpoint.CanonicalBaseUrl, "/"),
							})
						}
					}
				}
			} else {
				r := video.RichItemRenderer.Content.VideoRenderer
				if r.VideoId != "" {
					videos = append(videos, Video{
						VideoID:      r.VideoId,
						Title:        r.Title.Runs[0].Text,
						Length:       r.LengthText.SimpleText,
						Views:        r.ViewCountText.SimpleText,
						Date:         r.PublishedTimeText.SimpleText,
						Username:     r.LongBylineText.Runs[0].Text,
						ChannelID:    r.LongBylineText.Runs[0].NavigationEndpoint.BrowseEndpoint.BrowseId,
						NewChannelID: strings.TrimPrefix(r.LongBylineText.Runs[0].NavigationEndpoint.BrowseEndpoint.CanonicalBaseUrl, "/"),
					})
				}
			}
		}

		h.InitialComplete = true
	} else {
		var resp *http.Response
		resp, err = http.Post("https://www.youtube.com/youtubei/v1/browse", "application/json", bytes.NewReader(h.ContinueInputJson))
		if err != nil {
			return
		}
		h.ContinueInputJson = []byte{}

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
					h.ContinueInput.Continuation = rawVideo.ContinuationItemRenderer.ContinuationEndpoint.ContinuationCommand.Token

					h.ContinueInputJson, err = json.Marshal(h.ContinueInput)
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
						Username:     r.LongBylineText.Runs[0].Text,
						ChannelID:    r.LongBylineText.Runs[0].NavigationEndpoint.BrowseEndpoint.BrowseId,
						NewChannelID: strings.TrimPrefix(r.LongBylineText.Runs[0].NavigationEndpoint.BrowseEndpoint.CanonicalBaseUrl, "/"),
					})
				}
			}
		}
	}

	return
}
