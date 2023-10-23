package video

import scraper "git.catnip.ee/miisu/youtube_scraper"

// youtube json commentRenderer type
type commentRenderer struct {
	NewChannelID  string   `rjson:"authorText.simpleText"`
	CommentID     string   `rjson:"commentId"`
	Content       []string `rjson:"contentText.runs[].text"`
	PublishedTime string   `rjson:"publishedTimeText.runs[0].text"` // ends with "(edited)" if the comment has been edited
	LikeAmount    string   `rjson:"voteCount.simpleText"`           // 3K
	Pinned        []string `rjson:"pinnedCommentBadge.pinnedCommentBadgeRenderer.label.runs[].text"`
	IsHearted     bool     `rjson:"actionButtons.commentActionButtonsRenderer.creatorHeart.creatorHeartRenderer.isHearted"`
}

type subCommentsContinueOutput struct {
	Comments      []commentRenderer `rjson:"onResponseReceivedEndpoints[0]appendContinuationItemsAction.continuationItems[].commentRenderer"`
	ContinueToken string            `rjson:"onResponseReceivedEndpoints[0]appendContinuationItemsAction.continuationItems[-].continuationItemRenderer.button.buttonRenderer.command.continuationCommand.token"`
}

// commentThreadRenderer json type
type commentContinueOutputComment struct {
	Comment       commentRenderer `rjson:"comment.commentRenderer"`
	RepliesAmount string          `rjson:"replies.commentRepliesRenderer.viewReplies.buttonRenderer.text.runs[0].text"`
	RepliesToken  string          `rjson:"replies.commentRepliesRenderer.contents[-].continuationItemRenderer.continuationEndpoint.continuationCommand.token"`
}

type commentsContinueOutputInitial struct {
	Comments      []commentContinueOutputComment `rjson:"onResponseReceivedEndpoints[1]reloadContinuationItemsCommand.continuationItems[].commentThreadRenderer"`
	ContinueToken string                         `rjson:"onResponseReceivedEndpoints[1]reloadContinuationItemsCommand.continuationItems[-]continuationItemRenderer.continuationEndpoint.continuationCommand.token"`
}
type commentsContinueOutput struct {
	Comments      []commentContinueOutputComment `rjson:"onResponseReceivedEndpoints[0]appendContinuationItemsAction.continuationItems[].commentThreadRenderer"`
	ContinueToken string                         `rjson:"onResponseReceivedEndpoints[0]appendContinuationItemsAction.continuationItems[-]continuationItemRenderer.continuationEndpoint.continuationCommand.token"`
}

type compactVideoRenderer struct {
	VideoID         string                 `rjson:"videoId"`
	Title           string                 `rjson:"title.simpleText"`
	Username        string                 `rjson:"longBylineText.runs[0].text"`
	ChannelID       string                 `rjson:"longBylineText.runs[0].navigationEndpoint.browseEndpoint.browseId"`
	RawNewChannelID string                 `rjson:"longBylineText.runs[0].navigationEndpoint.browseEndpoint.canonicalBaseUrl"` // has "/" at start that must be trimmed
	Date            string                 `rjson:"publishedTimeText.simpleText"`
	Views           string                 `rjson:"viewCountText.simpleText"`
	Viewers         string                 `rjson:"viewCountText.runs[0].text"`
	Length          string                 `rjson:"lengthText.simpleText"`
	Badges          []string               `rjson:"badges[].metadataBadgeRenderer.label"`        // example of badge "New" or "CC"
	OwnerBadges     []string               `rjson:"ownerBadges[].metadataBadgeRenderer.tooltip"` // example of owner badge "Verified" or "Official Artist Channel"
	Thumbnails      []scraper.YoutubeImage `rjson:"thumbnail.thumbnails"`
}

type compactPlaylistRenderer struct {
	PlaylistID      string `rjson:"playlistId"`
	Title           string `rjson:"title.simpleText"`
	Username        string `rjson:"shortBylineText.runs[0].text"`
	ChannelID       string `rjson:"shortBylineText.runs[0].navigationEndpoint.browseEndpoint.browseId"`
	RawNewChannelID string `rjson:"shortBylineText.runs[0].navigationEndpoint.browseEndpoint.canonicalBaseUrl"` // has "/" at start that must be trimmed

	VideosAmount     string                 `rjson:"videoCountShortText.simpleText"`
	ThumbnailVideoID string                 `rjson:"navigationEndpoint.watchEndpoint.videoId"`
	Thumbnails       []scraper.YoutubeImage `rjson:"thumbnail.thumbnails"`
}

type compactRadioRenderer struct {
	RadioPlaylistID string `rjson:"playlistId"`
	Title           string `rjson:"title.simpleText"`
	SecondaryTitle  string `rjson:"longBylineText.simpleText"`

	ThumbnailVideoID string                 `rjson:"navigationEndpoint.watchEndpoint.videoId"`
	VideosAmount     string                 `rjson:"videoCountShortText.runs[0].text"`
	Thumbnails       []scraper.YoutubeImage `rjson:"thumbnail.thumbnails"`
}

type rawSidebarEntry struct {
	Video    compactVideoRenderer    `rjson:"compactVideoRenderer"`
	Playlist compactPlaylistRenderer `rjson:"compactPlaylistRenderer"`
	Radio    compactRadioRenderer    `rjson:"compactRadioRenderer"`
}

