package main

import (
	"fmt"
	"log"
)

func main() {
	c := NewChannelVideosScraper("@TomScottGo")

	for {
		videos, err := c.NextPage()
		if err != nil {
			log.Fatal(err)
		} else if len(videos) == 0 {
			break
		}

		for _, video := range videos {
			fmt.Println(video.VideoID, video.Title, video.Date)
		}
		fmt.Println()
	}
}

// 11 years ago
// 2 years ago
// 1 year ago

// 11 months ago
// 2 months ago
// 1 month ago

// 4 weeks ago
// 3 weeks ago
// 2 weeks ago

// 13 days ago
// 2 days ago
// 1 day ago

// 23 hours ago
// 21 hours ago
// 20 hours ago
// 5 hours ago
// 3 hours ago
// 2 hours ago
// 1 hour ago

// 59 minutes ago
// 2 minutes ago
// 1 minute ago

// 59 seconds ago
// 2 seconds ago
// 1 second ago
