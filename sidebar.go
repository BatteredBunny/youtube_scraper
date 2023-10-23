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

type compactVideoRenderer struct {
	VideoID         string   `rjson:"videoId"`
	Title           string   `rjson:"title.simpleText"`
	Username        string   `rjson:"longBylineText.runs[0].text"`
	ChannelID       string   `rjson:"longBylineText.runs[0].navigationEndpoint.browseEndpoint.browseId"`
	RawNewChannelID string   `rjson:"longBylineText.runs[0].navigationEndpoint.browseEndpoint.canonicalBaseUrl"` // has "/" at start that must be trimmed
	Date            string   `rjson:"publishedTimeText.simpleText"`
	Views           string   `rjson:"viewCountText.simpleText"`
	Length          string   `rjson:"lengthText.simpleText"`
	Badges          []string `rjson:"badges[].metadataBadgeRenderer.label"`        // example of badge "New"
	OwnerBadges     []string `rjson:"ownerBadges[].metadataBadgeRenderer.tooltip"` // example of owner badge "Verified"
}

type compactPlaylistRenderer struct {
	PlaylistID      string `rjson:"playlistId"`
	Title           string `rjson:"title.simpleText"`
	Username        string `rjson:"shortBylineText.runs[0].text"`
	ChannelID       string `rjson:"shortBylineText.runs[0].navigationEndpoint.browseEndpoint.browseId"`
	RawNewChannelID string `rjson:"shortBylineText.runs[0].navigationEndpoint.browseEndpoint.canonicalBaseUrl"` // has "/" at start that must be trimmed

	VideosAmount     string `rjson:"videoCountShortText.simpleText"`
	ThumbnailVideoID string `rjson:"navigationEndpoint.watchEndpoint.videoId"`
}

type compactRadioRenderer struct {
	RadioPlaylistID string `rjson:"playlistId"`
	Title           string `rjson:"title.simpleText"`
	SecondaryTitle  string `rjson:"longBylineText.simpleText"`

	ThumbnailVideoID string `rjson:"navigationEndpoint.watchEndpoint.videoId"`
	VideosAmount     string `rjson:"videoCountShortText.runs[0].text"`
}

type rawSidebarEntry struct {
	Video    compactVideoRenderer    `rjson:"compactVideoRenderer"`
	Playlist compactPlaylistRenderer `rjson:"compactPlaylistRenderer"`
	Radio    compactRadioRenderer    `rjson:"compactRadioRenderer"`
}

func (sidebarEntry rawSidebarEntry) ToSidebarEntry() (s SidebarEntry) {
	if sidebarEntry.Video.VideoID != "" {
		var isNew bool
		for _, badge := range sidebarEntry.Video.Badges {
			switch badge {
			case "New":
				isNew = true
			}
		}

		var authorIsVerifiedArtist bool
		var authorIsVerified bool
		for _, ownerBadge := range sidebarEntry.Video.OwnerBadges {
			switch ownerBadge {
			case "Verified":
				authorIsVerified = true
			case "Official Artist Channel":
				authorIsVerifiedArtist = true
			}
		}

		date, wasLive := strings.CutPrefix(sidebarEntry.Video.Date, "Streamed ")
		s = SidebarEntry{
			Type: SidebarEntryVideo,
			Entry: SidebarVideo{
				VideoID:                sidebarEntry.Video.VideoID,
				Title:                  sidebarEntry.Video.Title,
				Username:               sidebarEntry.Video.Username,
				ChannelID:              sidebarEntry.Video.ChannelID,
				RawNewChannelID:        strings.TrimPrefix(sidebarEntry.Video.RawNewChannelID, "/"),
				Date:                   date,
				Views:                  sidebarEntry.Video.Views,
				Length:                 sidebarEntry.Video.Length,
				AuthorIsVerified:       authorIsVerified,
				AuthorIsVerifiedArtist: authorIsVerifiedArtist,
				IsNew:                  isNew,
				IsLive:                 len(sidebarEntry.Video.Date) == 0,
				WasLive:                wasLive,
			},
		}
	} else if sidebarEntry.Playlist.PlaylistID != "" {
		s = SidebarEntry{
			Type: SidebarEntryPlaylist,
			Entry: SidebarPlaylist{
				PlaylistID:       sidebarEntry.Playlist.PlaylistID,
				Title:            sidebarEntry.Playlist.Title,
				Username:         sidebarEntry.Playlist.Username,
				ChannelId:        sidebarEntry.Playlist.ChannelID,
				NewChannelID:     strings.TrimPrefix(sidebarEntry.Playlist.RawNewChannelID, "/"),
				VideosAmount:     sidebarEntry.Playlist.VideosAmount,
				ThumbnailVideoID: sidebarEntry.Playlist.ThumbnailVideoID,
			},
		}
	} else if sidebarEntry.Radio.RadioPlaylistID != "" {
		s = SidebarEntry{
			Type: SidebarEntryRadio,
			Entry: SidebarRadio{
				PlaylistID:       sidebarEntry.Radio.RadioPlaylistID,
				Title:            sidebarEntry.Radio.Title,
				SecondaryTitle:   sidebarEntry.Radio.SecondaryTitle,
				VideosAmount:     sidebarEntry.Radio.VideosAmount,
				ThumbnailVideoID: sidebarEntry.Radio.ThumbnailVideoID,
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

	debugFileOutput(body, "sidebar_videos_%s.json", v.sidebarToken)

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
		if sidebarEntry.Video.VideoID != "" || sidebarEntry.Playlist.PlaylistID != "" || sidebarEntry.Radio.RadioPlaylistID != "" {
			sidebarEntries = append(sidebarEntries, sidebarEntry.ToSidebarEntry())
		}
	}

	return
}
