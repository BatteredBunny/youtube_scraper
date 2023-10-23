package scraper

import (
	"errors"
	"fmt"
	"github.com/ayes-web/rjson"
	"log"
	"os"
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

type FullVideo struct {
	VideoID       string `json:"VideoID"`
	Title         string `json:"Title"`
	Description   string `json:"Description"`
	Views         string `json:"Views"`
	Date          string `json:"Date"`
	Likes         string `json:"Likes"`
	CommentsCount string `json:"CommentsCount"`
	Category      string `json:"Category"`

	Username           string `json:"Username"`
	ChannelID          string `json:"ChannelID"`
	NewChannelID       string `json:"NewChannelID"`
	ChannelSubscribers string `json:"ChannelSubscribers"`
}

type VideoInitialOutput struct {
	Title              string `rjson:"playerOverlays.playerOverlayRenderer.videoDetails.playerOverlayVideoDetailsRenderer.title.simpleText"`
	Description        string `rjson:"engagementPanels[1].engagementPanelSectionListRenderer.content.structuredDescriptionContentRenderer.items[1].expandableVideoDescriptionBodyRenderer.attributedDescriptionBodyText.content"`
	Views              string `rjson:"playerOverlays.playerOverlayRenderer.videoDetails.playerOverlayVideoDetailsRenderer.subtitle.runs[2].text"`
	Date               string `rjson:"engagementPanels[1].engagementPanelSectionListRenderer.content.structuredDescriptionContentRenderer.items[0].videoDescriptionHeaderRenderer.publishDate.simpleText"`
	Username           string `rjson:"playerOverlays.playerOverlayRenderer.videoDetails.playerOverlayVideoDetailsRenderer.subtitle.runs[0].text"`
	ChannelID          string `rjson:"engagementPanels[1].engagementPanelSectionListRenderer.content.structuredDescriptionContentRenderer.items[0].videoDescriptionHeaderRenderer.channelNavigationEndpoint.browseEndpoint.browseId"`
	RawNewChannelID    string `rjson:"engagementPanels[1].engagementPanelSectionListRenderer.content.structuredDescriptionContentRenderer.items[0].videoDescriptionHeaderRenderer.channelNavigationEndpoint.browseEndpoint.canonicalBaseUrl"`
	Likes              string `rjson:"engagementPanels[1].engagementPanelSectionListRenderer.content.structuredDescriptionContentRenderer.items[0].videoDescriptionHeaderRenderer.factoid[0].factoidRenderer.value.simpleText"`
	ChannelSubscribers string `rjson:"engagementPanels[1].engagementPanelSectionListRenderer.content.structuredDescriptionContentRenderer.items[2].videoDescriptionInfocardsSectionRenderer.sectionSubtitle.simpleText"`
	CommentsCount      string `rjson:"engagementPanels[2].engagementPanelSectionListRenderer.header.engagementPanelTitleHeaderRenderer.contextualInfo.runs[0].text"`
	Category           string `rjson:"contents.twoColumnWatchNextResults.results.results.contents[1].videoSecondaryInfoRenderer.metadataRowContainer.metadataRowContainerRenderer.rows[0].richMetadataRowRenderer.contents[1].richMetadataRenderer.title.runs[0].text"`

	CommentsTopToken    string `rjson:"engagementPanels[2].engagementPanelSectionListRenderer.header.engagementPanelTitleHeaderRenderer.menu.sortFilterSubMenuRenderer.subMenuItems[0].serviceEndpoint.continuationCommand.token"`
	CommentsNewestToken string `rjson:"engagementPanels[2].engagementPanelSectionListRenderer.header.engagementPanelTitleHeaderRenderer.menu.sortFilterSubMenuRenderer.subMenuItems[1].serviceEndpoint.continuationCommand.token"`
}

func NewVideoScraper(id string) (v VideoScraper, err error) {
	v.url = fmt.Sprintf("https://www.youtube.com/watch?v=%s&hl=en", id)

	var rawJson string
	rawJson, err = ExtractInitialData(v.url)
	if err != nil {
		return
	}

	if Debug {
		fmt.Println("writing initial output to \"video_initial.json\"")
		os.WriteFile("video_initial.json", []byte(rawJson), 0777)
	}

	var output VideoInitialOutput
	if err = rjson.Unmarshal([]byte(rawJson), &output); err != nil {
		if errors.Unwrap(err) == rjson.ErrCantFindField {
			if Debug {
				log.Println("WARNING:", err)
			}
			err = nil
		}
		return
	}

	v.commentsNewestToken = output.CommentsNewestToken
	v.commentsNewestContinueInputJson, err = continueInput{Continuation: v.commentsNewestToken}.FillGenericInfo().Construct()
	if err != nil {
		return
	}

	v.commentsTopToken = output.CommentsTopToken
	v.commentsTopContinueInputJson, err = continueInput{Continuation: v.commentsTopToken}.FillGenericInfo().Construct()
	if err != nil {
		return
	}

	v.video = FullVideo{
		VideoID:            id,
		Title:              output.Title,
		Description:        output.Description,
		Views:              output.Views,
		Date:               output.Date,
		Likes:              output.Likes,
		CommentsCount:      output.CommentsCount,
		Category:           output.Category,
		Username:           output.Username,
		ChannelID:          output.ChannelID,
		NewChannelID:       strings.TrimPrefix(output.RawNewChannelID, "/"),
		ChannelSubscribers: output.ChannelSubscribers,
	}

	return
}

func (v *VideoScraper) GetVideo() FullVideo {
	return v.video
}
