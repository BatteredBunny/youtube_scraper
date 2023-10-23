package scraper

import (
	"bytes"
	"errors"
	"github.com/ayes-web/rjson"
	"io"
	"log"
	"net/http"
	"strings"
)

type Comment struct {
	NewChannelID  string
	CommentID     string
	Content       string
	PublishedTime string // "5 days ago"
	LikeAmount    string // Empty means 0 likes
	PinnedBy      string // "Pinned by Username"
	IsPinned      bool
	IsHearted     bool
	WasEdited     bool // if the comment was edited
	RepliesAmount string

	repliesToken             string
	repliesContinueInputJson []byte
}

// HasSubComments returns if the comment has any replies
func (c *Comment) HasSubComments() bool {
	return c.repliesToken != ""
}

type subCommentsContinueOutput struct {
	Comments []struct {
		NewChannelID  string   `rjson:"commentRenderer.authorText.simpleText"`
		CommentID     string   `rjson:"commentRenderer.commentId"`
		Content       []string `rjson:"commentRenderer.contentText.runs[].text"`
		PublishedTime string   `rjson:"commentRenderer.publishedTimeText.runs[0].text"` // ends with "(edited)" if the comment has been edited
		LikeAmount    string   `rjson:"commentRenderer.voteCount.simpleText"`           // 3K
		Pinned        []string `rjson:"commentRenderer.pinnedCommentBadge.pinnedCommentBadgeRenderer.label.runs[].text"`
		IsHearted     bool     `rjson:"commentRenderer.actionButtons.commentActionButtonsRenderer.creatorHeart.creatorHeartRenderer.isHearted"`
	} `rjson:"onResponseReceivedEndpoints[0]appendContinuationItemsAction.continuationItems"`
	ContinueToken string `rjson:"onResponseReceivedEndpoints[0]appendContinuationItemsAction.continuationItems[-].continuationItemRenderer.button.buttonRenderer.command.continuationCommand.token"`
}

// NextSubCommentPage returns comment replies in chunks. Check with HasSubComments if there are replies
func (c *Comment) NextSubCommentPage() (comments []Comment, err error) {
	var resp *http.Response
	resp, err = http.Post("https://www.youtube.com/youtubei/v1/next", "application/json", bytes.NewReader(c.repliesContinueInputJson))
	if err != nil {
		return
	}
	c.repliesContinueInputJson = []byte{}

	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	debugFileOutput(body, "subcomment_%s.json", c.repliesToken)

	var output subCommentsContinueOutput
	if err = rjson.Unmarshal(body, &output); err != nil {
		if errors.Unwrap(err) == rjson.ErrCantFindField {
			if Debug {
				log.Println("WARNING:", err)
			}
			err = nil
		}
		return
	}

	c.repliesToken = output.ContinueToken
	c.repliesContinueInputJson, err = continueInput{Continuation: c.repliesToken}.FillGenericInfo().Construct()
	if err != nil {
		return
	}

	for _, comment := range output.Comments {
		if comment.CommentID == "" {
			continue
		}

		publishedTime, wasEdited := strings.CutSuffix(comment.PublishedTime, " (edited)")
		comments = append(comments, Comment{
			NewChannelID:  comment.NewChannelID,
			CommentID:     comment.CommentID,
			Content:       strings.Join(comment.Content, ""),
			PublishedTime: publishedTime,
			WasEdited:     wasEdited,
			LikeAmount:    comment.LikeAmount,
			PinnedBy:      strings.Join(comment.Pinned, ""),
			IsPinned:      len(comment.Pinned) > 0,
			IsHearted:     comment.IsHearted,
		})
	}

	return
}

type commentContinueOutputComment struct {
	NewChannelID  string   `rjson:"commentThreadRenderer.comment.commentRenderer.authorText.simpleText"`
	CommentID     string   `rjson:"commentThreadRenderer.comment.commentRenderer.commentId"`
	Content       []string `rjson:"commentThreadRenderer.comment.commentRenderer.contentText.runs[].text"`
	PublishedTime string   `rjson:"commentThreadRenderer.comment.commentRenderer.publishedTimeText.runs[0].text"` // ends with "(edited)" if the comment has been edited
	LikeAmount    string   `rjson:"commentThreadRenderer.comment.commentRenderer.voteCount.simpleText"`           // 3K
	Pinned        []string `rjson:"commentThreadRenderer.comment.commentRenderer.pinnedCommentBadge.pinnedCommentBadgeRenderer.label.runs[].text"`
	IsHearted     bool     `rjson:"commentThreadRenderer.comment.commentRenderer.actionButtons.commentActionButtonsRenderer.creatorHeart.creatorHeartRenderer.isHearted"`
	RepliesAmount string   `rjson:"commentThreadRenderer.replies.commentRepliesRenderer.viewReplies.buttonRenderer.text.runs[0].text"`
	RepliesToken  string   `rjson:"commentThreadRenderer.replies.commentRepliesRenderer.contents[-].continuationItemRenderer.continuationEndpoint.continuationCommand.token"`
}

