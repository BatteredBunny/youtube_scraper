package examples

import (
	"testing"

	scraper "git.catnip.ee/miisu/youtube_scraper"
)

func TestSearch(t *testing.T) {
	searchResults, err := scraper.NewSearch("test")
	if err != nil {
		t.Fatal(err)
	}

	for _, result := range searchResults.Results {
		t.Log(result)
	}
}
