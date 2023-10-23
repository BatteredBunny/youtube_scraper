package youtube_scraper

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

func ExtractInitialData(url string) (rawjson string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	var doc *goquery.Document
	doc, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}

	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		if cut, valid := strings.CutPrefix(s.Text(), "var ytInitialData = "); valid {
			rawjson, _ = strings.CutSuffix(cut, ";")
		}
	})

	return
}
