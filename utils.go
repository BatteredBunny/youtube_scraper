package scraper

import (
	"fmt"
	"log"
	"strings"

	"github.com/dustin/go-humanize"
)

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}

func splice(s string, n int) string {
	a := []rune(s)
	a = a[n:]
	return string(a)
}

func GetVideoThumbnail(id string) string {
	return fmt.Sprintf("https://i.ytimg.com/vi/%s/maxresdefault.jpg", id)
}

// humanize library doesnt seem to understand that "10K" and "10k" are the same thing
func FixUnit(s string) string {
	if strings.HasSuffix(s, "K") {
		s = strings.TrimSuffix(s, "K") + "k"
	}

	return s
}

// Parses views from youtube outputs
func ParseViews(rawViews string) (views float64, err error) {
	if rawViews != "" && rawViews != "No views" {
		rawViews = strings.TrimSuffix(rawViews, " views")
		rawViews = strings.TrimSuffix(rawViews, " view")
		rawViews = strings.ReplaceAll(rawViews, ",", "")
		rawViews = FixUnit(rawViews)

		var unit string
		views, unit, err = humanize.ParseSI(rawViews)
		if err != nil {
			return
		} else if unit != "" {
			log.Printf("WARNING: possibly wrong number for views: %f%s\n", views, unit)
		}
	}

	return
}
