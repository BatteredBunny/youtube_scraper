package scraper

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/ayes-web/rjson"
	"io"
	"log"
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

func NewHomeVideosScraper() (h HomeVideosScraper) {
	h.url = "https://www.youtube.com/?hl=en"

	return
}

type homeInitialOutputVideo struct {
	VideoID         string `rjson:"videoId"`
	Title           string `rjson:"title.runs[0].text"`
	Length          string `rjson:"lengthText.simpleText"`
	Views           string `rjson:"viewCountText.simpleText"`
	Date            string `rjson:"publishedTimeText.simpleText"`
	Username        string `rjson:"longBylineText.runs[0].text"`
	ChannelID       string `rjson:"longBylineText.runs[0].navigationEndpoint.browseEndpoint.browseId"`
	RawNewChannelID string `rjson:"longBylineText.runs[0].navigationEndpoint.browseEndpoint.canonicalBaseUrl"`
}

type homeInitialOutput struct {
	VisitorData       string `rjson:"responseContext.webResponseContextExtensionData.ytConfigData.visitorData"`
	ContinuationToken string `rjson:"contents.twoColumnBrowseResultsRenderer.tabs[0].tabRenderer.content.richGridRenderer.contents[-].continuationItemRenderer.continuationEndpoint.continuationCommand.token"`

	Videos []struct {
		Video homeInitialOutputVideo `rjson:"richItemRenderer.content.videoRenderer"`

		ShelfName  string `rjson:"richSectionRenderer.content.richShelfRenderer.title.runs[0].text"`
		ShelfItems []struct {
			Short struct {
				VideoId string `rjson:"videoId"`
				Title   string `rjson:"headline.simpleText"`
				Views   string `rjson:"viewCountText.simpleText"`
				//Possibly parse length from here, example: Daily dose of cute animals for you ❤️v29 Chill Lofi - 1 minute - play video
				//Length string `rjson:"accessibility.accessibilityData.label"`
			} `rjson:"richItemRenderer.content.reelItemRenderer"`
			Video homeInitialOutputVideo `rjson:"richItemRenderer.content.videoRenderer"`
		} `rjson:"richSectionRenderer.content.richShelfRenderer.contents"`
	} `rjson:"contents.twoColumnBrowseResultsRenderer.tabs[0].tabRenderer.content.richGridRenderer.contents"`
}

func (h *HomeVideosScraper) runInitial() (videos []Video, err error) {
	var rawjson string
	rawjson, err = ExtractInitialData(h.url)
	if err != nil {
		return
	}

	if Debug {
		os.WriteFile("home_initial.json", []byte(rawjson), 0777)
	}

	var output homeInitialOutput
	if err = rjson.Unmarshal([]byte(rawjson), &output); err != nil {
		if errors.Unwrap(err) == rjson.ErrCantFindField {
			if Debug {
				log.Println("WARNING:", err)
			}
			err = nil
		}
		return
	}

	h.ContinueInput = continueInput{
		BrowseId:            "FEwhat_to_watch",
		InlineSettingStatus: "INLINE_SETTING_STATUS_ON",
		Continuation:        output.ContinuationToken,
	}.FillGenericInfo()

	h.ContinueInput.Context.Client.VisitorData = output.VisitorData
	h.ContinueInputJson, err = h.ContinueInput.Construct()
	if err != nil {
		return
	}

	for _, video := range output.Videos {
		if video.ShelfName != "" {
			// TODO: return shelves in api
			//fmt.Println("Reading shelf", video.ShelfName)
			//isShortsShelf := video.ShelfItems[0].Short.VideoId != ""
			//for _, gen := range video.ShelfItems {
			//	if isShortsShelf {
			//		fmt.Println(gen.Short)
			//	} else {
			//		fmt.Println(gen.Video)
			//	}
			//}
		} else {
			if video.Video.VideoID == "" {
				continue
			}

			videos = append(videos, Video{
				VideoID:      video.Video.VideoID,
				Title:        video.Video.Title,
				Length:       video.Video.Length,
				Views:        video.Video.Views,
				Date:         video.Video.Date,
				Username:     video.Video.Username,
				ChannelID:    video.Video.ChannelID,
				NewChannelID: strings.TrimPrefix(video.Video.RawNewChannelID, "/"),
			})
		}
	}

	h.InitialComplete = true
	return
}

type homeContinueOutput struct {
	Videos []struct {
		VideoID         string `rjson:"richItemRenderer.content.videoRenderer.videoId"`
		Title           string `rjson:"richItemRenderer.content.videoRenderer.title.runs[0].text"`
		Length          string `rjson:"richItemRenderer.content.videoRenderer.lengthText.simpleText"`
		Views           string `rjson:"richItemRenderer.content.videoRenderer.viewCountText.simpleText"`
		Date            string `rjson:"richItemRenderer.content.videoRenderer.publishedTimeText.simpleText"`
		Username        string `rjson:"richItemRenderer.content.videoRenderer.longBylineText.runs[0].text"`
		ChannelID       string `rjson:"richItemRenderer.content.videoRenderer.longBylineText.runs[0].navigationEndpoint.browseEndpoint.browseId"`
		RawNewChannelID string `rjson:"richItemRenderer.content.videoRenderer.longBylineText.runs[0].navigationEndpoint.browseEndpoint.canonicalBaseUrl"`
	} `rjson:"onResponseReceivedActions[0].appendContinuationItemsAction.continuationItems"`

	ContinueToken string `rjson:"onResponseReceivedActions[0].appendContinuationItemsAction.continuationItems[-]continuationItemRenderer.continuationEndpoint.continuationCommand.token"`
}

func (h *HomeVideosScraper) NextPage() (videos []Video, err error) {
	if !h.InitialComplete {
		return h.runInitial()
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

		if Debug {
			os.WriteFile("home_videos.json", body, 0777)
		}

		var output homeContinueOutput
		if err = rjson.Unmarshal(body, &output); err != nil {
			if errors.Unwrap(err) == rjson.ErrCantFindField {
				if Debug {
					log.Println("WARNING:", err)
				}
				err = nil
			}
			return
		}

		h.ContinueInput.Continuation = output.ContinueToken
		h.ContinueInputJson, err = json.Marshal(h.ContinueInput)
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
				Username:     video.Username,
				ChannelID:    video.ChannelID,
				NewChannelID: strings.TrimPrefix(video.RawNewChannelID, "/"),
			})
		}
	}

	return
}
