package examples

import (
	"encoding/json"
	scraper "git.catnip.ee/miisu/youtube_scraper"
	"testing"
)

func TestPlaylist(t *testing.T) {
	p := scraper.NewPlaylistScraper("PLMC9KNkIncKtPzgY-5rmhvj7fax8fdxoj")
	p.GetPlaylistInfo()

	var (
		videos []scraper.PlaylistVideo
		err    error
		r      []byte
	)
	for {
		videos, err = p.NextPage()
		if err != nil {
			t.Fatal(err)
		} else if len(videos) == 0 {
			break
		}

		for _, video := range videos {
			r, err = json.MarshalIndent(video, "", "	")
			if err != nil {
				t.Fatal(err)
			}
			t.Log("video: ", string(r))
		}
		t.Log("-----------------")
	}
}
