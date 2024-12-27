package video

import (
	"net/http"
	"testing"
	"time"

	assert "github.com/BatteredBunny/testingassert"
)

func TestLttVideo(t *testing.T) {
	scraper, err := NewVideoScraper("FdbvrqC6lOY") // normal VideoInfo
	if err != nil {
		t.Fatal(err)
	}

	var date time.Time
	date, err = time.Parse(YoutubeVideoDateLayout, "Jul 13, 2023")
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
	assert.NotEquals(video.Likes, 0)
	assert.Equals(video.ChannelSubscribers, 16100000) // TODO: please do not relay on channel subscribers, bad bad bad
	assert.NotEquals(video.Views, 0)
	assert.NotEquals(video.Title, "")
	assert.NotEquals(video.Description, "", "description is empty")
	assert.Equals(video.Date, date)
	assert.Assert(!video.WasLive, "marked as was live")
	assert.Assert(!video.IsLive, "marked as live")
	assert.NotEquals(video.CommentsCount, "")
	assert.Assert(video.ChannelIsVerified, "channel isnt verified")
	assert.Assert(!video.ChannelIsVerifiedArtist, "channel shouldnt be an artist")
}

func TestPastLttLivestream(t *testing.T) {
	scraper, err := NewVideoScraper("nfCUTZWwlvo") // past/ended livestream
	if err != nil {
		t.Fatal(err)
	}

	var date time.Time
	date, err = time.Parse(YoutubeVideoDateLayout, "Jul 7, 2023")
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
	assert.NotEquals(video.Likes, 0)
	assert.Equals(video.ChannelSubscribers, 16100000) // TODO: please do not relay on channel subscribers, bad bad bad
	assert.NotEquals(video.Views, 0)
	assert.NotEquals(video.Title, "")
	assert.NotEquals(video.Description, "", "description is empty")
	assert.Equals(video.Date, date)
	// TODO: fix
	// assert.NotEquals(video.CommentsCount, 0)
	assert.Assert(video.ChannelIsVerified, "channel isnt verified")
	assert.Assert(!video.ChannelIsVerifiedArtist, "channel shouldnt be an artist")
}

// TODO: find some other forever running livestream
// func TestLttLivestream(t *testing.T) {
// 	scraper.Debug = true
// 	scraper, err := NewVideoScraper("DzLdFmPncms") // id to a running livestream
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	video := scraper.VideoInfo
// 	assert.TestState = t
// 	assert.HideSuccess = true
// 	assert.Equals(video.VideoID, "DzLdFmPncms")
// 	assert.Equals(video.ChannelID, "UCXuqSBlHAE6Xw-yeJA0Tunw")
// 	assert.Equals(video.Username, "Linus Tech Tips")
// 	assert.Equals(video.NewChannelID, "@LinusTechTips")
// 	assert.NotEquals(video.Likes, 0)
// 	assert.Equals(video.ChannelSubscribers, 15600000)
// 	assert.Equals(video.CommentsCount, 0)
// 	assert.NotEquals(video.Views, 0)
// 	assert.NotEquals(video.Title, "")
// 	assert.NotEquals(video.Description, "", "description is empty")
// 	assert.NotEquals(video.Date, "")
// 	assert.Assert(video.IsLive, "fail: not live")
// 	assert.Assert(video.ChannelIsVerified, "channel isnt verified")
// 	assert.Assert(!video.ChannelIsVerifiedArtist, "channel shouldnt be an artist")
// }

func TestVerified(t *testing.T) {
	scraper, err := NewVideoScraper("twHFPMoJNXE")
	if err != nil {
		t.Fatal(err)
	}

	video := scraper.VideoInfo
	assert.TestState = t
	assert.HideSuccess = true
	assert.Equals(video.ChannelID, "UCwbRile4jo-LcW_PQwmMdBw")
	assert.Equals(video.VideoID, "twHFPMoJNXE")
	assert.Equals(video.Username, "Captain KRB")
	assert.Equals(video.NewChannelID, "@CaptainKRB")
	assert.Assert(video.ChannelIsVerified, "channel should be verified")
	assert.Assert(!video.ChannelIsVerifiedArtist, "channel shouldnt be an artist")
}

func TestArtistVideo(t *testing.T) {
	scraper, err := NewVideoScraper("U3ASj1L6_sY")
	if err != nil {
		t.Fatal(err)
	}

	video := scraper.VideoInfo
	assert.TestState = t
	assert.HideSuccess = true
	assert.Equals(video.ChannelID, "UCsRM0YB_dabtEPGPTKo-gcw")
	assert.Equals(video.VideoID, "U3ASj1L6_sY")
	assert.Equals(video.Title, "Adele - Easy On Me (Official Video)")
	assert.Assert(!video.ChannelIsVerified, "channel shouldnt be verified")
	assert.Assert(video.ChannelIsVerifiedArtist, "channel should be an artist")
}

func TestUnlistedVideo(t *testing.T) {
	scraper, err := NewVideoScraper("NkpskWvac3U")
	if err != nil {
		t.Fatal(err)
	}

	var date time.Time
	date, err = time.Parse(YoutubeVideoDateLayout, "Apr 27, 2019")
	if err != nil {
		t.Fatal(err)
	}

	video := scraper.VideoInfo
	assert.TestState = t
	assert.HideSuccess = true
	assert.Equals(video.ChannelID, "UCFQMnBA3CS502aghlcr0_aw")
	assert.Equals(video.VideoID, "NkpskWvac3U")
	assert.Equals(video.Username, "Coffeezilla")
	assert.Equals(video.NewChannelID, "@Coffeezilla")
	assert.Equals(video.Title, "Pewdiepie's Last Hope - Save a Swede By Going AFK")
	assert.Equals(video.Date, date)
	assert.Assert(video.ChannelIsVerified, "channel should be verified")
	assert.Assert(!video.ChannelIsVerifiedArtist, "channel shouldnt be an artist")
	assert.Assert(video.IsUnlisted, "should be unlisted")
}

// TODO: make it actually check if the media url is valid
func TestMediaUrl(t *testing.T) {
	assert.TestState = t
	v, err := NewVideoScraper("UXqq0ZvbOnk")
	if err != nil {
		t.Fatal(err)
	}

	output, err := v.ExtractMediaFormats()
	if err != nil {
		t.Fatal(err)
	}

	var bestMediaFormat MediaFormat
	for _, format := range output.AdaptiveFormats {
		if format.Bitrate > bestMediaFormat.Bitrate {
			bestMediaFormat = format
		}
	}

	out, err := bestMediaFormat.GetMediaUrl(&v)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEquals(out, "")

	resp, err := http.Get(out)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equals(resp.StatusCode, http.StatusOK)
}

func TestMediaUrlDrm(t *testing.T) {
	assert.TestState = t
	v, err := NewVideoScraper("rYEDA3JcQqw")
	if err != nil {
		t.Fatal(err)
	}

	output, err := v.ExtractMediaFormats()
	if err != nil {
		t.Fatal(err)
	}

	var bestMediaFormat MediaFormat
	for _, format := range output.AdaptiveFormats {
		if format.Bitrate > bestMediaFormat.Bitrate {
			bestMediaFormat = format
		}
	}

	out, err := bestMediaFormat.GetMediaUrl(&v)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEquals(out, "")

	resp, err := http.Get(out)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equals(resp.StatusCode, http.StatusOK)
}
