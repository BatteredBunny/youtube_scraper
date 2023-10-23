package scraper

import (
	"github.com/ayes-web/testingassert"
	"testing"
)

func TestLttVideo(t *testing.T) {
	scraper, err := NewVideoScraper("FdbvrqC6lOY") // normal VideoInfo
	if err != nil {
		t.Fatal(err)
	}

	video := scraper.VideoInfo
	assert.TestState = t
	assert.HideSuccess = true
	assert.Equals(video.VideoID, "FdbvrqC6lOY")
	assert.Equals(video.ChannelID, "UCXuqSBlHAE6Xw-yeJA0Tunw")
	assert.Equals(video.Username, "Linus Tech Tips")
	assert.Equals(video.NewChannelID, "@LinusTechTips")
	assert.NotEquals(video.Likes, "")
	assert.Equals(video.ChannelSubscribers, "15.6M")
	assert.NotEquals(video.Views, "")
	assert.NotEquals(video.Title, "")
	assert.NotEquals(video.Description, "", "description is empty")
	assert.Equals(video.Date, "Jul 13, 2023")
	assert.Assert(!video.WasLive, "marked as was live")
	assert.Assert(!video.IsLive, "marked as live")
	assert.NotEquals(video.CommentsCount, "")
}

func TestPastLttLivestream(t *testing.T) {
	scraper, err := NewVideoScraper("nfCUTZWwlvo") // past/ended livestream
	if err != nil {
		t.Fatal(err)
	}

	video := scraper.VideoInfo
	assert.TestState = t
	assert.HideSuccess = true
	assert.Assert(!video.IsLive, "marked as live")
	assert.Assert(video.WasLive, "fail: was not live")
	assert.Equals(video.VideoID, "nfCUTZWwlvo")
	assert.Equals(video.ChannelID, "UCXuqSBlHAE6Xw-yeJA0Tunw")
	assert.Equals(video.Username, "Linus Tech Tips")
	assert.Equals(video.NewChannelID, "@LinusTechTips")
	assert.NotEquals(video.Likes, "")
	assert.Equals(video.ChannelSubscribers, "15.6M")
	assert.NotEquals(video.Views, "")
	assert.NotEquals(video.Title, "")
	assert.NotEquals(video.Description, "", "description is empty")
	assert.Equals(video.Date, "Jul 7, 2023")
	assert.NotEquals(video.CommentsCount, "")
}

func TestLttLivestream(t *testing.T) {
	scraper, err := NewVideoScraper("DzLdFmPncms") // id to a running livestream
	if err != nil {
		t.Fatal(err)
	}

	video := scraper.VideoInfo
	assert.TestState = t
	assert.HideSuccess = true
	assert.Equals(video.VideoID, "DzLdFmPncms")
	assert.Equals(video.ChannelID, "UCXuqSBlHAE6Xw-yeJA0Tunw")
	assert.Equals(video.Username, "Linus Tech Tips")
	assert.Equals(video.NewChannelID, "@LinusTechTips")
	assert.NotEquals(video.Likes, "")
	assert.Equals(video.ChannelSubscribers, "15.6M")
	assert.NotEquals(video.Views, "")
	assert.NotEquals(video.Title, "")
	assert.NotEquals(video.Description, "", "description is empty")
	assert.NotEquals(video.Date, "")
	assert.Assert(video.IsLive, "fail: not live")
}
