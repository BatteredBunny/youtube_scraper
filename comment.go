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

type Comment struct {
	NewChannelID  string
	CommentID     string
	Content       string
	PublishedTime string
	LikeAmount    string // Empty means 0 likes
	PinnedBy      string // "Pinned by Username"
	IsPinned      bool
	IsHearted     bool
}

type commentContinueOutputComment struct {
	NewChannelID  string   `rjson:"commentThreadRenderer.comment.commentRenderer.authorText.simpleText"`
	CommentID     string   `rjson:"commentThreadRenderer.comment.commentRenderer.commentId"`
	Content       []string `rjson:"commentThreadRenderer.comment.commentRenderer.contentText.runs[].text"`
	PublishedTime string   `rjson:"commentThreadRenderer.comment.commentRenderer.publishedTimeText.runs[0].text"` // ends with "(edited)" if the comment has been edited
	LikeAmount    string   `rjson:"commentThreadRenderer.comment.commentRenderer.voteCount.simpleText"`           // 3K
	Pinned        []string `rjson:"commentThreadRenderer.comment.commentRenderer.pinnedCommentBadge.pinnedCommentBadgeRenderer.label.runs[].text"`
	IsHearted     bool     `rjson:"commentThreadRenderer.comment.commentRenderer.actionButtons.commentActionButtonsRenderer.creatorHeart.creatorHeartRenderer.isHearted"`
}

type commentsContinueOutputInitial struct {
	Comments      []commentContinueOutputComment `rjson:"onResponseReceivedEndpoints[1]reloadContinuationItemsCommand.continuationItems"`
	ContinueToken string                         `rjson:"onResponseReceivedEndpoints[1]reloadContinuationItemsCommand.continuationItems[-]continuationItemRenderer.continuationEndpoint.continuationCommand.token"`
}

func (n commentsContinueOutputInitial) GetComments() []commentContinueOutputComment {
	return n.Comments
}
func (n commentsContinueOutputInitial) GetContinueToken() string {
	return n.ContinueToken
}

type commentsContinueOutput struct {
	Comments      []commentContinueOutputComment `rjson:"onResponseReceivedEndpoints[0]appendContinuationItemsAction.continuationItems"`
	ContinueToken string                         `rjson:"onResponseReceivedEndpoints[0]appendContinuationItemsAction.continuationItems[-]continuationItemRenderer.continuationEndpoint.continuationCommand.token"`
}

func (n commentsContinueOutput) GetComments() []commentContinueOutputComment {
	return n.Comments
}
func (n commentsContinueOutput) GetContinueToken() string {
	return n.ContinueToken
}

type commentsContinueOutputCommon interface {
	GetComments() []commentContinueOutputComment
	GetContinueToken() string
}

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
			IsHearted:     comment.IsHearted,
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
