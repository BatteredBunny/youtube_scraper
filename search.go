package scraper

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/ayes-web/rjson"
)

type searchVideoRenderer struct {
	VideoID       string         `rjson:"videoId"`
	Title         string         `rjson:"title.runs[0].text"`
	Date          string         `rjson:"publishedTimeText.simpleText"`
	Length        string         `rjson:"lengthText.simpleText"`
	Views         string         `rjson:"viewCountText.simpleText"`
	Viewers       string         `rjson:"viewCountText.runs[0].text"`
	Thumbnails    []YoutubeImage `rjson:"thumbnail.thumbnails"`
	ChannelAvatar string         `rjson:"channelThumbnailSupportedRenderers.channelThumbnailWithLinkRenderer.thumbnail.thumbnails[0].url"`

	Badges      []string `rjson:"badges[].metadataBadgeRenderer.label"`        // example of badge "New", "CC", "4K",
	OwnerBadges []string `rjson:"ownerBadges[].metadataBadgeRenderer.tooltip"` // example of owner badge "Verified" or "Official Artist Channel"

	Username            string   `rjson:"ownerText.runs[0].text"`
	ChannelID           string   `rjson:"ownerText.runs[0].navigationEndpoint.browseEndpoint.browseId"`
	NewChannelID        string   `rjson:"ownerText.runs[0].navigationEndpoint.browseEndpoint.canonicalBaseUrl"`
	DescriptionSnippets []string `rjson:"detailedMetadataSnippets[0]snippetText.runs[].text"`
}

func (rawVideo searchVideoRenderer) ToVideo() (video SearchVideo, err error) {
	var truncatedDescription string
	for _, snippet := range rawVideo.DescriptionSnippets {
		truncatedDescription += snippet
	}

	var hasNewBadge, hasCCBadge, has4kBadge bool
	for _, badge := range rawVideo.Badges {
		switch badge {
		case VideoBadge4k:
			has4kBadge = true
		case VideoBadgeCC:
			hasCCBadge = true
		case VideoBadgeNew:
			hasNewBadge = true
		}
	}

	var isVerified, isVerifiedArtist bool
	for _, badge := range rawVideo.OwnerBadges {
		switch badge {
		case ChannelBadgeVerified:
			isVerified = true
		case ChannelBadgeVerifiedArtistChannel:
			isVerifiedArtist = true
		}
	}

	var viewers int
	if strings.HasSuffix(rawVideo.Viewers, " watching") {
		viewers, err = strconv.Atoi(strings.TrimSuffix(rawVideo.Viewers, " watching"))
		if err != nil {
			return
		}
	} else if rawVideo.Viewers != "" {
		viewers, err = strconv.Atoi(strings.ReplaceAll(rawVideo.Viewers, ",", ""))
		if err != nil {
			return
		}
	}

	var views int
	if rawVideo.Views != "" {
		views, err = strconv.Atoi(strings.ReplaceAll(strings.TrimSuffix(rawVideo.Views, " views"), ",", ""))
		if err != nil {
			return
		}
	}

	video = SearchVideo{
		VideoID: rawVideo.VideoID,
		Title:   rawVideo.Title,
		Date:    rawVideo.Date,
		Length:  rawVideo.Length,

		Views:         views,
		Viewers:       viewers,
		IsLive:        rawVideo.Date == "" || rawVideo.Length == "" || viewers > 0,
		Thumbnails:    rawVideo.Thumbnails,
		ChannelAvatar: rawVideo.ChannelAvatar,

		Username:             rawVideo.Username,
		ChannelID:            rawVideo.ChannelID,
		TruncatedDescription: truncatedDescription,
		NewChannelID:         strings.TrimPrefix(rawVideo.NewChannelID, "/"),

		HasNewBadge: hasNewBadge, HasCCBadge: hasCCBadge, Has4kBadge: has4kBadge,
		IsVerified: isVerified, IsVerifiedArtist: isVerifiedArtist,
	}
	return
}

type SearchVideo struct {
	VideoID        string
	Title          string
	Date           string // e.g "2 years ago" or "5 hours ago", will be empty when its a livestream
	Length         string // e.g "15:54", will be empty when its a livestream
	Views, Viewers int
	IsLive         bool
	Thumbnails     []YoutubeImage
	ChannelAvatar  string

	HasNewBadge, HasCCBadge, Has4kBadge bool
	IsVerified, IsVerifiedArtist        bool

	Username, ChannelID, NewChannelID string

	TruncatedDescription string
}

