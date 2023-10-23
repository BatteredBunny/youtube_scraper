package examples

import (
	"git.catnip.ee/miisu/youtube_scraper"
	"testing"
)

func TestVideoCommentNewestScraper(t *testing.T) {
	v, err := scraper.NewVideoScraper("_jgcDuRbM_w")
	if err != nil {
		t.Fatal(err)
	}

	var comments []scraper.Comment
	for {
		comments, err = v.NextNewestCommentsPage()
		if err != nil {
			t.Fatal(err)
		} else if len(comments) == 0 {
			break
		}

		for _, comment := range comments {
			t.Log("id:", comment.CommentID, "content:", comment.Content, "date:", comment.PublishedTime)
		}
	}
}
func TestVideoCommentTopScraper(t *testing.T) {
	v, err := scraper.NewVideoScraper("FdbvrqC6lOY")
	if err != nil {
		t.Fatal(err)
	}

	var comments []scraper.Comment
	comments, err = v.NextTopCommentsPage()
	if err != nil {
		t.Fatal(err)
	}

	for _, comment := range comments {
		t.Log("id:", comment.CommentID, "likes:", comment.LikeAmount, "content:", comment.Content)
	}
}

func TestSubcommentSection(t *testing.T) {
	v, err := scraper.NewVideoScraper("FdbvrqC6lOY")
	if err != nil {
		t.Fatal(err)
	}

	var comments []scraper.Comment
	comments, err = v.NextTopCommentsPage()
	if err != nil {
		t.Fatal(err)
	}

	for _, comment := range comments {
		if comment.HasSubComments() {
			var subComments []scraper.Comment
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
