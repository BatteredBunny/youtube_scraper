package examples

import (
	"encoding/json"
	"testing"

	scraper "github.com/ayes-web/youtube_scraper"
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
		t.Log("-------------")
	}
}

func TestHomeContinue(t *testing.T) {
	h := scraper.NewHomeVideosScraper()
	h.NextPage()

	var err error
	h, err = scraper.HomeVideosScraperFromExport(h.Export())
	if err != nil {
		t.Fatal(err)
	}

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
		t.Log("-------------")
	}
}