// TODO: add option to use filters
type searchInitialOutput struct {
	FilterGroups []struct {
		Title   string `rjson:"title.simpleText"`
		Filters []struct {
			Title      string `rjson:"label.simpleText"`
			QueryValue string `rjson:"navigationEndpoint.searchEndpoint.params"`
		} `rjson:"filters[].searchFilterRenderer"`
	} `rjson:"header.searchHeaderRenderer.searchFilterButton.buttonRenderer.command.openPopupAction.popup.searchFilterOptionsDialogRenderer.groups[].searchFilterGroupRenderer"`

	ChipFilters []struct {
		Title             string `rjson:"text.simpleText"`
		ContinuationToken string `rjson:"navigationEndpoint.continuationCommand.token"`
	} `rjson:"header.searchHeaderRenderer.chipBar.chipCloudRenderer.chips[].chipCloudChipRenderer"`

	// TODO: add reelShelfRenderer from search results
	Results       []searchVideoRenderer `rjson:"contents.twoColumnSearchResultsRenderer.primaryContents.sectionListRenderer.contents[0].itemSectionRenderer.contents[].videoRenderer"`
	ContinueToken string                `rjson:"contents.twoColumnSearchResultsRenderer.primaryContents.sectionListRenderer.contents[1].continuationItemRenderer.continuationEndpoint.continuationCommand.token"`
}

// TODO: add option to use filters
type searchContinueOutput struct {
	FilterGroups []struct {
		Title   string `rjson:"title.simpleText"`
		Filters []struct {
			Title      string `rjson:"label.simpleText"`
			QueryValue string `rjson:"navigationEndpoint.searchEndpoint.params"`
		} `rjson:"filters[].searchFilterRenderer"`
	} `rjson:"header.searchHeaderRenderer.searchFilterButton.buttonRenderer.command.openPopupAction.popup.searchFilterOptionsDialogRenderer.groups[].searchFilterGroupRenderer"`

	// TODO: add reelShelfRenderer from search results
	Results       []searchVideoRenderer `rjson:"onResponseReceivedCommands[0]appendContinuationItemsAction.continuationItems[0]itemSectionRenderer.contents[].videoRenderer"`
	ContinueToken string                `rjson:"onResponseReceivedCommands[0]appendContinuationItemsAction.continuationItems[1]continuationItemRenderer.continuationEndpoint.continuationCommand.token"`
}

type SearchScraper struct {
	url             string
	initialComplete bool

	searchContinueInput     continueInput
	searchContinueInputJson []byte
}

type SearchScraperExport struct {
	Url             string
	InitialComplete bool
	Token           string
}

func (s SearchScraper) Export() SearchScraperExport {
	return SearchScraperExport{
		Url:             s.url,
		InitialComplete: s.initialComplete,
		Token:           s.searchContinueInput.Continuation,
	}
}

func SearchScraperFromExport(export SearchScraperExport) (s SearchScraper, err error) {
	s.url = export.Url
	s.initialComplete = export.InitialComplete

	s.searchContinueInput = continueInput{Continuation: export.Token}.FillGenericInfo()
	s.searchContinueInputJson, err = s.searchContinueInput.Construct()
	if err != nil {
		return
	}

	return
}

func NewSearchScraper(query string) (s SearchScraper, err error) {
	rawUrl, err := url.Parse("https://www.youtube.com/results")
	if err != nil {
		return
	}

	q := rawUrl.Query()
	q.Set("search_query", query)
	q.Set("hl", "en")
	rawUrl.RawQuery = q.Encode()
	s.url = rawUrl.String()
	return
}

func (s *SearchScraper) NextPage() (videos []SearchVideo, err error) {
	if !s.initialComplete {
		var rawJson string
		rawJson, err = extractInitialData(s.url)
		if err != nil {
			return
		}

		debugFileOutput([]byte(rawJson), "search_initial.json")

		var out searchInitialOutput
		if err = rjson.Unmarshal([]byte(rawJson), &out); err != nil {
			if errors.Is(err, rjson.ErrCantFindField) {
				if Debug {
					log.Println("WARNING:", err)
				}
				err = nil
			}
			return
		}

		for _, v := range out.Results {
			if video, err := v.ToVideo(); err != nil {
				log.Fatal(err)
				log.Println("WARNING error while converting search result:", err)
			} else {
				videos = append(videos, video)
			}
		}

		s.searchContinueInput = continueInput{Continuation: out.ContinueToken}.FillGenericInfo()
		s.searchContinueInputJson, err = s.searchContinueInput.Construct()
		if err != nil {
			return
		}

		s.initialComplete = true
	} else {
		var resp *http.Response
		resp, err = http.Post("https://www.youtube.com/youtubei/v1/search", "application/json", bytes.NewReader(s.searchContinueInputJson))
		if err != nil {
			return
		}
		s.searchContinueInputJson = []byte{}

		var body []byte
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return
		}

		debugFileOutput(body, "search_continue_%s.json", s.searchContinueInput.Continuation)

		var out searchContinueOutput
		if err = rjson.Unmarshal(body, &out); err != nil {
			if errors.Is(err, rjson.ErrCantFindField) {
				if Debug {
					log.Println("WARNING:", err)
				}
				err = nil
			}
			return
		}

		for _, v := range out.Results {
			if video, err := v.ToVideo(); err != nil {
				log.Println("WARNING error while converting search result:", err)
			} else {
				videos = append(videos, video)
			}
		}

		s.searchContinueInput = continueInput{Continuation: out.ContinueToken}.FillGenericInfo()
		s.searchContinueInputJson, err = s.searchContinueInput.Construct()
		if err != nil {
			return
		}
	}
	return
}
