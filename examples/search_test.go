package examples

import (
	"encoding/json"
	"testing"

	scraper "git.catnip.ee/miisu/youtube_scraper"
)

func TestSearch(t *testing.T) {
	searchScraper, err := scraper.NewSearchScraper("livestream")
	if err != nil {
		t.Fatal(err)
	}

	searchScraper.NextPage()

	var results []scraper.SearchVideo
	for i := 0; i <= 3; i++ {
		results, err = searchScraper.NextPage()
		if err != nil {
			t.Fatal(err)
		} else if len(results) == 0 {
			break
		}

		for _, result := range results {
			bs, err := json.MarshalIndent(result, "", "	")
			if err != nil {
				t.Fatal(err)
			}
			t.Log(string(bs))
		}
	}
}
