package scraper

import (
	"encoding/json"
	"testing"
)

func TestSidebarVideos(t *testing.T) {
	Debug = true
	scraper, err := NewVideoScraper("FdbvrqC6lOY")
	if err != nil {
		t.Fatal(err)
	}

	var sidebarVideos []SidebarVideo
	for i := 0; i <= 3; i++ {
		sidebarVideos, err = scraper.NextSidebarVideosPage()
		if err != nil {
			t.Fatal(err)
		} else if len(sidebarVideos) == 0 {
			break
		}

		for _, sidebarVideo := range sidebarVideos {
			bs, err := json.MarshalIndent(sidebarVideo, "", "	")
			if err != nil {
				t.Fatal(err)
			}

			t.Log(string(bs))
		}
	}
}
