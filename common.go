package scraper

import (
	"bytes"
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"strings"
)

func extractInitialDataBytes(body []byte) (rawJson string, err error) {
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

func extractInitialData(url string) (rawJson string, err error) {
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

	return extractInitialDataBytes(body)
}

func (ci continueInput) FillGenericInfo() continueInput {
	ci.Context.Client.Hl = "en"
	ci.Context.Client.Gl = "GB"
	ci.Context.Client.ClientName = "WEB"
	ci.Context.Client.ClientVersion = "2.20230714.00.00"

	return ci
}

func (ci continueInput) Construct() ([]byte, error) {
	return json.Marshal(ci)
}
