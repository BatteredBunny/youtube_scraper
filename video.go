package scraper

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ayes-web/rjson"
	"github.com/dustin/go-humanize"
)

type VideoScraper struct {
	VideoInfo             FullVideo // Initial info about the video, contains info by default
	InitialSidebarEntries []SidebarEntry

	mediaUrlJs string
	url        string

	commentsNewestPassedInitial     bool
	commentsNewestContinueInput     ContinueInput
	commentsNewestContinueInputJson []byte

	commentsTopPassedInitial     bool
	commentsTopContinueInput     ContinueInput
	commentsTopContinueInputJson []byte

	sidebarContinueInput     ContinueInput
	sidebarContinueInputJson []byte
}

// FullVideo has the full metadata unlike Video which is fetched from Video lists
type FullVideo struct {
	VideoID        string
	Title          string
	Description    string
	Views          int // Displays number of video views exept in a livestream where it will display number of viewers
	IsLive         bool
	WasLive        bool      // if this video was live in the past
	Date           time.Time // video upload date
	Likes          int
	CommentsCount  int
	Category       string // video category
	IsUnlisted     bool
	VideoPremiered bool // if the video was premiered in the past

	Username           string
	ChannelID          string
	NewChannelID       string
	ChannelSubscribers int
	ChannelAvatars     []YoutubeImage

	ChannelIsVerified       bool
	ChannelIsVerifiedArtist bool
}

type videoInitialOutput struct {
	Title              string         `rjson:"playerOverlays.playerOverlayRenderer.videoDetails.playerOverlayVideoDetailsRenderer.title.simpleText"`
	Description        string         `rjson:"contents.twoColumnWatchNextResults.results.results.contents[1].videoSecondaryInfoRenderer.attributedDescription.content"`
	Views              string         `rjson:"playerOverlays.playerOverlayRenderer.videoDetails.playerOverlayVideoDetailsRenderer.subtitle.runs[2].text"`
	IsLive             bool           `rjson:"contents.twoColumnWatchNextResults.results.results.contents[0].videoPrimaryInfoRenderer.viewCount.videoViewCountRenderer.isLive"`
	Date               string         `rjson:"contents.twoColumnWatchNextResults.results.results.contents[0].videoPrimaryInfoRenderer.dateText.simpleText"`
	Username           string         `rjson:"playerOverlays.playerOverlayRenderer.videoDetails.playerOverlayVideoDetailsRenderer.subtitle.runs[0].text"`
	ChannelID          string         `rjson:"contents.twoColumnWatchNextResults.results.results.contents[1].videoSecondaryInfoRenderer.owner.videoOwnerRenderer.title.runs[0].navigationEndpoint.browseEndpoint.browseId"`
	RawNewChannelID    string         `rjson:"contents.twoColumnWatchNextResults.results.results.contents[1].videoSecondaryInfoRenderer.owner.videoOwnerRenderer.title.runs[0].navigationEndpoint.browseEndpoint.canonicalBaseUrl"`
	Likes              string         `rjson:"contents.twoColumnWatchNextResults.results.results.contents[0].videoPrimaryInfoRenderer.videoActions.menuRenderer.topLevelButtons[0].segmentedLikeDislikeButtonRenderer.likeButton.toggleButtonRenderer.defaultText.simpleText"`
	ChannelSubscribers string         `rjson:"contents.twoColumnWatchNextResults.results.results.contents[1].videoSecondaryInfoRenderer.owner.videoOwnerRenderer.subscriberCountText.simpleText"`
	CommentsCount      string         `rjson:"contents.twoColumnWatchNextResults.results.results.contents[2].itemSectionRenderer.contents[0].commentsEntryPointHeaderRenderer.commentCount.simpleText"`
	Category           string         `rjson:"contents.twoColumnWatchNextResults.results.results.contents[1].videoSecondaryInfoRenderer.metadataRowContainer.metadataRowContainerRenderer.rows[0].richMetadataRowRenderer.contents[1].richMetadataRenderer.title.runs[0].text"`
	OwnerBadges        []string       `rjson:"contents.twoColumnWatchNextResults.results.results.contents[1].videoSecondaryInfoRenderer.owner.videoOwnerRenderer.badges[].metadataBadgeRenderer.tooltip"`
	Badges             []string       `rjson:"contents.twoColumnWatchNextResults.results.results.contents[0].videoPrimaryInfoRenderer.badges[].metadataBadgeRenderer.label"`
	ChannelAvatars     []YoutubeImage `rjson:"contents.twoColumnWatchNextResults.results.results.contents[1].videoSecondaryInfoRenderer.owner.videoOwnerRenderer.thumbnail.thumbnails"`

	Tokens []struct {
		Title string `rjson:"title"`
		Token string `rjson:"serviceEndpoint.continuationCommand.token"`
	} `rjson:"engagementPanels[].engagementPanelSectionListRenderer.header.engagementPanelTitleHeaderRenderer.menu.sortFilterSubMenuRenderer.subMenuItems[0]"`

	SidebarEntries []rawSidebarEntry `rjson:"contents.twoColumnWatchNextResults.secondaryResults.secondaryResults.results"`
	SidebarToken   string            `rjson:"contents.twoColumnWatchNextResults.secondaryResults.secondaryResults.results[-].continuationItemRenderer.button.buttonRenderer.command.continuationCommand.token"`
}

