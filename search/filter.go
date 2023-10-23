package search

import scraper "git.catnip.ee/miisu/youtube_scraper"

// returns chip filter options that can be applied, chips will become available after first page
func (s *SearchScraper) GetChipFilters() (filters []string) {
	for _, filter := range s.chipFilters {
		filters = append(filters, filter.Title)
	}

	return
}

// chips will become available after first page
func (s *SearchScraper) ApplyChipFilter(filterName string) (err error) {
	for _, filter := range s.chipFilters {
		if filter.Title == filterName {
			s.searchContinueInput = scraper.ContinueInput{Continuation: filter.ContinuationToken}.FillGenericInfo()
			s.searchContinueInputJson, err = s.searchContinueInput.Construct()
			if err != nil {
				return
			}

			break
		}
	}

	return
}

// filters will become available after first page
func (s *SearchScraper) GetFilters() (filters []string) {
	for _, filter := range s.filters {
		filters = append(filters, filter.Title)
	}

	return
}

// filters will become available after first page
func (s *SearchScraper) ApplyFilter(filterName string) (err error) {
	for _, filter := range s.filters {
		if filter.Title == filterName {
			*s, err = NewSearchScraper(filter.Query, filter.Param)
			if err != nil {
				return
			}

			break
		}
	}

	return
}
