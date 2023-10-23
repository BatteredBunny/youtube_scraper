# youtube_scraper

Youtube metadata scraper library for golang

```go
package main

import (
	youtube "git.catnip.ee/miisu/youtube_scraper"
	"log"
)

func main() {
	c := youtube.NewChannelVideosScraper("@TomScottGo")

	var (
		videos         []youtube.Video
		err            error
		printedChannel bool
	)
	for {
		videos, err = c.NextPage()
		if err != nil {
			log.Fatal(err)
		} else if len(videos) == 0 {
			break
		}
		
		if !printedChannel {
			if available, channel := c.GetChannelInfo(); available {
				log.Println(channel)
			}
        }
		
		for _, video := range videos {
			log.Println(video.VideoID, video.Title, video.Views)
		}
	}
}
```