var mediaUrlJsRegex = regexp.MustCompile(`src="(/s/player/[^\\/]+/player_ias[^\\/]+/en_US/base.js)"`)

func NewVideoScraper(id string) (v VideoScraper, err error) {
	rawUrl, err := url.Parse("https://www.youtube.com/watch")
	if err != nil {
		return
	}

	q := rawUrl.Query()
	q.Set("v", id)
	q.Set("hl", "en")
	rawUrl.RawQuery = q.Encode()
	v.url = rawUrl.String()

	resp, err := http.Get(v.url)
	if err != nil {
		return
	}

	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	v.mediaUrlJs = string(mediaUrlJsRegex.FindSubmatch(body)[1])

	var rawJson string
	rawJson, err = ExtractInitialDataBytes(body)
	if err != nil {
		return
	}

	DebugFileOutput([]byte(rawJson), "video_initial.json")

	var output videoInitialOutput
	if err = rjson.Unmarshal([]byte(rawJson), &output); err != nil {
		if errors.Is(err, rjson.ErrCantFindField) {
			if Debug {
				log.Println("WARNING:", err)
			}
			err = nil
		}
		return
	}

	var channelIsVerified bool
	var channelIsVerifiedArtist bool
	for _, badge := range output.OwnerBadges {
		switch badge {
		case ChannelBadgeVerified:
			channelIsVerified = true
		case ChannelBadgeVerifiedArtistChannel:
			channelIsVerifiedArtist = true
		}
	}

	var videoIsUnlisted bool
	for _, badge := range output.Badges {
		switch badge {
		case "Unlisted":
			videoIsUnlisted = true
		}
	}

	for _, token := range output.Tokens {
		switch token.Title {
		case "Top comments":
			v.commentsTopContinueInput = ContinueInput{Continuation: token.Token}.FillGenericInfo()
			v.commentsTopContinueInputJson, err = v.commentsNewestContinueInput.Construct()
			if err != nil {
				return
			}
		case "Newest first":
			v.commentsNewestContinueInput = ContinueInput{Continuation: token.Token}.FillGenericInfo()
			v.commentsNewestContinueInputJson, err = v.commentsTopContinueInput.Construct()
			if err != nil {
				return
			}
		}
	}

	v.sidebarContinueInput = ContinueInput{Continuation: output.SidebarToken}.FillGenericInfo()
	v.sidebarContinueInputJson, err = v.sidebarContinueInput.Construct()
	if err != nil {
		return
	}

	for _, sidebarEntry := range output.SidebarEntries {
		if sidebarEntry.Video.VideoID != "" || sidebarEntry.Playlist.PlaylistID != "" || sidebarEntry.Radio.RadioPlaylistID != "" {
			if entry, err := sidebarEntry.ToSidebarEntry(); err != nil {
				log.Println("WARNING converting to sidebar failed:", err)
			} else {
				v.InitialSidebarEntries = append(v.InitialSidebarEntries, entry)
			}
		}
	}

	dateText, premiered := strings.CutPrefix(output.Date, "Premiered ")
	dateText, wasLive := strings.CutPrefix(dateText, "Streamed live on ")
	dateText, isLive := strings.CutPrefix(dateText, "Started streaming on ")

	var date time.Time
	date, err = time.Parse(YoutubeVideoDateLayout, dateText)
	if err != nil {
		return
	}

	views, err := strconv.Atoi(strings.ReplaceAll(strings.TrimSuffix(output.Views, " views"), ",", ""))
	if err != nil {
		return
	}

	likes, unit, err := humanize.ParseSI(FixUnit(output.Likes))
	if err != nil {
		return
	} else if unit != "" {
		log.Printf("WARNING: possibly wrong number for likes: %f%s\n", likes, unit)
	}

	var comments float64
	if output.CommentsCount != "" {
		comments, unit, err = humanize.ParseSI(FixUnit(output.CommentsCount))
		if err != nil {
			return
		} else if unit != "" {
			log.Printf("WARNING: possibly wrong number for comments count: %f%s\n", comments, unit)
		}
	}

	channelSubscribers, unit, err := humanize.ParseSI(FixUnit(strings.TrimSuffix(output.ChannelSubscribers, " subscribers")))
	if err != nil {
		return
	} else if unit != "" {
		log.Printf("WARNING: possibly wrong number for channel subscribers count: %f%s\n", channelSubscribers, unit)
	}

	v.VideoInfo = FullVideo{
		VideoID:                 id,
		Title:                   output.Title,
		Description:             output.Description,
		Views:                   views,
		IsLive:                  output.IsLive || isLive,
		WasLive:                 wasLive,
		Date:                    date,
		Likes:                   int(likes),
		CommentsCount:           int(comments),
		Category:                output.Category,
		Username:                output.Username,
		ChannelID:               output.ChannelID,
		NewChannelID:            strings.TrimPrefix(output.RawNewChannelID, "/"),
		ChannelSubscribers:      int(channelSubscribers),
		ChannelIsVerified:       channelIsVerified,
		ChannelIsVerifiedArtist: channelIsVerifiedArtist,
		IsUnlisted:              videoIsUnlisted,
		VideoPremiered:          premiered,
		ChannelAvatars:          output.ChannelAvatars,
	}

	return
}

