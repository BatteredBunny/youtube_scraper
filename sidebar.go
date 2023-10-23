package scraper

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/ayes-web/rjson"
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
	Views           int
	Viewers         int
	Length          string
	Thumbnails      []YoutubeImage

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

	VideosAmount     int
	ThumbnailVideoID string
	Thumbnails       []YoutubeImage
}
type SidebarRadio struct {
	PlaylistID     string
	Title          string
	SecondaryTitle string

	VideosAmount     int
	ThumbnailVideoID string
	Thumbnails       []YoutubeImage
}

type compactVideoRenderer struct {
	VideoID         string         `rjson:"videoId"`
	Title           string         `rjson:"title.simpleText"`
	Username        string         `rjson:"longBylineText.runs[0].text"`
	ChannelID       string         `rjson:"longBylineText.runs[0].navigationEndpoint.browseEndpoint.browseId"`
	RawNewChannelID string         `rjson:"longBylineText.runs[0].navigationEndpoint.browseEndpoint.canonicalBaseUrl"` // has "/" at start that must be trimmed
	Date            string         `rjson:"publishedTimeText.simpleText"`
	Views           string         `rjson:"viewCountText.simpleText"`
	Viewers         string         `rjson:"viewCountText.runs[0].text"`
	Length          string         `rjson:"lengthText.simpleText"`
	Badges          []string       `rjson:"badges[].metadataBadgeRenderer.label"`        // example of badge "New" or "CC"
	OwnerBadges     []string       `rjson:"ownerBadges[].metadataBadgeRenderer.tooltip"` // example of owner badge "Verified" or "Official Artist Channel"
	Thumbnails      []YoutubeImage `rjson:"thumbnail.thumbnails"`
}

type compactPlaylistRenderer struct {
	PlaylistID      string `rjson:"playlistId"`
	Title           string `rjson:"title.simpleText"`
	Username        string `rjson:"shortBylineText.runs[0].text"`
	ChannelID       string `rjson:"shortBylineText.runs[0].navigationEndpoint.browseEndpoint.browseId"`
	RawNewChannelID string `rjson:"shortBylineText.runs[0].navigationEndpoint.browseEndpoint.canonicalBaseUrl"` // has "/" at start that must be trimmed

	VideosAmount     string         `rjson:"videoCountShortText.simpleText"`
	ThumbnailVideoID string         `rjson:"navigationEndpoint.watchEndpoint.videoId"`
	Thumbnails       []YoutubeImage `rjson:"thumbnail.thumbnails"`
}

type compactRadioRenderer struct {
	RadioPlaylistID string `rjson:"playlistId"`
	Title           string `rjson:"title.simpleText"`
	SecondaryTitle  string `rjson:"longBylineText.simpleText"`

	ThumbnailVideoID string         `rjson:"navigationEndpoint.watchEndpoint.videoId"`
	VideosAmount     string         `rjson:"videoCountShortText.runs[0].text"`
	Thumbnails       []YoutubeImage `rjson:"thumbnail.thumbnails"`
}

type rawSidebarEntry struct {
	Video    compactVideoRenderer    `rjson:"compactVideoRenderer"`
	Playlist compactPlaylistRenderer `rjson:"compactPlaylistRenderer"`
	Radio    compactRadioRenderer    `rjson:"compactRadioRenderer"`
}

