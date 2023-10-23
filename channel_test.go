package scraper

import (
	"log"
	"testing"
)

func TestChannelVideosScraper(t *testing.T) {
	scraper := NewChannelScraper("@TomScottGo")

	var (
		videos         []Video
		err            error
		printedChannel bool
	)
	for {
		videos, err = scraper.NextVideosPage()
		if err != nil {
			t.Fatal(err)
		} else if len(videos) == 0 {
			break
		}

		if !printedChannel {
			if available, channel := scraper.GetChannelInfo(); available {
				log.Println(channel)
			}

			printedChannel = true
		}

		for _, video := range videos {
			t.Log(video.VideoID, video.Title, video.Views)
		}
	}
}
