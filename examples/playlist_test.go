package examples

import (
	"testing"

	scraper "github.com/ayes-web/youtube_scraper"
)

func TestPlaylist(t *testing.T) {
	p, err := scraper.NewPlaylistScraper("PLJV9alv4vklceQIPKDx9X-9nZsINLQDqF")
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
