package youtube_scraper

import (
	"log"
	"testing"
)

func TestChannelVideosScraper(t *testing.T) {
	c := NewChannelVideosScraper("@TomScottGo")

	var (
		videos         []Video
		err            error
		printedChannel bool
	)
	for {
		videos, err = c.NextPage()
		if err != nil {
			t.Fatal(err)
		} else if len(videos) == 0 {
			break
		}

		if !printedChannel {
			if available, channel := c.GetChannelInfo(); available {
				log.Println(channel)
			}
		}

		for _, video := range videos {
			t.Log(video.VideoID, video.Title, video.Views)
		}
	}
}
