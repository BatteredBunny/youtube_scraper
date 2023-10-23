# youtube_scraper

Youtube metadata scraper library for golang

```go
package main

import (
	"fmt"
	youtube "git.catnip.ee/miisu/youtube_scraper"
	"log"
)

func main() {
	c := youtube.NewChannelVideosScraper("@TomScottGo")

	for {
		videos, err := c.NextPage()
		if err != nil {
			log.Fatal(err)
		} else if len(videos) == 0 {
			break
		}

		for _, video := range videos {
			fmt.Println(video.VideoID, video.Title, video.Views)
		}
	}
}
```