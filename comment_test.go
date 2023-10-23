package scraper

import "testing"

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
	scraper, err := NewVideoScraper("FdbvrqC6lOY")
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
			t.Log("id:", comment.CommentID, "replies amount:", comment.RepliesAmount)
		}
	}
}

func TestSubcommentSection(t *testing.T) {
	scraper, err := NewVideoScraper("FdbvrqC6lOY")
	if err != nil {
		t.Fatal(err)
	}

	var comments []Comment
	comments, err = scraper.NextTopCommentsPage()
	if err != nil {
		t.Fatal(err)
	}

	for _, comment := range comments {
		if comment.HasSubComments() {
			var subcomments []Comment
			for {
				subcomments, err = comment.NextSubCommentPage()
				if err != nil {
					t.Fatal(err)
				} else if len(subcomments) == 0 {
					break
				}

				for _, subcomment := range subcomments {
					t.Log(subcomment.Content)
				}
			}
			break
		}
	}
}
