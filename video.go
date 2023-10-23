package youtube_scraper

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type VideoScraper struct {
	url   string
	video FullVideo
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

func NewVideoScraper(id string) (v VideoScraper, err error) {
	v.url = fmt.Sprintf("https://www.youtube.com/watch?v=%s&hl=en", id)

	var rawjson string
	rawjson, err = ExtractInitialData(v.url)
	if err != nil {
		return
	}

	if DEBUG {
		os.WriteFile("video_initial.json", []byte(rawjson), 0777)
	}

	var a initialData
	if err = json.Unmarshal([]byte(rawjson), &a); err != nil {
		return
	}

	v.video = FullVideo{
		VideoID:            id,
		Title:              a.PlayerOverlays.PlayerOverlayRenderer.VideoDetails.PlayerOverlayVideoDetailsRenderer.Title.SimpleText,
		Description:        a.EngagementPanels[1].EngagementPanelSectionListRenderer.Content.StructuredDescriptionContentRenderer.Items[1].ExpandableVideoDescriptionBodyRenderer.AttributedDescriptionBodyText.Content,
		Views:              a.PlayerOverlays.PlayerOverlayRenderer.VideoDetails.PlayerOverlayVideoDetailsRenderer.Subtitle.Runs[2].Text,
		Date:               a.EngagementPanels[1].EngagementPanelSectionListRenderer.Content.StructuredDescriptionContentRenderer.Items[0].VideoDescriptionHeaderRenderer.PublishDate.SimpleText,
		Username:           a.PlayerOverlays.PlayerOverlayRenderer.VideoDetails.PlayerOverlayVideoDetailsRenderer.Subtitle.Runs[0].Text,
		ChannelID:          a.EngagementPanels[1].EngagementPanelSectionListRenderer.Content.StructuredDescriptionContentRenderer.Items[0].VideoDescriptionHeaderRenderer.ChannelNavigationEndpoint.BrowseEndpoint.BrowseId,
		NewChannelID:       strings.TrimPrefix(a.EngagementPanels[1].EngagementPanelSectionListRenderer.Content.StructuredDescriptionContentRenderer.Items[0].VideoDescriptionHeaderRenderer.ChannelNavigationEndpoint.BrowseEndpoint.CanonicalBaseUrl, "/"),
		Likes:              a.EngagementPanels[1].EngagementPanelSectionListRenderer.Content.StructuredDescriptionContentRenderer.Items[0].VideoDescriptionHeaderRenderer.Factoid[0].FactoidRenderer.Value.SimpleText,
		ChannelSubscribers: a.EngagementPanels[1].EngagementPanelSectionListRenderer.Content.StructuredDescriptionContentRenderer.Items[2].VideoDescriptionInfocardsSectionRenderer.SectionSubtitle.SimpleText,
		CommentsCount:      a.EngagementPanels[2].EngagementPanelSectionListRenderer.Header.EngagementPanelTitleHeaderRenderer.ContextualInfo.Runs[0].Text,
		Category:           a.Contents.TwoColumnWatchNextResults.Results.Results.Contents[1].VideoSecondaryInfoRenderer.MetadataRowContainer.MetadataRowContainerRenderer.Rows[0].RichMetadataRowRenderer.Contents[1].RichMetadataRenderer.Title.Runs[0].Text,
	}

	return
}

func (v VideoScraper) GetVideo() FullVideo {
	return v.video
}