type MediaFormat struct {
	Bitrate int `rjson:"bitrate"`
	Width   int `rjson:"width"`
	Height  int `rjson:"height"`

	Url             string `rjson:"url"`
	MimeType        string `rjson:"mimeType"` // e.g "audio/mp4; codecs=\"mp4a.40.2\"" or "video/mp4; codecs=\"av01.0.00M.08\""
	QualityLabel    string `rjson:"qualityLabel"`
	SignatureCipher string `rjson:"signatureCipher"` // DRM

	// Videos can have dubs
	AudioTrack struct {
		DisplayName    string `rjson:"displayName"`
		AudioIsDefault bool   `rjson:"audioIsDefault"`
	} `rjson:"audioTrack"`
}

// GetMediaUrl is a generic function to get the media url, doesnt matter if it has DRM or not
func (m *MediaFormat) GetMediaUrl(v *VideoScraper) (out string, err error) {
	if m.Url == "" && m.SignatureCipher != "" {
		var q url.Values
		q, err = url.ParseQuery(m.SignatureCipher)
		if err != nil {
			return
		}

		m.Url, err = v.decryptSignature(q)
		if err != nil {
			return
		}
	}

	out = m.Url
	return
}

type ExtractMediaOutput struct {
	Formats []MediaFormat `rjson:"streamingData.formats"`

	// Best quality formats are here, but the video and audio tracks will be separate
	AdaptiveFormats []MediaFormat `rjson:"streamingData.adaptiveFormats"`
}

