package examples

import (
	scraper "git.catnip.ee/miisu/youtube_scraper"
	"testing"
)

func TestSearch(t *testing.T) {
	scraper.Debug = true
	searchResults, err := scraper.NewSearch("test")
	if err != nil {
		t.Fatal(err)
	}

	for _, result := range searchResults.Results {
		t.Log(result)
	}
}
