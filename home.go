package scraper

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/ayes-web/rjson"
)

type HomeVideosScraper struct {
	url string

	initialComplete   bool
	continueInput     ContinueInput
	continueInputJson []byte
}

func HomeVideosScraperFromExport(export HomeVideosExport) (h HomeVideosScraper, err error) {
	h.initialComplete = export.InitialComplete
	h.url = "https://www.youtube.com/?hl=en"
	h.continueInput = ContinueInput{
		BrowseId:            "FEwhat_to_watch",
		InlineSettingStatus: "INLINE_SETTING_STATUS_ON",
		Continuation:        export.ContinueToken,
	}.FillGenericInfo()

	h.continueInput.Context.Client.VisitorData = export.VisitorData
	h.continueInputJson, err = h.continueInput.Construct()
	if err != nil {
		return
	}

	return
}

func NewHomeVideosScraper() (h HomeVideosScraper) {
	h.url = "https://www.youtube.com/?hl=en"

	return
}

type HomeVideosExport struct {
	ContinueToken   string
	VisitorData     string
	InitialComplete bool
}

func (h *HomeVideosScraper) Export() HomeVideosExport {
	return HomeVideosExport{
		ContinueToken:   h.continueInput.Continuation,
		VisitorData:     h.continueInput.Context.Client.VisitorData,
		InitialComplete: h.initialComplete,
	}
}

// home has a modified version of videoRenderer with few additional lines of info, maybe best to merge them fully?
type homeVideo struct {
	VideoID    string         `rjson:"videoId"`
	Title      string         `rjson:"title.runs[0].text"`
	Length     string         `rjson:"lengthText.simpleText"`
	Views      string         `rjson:"viewCountText.simpleText"`
	Viewers    string         `rjson:"viewCountText.runs[0].text"`
	Date       string         `rjson:"publishedTimeText.simpleText"`
	Thumbnails []YoutubeImage `rjson:"thumbnail.thumbnails"`

	ChannelAvatar   string   `rjson:"channelThumbnailSupportedRenderers.channelThumbnailWithLinkRenderer.thumbnail.thumbnails[0].url"`
	Username        string   `rjson:"longBylineText.runs[0].text"`
	ChannelID       string   `rjson:"longBylineText.runs[0].navigationEndpoint.browseEndpoint.browseId"`
	RawNewChannelID string   `rjson:"longBylineText.runs[0].navigationEndpoint.browseEndpoint.canonicalBaseUrl"` // comes with "/" at start, make sure to trim it
	OwnerBadges     []string `rjson:"ownerBadges[].metadataBadgeRenderer.tooltip"`
}

func (video homeVideo) ToVideo() (v Video, err error) {
	var authorIsVerified, authorIsVerifiedArtist bool
	for _, ownerBadge := range video.OwnerBadges {
		switch ownerBadge {
		case ChannelBadgeVerified:
			authorIsVerified = true
		case ChannelBadgeVerifiedArtistChannel:
			authorIsVerifiedArtist = true
		}
	}

	date, wasLive := strings.CutPrefix(video.Date, "Streamed ")

	views, err := ParseViews(video.Views)
	if err != nil {
		return
	}

	var viewers int
	if video.Viewers != "" {
		viewers, err = strconv.Atoi(FixUnit(strings.ReplaceAll(strings.TrimSuffix(video.Viewers, " watching"), ",", "")))
		if err != nil {
			return
		}
	}

	v = Video{
		VideoID:                video.VideoID,
		Title:                  video.Title,
		Length:                 video.Length,
		Views:                  int(views),
		Viewers:                viewers,
		Date:                   date,
		Username:               video.Username,
		ChannelID:              video.ChannelID,
		NewChannelID:           strings.TrimPrefix(video.RawNewChannelID, "/"),
		IsLive:                 len(video.Viewers) > 0,
		WasLive:                wasLive,
		AuthorIsVerified:       authorIsVerified,
		AuthorIsVerifiedArtist: authorIsVerifiedArtist,
		ChannelAvatar:          video.ChannelAvatar,
		Thumbnails:             video.Thumbnails,
	}

	return
}

