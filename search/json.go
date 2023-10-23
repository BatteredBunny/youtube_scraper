package search

import scraper "git.catnip.ee/miisu/youtube_scraper"

type searchVideoRenderer struct {
	VideoID       string                 `rjson:"videoId"`
	Title         string                 `rjson:"title.runs[0].text"`
	Date          string                 `rjson:"publishedTimeText.simpleText"`
	Length        string                 `rjson:"lengthText.simpleText"`
	Views         string                 `rjson:"viewCountText.simpleText"`
	Viewers       string                 `rjson:"viewCountText.runs[0].text"`
	Thumbnails    []scraper.YoutubeImage `rjson:"thumbnail.thumbnails"`
	ChannelAvatar string                 `rjson:"channelThumbnailSupportedRenderers.channelThumbnailWithLinkRenderer.thumbnail.thumbnails[0].url"`

	Badges      []string `rjson:"badges[].metadataBadgeRenderer.label"`        // example of badge "New", "CC", "4K",
	OwnerBadges []string `rjson:"ownerBadges[].metadataBadgeRenderer.tooltip"` // example of owner badge "Verified" or "Official Artist Channel"

	Username            string   `rjson:"ownerText.runs[0].text"`
	ChannelID           string   `rjson:"ownerText.runs[0].navigationEndpoint.browseEndpoint.browseId"`
	NewChannelID        string   `rjson:"ownerText.runs[0].navigationEndpoint.browseEndpoint.canonicalBaseUrl"`
	DescriptionSnippets []string `rjson:"detailedMetadataSnippets[0]snippetText.runs[].text"`
}

type searchChannelRenderer struct {
	ChannelID    string                 `rjson:"channelId"`
	Username     string                 `rjson:"title.simpleText"`
	NewChannelID string                 `rjson:"subscriberCountText.simpleText"` // e.g "@NewChannelID"
	Avatars      []scraper.YoutubeImage `rjson:"thumbnail.thumbnails"`

	RawBioSnippet  []string `rjson:"descriptionSnippet.runs[].text"`
	RawSubscribers string   `rjson:"videoCountText.simpleText"` // e.g. "2.04M subscribers"

	OwnerBadges []string `rjson:"ownerBadges[].metadataBadgeRenderer.tooltip"`
}

type searchPlaylistRenderer struct {
	PlaylistID    string                 `rjson:"playlistId"`
	Title         string                 `rjson:"title.simpleText"`
	Thumbnails    []scraper.YoutubeImage `rjson:"thumbnails[0]thumbnails"`
	RawVideoCount string                 `rjson:"videoCount"`

	Username     string `rjson:"shortBylineText.runs[0].text"`
	ChannelID    string `rjson:"shortBylineText.runs[0].navigationEndpoint.browseEndpoint.browseId"`
	NewChannelID string `rjson:"shortBylineText.runs[0].navigationEndpoint.browseEndpoint.canonicalBaseUrl"` // e.g "/@NewChannelID"
}

type rawSearchEntry struct {
	Video    searchVideoRenderer    `rjson:"videoRenderer"`
	Channel  searchChannelRenderer  `rjson:"channelRenderer"`
	Playlist searchPlaylistRenderer `rjson:"playlistRenderer"`
}

type chipFilter struct {
	Title             string `rjson:"text.simpleText"`
	ContinuationToken string `rjson:"navigationEndpoint.continuationCommand.token"`
}

type searchFilter struct {
	Title string `rjson:"label.simpleText"`
	Query string `rjson:"navigationEndpoint.searchEndpoint.query"`
	Param string `rjson:"navigationEndpoint.searchEndpoint.params"`
}

type searchInitialOutput struct {
	FilterGroups []struct {
		Title   string         `rjson:"title.simpleText"`
		Filters []searchFilter `rjson:"filters[].searchFilterRenderer"`
	} `rjson:"header.searchHeaderRenderer.searchFilterButton.buttonRenderer.command.openPopupAction.popup.searchFilterOptionsDialogRenderer.groups[].searchFilterGroupRenderer"`

	ChipFilters []chipFilter `rjson:"header.searchHeaderRenderer.chipBar.chipCloudRenderer.chips[].chipCloudChipRenderer"`

	// TODO: add reelShelfRenderer from search results
	Results       []rawSearchEntry `rjson:"contents.twoColumnSearchResultsRenderer.primaryContents.sectionListRenderer.contents[0].itemSectionRenderer.contents"`
	ContinueToken string           `rjson:"contents.twoColumnSearchResultsRenderer.primaryContents.sectionListRenderer.contents[1].continuationItemRenderer.continuationEndpoint.continuationCommand.token"`
}

type searchContinueOutput struct {
	FilterGroups []struct {
		Title   string         `rjson:"title.simpleText"`
		Filters []searchFilter `rjson:"filters[].searchFilterRenderer"`
	} `rjson:"header.searchHeaderRenderer.searchFilterButton.buttonRenderer.command.openPopupAction.popup.searchFilterOptionsDialogRenderer.groups[].searchFilterGroupRenderer"`

	// TODO: add reelShelfRenderer from search results
	AppendResults       []rawSearchEntry `rjson:"onResponseReceivedCommands[0]appendContinuationItemsAction.continuationItems[0]itemSectionRenderer.contents"`
	AppendContinueToken string           `rjson:"onResponseReceivedCommands[0]appendContinuationItemsAction.continuationItems[1]continuationItemRenderer.continuationEndpoint.continuationCommand.token"`

	ReloadResults       []rawSearchEntry `rjson:"onResponseReceivedCommands[0]reloadContinuationItemsCommand.continuationItems[0]twoColumnSearchResultsRenderer.primaryContents.sectionListRenderer.contents[0]itemSectionRenderer.contents"`
	ReloadContinueToken string           `rjson:"onResponseReceivedCommands[0]reloadContinuationItemsCommand.continuationItems[0]twoColumnSearchResultsRenderer.primaryContents.sectionListRenderer.contents[1]continuationItemRenderer.continuationEndpoint.continuationCommand.token"`
}
