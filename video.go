package scraper

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/ayes-web/rjson"
	"io"
	"log"
	"net/http"
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

type Comment struct {
	NewChannelID  string
	CommentID     string
	Content       string
	PublishedTime string
	LikeAmount    string // Empty means 0 likes
	PinnedBy      string // "Pinned by Username"
	IsPinned      bool
}

type newestCommentContinueOutputComment struct {
	NewChannelID  string   `rjson:"commentThreadRenderer.comment.commentRenderer.authorText.simpleText"`
	CommentID     string   `rjson:"commentThreadRenderer.comment.commentRenderer.commentId"`
	Content       []string `rjson:"commentThreadRenderer.comment.commentRenderer.contentText.runs[].text"`
	PublishedTime string   `rjson:"commentThreadRenderer.comment.commentRenderer.publishedTimeText.runs[0].text"` // ends with "(edited)" if the comment has been edited
	LikeAmount    string   `rjson:"commentThreadRenderer.comment.commentRenderer.voteCount.simpleText"`           // 3K
	Pinned        []string `rjson:"commentThreadRenderer.comment.commentRenderer.pinnedCommentBadge.pinnedCommentBadgeRenderer.label.runs[].text"`
}

type commentsContinueOutputInitial struct {
	Comments      []newestCommentContinueOutputComment `rjson:"onResponseReceivedEndpoints[1]reloadContinuationItemsCommand.continuationItems"`
	ContinueToken string                               `rjson:"onResponseReceivedEndpoints[1]reloadContinuationItemsCommand.continuationItems[-]continuationItemRenderer.continuationEndpoint.continuationCommand.token"`
}

func (n commentsContinueOutputInitial) GetComments() []newestCommentContinueOutputComment {
	return n.Comments
}
func (n commentsContinueOutputInitial) GetContinueToken() string {
	return n.ContinueToken
}

type commentsContinueOutput struct {
	Comments      []newestCommentContinueOutputComment `rjson:"onResponseReceivedEndpoints[0]appendContinuationItemsAction.continuationItems"`
	ContinueToken string                               `rjson:"onResponseReceivedEndpoints[0]appendContinuationItemsAction.continuationItems[-]continuationItemRenderer.continuationEndpoint.continuationCommand.token"`
}

func (n commentsContinueOutput) GetComments() []newestCommentContinueOutputComment {
	return n.Comments
}
func (n commentsContinueOutput) GetContinueToken() string {
	return n.ContinueToken
}

type commentsContinueOutputCommon interface {
	GetComments() []newestCommentContinueOutputComment
	GetContinueToken() string
}

// TODO: add field indicating if comment is liked by creator badge
// TODO: comment subsection/comment replies
func genericNextCommentsPage(token *string, continueInputJson *[]byte, commentsPassedInitial *bool) (comments []Comment, err error) {
	var resp *http.Response
	resp, err = http.Post("https://www.youtube.com/youtubei/v1/next", "application/json", bytes.NewReader(*continueInputJson))
	if err != nil {
		return
	}
	*continueInputJson = []byte{}

	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if Debug {
		fmt.Printf("writing current api output to \"comment_%s.json\"\n", *token)
		os.WriteFile(fmt.Sprintf("comment_%s.json", *token), body, 0777)
	}

	var output commentsContinueOutputCommon
	if !*commentsPassedInitial {
		var tempout commentsContinueOutputInitial
		if err = rjson.Unmarshal(body, &tempout); err != nil {
			if errors.Unwrap(err) == rjson.ErrCantFindField {
				if Debug {
					log.Println("WARNING:", err)
				}
				err = nil
			}
			return
		}

		*commentsPassedInitial = true
		output = tempout
	} else {
		var tempout commentsContinueOutput
		if err = rjson.Unmarshal(body, &tempout); err != nil {
			if errors.Unwrap(err) == rjson.ErrCantFindField {
				if Debug {
					log.Println("WARNING:", err)
				}
				err = nil
			}
			return
		}
		output = tempout
	}

	*token = output.GetContinueToken()
	*continueInputJson, err = continueInput{Continuation: *token}.FillGenericInfo().Construct()
	if err != nil {
		return
	}

	for _, comment := range output.GetComments() {
		if comment.CommentID == "" {
			continue
		}

		comments = append(comments, Comment{
			NewChannelID:  comment.NewChannelID,
			CommentID:     comment.CommentID,
			Content:       strings.Join(comment.Content, ""),
			PublishedTime: comment.PublishedTime,
			LikeAmount:    comment.LikeAmount,
			PinnedBy:      strings.Join(comment.Pinned, ""),
			IsPinned:      len(comment.Pinned) > 0,
		})
	}

	return
}

func (v *VideoScraper) NextNewestCommentsPage() (comments []Comment, err error) {
	return genericNextCommentsPage(&v.commentsNewestToken, &v.commentsNewestContinueInputJson, &v.commentsNewestPassedInitial)
}
func (v *VideoScraper) NextTopCommentsPage() (comments []Comment, err error) {
	return genericNextCommentsPage(&v.commentsTopToken, &v.commentsTopContinueInputJson, &v.commentsTopPassedInitial)
}
