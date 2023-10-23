package examples

import (
	"encoding/json"
	"git.catnip.ee/miisu/youtube_scraper"
	"testing"
)

func TestHomeVideosScraper(t *testing.T) {
	h := scraper.NewHomeVideosScraper()

	// Gets 5 pages of home page feed
	for i := 0; i <= 5; i++ {
		videos, err := h.NextPage()
		if err != nil {
			t.Fatal(err)
		} else if len(videos) == 0 {
			break
		}

		for _, video := range videos {
			bs, err := json.MarshalIndent(video, "", "	")
			if err != nil {
				t.Fatal(err)
			}
			t.Log(string(bs))
		}
	}
}