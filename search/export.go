package search

import scraper "git.catnip.ee/miisu/youtube_scraper"

type SearchScraperExport struct {
	Url             string
	InitialComplete bool
	Token           string
	ChipFilters     []chipFilter
}

func (s SearchScraper) Export() SearchScraperExport {
	return SearchScraperExport{
		Url:             s.url,
		InitialComplete: s.initialComplete,
		Token:           s.searchContinueInput.Continuation,
		ChipFilters:     s.chipFilters,
	}
}

func SearchScraperFromExport(export SearchScraperExport) (s SearchScraper, err error) {
	s.url = export.Url
	s.initialComplete = export.InitialComplete
	s.chipFilters = export.ChipFilters

	s.searchContinueInput = scraper.ContinueInput{Continuation: export.Token}.FillGenericInfo()
	s.searchContinueInputJson, err = s.searchContinueInput.Construct()
	if err != nil {
		return
	}

	return
}
