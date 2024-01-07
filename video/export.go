package video

import scraper "github.com/BatteredBunny/youtube_scraper"

type VideoScraperExport struct {
	MediaUrlJs string
	Url        string

	CommentsNewestPassedInitial bool
	CommentsNewestToken         string

	CommentsTopPassedInitial bool
	CommentsTopToken         string

	SidebarToken string
}

func (v VideoScraper) Export() VideoScraperExport {
	return VideoScraperExport{
		MediaUrlJs:                  v.mediaUrlJs,
		Url:                         v.url,
		CommentsNewestPassedInitial: v.commentsNewestPassedInitial,
		CommentsNewestToken:         v.commentsNewestContinueInput.Continuation,
		CommentsTopPassedInitial:    v.commentsTopPassedInitial,
		CommentsTopToken:            v.commentsTopContinueInput.Continuation,
		SidebarToken:                v.sidebarContinueInput.Continuation,
	}
}

func VideoScraperFromExport(export VideoScraperExport) (v VideoScraper, err error) {
	v.url = export.Url
	v.mediaUrlJs = export.Url

	v.commentsNewestPassedInitial = export.CommentsNewestPassedInitial
	v.commentsNewestContinueInput = scraper.ContinueInput{Continuation: export.CommentsNewestToken}.FillGenericInfo()
	v.commentsNewestContinueInputJson, err = v.commentsNewestContinueInput.Construct()
	if err != nil {
		return
	}

	v.commentsTopPassedInitial = export.CommentsTopPassedInitial
	v.commentsTopContinueInput = scraper.ContinueInput{Continuation: export.CommentsTopToken}.FillGenericInfo()
	v.commentsTopContinueInputJson, err = v.commentsTopContinueInput.Construct()
	if err != nil {
		return
	}

	v.sidebarContinueInput = scraper.ContinueInput{Continuation: export.SidebarToken}.FillGenericInfo()
	v.sidebarContinueInputJson, err = v.sidebarContinueInput.Construct()
	if err != nil {
		return
	}

	return
}
