package examples

import (
	"testing"

	scraper "github.com/BatteredBunny/youtube_scraper"
	"github.com/BatteredBunny/youtube_scraper/video"
)

func TestVideoCommentNewestScraper(t *testing.T) {
	v, err := video.NewVideoScraper("_jgcDuRbM_w")
	if err != nil {
		t.Fatal(err)
	}

	var comments []video.Comment
	for {
		comments, err = v.NextNewestCommentsPage()
		if err != nil {
			t.Fatal(err)
		} else if len(comments) == 0 {
			break
		}

		for _, comment := range comments {
			t.Log("id:", comment.CommentID, "likes:", comment.Likes, "replies:", comment.RepliesAmount)
		}
		t.Log("-------------")
	}
}
func TestVideoCommentTopScraper(t *testing.T) {
	scraper.Debug = true
	v, err := video.NewVideoScraper("FdbvrqC6lOY")
	if err != nil {
		t.Fatal(err)
	}

	var comments []video.Comment
	comments, err = v.NextTopCommentsPage()
	if err != nil {
		t.Fatal(err)
	}

	for _, comment := range comments {
		t.Log("id:", comment.CommentID, "likes:", comment.Likes, "replies:", comment.RepliesAmount)
	}
}

func TestSubcommentSection(t *testing.T) {
	v, err := video.NewVideoScraper("FdbvrqC6lOY")
	if err != nil {
		t.Fatal(err)
	}

	var comments []video.Comment
	comments, err = v.NextTopCommentsPage()
	if err != nil {
		t.Fatal(err)
	}

	for _, comment := range comments {
		if comment.HasSubComments() {
			var subComments []video.Comment
			for {
				subComments, err = comment.NextSubCommentPage()
				if err != nil {
					t.Fatal(err)
				} else if len(subComments) == 0 {
					break
				}

				for _, subComment := range subComments {
					t.Log(subComment.Content)
				}
			}
			break
		}
	}
}
