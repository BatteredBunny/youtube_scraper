package youtube_scraper

import "testing"

func TestVideoScraper(t *testing.T) {
	scraper, err := NewVideoScraper("51o5J0XVGoc")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(scraper.GetVideo())
}
