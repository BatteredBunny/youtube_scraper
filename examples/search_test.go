package examples

import (
	"encoding/json"
	"log"
	"testing"

	search "github.com/BatteredBunny/youtube_scraper/search"
)

func TestSearch(t *testing.T) {
	searchScraper, err := search.NewSearchScraper("test", "")
	if err != nil {
		t.Fatal(err)
	}

	var results []search.SearchEntry
	for i := 0; i <= 2; i++ {
		results, err = searchScraper.NextPage()
		if err != nil {
			t.Fatal(err)
		} else if len(results) == 0 {
			break
		}

		HandleSearchEntries(t, results)
	}
}

func TestSearchFilters(t *testing.T) {
	searchScraper, err := search.NewSearchScraper("test", "")
	if err != nil {
		t.Fatal(err)
	}

	searchScraper.NextPage()

	bs, _ := json.Marshal(searchScraper.GetFilters())
	log.Println(string(bs))
	searchScraper.ApplyFilter("Channel")

	var results []search.SearchEntry
	for i := 0; i <= 2; i++ {
		results, err = searchScraper.NextPage()
		if err != nil {
			t.Fatal(err)
		} else if len(results) == 0 {
			break
		}

		HandleSearchEntries(t, results)
	}
}

func HandleSearchEntries(t *testing.T, searchEntries []search.SearchEntry) {
	for _, searchEntry := range searchEntries {
		switch searchEntry.Type {
		case search.SearchEntryTypeVideo:
			searchVideo := searchEntry.Entry.(search.SearchVideo)
			t.Log("video:", searchVideo)
		case search.SearchEntryTypeChannel:
			searchChannel := searchEntry.Entry.(search.SearchChannel)
			t.Log("channel:", searchChannel)
		case search.SearchEntryTypePlaylist:
			searchPlaylist := searchEntry.Entry.(search.SearchPlaylist)
			t.Log("playlist:", searchPlaylist)
		}
	}
	t.Log("-------------")
}
