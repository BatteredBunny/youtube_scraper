package youtube_scraper

import (
	"testing"
)

func TestChannelVideosScraper(t *testing.T) {
	c := NewChannelVideosScraper("@TomScottGo")
	videos, err := c.NextPage()
	if err != nil {
		t.Fatal(err)
	}

	for _, video := range videos {
		t.Log(video.VideoID, video.Title, video.Views)
	}

	if available, channel := c.GetChannelInfo(); available {
		t.Log(channel)
	}

	for {
		videos, err := c.NextPage()
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
