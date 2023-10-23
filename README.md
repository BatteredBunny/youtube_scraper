# youtube_scraper

Youtube metadata scraping library for golang

## Features
- Fetch channel videos
- Fetch home page videos
- Fetch video metadata
- Fetch comments and subcomments

## Example
```go
package main

import (
	"git.catnip.ee/miisu/youtube_scraper"
	"log"
)

func main() {
	c := scraper.NewChannelVideosScraper("@TomScottGo")

	var (
		videos         []scraper.Video
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