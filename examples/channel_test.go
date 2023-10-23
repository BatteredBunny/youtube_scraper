package examples

import (
	"encoding/json"
	"testing"

	scraper "git.catnip.ee/miisu/youtube_scraper"
)

func TestJustChannelInfo(t *testing.T) {
	c, err := scraper.NewChannelScraper("@TomScottGo")
	if err != nil {
		t.Fatal(err)
	}

	if _, err = c.NextVideosPage(); err != nil {
		t.Fatal(err)
	}

	_, info := c.GetChannelInfo()

	bs, err := json.MarshalIndent(info, "", "	")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("info:", string(bs))
}

func TestChannelVideosScraper(t *testing.T) {
	c, err := scraper.NewChannelScraper("@TomScottGo")
	if err != nil {
		t.Fatal(err)
	}

	var printedChannel bool
	for {
		videos, err := c.NextVideosPage()
		if err != nil {
			t.Fatal(err)
		} else if len(videos) == 0 {
			break
		}

		if !printedChannel {
			if available, channel := c.GetChannelInfo(); available {
				bs, err := json.MarshalIndent(channel, "", "	")
				if err != nil {
					t.Fatal(err)
				}
				t.Log(string(bs))
			}

			printedChannel = true
		}

		for _, video := range videos {
			t.Log(video.VideoID, video.Title, video.Views)
		}
	}
}

func TestChannelStreamsScraper(t *testing.T) {
	c, err := scraper.NewChannelScraper("@LinusTechTips")
	if err != nil {
		t.Fatal(err)
	}

	var printedChannel bool
	for {
		videos, err := c.NextStreamsPage()
		if err != nil {
			t.Fatal(err)
		} else if len(videos) == 0 {
			break
		}

		if !printedChannel {
			if available, channel := c.GetChannelInfo(); available {
				bs, err := json.MarshalIndent(channel, "", "	")
				if err != nil {
					t.Fatal(err)
				}
				t.Log(string(bs))
			}

			printedChannel = true
		}

		for _, video := range videos {
			t.Log(video.IsLive, video.Viewers, video.WasLive, video.VideoID, video.Title, video.Views)
		}
	}
}
