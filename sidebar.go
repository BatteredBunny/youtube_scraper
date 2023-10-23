package scraper

import (
	"bytes"
	"github.com/ayes-web/rjson"
	"io"
	"net/http"
	"strings"
)

type SidebarVideo struct {
	VideoID         string
	Title           string
	Username        string
	ChannelID       string
	RawNewChannelID string
	Date            string
	Views           string
	Length          string

	AuthorIsVerified bool
	IsNew            bool
	IsLive           bool
	WasLive          bool
}

type rawSidebarVideo struct {
	VideoID         string `rjson:"compactVideoRenderer.videoId"`
	Title           string `rjson:"compactVideoRenderer.title.simpleText"`
	Username        string `rjson:"compactVideoRenderer.longBylineText.runs[0].text"`
	ChannelID       string `rjson:"compactVideoRenderer.longBylineText.runs[0].navigationEndpoint.browseEndpoint.browseId"`
	RawNewChannelID string `rjson:"compactVideoRenderer.longBylineText.runs[0].navigationEndpoint.browseEndpoint.canonicalBaseUrl"`
	Date            string `rjson:"compactVideoRenderer.publishedTimeText.simpleText"`
	Views           string `rjson:"compactVideoRenderer.viewCountText.simpleText"`
	Length          string `rjson:"compactVideoRenderer.lengthText.simpleText"`

	Badges      []string `rjson:"compactVideoRenderer.badges[].metadataBadgeRenderer.label"`        // example of badge "New"
	OwnerBadges []string `rjson:"compactVideoRenderer.ownerBadges[].metadataBadgeRenderer.tooltip"` // example of owner badge "Verified"

	// TODO: parse playlist info
	// compactPlaylistRenderer
}

func (sidebarVideo rawSidebarVideo) ToSidebarVideo() SidebarVideo {
	var isNew bool
	for _, badge := range sidebarVideo.Badges {
		switch badge {
		case "New":
			isNew = true
		}
	}

	var authorIsVerified bool
	for _, ownerBadge := range sidebarVideo.OwnerBadges {
		switch ownerBadge {
		case "Verified":
			authorIsVerified = true
		}
	}

	date, wasLive := strings.CutPrefix(sidebarVideo.Date, "Streamed ")
	return SidebarVideo{
		VideoID:          sidebarVideo.VideoID,
		Title:            sidebarVideo.Title,
		Username:         sidebarVideo.Username,
		ChannelID:        sidebarVideo.ChannelID,
		RawNewChannelID:  strings.TrimPrefix(sidebarVideo.RawNewChannelID, "/"),
		Date:             date,
		Views:            sidebarVideo.Views,
		Length:           sidebarVideo.Length,
		AuthorIsVerified: authorIsVerified,
		IsNew:            isNew,
		IsLive:           len(sidebarVideo.Date) == 0,
		WasLive:          wasLive,
	}
}

type sidebarOutput struct {
	SidebarVideos []rawSidebarVideo `rjson:"onResponseReceivedEndpoints[0].appendContinuationItemsAction.continuationItems"`
	ContinueToken string            `rjson:"onResponseReceivedEndpoints[0].appendContinuationItemsAction.continuationItems[-].continuationItemRenderer.continuationEndpoint.continuationCommand.token"`
}

func (v *VideoScraper) NextSidebarVideosPage() (sidebarVideos []SidebarVideo, err error) {
	var resp *http.Response
	resp, err = http.Post("https://www.youtube.com/youtubei/v1/next", "application/json", bytes.NewReader(v.sidebarContinueInputJson))
	if err != nil {
		return
	}
	v.sidebarContinueInputJson = []byte{}

	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	debugFileOutput(body, "sidebarvideos_%s.json", v.sidebarToken)

	var output sidebarOutput
	if err = rjson.Unmarshal(body, &output); err != nil {
		return
	}

	v.sidebarToken = output.ContinueToken
	v.sidebarContinueInputJson, err = continueInput{Continuation: v.sidebarToken}.FillGenericInfo().Construct()
	if err != nil {
		return
	}

	for _, sidebarVideo := range output.SidebarVideos {
		if sidebarVideo.VideoID == "" {
			continue
		}

		sidebarVideos = append(sidebarVideos, sidebarVideo.ToSidebarVideo())
	}

	return
}