type commentsContinueOutputInitial struct {
	Comments      []commentContinueOutputComment `rjson:"onResponseReceivedEndpoints[1]reloadContinuationItemsCommand.continuationItems"`
	ContinueToken string                         `rjson:"onResponseReceivedEndpoints[1]reloadContinuationItemsCommand.continuationItems[-]continuationItemRenderer.continuationEndpoint.continuationCommand.token"`
}
type commentsContinueOutput struct {
	Comments      []commentContinueOutputComment `rjson:"onResponseReceivedEndpoints[0]appendContinuationItemsAction.continuationItems"`
	ContinueToken string                         `rjson:"onResponseReceivedEndpoints[0]appendContinuationItemsAction.continuationItems[-]continuationItemRenderer.continuationEndpoint.continuationCommand.token"`
}

func genericNextCommentsPage(token *string, continueInputJson *[]byte, outputGeneric func(rawJson []byte) (rawToken string, rawComments []commentContinueOutputComment, err error)) (comments []Comment, err error) {
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

	rawToken, rawComments, err := outputGeneric(body)
	if err != nil {
		return
	}

	*token = rawToken
	*continueInputJson, err = continueInput{Continuation: *token}.FillGenericInfo().Construct()
	if err != nil {
		return
	}

	for _, comment := range rawComments {
		if comment.CommentID == "" {
			continue
		}

		var (
			repliesToken             string
			repliesContinueInputJson []byte
		)
		if comment.RepliesToken != "" {
			repliesToken = comment.RepliesToken
			repliesContinueInputJson, err = continueInput{Continuation: repliesToken}.FillGenericInfo().Construct()
			if err != nil {
				return
			}
		}

		publishedTime, wasEdited := strings.CutSuffix(comment.PublishedTime, " (edited)")
		comments = append(comments, Comment{
			NewChannelID:             comment.NewChannelID,
			CommentID:                comment.CommentID,
			Content:                  strings.Join(comment.Content, ""),
			PublishedTime:            publishedTime,
			WasEdited:                wasEdited,
			LikeAmount:               comment.LikeAmount,
			PinnedBy:                 strings.Join(comment.Pinned, ""),
			IsPinned:                 len(comment.Pinned) > 0,
			IsHearted:                comment.IsHearted,
			RepliesAmount:            strings.TrimSuffix(strings.TrimSuffix(comment.RepliesAmount, " replies"), " reply"),
			repliesToken:             repliesToken,
			repliesContinueInputJson: repliesContinueInputJson,
		})
	}

	return
}

// NextNewestCommentsPage returns comments in chunks sorted by newest
func (v *VideoScraper) NextNewestCommentsPage() (comments []Comment, err error) {
	return genericNextCommentsPage(&v.commentsNewestToken, &v.commentsNewestContinueInputJson, func(rawJson []byte) (rawToken string, rawComments []commentContinueOutputComment, err error) {
		if !v.commentsNewestPassedInitial {
			debugFileOutput(rawJson, "comment_newest_initial_%s.json", v.commentsNewestToken)

			var output commentsContinueOutputInitial
			if err = rjson.Unmarshal(rawJson, &output); err != nil {
				if errors.Unwrap(err) == rjson.ErrCantFindField {
					if Debug {
						log.Println("WARNING:", err)
					}
					err = nil
				}
				return
			}

			rawToken = output.ContinueToken
			rawComments = output.Comments
			v.commentsNewestPassedInitial = true
		} else {
			debugFileOutput(rawJson, "comment_newest_%s.json", v.commentsNewestToken)

			var output commentsContinueOutput
			if err = rjson.Unmarshal(rawJson, &output); err != nil {
				if errors.Unwrap(err) == rjson.ErrCantFindField {
					if Debug {
						log.Println("WARNING:", err)
					}
					err = nil
				}
				return
			}

			rawToken = output.ContinueToken
			rawComments = output.Comments
		}

		return
	})
}

// NextTopCommentsPage returns comments in chunks sorted by most popular
func (v *VideoScraper) NextTopCommentsPage() (comments []Comment, err error) {
	return genericNextCommentsPage(&v.commentsTopToken, &v.commentsTopContinueInputJson, func(rawJson []byte) (rawToken string, rawComments []commentContinueOutputComment, err error) {
		debugFileOutput(rawJson, "comment_top_%s.json", v.commentsTopToken)

		if !v.commentsTopPassedInitial {
			var output commentsContinueOutputInitial
			if err = rjson.Unmarshal(rawJson, &output); err != nil {
				if errors.Unwrap(err) == rjson.ErrCantFindField {
					if Debug {
						log.Println("WARNING:", err)
					}
					err = nil
				}
				return
			}

			rawToken = output.ContinueToken
			rawComments = output.Comments
			v.commentsTopPassedInitial = true
		} else {
			var output commentsContinueOutput
			if err = rjson.Unmarshal(rawJson, &output); err != nil {
				if errors.Unwrap(err) == rjson.ErrCantFindField {
					if Debug {
						log.Println("WARNING:", err)
					}
					err = nil
				}
				return
			}

			rawToken = output.ContinueToken
			rawComments = output.Comments
		}

		return
	})
}
