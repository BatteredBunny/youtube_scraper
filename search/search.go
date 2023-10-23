package search

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/ayes-web/rjson"
	scraper "github.com/ayes-web/youtube_scraper"
)

type SearchEntryType = int

const (
	SearchEntryTypeVideo SearchEntryType = iota
	SearchEntryTypeChannel
	SearchEntryTypePlaylist
)

type SearchEntry struct {
	Type  SearchEntryType
	Entry any
}

type SearchVideo struct {
	VideoID        string
	Title          string
	Date           string // e.g "2 years ago" or "5 hours ago", will be empty when its a livestream
	Length         string // e.g "15:54", will be empty when its a livestream
	Views, Viewers int
	IsLive         bool
	Thumbnails     []scraper.YoutubeImage
	ChannelAvatar  string

	HasNewBadge, HasCCBadge, Has4kBadge      bool
	AuthorIsVerified, AuthorIsVerifiedArtist bool

	Username, ChannelID, NewChannelID string

	TruncatedDescription string
}

type SearchChannel struct {
	ChannelID    string
	Username     string
	NewChannelID string
	Avatars      []scraper.YoutubeImage

	BioSnippet  string
	Subscribers int // e.g. "2.04M subscribers"

	IsVerified, IsVerifiedAuthor bool
}

type SearchPlaylist struct {
	PlaylistID   string
	Username     string
	NewChannelID string
	Thumbnails   []scraper.YoutubeImage

	VideoCount int
	Title      string
}

type SearchScraper struct {
	url             string
	initialComplete bool

	chipFilters []chipFilter
	filters     []searchFilter

	searchContinueInput     scraper.ContinueInput
	searchContinueInputJson []byte
}

// Feel free to leave filter option as empty
func NewSearchScraper(query string, filter string) (s SearchScraper, err error) {
	rawUrl, err := url.Parse("https://www.youtube.com/results")
	if err != nil {
		return
	}

	q := rawUrl.Query()
	q.Set("search_query", query)
	q.Set("hl", "en")
	if filter != "" {
		q.Set("sp", filter)
	}

	rawUrl.RawQuery = q.Encode()
	s.url = rawUrl.String()
	return
}

func (s *SearchScraper) NextPage() (searchEntries []SearchEntry, err error) {
	if !s.initialComplete {
		var rawJson string
		rawJson, err = scraper.ExtractInitialData(s.url)
		if err != nil {
			return
		}

		scraper.DebugFileOutput([]byte(rawJson), "search_initial.json")

		var out searchInitialOutput
		if err = rjson.Unmarshal([]byte(rawJson), &out); err != nil {
			if errors.Is(err, rjson.ErrCantFindField) {
				if scraper.Debug {
					log.Println("WARNING:", err)
				}
				err = nil
			}
			return
		}

		for _, rawEntry := range out.Results {
			if rawEntry.Video.VideoID != "" || rawEntry.Playlist.PlaylistID != "" || rawEntry.Channel.ChannelID != "" {
				if entry, err := rawEntry.ToSearchEntry(); err != nil {
					log.Println("WARNING error while converting search result to entry:", err)
				} else {
					searchEntries = append(searchEntries, entry)
				}
			}
		}

		s.searchContinueInput = scraper.ContinueInput{Continuation: out.ContinueToken}.FillGenericInfo()
		s.searchContinueInputJson, err = s.searchContinueInput.Construct()
		if err != nil {
			return
		}

		s.chipFilters = out.ChipFilters

		s.filters = []searchFilter{}
		for _, group := range out.FilterGroups {
			s.filters = append(s.filters, group.Filters...)
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

		scraper.DebugFileOutput(body, "search_continue_%s.json", s.searchContinueInput.Continuation)

		var out searchContinueOutput
		if err = rjson.Unmarshal(body, &out); err != nil {
			if errors.Is(err, rjson.ErrCantFindField) {
				if scraper.Debug {
					log.Println("WARNING:", err)
				}
				err = nil
			}
			return
		}

		for _, rawEntry := range out.AppendResults {
			if rawEntry.Video.VideoID != "" || rawEntry.Playlist.PlaylistID != "" || rawEntry.Channel.ChannelID != "" {
				if entry, err := rawEntry.ToSearchEntry(); err != nil {
					log.Println("WARNING error while converting search result to entry:", err)
				} else {
					searchEntries = append(searchEntries, entry)
				}
			}
		}

		for _, rawEntry := range out.ReloadResults {
			if rawEntry.Video.VideoID != "" || rawEntry.Playlist.PlaylistID != "" || rawEntry.Channel.ChannelID != "" {
				if entry, err := rawEntry.ToSearchEntry(); err != nil {
					log.Println("WARNING error while converting search result to entry:", err)
				} else {
					searchEntries = append(searchEntries, entry)
				}
			}
		}

		if out.AppendContinueToken != "" {
			s.searchContinueInput = scraper.ContinueInput{Continuation: out.AppendContinueToken}.FillGenericInfo()
			s.searchContinueInputJson, err = s.searchContinueInput.Construct()
			if err != nil {
				return
			}
		} else if out.ReloadContinueToken != "" {
			s.searchContinueInput = scraper.ContinueInput{Continuation: out.ReloadContinueToken}.FillGenericInfo()
			s.searchContinueInputJson, err = s.searchContinueInput.Construct()
			if err != nil {
				return
			}
		}

		s.filters = []searchFilter{}
		for _, group := range out.FilterGroups {
			s.filters = append(s.filters, group.Filters...)
		}
	}
	return
}
