package search

import (
	"log"
	"strconv"
	"strings"

	scraper "github.com/ayes-web/youtube_scraper"
	"github.com/dustin/go-humanize"
)

func (rawVideo searchVideoRenderer) ToVideo() (video SearchVideo, err error) {
	var truncatedDescription string
	for _, snippet := range rawVideo.DescriptionSnippets {
		truncatedDescription += snippet
	}

	var hasNewBadge, hasCCBadge, has4kBadge bool
	for _, badge := range rawVideo.Badges {
		switch badge {
		case scraper.VideoBadge4k:
			has4kBadge = true
		case scraper.VideoBadgeCC:
			hasCCBadge = true
		case scraper.VideoBadgeNew:
			hasNewBadge = true
		}
	}

	var isVerified, isVerifiedArtist bool
	for _, badge := range rawVideo.OwnerBadges {
		switch badge {
		case scraper.ChannelBadgeVerified:
			isVerified = true
		case scraper.ChannelBadgeVerifiedArtistChannel:
			isVerifiedArtist = true
		}
	}

	var viewers int
	if rawVideo.Viewers != "" {
		viewers, err = strconv.Atoi(strings.ReplaceAll(strings.TrimSuffix(rawVideo.Viewers, " watching"), ",", ""))
		if err != nil {
			return
		}
	}

	var views float64
	views, err = scraper.ParseViews(rawVideo.Views)
	if err != nil {
		return
	}

	video = SearchVideo{
		VideoID: rawVideo.VideoID,
		Title:   rawVideo.Title,
		Date:    rawVideo.Date,
		Length:  rawVideo.Length,

		Views:         int(views),
		Viewers:       viewers,
		IsLive:        rawVideo.Date == "" || rawVideo.Length == "" || viewers > 0,
		Thumbnails:    rawVideo.Thumbnails,
		ChannelAvatar: rawVideo.ChannelAvatar,

		Username:             rawVideo.Username,
		ChannelID:            rawVideo.ChannelID,
		TruncatedDescription: truncatedDescription,
		NewChannelID:         strings.TrimPrefix(rawVideo.NewChannelID, "/"),

		HasNewBadge: hasNewBadge, HasCCBadge: hasCCBadge, Has4kBadge: has4kBadge,
		AuthorIsVerified: isVerified, AuthorIsVerifiedArtist: isVerifiedArtist,
	}
	return
}

func (rawChannel searchChannelRenderer) ToChannel() (channel SearchChannel, err error) {
	var subscribers float64
	if rawChannel.RawSubscribers != "" {
		var unit string
		subscribers, unit, err = humanize.ParseSI(scraper.FixUnit(strings.TrimSuffix(rawChannel.RawSubscribers, " subscribers")))
		if err != nil {
			log.Println("meow:", rawChannel.RawSubscribers)
			return
		} else if unit != "" {
			log.Printf("WARNING: possibly wrong number for channel subscribers count: %f%s\n", subscribers, unit)
		}
	}

	var bioSnippet string
	for _, snippet := range rawChannel.RawBioSnippet {
		bioSnippet += snippet
	}

	var isVerified, isVerifiedArtist bool
	for _, badge := range rawChannel.OwnerBadges {
		switch badge {
		case scraper.ChannelBadgeVerified:
			isVerified = true
		case scraper.ChannelBadgeVerifiedArtistChannel:
			isVerifiedArtist = true
		}
	}

	channel = SearchChannel{
		ChannelID:        rawChannel.ChannelID,
		Username:         rawChannel.Username,
		NewChannelID:     rawChannel.NewChannelID,
		Avatars:          rawChannel.Avatars,
		BioSnippet:       bioSnippet,
		Subscribers:      int(subscribers),
		IsVerified:       isVerified,
		IsVerifiedAuthor: isVerifiedArtist,
	}
	return
}

func (rawPlaylist searchPlaylistRenderer) ToPlaylist() (playlist SearchPlaylist, err error) {
	var videoCount int
	videoCount, err = strconv.Atoi(strings.ReplaceAll(rawPlaylist.RawVideoCount, ",", ""))
	if err != nil {
		return
	}

	newChannelID := strings.TrimPrefix(rawPlaylist.NewChannelID, "/")

	playlist = SearchPlaylist{
		PlaylistID:   rawPlaylist.PlaylistID,
		Username:     rawPlaylist.Username,
		NewChannelID: newChannelID,
		Thumbnails:   rawPlaylist.Thumbnails,
		VideoCount:   videoCount,
		Title:        rawPlaylist.Title,
	}
	return
}

func (searchEntry rawSearchEntry) ToSearchEntry() (s SearchEntry, err error) {
	if searchEntry.Video.VideoID != "" {
		if video, err := searchEntry.Video.ToVideo(); err != nil {
			log.Println("WARNING error while converting search result video:", err)
		} else {
			s = SearchEntry{
				Type:  SearchEntryTypeVideo,
				Entry: video,
			}
		}
	} else if searchEntry.Channel.ChannelID != "" {
		if channel, err := searchEntry.Channel.ToChannel(); err != nil {
			log.Println("WARNING error while converting search result channel:", err)
		} else {
			s = SearchEntry{
				Type:  SearchEntryTypeChannel,
				Entry: channel,
			}
		}
	} else if searchEntry.Playlist.PlaylistID != "" {
		if playlist, err := searchEntry.Playlist.ToPlaylist(); err != nil {
			log.Println("WARNING error while converting search result playlist:", err)
		} else {
			s = SearchEntry{
				Type:  SearchEntryTypePlaylist,
				Entry: playlist,
			}
		}
	}

	return
}