func (sidebarEntry rawSidebarEntry) ToSidebarEntry() (s SidebarEntry, err error) {
	if sidebarEntry.Video.VideoID != "" {
		var isNew bool
		for _, badge := range sidebarEntry.Video.Badges {
			switch badge {
			case VideoBadgeNew:
				isNew = true
			}
		}

		var authorIsVerifiedArtist bool
		var authorIsVerified bool
		for _, ownerBadge := range sidebarEntry.Video.OwnerBadges {
			switch ownerBadge {
			case BadgeChannelVerified:
				authorIsVerified = true
			case BadgeChannelVerifiedArtistChannel:
				authorIsVerifiedArtist = true
			}
		}

		date, wasLive := strings.CutPrefix(sidebarEntry.Video.Date, "Streamed ")

		var views int
		if sidebarEntry.Video.Views != "" {
			views, err = strconv.Atoi(strings.ReplaceAll(strings.TrimSuffix(sidebarEntry.Video.Views, " views"), ",", ""))
			if err != nil {
				return
			}
		}

		var viewers int
		if sidebarEntry.Video.Viewers != "" {
			viewers, err = strconv.Atoi(strings.ReplaceAll(sidebarEntry.Video.Viewers, ",", ""))
			if err != nil {
				return
			}
		}

		s = SidebarEntry{
			Type: SidebarEntryVideo,
			Entry: SidebarVideo{
				VideoID:                sidebarEntry.Video.VideoID,
				Title:                  sidebarEntry.Video.Title,
				Username:               sidebarEntry.Video.Username,
				ChannelID:              sidebarEntry.Video.ChannelID,
				RawNewChannelID:        strings.TrimPrefix(sidebarEntry.Video.RawNewChannelID, "/"),
				Date:                   date,
				Views:                  views,
				Viewers:                viewers,
				Length:                 sidebarEntry.Video.Length,
				AuthorIsVerified:       authorIsVerified,
				AuthorIsVerifiedArtist: authorIsVerifiedArtist,
				IsNew:                  isNew,
				IsLive:                 len(sidebarEntry.Video.Date) == 0,
				WasLive:                wasLive,
				Thumbnails:             sidebarEntry.Video.Thumbnails,
			},
		}
	} else if sidebarEntry.Playlist.PlaylistID != "" {
		var videosAmount int
		videosAmount, err = strconv.Atoi(fixUnit(strings.ReplaceAll(sidebarEntry.Playlist.VideosAmount, ",", "")))
		if err != nil {
			return
		}

		s = SidebarEntry{
			Type: SidebarEntryPlaylist,
			Entry: SidebarPlaylist{
				PlaylistID:       sidebarEntry.Playlist.PlaylistID,
				Title:            sidebarEntry.Playlist.Title,
				Username:         sidebarEntry.Playlist.Username,
				ChannelId:        sidebarEntry.Playlist.ChannelID,
				NewChannelID:     strings.TrimPrefix(sidebarEntry.Playlist.RawNewChannelID, "/"),
				VideosAmount:     videosAmount,
				ThumbnailVideoID: sidebarEntry.Playlist.ThumbnailVideoID,
				Thumbnails:       sidebarEntry.Playlist.Thumbnails,
			},
		}
	} else if sidebarEntry.Radio.RadioPlaylistID != "" {
		var videosAmount int
		videosAmount, err = strconv.Atoi(fixUnit(strings.ReplaceAll(sidebarEntry.Playlist.VideosAmount, ",", "")))
		if err != nil {
			return
		}

		s = SidebarEntry{
			Type: SidebarEntryRadio,
			Entry: SidebarRadio{
				PlaylistID:       sidebarEntry.Radio.RadioPlaylistID,
				Title:            sidebarEntry.Radio.Title,
				SecondaryTitle:   sidebarEntry.Radio.SecondaryTitle,
				VideosAmount:     videosAmount,
				ThumbnailVideoID: sidebarEntry.Radio.ThumbnailVideoID,
				Thumbnails:       sidebarEntry.Radio.Thumbnails,
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

	debugFileOutput(body, "sidebar_videos_%s.json", v.sidebarContinueInput.Continuation)

	var output sidebarOutput
	if err = rjson.Unmarshal(body, &output); err != nil {
		return
	}

	v.sidebarContinueInput = continueInput{Continuation: output.ContinueToken}.FillGenericInfo()
	v.sidebarContinueInputJson, err = v.sidebarContinueInput.Construct()
	if err != nil {
		return
	}

	for _, sidebarEntry := range output.SidebarEntries {
		if sidebarEntry.Video.VideoID != "" || sidebarEntry.Playlist.PlaylistID != "" || sidebarEntry.Radio.RadioPlaylistID != "" {
			if entry, err := sidebarEntry.ToSidebarEntry(); err != nil {
				log.Println("WARNING converting to sidebar failed:", err)
			} else {
				sidebarEntries = append(sidebarEntries, entry)
			}
		}
	}

	return
}
