package scraper

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ExtractInitialDataBytes(body []byte) (rawJson string, err error) {
	var doc *goquery.Document
	doc, err = goquery.NewDocumentFromReader(bytes.NewReader(body))
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

// Helper function that scraper json from html page
func ExtractInitialData(url string) (rawJson string, err error) {
	var resp *http.Response
	resp, err = http.Get(url)
	if err != nil {
		return
	}

	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return ExtractInitialDataBytes(body)
}

func (ci ContinueInput) FillGenericInfo() ContinueInput {
	ci.Context.Client.Hl = "en"
	ci.Context.Client.Gl = "GB"
	ci.Context.Client.ClientName = "WEB"
	ci.Context.Client.ClientVersion = "2.20230714.00.00"

	return ci
}

func (ci ContinueInput) Construct() ([]byte, error) {
	return json.Marshal(ci)
}
