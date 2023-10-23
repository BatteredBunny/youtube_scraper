package video

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	scraper "git.catnip.ee/miisu/youtube_scraper"
	"github.com/ayes-web/rjson"
)

type sidebarEntryType = int

const (
	SidebarEntryTypeVideo sidebarEntryType = iota
	SidebarEntryTypePlaylist
	SidebarEntryTypeRadio
)

type SidebarEntry struct {
	Type  sidebarEntryType
	Entry any
}

func (s SidebarEntry) IsVideo() bool {
	return s.Type == SidebarEntryTypeVideo
}

func (s SidebarEntry) IsPlaylist() bool {
	return s.Type == SidebarEntryTypePlaylist
}

func (s SidebarEntry) IsRadio() bool {
	return s.Type == SidebarEntryTypeRadio
}

type SidebarVideo struct {
	VideoID      string
	Title        string
	Username     string
	ChannelID    string
	NewChannelID string
	Date         string
	Views        int
	Viewers      int
	Length       string
	Thumbnails   []scraper.YoutubeImage

	AuthorIsVerified, AuthorIsVerifiedArtist bool
	HasNewBadge, HasCCBadge, Has4kBadge      bool

	IsLive, WasLive bool
}
type SidebarPlaylist struct {
	PlaylistID string
	Title      string

	Username     string
	ChannelId    string
	NewChannelID string

	VideosAmount     int
	ThumbnailVideoID string
	Thumbnails       []scraper.YoutubeImage
}
type SidebarRadio struct {
	PlaylistID     string
	Title          string
	SecondaryTitle string

	VideosAmount     int
	ThumbnailVideoID string
	Thumbnails       []scraper.YoutubeImage
}

func (sidebarEntry rawSidebarEntry) ToSidebarEntry() (s SidebarEntry, err error) {
	if sidebarEntry.Video.VideoID != "" {
		var hasNewBadge, hasCCBadge, has4kBadge bool
		for _, badge := range sidebarEntry.Video.Badges {
			switch badge {
			case scraper.VideoBadgeNew:
				hasNewBadge = true
			case scraper.VideoBadgeCC:
				hasCCBadge = true
			case scraper.VideoBadge4k:
				has4kBadge = true
			}
		}

		var authorIsVerifiedArtist, authorIsVerified bool
		for _, ownerBadge := range sidebarEntry.Video.OwnerBadges {
			switch ownerBadge {
			case scraper.ChannelBadgeVerified:
				authorIsVerified = true
			case scraper.ChannelBadgeVerifiedArtistChannel:
				authorIsVerifiedArtist = true
			}
		}

		date, wasLive := strings.CutPrefix(sidebarEntry.Video.Date, "Streamed ")

		var views float64
		views, err = scraper.ParseViews(sidebarEntry.Video.Views)
		if err != nil {
			return
		}

		var viewers int
		if sidebarEntry.Video.Viewers != "" {
			viewers, err = strconv.Atoi(strings.ReplaceAll(strings.TrimSuffix(sidebarEntry.Video.Viewers, " watching"), ",", ""))
			if err != nil {
				return
			}
		}

		s = SidebarEntry{
			Type: SidebarEntryTypeVideo,
			Entry: SidebarVideo{
				VideoID:      sidebarEntry.Video.VideoID,
				Title:        sidebarEntry.Video.Title,
				Username:     sidebarEntry.Video.Username,
				ChannelID:    sidebarEntry.Video.ChannelID,
				NewChannelID: strings.TrimPrefix(sidebarEntry.Video.RawNewChannelID, "/"),
				Date:         date,
				Views:        int(views),
				Viewers:      viewers,
				Length:       sidebarEntry.Video.Length,

				AuthorIsVerified: authorIsVerified, AuthorIsVerifiedArtist: authorIsVerifiedArtist,
				HasNewBadge: hasNewBadge, HasCCBadge: hasCCBadge, Has4kBadge: has4kBadge,

				IsLive:     len(sidebarEntry.Video.Date) == 0,
				WasLive:    wasLive,
				Thumbnails: sidebarEntry.Video.Thumbnails,
			},
		}
	} else if sidebarEntry.Playlist.PlaylistID != "" {
		var videosAmount int
		videosAmount, err = strconv.Atoi(scraper.FixUnit(strings.ReplaceAll(sidebarEntry.Playlist.VideosAmount, ",", "")))
		if err != nil {
			return
		}

		s = SidebarEntry{
			Type: SidebarEntryTypePlaylist,
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
		videosAmount, err = strconv.Atoi(scraper.FixUnit(strings.ReplaceAll(sidebarEntry.Playlist.VideosAmount, ",", "")))
		if err != nil {
			return
		}

		s = SidebarEntry{
			Type: SidebarEntryTypeRadio,
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

	scraper.DebugFileOutput(body, "sidebar_videos_%s.json", v.sidebarContinueInput.Continuation)

	var output sidebarOutput
	if err = rjson.Unmarshal(body, &output); err != nil {
		return
	}

	v.sidebarContinueInput = scraper.ContinueInput{Continuation: output.ContinueToken}.FillGenericInfo()
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
