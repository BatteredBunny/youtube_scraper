package examples

import (
	scraper "git.catnip.ee/miisu/youtube_scraper"
	"testing"
)

func TestPlaylist(t *testing.T) {
	p, err := scraper.NewPlaylistScraper("PLMC9KNkIncKtPzgY-5rmhvj7fax8fdxoj")
	if err != nil {
		t.Fatal(err)
	}

	info, err := p.GetPlaylistInfo()
	if err != nil {
		t.Fatal(err)
	}

	handlePlaylistVideos(info.Videos, t)

	var (
		videos []scraper.PlaylistVideo
	)
	for {
		videos, err = p.NextPage()
		if err != nil {
			t.Fatal(err)
		} else if len(videos) == 0 {
			break
		}

		handlePlaylistVideos(videos, t)
	}
}

func handlePlaylistVideos(videos []scraper.PlaylistVideo, t *testing.T) {
	for _, video := range videos {
		t.Logf("videoID: %s, pos: %d", video.VideoID, video.PlaylistPosition)
	}
	t.Log("-----------------")
}