type sidebarOutput struct {
	SidebarEntries []rawSidebarEntry `rjson:"onResponseReceivedEndpoints[0].appendContinuationItemsAction.continuationItems"`
	ContinueToken  string            `rjson:"onResponseReceivedEndpoints[0].appendContinuationItemsAction.continuationItems[-].continuationItemRenderer.continuationEndpoint.continuationCommand.token"`
}

type videoInitialOutput struct {
	Title              string                 `rjson:"playerOverlays.playerOverlayRenderer.videoDetails.playerOverlayVideoDetailsRenderer.title.simpleText"`
	Description        string                 `rjson:"contents.twoColumnWatchNextResults.results.results.contents[1].videoSecondaryInfoRenderer.attributedDescription.content"`
	Views              string                 `rjson:"playerOverlays.playerOverlayRenderer.videoDetails.playerOverlayVideoDetailsRenderer.subtitle.runs[2].text"`
	IsLive             bool                   `rjson:"contents.twoColumnWatchNextResults.results.results.contents[0].videoPrimaryInfoRenderer.viewCount.videoViewCountRenderer.isLive"`
	Date               string                 `rjson:"contents.twoColumnWatchNextResults.results.results.contents[0].videoPrimaryInfoRenderer.dateText.simpleText"`
	Username           string                 `rjson:"playerOverlays.playerOverlayRenderer.videoDetails.playerOverlayVideoDetailsRenderer.subtitle.runs[0].text"`
	ChannelID          string                 `rjson:"contents.twoColumnWatchNextResults.results.results.contents[1].videoSecondaryInfoRenderer.owner.videoOwnerRenderer.title.runs[0].navigationEndpoint.browseEndpoint.browseId"`
	RawNewChannelID    string                 `rjson:"contents.twoColumnWatchNextResults.results.results.contents[1].videoSecondaryInfoRenderer.owner.videoOwnerRenderer.title.runs[0].navigationEndpoint.browseEndpoint.canonicalBaseUrl"`
	Likes              string                 `rjson:"contents.twoColumnWatchNextResults.results.results.contents[0].videoPrimaryInfoRenderer.videoActions.menuRenderer.topLevelButtons[0].segmentedLikeDislikeButtonRenderer.likeButton.toggleButtonRenderer.defaultText.simpleText"`
	ChannelSubscribers string                 `rjson:"contents.twoColumnWatchNextResults.results.results.contents[1].videoSecondaryInfoRenderer.owner.videoOwnerRenderer.subscriberCountText.simpleText"`
	CommentsCount      string                 `rjson:"contents.twoColumnWatchNextResults.results.results.contents[2].itemSectionRenderer.contents[0].commentsEntryPointHeaderRenderer.commentCount.simpleText"`
	Category           string                 `rjson:"contents.twoColumnWatchNextResults.results.results.contents[1].videoSecondaryInfoRenderer.metadataRowContainer.metadataRowContainerRenderer.rows[0].richMetadataRowRenderer.contents[1].richMetadataRenderer.title.runs[0].text"`
	OwnerBadges        []string               `rjson:"contents.twoColumnWatchNextResults.results.results.contents[1].videoSecondaryInfoRenderer.owner.videoOwnerRenderer.badges[].metadataBadgeRenderer.tooltip"`
	Badges             []string               `rjson:"contents.twoColumnWatchNextResults.results.results.contents[0].videoPrimaryInfoRenderer.badges[].metadataBadgeRenderer.label"`
	ChannelAvatars     []scraper.YoutubeImage `rjson:"contents.twoColumnWatchNextResults.results.results.contents[1].videoSecondaryInfoRenderer.owner.videoOwnerRenderer.thumbnail.thumbnails"`

	Tokens []struct {
		Title string `rjson:"title"`
		Token string `rjson:"serviceEndpoint.continuationCommand.token"`
	} `rjson:"engagementPanels[].engagementPanelSectionListRenderer.header.engagementPanelTitleHeaderRenderer.menu.sortFilterSubMenuRenderer.subMenuItems[0]"`

	SidebarEntries []rawSidebarEntry `rjson:"contents.twoColumnWatchNextResults.secondaryResults.secondaryResults.results"`
	SidebarToken   string            `rjson:"contents.twoColumnWatchNextResults.secondaryResults.secondaryResults.results[-].continuationItemRenderer.button.buttonRenderer.command.continuationCommand.token"`
}

type MediaFormat struct {
	Bitrate int `rjson:"bitrate"`
	Width   int `rjson:"width"`
	Height  int `rjson:"height"`

	Url             string `rjson:"url"`
	MimeType        string `rjson:"mimeType"` // e.g "audio/mp4; codecs=\"mp4a.40.2\"" or "video/mp4; codecs=\"av01.0.00M.08\""
	QualityLabel    string `rjson:"qualityLabel"`
	SignatureCipher string `rjson:"signatureCipher"` // DRM

	// Videos can have dubs
	AudioTrack struct {
		DisplayName    string `rjson:"displayName"`
		AudioIsDefault bool   `rjson:"audioIsDefault"`
	} `rjson:"audioTrack"`
}

type ExtractMediaOutput struct {
	Formats []MediaFormat `rjson:"streamingData.formats"`

	// Best quality formats are here, but the video and audio tracks will be separate
	AdaptiveFormats []MediaFormat `rjson:"streamingData.adaptiveFormats"`
}
