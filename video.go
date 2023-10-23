package scraper

import (
	"errors"
	"fmt"
	"github.com/ayes-web/rjson"
	"log"
	"strings"
)

type VideoScraper struct {
	url   string
	video FullVideo

	commentsNewestPassedInitial     bool
	commentsNewestToken             string
	commentsNewestContinueInputJson []byte

	commentsTopPassedInitial     bool
	commentsTopToken             string
	commentsTopContinueInputJson []byte
}

// FullVideo has the full metadata unlike Video which is fetched from video lists
type FullVideo struct {
	VideoID       string `json:"VideoID"`
	Title         string `json:"Title"`
	Description   string `json:"Description"`
	Views         string `json:"Views"` // if its live this will display number of viewers instead
	IsLive        bool   `json:"IsLive"`
	WasLive       bool   `json:"WasLive"` // if this video was live
	Date          string `json:"Date"`    // Date will be in this format: "Jul 12, 2023"
	Likes         string `json:"Likes"`
	CommentsCount string `json:"CommentsCount"`
	Category      string `json:"Category"`

	Username           string `json:"Username"`
	ChannelID          string `json:"ChannelID"`
	NewChannelID       string `json:"NewChannelID"`
	ChannelSubscribers string `json:"ChannelSubscribers"`
}

type videoInitialOutput struct {
	Title              string `rjson:"playerOverlays.playerOverlayRenderer.videoDetails.playerOverlayVideoDetailsRenderer.title.simpleText"`
	Description        string `rjson:"contents.twoColumnWatchNextResults.results.results.contents[1].videoSecondaryInfoRenderer.attributedDescription.content"`
	Views              string `rjson:"playerOverlays.playerOverlayRenderer.videoDetails.playerOverlayVideoDetailsRenderer.subtitle.runs[2].text"`
	IsLive             bool   `rjson:"contents.twoColumnWatchNextResults.results.results.contents[0].videoPrimaryInfoRenderer.viewCount.videoViewCountRenderer.isLive"`
	Date               string `rjson:"contents.twoColumnWatchNextResults.results.results.contents[0].videoPrimaryInfoRenderer.dateText.simpleText"`
	Username           string `rjson:"playerOverlays.playerOverlayRenderer.videoDetails.playerOverlayVideoDetailsRenderer.subtitle.runs[0].text"`
	ChannelID          string `rjson:"contents.twoColumnWatchNextResults.results.results.contents[1].videoSecondaryInfoRenderer.owner.videoOwnerRenderer.title.runs[0].navigationEndpoint.browseEndpoint.browseId"`
	RawNewChannelID    string `rjson:"contents.twoColumnWatchNextResults.results.results.contents[1].videoSecondaryInfoRenderer.owner.videoOwnerRenderer.title.runs[0].navigationEndpoint.browseEndpoint.canonicalBaseUrl"`
	Likes              string `rjson:"contents.twoColumnWatchNextResults.results.results.contents[0].videoPrimaryInfoRenderer.videoActions.menuRenderer.topLevelButtons[0].segmentedLikeDislikeButtonRenderer.likeButton.toggleButtonRenderer.defaultText.simpleText"`
	ChannelSubscribers string `rjson:"contents.twoColumnWatchNextResults.results.results.contents[1].videoSecondaryInfoRenderer.owner.videoOwnerRenderer.subscriberCountText.simpleText"`
	CommentsCount      string `rjson:"contents.twoColumnWatchNextResults.results.results.contents[2].itemSectionRenderer.contents[0].commentsEntryPointHeaderRenderer.commentCount.simpleText"`
	Category           string `rjson:"contents.twoColumnWatchNextResults.results.results.contents[1].videoSecondaryInfoRenderer.metadataRowContainer.metadataRowContainerRenderer.rows[0].richMetadataRowRenderer.contents[1].richMetadataRenderer.title.runs[0].text"`

	Tokens []struct {
		Title string `rjson:"title"`
		Token string `rjson:"serviceEndpoint.continuationCommand.token"`
	} `rjson:"engagementPanels[].engagementPanelSectionListRenderer.header.engagementPanelTitleHeaderRenderer.menu.sortFilterSubMenuRenderer.subMenuItems[0]"`
}

func NewVideoScraper(id string) (v VideoScraper, err error) {
	v.url = fmt.Sprintf("https://www.youtube.com/watch?v=%s&hl=en", id)

	var rawJson string
	rawJson, err = extractInitialData(v.url)
	if err != nil {
		return
	}

	debugFileOutput([]byte(rawJson), "video_initial.json")

	var output videoInitialOutput
	if err = rjson.Unmarshal([]byte(rawJson), &output); err != nil {
		if errors.Unwrap(err) == rjson.ErrCantFindField {
			if Debug {
				log.Println("WARNING:", err)
			}
			err = nil
		}
		return
	}

	for _, token := range output.Tokens {
		switch token.Title {
		case "Top comments":
			v.commentsTopToken = token.Token
		case "Newest first":
			v.commentsNewestToken = token.Token
		}
	}

	v.commentsNewestContinueInputJson, err = continueInput{Continuation: v.commentsNewestToken}.FillGenericInfo().Construct()
	if err != nil {
		return
	}
	v.commentsTopContinueInputJson, err = continueInput{Continuation: v.commentsTopToken}.FillGenericInfo().Construct()
	if err != nil {
		return
	}

	date, wasLive := strings.CutPrefix(output.Date, "Streamed live on ")
	date, isLive := strings.CutPrefix(date, "Started streaming on ")

	v.video = FullVideo{
		VideoID:            id,
		Title:              output.Title,
		Description:        output.Description,
		Views:              strings.TrimSuffix(output.Views, " views"),
		IsLive:             output.IsLive || isLive,
		WasLive:            wasLive,
		Date:               date,
		Likes:              output.Likes,
		CommentsCount:      output.CommentsCount,
		Category:           output.Category,
		Username:           output.Username,
		ChannelID:          output.ChannelID,
		NewChannelID:       strings.TrimPrefix(output.RawNewChannelID, "/"),
		ChannelSubscribers: strings.TrimSuffix(output.ChannelSubscribers, " subscribers"),
	}

	return
}

func (v *VideoScraper) GetVideo() FullVideo {
	return v.video
}
