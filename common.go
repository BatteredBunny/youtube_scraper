package scraper

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

func extractInitialData(url string) (rawJson string, err error) {
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
			rawJson, _ = strings.CutSuffix(cut, ";")
		}
	})

	return
}

func (ci continueInput) FillGenericInfo() continueInput {
	ci.Context.Client.Hl = "en"
	ci.Context.Client.Gl = "UK"
	ci.Context.Client.ClientName = "WEB"
	ci.Context.Client.ClientVersion = "2.20230706.00.00"

	return ci
}

func (ci continueInput) Construct() (continueInputJson []byte, err error) {
	return json.Marshal(ci)
}