func ExtractMediaFormats(id string) (output ExtractMediaOutput, err error) {
	var bs []byte
	bs, err = ContinueInput{VideoID: id}.FillGenericInfo().Construct()
	if err != nil {
		return
	}

	var resp *http.Response
	// using the web key
	resp, err = http.Post("https://youtubei.googleapis.com/youtubei/v1/player?key="+webKey, "application/json", bytes.NewReader(bs))
	if err != nil {
		return
	}

	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = rjson.Unmarshal(body, &output)

	return
}
func (v *VideoScraper) ExtractMediaFormats() (output ExtractMediaOutput, err error) {
	return ExtractMediaFormats(v.VideoInfo.VideoID)
}

var funcNameRegex = regexp.MustCompile("\n([^=]+)=function\\(\\w\\){\\w=\\w\\.split\\(\"\"\\);[^. ]+\\.[^( ]+")

type operationFunc = func(a string, b int) string

type decryptFunc struct {
	f operationFunc
	i int
}

// TODO: fix & cleanup
// TODO: add caching
func FetchDecryptFunction(mediaUrlJs string) (decryptFunctions []decryptFunc, err error) {
	var resp *http.Response
	resp, err = http.Get("https://www.youtube.com" + mediaUrlJs)
	if err != nil {
		return
	}

	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	name := funcNameRegex.FindSubmatch(body)[1]

	var re *regexp.Regexp
	re, err = regexp.Compile(fmt.Sprintf("\n%s=function\\(\\w\\){([^}]+)}", name))
	if err != nil {
		return
	}

	rawFuncBody := string(re.FindSubmatch(body)[1])

	t := strings.Split(rawFuncBody, ";")
	funcBody := t[1 : len(t)-1]

	varName := funcBody[0][0:2]
	re, err = regexp.Compile(fmt.Sprintf("var %s={([a-zA-Z:;%s(){},\\n0-9. =\\[\\]]*)};", varName, "%"))
	if err != nil {
		return
	}

	varBody := string(re.FindSubmatch(body)[1])
	operations := make(map[string]operationFunc)
	for _, row := range strings.Split(varBody, "},") {
		cleanedRow := strings.Trim(row, "\n")
		opName := regexp.MustCompile("^[^:]+").FindString(cleanedRow)
		opBody := regexp.MustCompile("\\{[^}]+").FindString(cleanedRow)

		switch opBody {
		case "{a.reverse()":
			operations[opName] = func(a string, b int) string {
				return reverse(a)
			}
		case "{a.splice(0,b)":
			operations[opName] = func(a string, b int) string {
				return splice(a, b)
			}
		default:
			operations[opName] = func(a string, b int) string {
				raw := []rune(a)
				c := raw[0]
				raw[0] = raw[b%len(a)]
				raw[b%len(a)] = c
				return string(raw)
			}
		}
	}

	for _, f := range funcBody {
		f = strings.TrimPrefix(f, varName)
		opName := strings.TrimPrefix(regexp.MustCompile("[^\\(]+").FindString(f), ".")

		var i int
		if _, err = fmt.Sscanf(regexp.MustCompile("\\(\\w,([\\d]+)\\)").FindString(f), "(a,%d)", &i); err != nil {
			return
		}

		decryptFunctions = append(decryptFunctions, decryptFunc{
			f: operations[opName],
			i: i,
		})
	}

	return
}

func (v *VideoScraper) decryptSignature(query url.Values) (out string, err error) {
	rawUrl := query.Get("url")
	signatureUrlName := query.Get("sp")

	var funcs []decryptFunc
	funcs, err = FetchDecryptFunction(v.mediaUrlJs)
	if err != nil {
		return
	}

	sig := query.Get("s")
	for _, f := range funcs {
		sig = f.f(sig, f.i)
	}

	var parsedUrl *url.URL
	parsedUrl, err = url.Parse(rawUrl)
	if err != nil {
		return
	}
	q := parsedUrl.Query()
	q.Set(signatureUrlName, sig)
	parsedUrl.RawQuery = q.Encode()

	out = parsedUrl.String()
	return
}
