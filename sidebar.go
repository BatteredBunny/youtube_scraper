package scraper

import (
	"bytes"
	"github.com/ayes-web/rjson"
	"io"
	"net/http"
	"strings"
)

type sidebarEntryType = int

const (
	SidebarEntryVideo sidebarEntryType = iota
	SidebarEntryPlaylist
	SidebarEntryRadio
)

type SidebarEntry struct {
	Type  sidebarEntryType
	Entry any
}

type SidebarVideo struct {
	VideoID         string
	Title           string
	Username        string
	ChannelID       string
	RawNewChannelID string
	Date            string
	Views           string
	Length          string

	AuthorIsVerified       bool
	AuthorIsVerifiedArtist bool
	IsNew                  bool
	IsLive                 bool
	WasLive                bool
}
type SidebarPlaylist struct {
	PlaylistID string
	Title      string

	Username     string
	ChannelId    string
	NewChannelID string

	VideosAmount     string
	ThumbnailVideoID string
}
type SidebarRadio struct {
	PlaylistID     string
	Title          string
	SecondaryTitle string

	VideosAmount     string
	ThumbnailVideoID string
}

type rawSidebarEntry struct {
	VideoID         string   `rjson:"compactVideoRenderer.videoId"`
	Title           string   `rjson:"compactVideoRenderer.title.simpleText"`
	Username        string   `rjson:"compactVideoRenderer.longBylineText.runs[0].text"`
	ChannelID       string   `rjson:"compactVideoRenderer.longBylineText.runs[0].navigationEndpoint.browseEndpoint.browseId"`
	RawNewChannelID string   `rjson:"compactVideoRenderer.longBylineText.runs[0].navigationEndpoint.browseEndpoint.canonicalBaseUrl"`
	Date            string   `rjson:"compactVideoRenderer.publishedTimeText.simpleText"`
	Views           string   `rjson:"compactVideoRenderer.viewCountText.simpleText"`
	Length          string   `rjson:"compactVideoRenderer.lengthText.simpleText"`
	Badges          []string `rjson:"compactVideoRenderer.badges[].metadataBadgeRenderer.label"`        // example of badge "New"
	OwnerBadges     []string `rjson:"compactVideoRenderer.ownerBadges[].metadataBadgeRenderer.tooltip"` // example of owner badge "Verified"

	// compactPlaylistRenderer
	PlaylistID               string `rjson:"compactPlaylistRenderer.playlistId"`
	PlaylistTitle            string `rjson:"compactPlaylistRenderer.title.simpleText"`
	PlaylistUsername         string `rjson:"compactPlaylistRenderer.shortBylineText.runs[0].text"`
	PlaylistChannelID        string `rjson:"compactPlaylistRenderer.shortBylineText.runs[0].navigationEndpoint.browseEndpoint.browseId"`
	PlaylistRawNewChannelID  string `rjson:"compactPlaylistRenderer.shortBylineText.runs[0].navigationEndpoint.browseEndpoint.canonicalBaseUrl"`
	PlaylistVideosAmount     string `rjson:"compactPlaylistRenderer.videoCountShortText.simpleText"`
	PlaylistThumbnailVideoID string `rjson:"compactPlaylistRenderer.navigationEndpoint.watchEndpoint.videoId"`

	// compactRadioRenderer
	RadioPlaylistID       string `rjson:"compactRadioRenderer.playlistId"`
	RadioTitle            string `rjson:"compactRadioRenderer.title.simpleText"`
	RadioSecondaryTitle   string `rjson:"compactRadioRenderer.longBylineText.simpleText"`
	RadioThumbnailVideoID string `rjson:"compactRadioRenderer.navigationEndpoint.watchEndpoint.videoId"`
	RadioVideosAmount     string `rjson:"compactRadioRenderer.videoCountShortText.runs[0].text"`
}

func (sidebarEntry rawSidebarEntry) ToSidebarEntry() (s SidebarEntry) {
	if sidebarEntry.VideoID != "" {
		var isNew bool
		for _, badge := range sidebarEntry.Badges {
			switch badge {
			case "New":
				isNew = true
			}
		}

		var authorIsVerifiedArtist bool
		var authorIsVerified bool
		for _, ownerBadge := range sidebarEntry.OwnerBadges {
			switch ownerBadge {
			case "Verified":
				authorIsVerified = true
			case "Official Artist Channel":
				authorIsVerifiedArtist = true
			}
		}

		date, wasLive := strings.CutPrefix(sidebarEntry.Date, "Streamed ")
		s = SidebarEntry{
			Type: SidebarEntryVideo,
			Entry: SidebarVideo{
				VideoID:                sidebarEntry.VideoID,
				Title:                  sidebarEntry.Title,
				Username:               sidebarEntry.Username,
				ChannelID:              sidebarEntry.ChannelID,
				RawNewChannelID:        strings.TrimPrefix(sidebarEntry.RawNewChannelID, "/"),
				Date:                   date,
				Views:                  sidebarEntry.Views,
				Length:                 sidebarEntry.Length,
				AuthorIsVerified:       authorIsVerified,
				AuthorIsVerifiedArtist: authorIsVerifiedArtist,
				IsNew:                  isNew,
				IsLive:                 len(sidebarEntry.Date) == 0,
				WasLive:                wasLive,
			},
		}
	} else if sidebarEntry.PlaylistID != "" {
		s = SidebarEntry{
			Type: SidebarEntryPlaylist,
			Entry: SidebarPlaylist{
				PlaylistID:       sidebarEntry.PlaylistID,
				Title:            sidebarEntry.PlaylistTitle,
				Username:         sidebarEntry.PlaylistUsername,
				ChannelId:        sidebarEntry.PlaylistChannelID,
				NewChannelID:     strings.TrimPrefix(sidebarEntry.PlaylistRawNewChannelID, "/"),
				VideosAmount:     sidebarEntry.PlaylistVideosAmount,
				ThumbnailVideoID: sidebarEntry.PlaylistThumbnailVideoID,
			},
		}
	} else if sidebarEntry.RadioPlaylistID != "" {
		s = SidebarEntry{
			Type: SidebarEntryRadio,
			Entry: SidebarRadio{
				PlaylistID:       sidebarEntry.RadioPlaylistID,
				Title:            sidebarEntry.RadioTitle,
				SecondaryTitle:   sidebarEntry.RadioSecondaryTitle,
				VideosAmount:     sidebarEntry.RadioVideosAmount,
				ThumbnailVideoID: sidebarEntry.RadioThumbnailVideoID,
			},
		}
	}

	return
}

type sidebarOutput struct {
	SidebarEntries []rawSidebarEntry `rjson:"onResponseReceivedEndpoints[0].appendContinuationItemsAction.continuationItems"`
	ContinueToken  string            `rjson:"onResponseReceivedEndpoints[0].appendContinuationItemsAction.continuationItems[-].continuationItemRenderer.continuationEndpoint.continuationCommand.token"`
}

func (v *VideoScraper) NextSidebarVideosPage() (sidebarEntries []SidebarEntry, err error) {
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

	for _, sidebarEntry := range output.SidebarEntries {
		if sidebarEntry.VideoID == "" && sidebarEntry.PlaylistID == "" && sidebarEntry.RadioPlaylistID == "" {
			continue
		}

		sidebarEntries = append(sidebarEntries, sidebarEntry.ToSidebarEntry())
	}

	return
}
