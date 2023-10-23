package scraper

import (
	"bytes"
	"github.com/ayes-web/rjson"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type PlaylistScraper struct {
	url   string
	state PlaylistInfo

	playlistContinueToken string
	playlistContinueInput []byte

	initialDone bool
}

func NewPlaylistScraper(playlistId string) (p PlaylistScraper, err error) {
	rawUrl, err := url.Parse("https://www.youtube.com/playlist")
	if err != nil {
		return
	}

	q := rawUrl.Query()
	q.Set("list", playlistId)
	q.Set("hl", "en")
	rawUrl.RawQuery = q.Encode()

	p.url = rawUrl.String()

	return
}

// youtube json type playlistVideoRenderer
type PlaylistVideo struct {
	VideoID              string `rjson:"videoId"`
	Title                string `rjson:"title.runs[0].text"`
	PlaylistPosition     int    `rjson:"index.simpleText"`
	ChannelName          string `rjson:"shortBylineText.runs[0].text"`
	ChannelID            string `rjson:"shortBylineText.runs[0].navigationEndpoint.browseEndpoint.browseId"`
	VideoLengthInSeconds int    `rjson:"lengthSeconds"`
	Views                string `rjson:"videoInfo.runs[0].text"`
	Date                 string `rjson:"videoInfo.runs[-].text"` // example: "8 years ago"
}

type PlaylistInfo struct {
	Title        string `rjson:"header.playlistHeaderRenderer.title.simpleText"`
	Description  string `rjson:"header.playlistHeaderRenderer.descriptionText.simpleText"`
	ChannelName  string `rjson:"header.playlistHeaderRenderer.ownerText.runs[0].text"`
	ChannelID    string `rjson:"header.playlistHeaderRenderer.ownerText.runs[0].navigationEndpoint.browseEndpoint.browseId"`
	NewChannelID string `rjson:"header.playlistHeaderRenderer.ownerText.runs[0].navigationEndpoint.browseEndpoint.canonicalBaseUrl"`
	VideosCount  string `rjson:"header.playlistHeaderRenderer.numVideosText.runs[0].text"`
	Views        string `rjson:"header.playlistHeaderRenderer.viewCountText.simpleText"`
	UpdateStatus string `rjson:"header.playlistHeaderRenderer.byline[2].playlistBylineRenderer.text.runs[0].text"` // example: "Updated today"

	ContinuationToken string          `rjson:"contents.twoColumnBrowseResultsRenderer.tabs[0].tabRenderer.content.sectionListRenderer.contents[0].itemSectionRenderer.contents[0].playlistVideoListRenderer.contents[-].continuationItemRenderer.continuationEndpoint.continuationCommand.token"`
	Videos            []PlaylistVideo `rjson:"contents.twoColumnBrowseResultsRenderer.tabs[0].tabRenderer.content.sectionListRenderer.contents[0].itemSectionRenderer.contents[0].playlistVideoListRenderer.contents[].playlistVideoRenderer"`
}

type PlaylistContinueOutput struct {
	ContinuationToken string          `rjson:"onResponseReceivedActions[0]appendContinuationItemsAction.continuationItems[-].continuationItemRenderer.continuationEndpoint.continuationCommand.token"`
	Videos            []PlaylistVideo `rjson:"onResponseReceivedActions[0]appendContinuationItemsAction.continuationItems[].playlistVideoRenderer"`
}

// GetPlaylistInfo returns the initial info from the page
func (p *PlaylistScraper) GetPlaylistInfo() (info PlaylistInfo, err error) {
	if !p.initialDone {
		if _, err = p.NextPage(); err != nil {
			return
		}
	}

	info = p.state
	return
}

func (p *PlaylistScraper) NextPage() (videos []PlaylistVideo, err error) {
	if !p.initialDone {
		var rawJson string
		rawJson, err = extractInitialData(p.url)
		if err != nil {
			return
		}

		debugFileOutput([]byte(rawJson), "initial_playlist_output.json")

		if err = rjson.Unmarshal([]byte(rawJson), &p.state); err != nil {
			return
		}

		p.state.ChannelName = strings.TrimPrefix(p.state.ChannelName, "by ")
		p.state.NewChannelID = strings.TrimPrefix(p.state.NewChannelID, "/")
		p.state.Views = strings.TrimSuffix(p.state.Views, " views")

		p.playlistContinueToken = p.state.ContinuationToken
		p.playlistContinueInput, err = continueInput{Continuation: p.playlistContinueToken}.FillGenericInfo().Construct()
		if err != nil {
			return
		}

		p.initialDone = true
		videos = p.state.Videos
	} else {
		if len(p.playlistContinueInput) == 0 {
			return
		}

		var resp *http.Response
		resp, err = http.Post("https://www.youtube.com/youtubei/v1/browse", "application/json", bytes.NewReader(p.playlistContinueInput))
		if err != nil {
			return
		}

		p.playlistContinueInput = []byte{}

		var body []byte
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return
		}

		debugFileOutput(body, "initial_playlist_%s.json", p.playlistContinueToken)

		var out PlaylistContinueOutput
		if err = rjson.Unmarshal(body, &out); err != nil {
			return
		}

		if out.ContinuationToken != "" {
			p.playlistContinueToken = out.ContinuationToken
			p.playlistContinueInput, err = continueInput{Continuation: p.playlistContinueToken}.FillGenericInfo().Construct()
			if err != nil {
				return
			}
		}

		videos = out.Videos
	}

	return
}