type homeInitialOutput struct {
	VisitorData       string `rjson:"responseContext.webResponseContextExtensionData.ytConfigData.visitorData"`
	ContinuationToken string `rjson:"contents.twoColumnBrowseResultsRenderer.tabs[0].tabRenderer.content.richGridRenderer.contents[-].continuationItemRenderer.continuationEndpoint.continuationCommand.token"`

	Videos []struct {
		Video homeVideo `rjson:"richItemRenderer.content.videoRenderer"`

		ShelfName  string `rjson:"richSectionRenderer.content.richShelfRenderer.title.runs[0].text"`
		ShelfItems []struct {
			Short struct {
				VideoId string `rjson:"videoId"`
				Title   string `rjson:"headline.simpleText"`
				Views   string `rjson:"viewCountText.simpleText"`
				//Possibly parse length from here, example: Daily dose of cute animals for you ❤️v29 Chill Lofi - 1 minute - play VideoInfo
				//Length string `rjson:"accessibility.accessibilityData.label"`
			} `rjson:"richItemRenderer.content.reelItemRenderer"`
			Video homeVideo `rjson:"richItemRenderer.content.videoRenderer"`
		} `rjson:"richSectionRenderer.content.richShelfRenderer.contents"`
	} `rjson:"contents.twoColumnBrowseResultsRenderer.tabs[0].tabRenderer.content.richGridRenderer.contents"`
}

func (h *HomeVideosScraper) runInitial() (videos []Video, err error) {
	var rawJson string
	rawJson, err = ExtractInitialData(h.url)
	if err != nil {
		return
	}

	DebugFileOutput([]byte(rawJson), "home_initial.json")

	var output homeInitialOutput
	if err = rjson.Unmarshal([]byte(rawJson), &output); err != nil {
		if errors.Is(err, rjson.ErrCantFindField) {
			if Debug {
				log.Println("WARNING:", err)
			}
			err = nil
		}
		return
	}

	h.continueInput = ContinueInput{
		BrowseId:            "FEwhat_to_watch",
		InlineSettingStatus: "INLINE_SETTING_STATUS_ON",
		Continuation:        output.ContinuationToken,
	}.FillGenericInfo()

	h.continueInput.Context.Client.VisitorData = output.VisitorData
	h.continueInputJson, err = h.continueInput.Construct()
	if err != nil {
		return
	}

	for _, video := range output.Videos {
		if video.ShelfName != "" {
			// TODO: return shelves in api
			//fmt.Println("Reading shelf", VideoInfo.ShelfName)
			//isShortsShelf := VideoInfo.ShelfItems[0].Short.VideoId != ""
			//for _, gen := range VideoInfo.ShelfItems {
			//	if isShortsShelf {
			//		fmt.Println(gen.Short)
			//	} else {
			//		fmt.Println(gen.VideoInfo)
			//	}
			//}
		} else if video.Video.VideoID != "" {
			if v, err := video.Video.ToVideo(); err != nil {
				log.Println("WARNING error while converting video:", err)
			} else {
				videos = append(videos, v)
			}
		}
	}

	h.initialComplete = true
	return
}

type homeContinueOutput struct {
	Videos        []homeVideo `rjson:"onResponseReceivedActions[0].appendContinuationItemsAction.continuationItems[].richItemRenderer.content.videoRenderer"`
	ContinueToken string      `rjson:"onResponseReceivedActions[0].appendContinuationItemsAction.continuationItems[-]continuationItemRenderer.continuationEndpoint.continuationCommand.token"`
}

func (h *HomeVideosScraper) NextPage() (videos []Video, err error) {
	if !h.initialComplete {
		return h.runInitial()
	} else {
		var resp *http.Response
		resp, err = http.Post("https://www.youtube.com/youtubei/v1/browse", "application/json", bytes.NewReader(h.continueInputJson))
		if err != nil {
			return
		}

		h.continueInputJson = []byte{}

		var body []byte
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return
		}

		DebugFileOutput(body, "home_videos_%s.json", h.continueInput.Continuation)

		var output homeContinueOutput
		if err = rjson.Unmarshal(body, &output); err != nil {
			if errors.Is(err, rjson.ErrCantFindField) {
				if Debug {
					log.Println("WARNING:", err)
				}
				err = nil
			}
			return
		}

		h.continueInput.Continuation = output.ContinueToken
		h.continueInputJson, err = json.Marshal(h.continueInput)
		if err != nil {
			return
		}

		for _, video := range output.Videos {
			if video.VideoID != "" {
				if v, err := video.ToVideo(); err != nil {
					log.Println("WARNING error while converting video:", err)
				} else {
					videos = append(videos, v)
				}
			}
		}
	}

	return
}
