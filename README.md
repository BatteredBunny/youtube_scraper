# youtube_scraper

Youtube metadata scraping library for golang

## Features
- Fetch channel's videos and streams tab
- Fetch homepage videos
- Fetch video metadata
- Fetch video sidebar recommendations
- Fetch comments and its reply threads

## Example
```go
package main

import (
	"git.catnip.ee/miisu/youtube_scraper"
	"log"
)

func main() {
	c := scraper.NewChannelScraper("@TomScottGo")

	var (
		videos         []scraper.Video
		err            error
		printedChannel bool
	)
	for {
		videos, err = c.NextVideosPage()
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