package scraper

import (
	"testing"
)

func TestVideoScraper(t *testing.T) {
	scraper, err := NewVideoScraper("51o5J0XVGoc")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(scraper.GetVideo())
}

func TestVideoCommentNewestScraper(t *testing.T) {
	scraper, err := NewVideoScraper("_jgcDuRbM_w")
	if err != nil {
		t.Fatal(err)
	}

	var comments []Comment
	for {
		comments, err = scraper.NextNewestCommentsPage()
		if err != nil {
			t.Fatal(err)
		} else if len(comments) == 0 {
			break
		}

		for _, comment := range comments {
			t.Log("id:", comment.CommentID, "content:", comment.Content)
		}
	}
}
func TestVideoCommentTopScraper(t *testing.T) {
	scraper, err := NewVideoScraper("_jgcDuRbM_w")
	if err != nil {
		t.Fatal(err)
	}

	var comments []Comment
	for {
		comments, err = scraper.NextTopCommentsPage()
		if err != nil {
			t.Fatal(err)
		} else if len(comments) == 0 {
			break
		}

		for _, comment := range comments {
			t.Log("id:", comment.CommentID, "content:", comment.Content, "likes:", comment.LikeAmount)
		}
	}
}
