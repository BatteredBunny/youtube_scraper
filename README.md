# youtube_scraper

Youtube metadata scraping library for golang

## Features
- Fetch basic info about channel
- Fetch channel videos and streams tabs
- Fetch homepage videos
- Fetch video metadata
- Fetch video sidebar recommendations (video, playlist, radio)
- Fetch comments and its reply threads

## Example
For more examples please look into the "examples" folder

```go
package main

import (
	"encoding/json"
	"git.catnip.ee/miisu/youtube_scraper"
	"log"
)

func main() {
	c := scraper.NewChannelScraper("@TomScottGo")

	var printedChannel bool
	for {
		videos, err := c.NextVideosPage()
		if err != nil {
			log.Fatal(err)
		} else if len(videos) == 0 {
			break
		}

		if !printedChannel {
			if available, channel := c.GetChannelInfo(); available {
				bs, err := json.MarshalIndent(channel, "", "	")
				if err != nil {
					log.Fatal(err)
				}
				log.Println(string(bs))
			}

			printedChannel = true
		}

		for _, video := range videos {
			log.Println(video.VideoID, video.Title, video.Views)
		}
	}
}
```