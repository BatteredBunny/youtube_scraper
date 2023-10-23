package youtube_scraper

import "testing"

func TestHomeVideosScraper(t *testing.T) {
	h := NewHomeVideosScraper()

	// Gets 5 pages of home page feed
	for i := 0; i <= 5; i++ {
		videos, err := h.NextPage()
		if err != nil {
			t.Fatal(err)
		} else if len(videos) == 0 {
			break
		}

		for _, video := range videos {
			t.Log(video.VideoID, video.Title, video.Views)
		}
	}
}
