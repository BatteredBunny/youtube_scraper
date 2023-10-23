package scraper

import (
	"errors"
	"github.com/ayes-web/rjson"
	"log"
	"net/url"
	"strings"
)

type searchVideoRenderer struct {
	VideoID string `rjson:"videoId"`
	Title   string `rjson:"title.runs[0].text"`
	Date    string `rjson:"publishedTimeText.simpleText"`
	Length  string `rjson:"lengthText.simpleText"`
	Views   string `rjson:"viewCountText.simpleText"`
	Viewers string `rjson:"viewCountText.runs[0].text"`

	Badges      []string `rjson:"badges[].metadataBadgeRenderer.label"`        // example of badge "New", "CC",
	OwnerBadges []string `rjson:"ownerBadges[].metadataBadgeRenderer.tooltip"` // example of owner badge "Verified"

	Username            string   `rjson:"ownerText.runs[0].text"`
	ChannelID           string   `rjson:"ownerText.runs[0].navigationEndpoint.browseEndpoint.browseId"`
	NewChannelID        string   `rjson:"ownerText.runs[0].navigationEndpoint.browseEndpoint.canonicalBaseUrl"`
	DescriptionSnippets []string `rjson:"detailedMetadataSnippets[0]snippetText.runs[].text"`
}

// TODO: add reelShelfRenderer from search results
type SearchInitialOutput struct {
	FilterGroups []struct {
		Title   string `rjson:"title.simpleText"`
		Filters []struct {
			Title      string `rjson:"label.simpleText"`
			QueryValue string `rjson:"navigationEndpoint.searchEndpoint.params"`
		} `rjson:"filters[].searchFilterRenderer"`
	} `rjson:"header.searchHeaderRenderer.searchFilterButton.buttonRenderer.command.openPopupAction.popup.searchFilterOptionsDialogRenderer.groups[].searchFilterGroupRenderer"`

	// TODO: add option to use it
	ChipFilters []struct {
		Title             string `rjson:"text.simpleText"`
		ContinuationToken string `rjson:"navigationEndpoint.continuationCommand.token"`
	} `rjson:"header.searchHeaderRenderer.chipBar.chipCloudRenderer.chips[].chipCloudChipRenderer"`

	Results       []searchVideoRenderer `rjson:"contents.twoColumnSearchResultsRenderer.primaryContents.sectionListRenderer.contents[0].itemSectionRenderer.contents[].videoRenderer"`
	ContinueToken string                `rjson:"contents.twoColumnSearchResultsRenderer.primaryContents.sectionListRenderer.contents[1].continuationItemRenderer.continuationEndpoint.continuationCommand.token"`
}

func NewSearch(query string) (s SearchInitialOutput, err error) {
	rawUrl, err := url.Parse("https://www.youtube.com/results")
	if err != nil {
		return
	}

	q := rawUrl.Query()
	q.Set("search_query", query)
	q.Set("hl", "en")
	rawUrl.RawQuery = q.Encode()
	parsedUrl := rawUrl.String()

	var rawJson string
	rawJson, err = extractInitialData(parsedUrl)
	if err != nil {
		return
	}

	debugFileOutput([]byte(rawJson), "search_initial.json")

	if err = rjson.Unmarshal([]byte(rawJson), &s); err != nil {
		if errors.Is(err, rjson.ErrCantFindField) {
			if Debug {
				log.Println("WARNING:", err)
			}
			err = nil
		}
		return
	}

	for i, v := range s.Results {
		s.Results[i].NewChannelID = strings.TrimPrefix(v.NewChannelID, "/")
	}

	return
}
