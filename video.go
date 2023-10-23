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
	"strings"

	"github.com/ayes-web/rjson"
)

type VideoScraper struct {
	VideoInfo             FullVideo
	InitialSidebarEntries []SidebarEntry

	mediaUrlJs string
	url        string

	commentsNewestPassedInitial     bool
	commentsNewestToken             string
	commentsNewestContinueInputJson []byte

	commentsTopPassedInitial     bool
	commentsTopToken             string
	commentsTopContinueInputJson []byte

	sidebarToken             string
	sidebarContinueInputJson []byte
}

// FullVideo has the full metadata unlike Video which is fetched from Video lists
type FullVideo struct {
	VideoID       string `json:"VideoID"`
	Title         string `json:"Title"`
	Description   string `json:"Description"`
	Views         string `json:"Views"` // if its live this will display number of viewers instead
	IsLive        bool   `json:"IsLive"`
	WasLive       bool   `json:"WasLive"` // if this Video was live
	Date          string `json:"Date"`    // Date will be in this format: "Jul 12, 2023"
	Likes         string `json:"Likes"`
	CommentsCount string `json:"CommentsCount"`
	Category      string `json:"Category"`

	Username           string `json:"Username"`
	ChannelID          string `json:"ChannelID"`
	NewChannelID       string `json:"NewChannelID"`
	ChannelSubscribers string `json:"ChannelSubscribers"`
}

type videoInitialOutput struct {
	Title              string `rjson:"playerOverlays.playerOverlayRenderer.videoDetails.playerOverlayVideoDetailsRenderer.title.simpleText"`
	Description        string `rjson:"contents.twoColumnWatchNextResults.results.results.contents[1].videoSecondaryInfoRenderer.attributedDescription.content"`
	Views              string `rjson:"playerOverlays.playerOverlayRenderer.videoDetails.playerOverlayVideoDetailsRenderer.subtitle.runs[2].text"`
	IsLive             bool   `rjson:"contents.twoColumnWatchNextResults.results.results.contents[0].videoPrimaryInfoRenderer.viewCount.videoViewCountRenderer.isLive"`
	Date               string `rjson:"contents.twoColumnWatchNextResults.results.results.contents[0].videoPrimaryInfoRenderer.dateText.simpleText"`
	Username           string `rjson:"playerOverlays.playerOverlayRenderer.videoDetails.playerOverlayVideoDetailsRenderer.subtitle.runs[0].text"`
	ChannelID          string `rjson:"contents.twoColumnWatchNextResults.results.results.contents[1].videoSecondaryInfoRenderer.owner.videoOwnerRenderer.title.runs[0].navigationEndpoint.browseEndpoint.browseId"`
	RawNewChannelID    string `rjson:"contents.twoColumnWatchNextResults.results.results.contents[1].videoSecondaryInfoRenderer.owner.videoOwnerRenderer.title.runs[0].navigationEndpoint.browseEndpoint.canonicalBaseUrl"`
	Likes              string `rjson:"contents.twoColumnWatchNextResults.results.results.contents[0].videoPrimaryInfoRenderer.videoActions.menuRenderer.topLevelButtons[0].segmentedLikeDislikeButtonRenderer.likeButton.toggleButtonRenderer.defaultText.simpleText"`
	ChannelSubscribers string `rjson:"contents.twoColumnWatchNextResults.results.results.contents[1].videoSecondaryInfoRenderer.owner.videoOwnerRenderer.subscriberCountText.simpleText"`
	CommentsCount      string `rjson:"contents.twoColumnWatchNextResults.results.results.contents[2].itemSectionRenderer.contents[0].commentsEntryPointHeaderRenderer.commentCount.simpleText"`
	Category           string `rjson:"contents.twoColumnWatchNextResults.results.results.contents[1].videoSecondaryInfoRenderer.metadataRowContainer.metadataRowContainerRenderer.rows[0].richMetadataRowRenderer.contents[1].richMetadataRenderer.title.runs[0].text"`

	Tokens []struct {
		Title string `rjson:"title"`
		Token string `rjson:"serviceEndpoint.continuationCommand.token"`
	} `rjson:"engagementPanels[].engagementPanelSectionListRenderer.header.engagementPanelTitleHeaderRenderer.menu.sortFilterSubMenuRenderer.subMenuItems[0]"`

	SidebarEntries []rawSidebarEntry `rjson:"contents.twoColumnWatchNextResults.secondaryResults.secondaryResults.results"`
	SidebarToken   string            `rjson:"contents.twoColumnWatchNextResults.secondaryResults.secondaryResults.results[-].continuationItemRenderer.button.buttonRenderer.command.continuationCommand.token"`
}

var mediaUrlJsRegex = regexp.MustCompile(`src="(/s/player/[^\\/]+/player_ias[^\\/]+/en_US/base.js)"`)

func NewVideoScraper(id string) (v VideoScraper, err error) {
	v.url = fmt.Sprintf("https://www.youtube.com/watch?v=%s&hl=en", id)

	resp, err := http.Get(v.url)
	if err != nil {
		return
	}

	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	debugFileOutput([]byte(body), "video_initial.html")

	v.mediaUrlJs = string(mediaUrlJsRegex.FindSubmatch(body)[1])

	var rawJson string
	rawJson, err = extractInitialDataBytes(body)
	if err != nil {
		return
	}

	debugFileOutput([]byte(rawJson), "video_initial.json")

	var output videoInitialOutput
	if err = rjson.Unmarshal([]byte(rawJson), &output); err != nil {
		if errors.Unwrap(err) == rjson.ErrCantFindField {
			if Debug {
				log.Println("WARNING:", err)
			}
			err = nil
		}
		return
	}

	for _, token := range output.Tokens {
		switch token.Title {
		case "Top comments":
			v.commentsTopToken = token.Token
		case "Newest first":
			v.commentsNewestToken = token.Token
		}
	}

	v.sidebarToken = output.SidebarToken

	v.commentsNewestContinueInputJson, err = continueInput{Continuation: v.commentsNewestToken}.FillGenericInfo().Construct()
	if err != nil {
		return
	}
	v.commentsTopContinueInputJson, err = continueInput{Continuation: v.commentsTopToken}.FillGenericInfo().Construct()
	if err != nil {
		return
	}
	v.sidebarContinueInputJson, err = continueInput{Continuation: v.sidebarToken}.FillGenericInfo().Construct()
	if err != nil {
		return
	}

	for _, sidebarEntry := range output.SidebarEntries {
		if sidebarEntry.VideoID == "" && sidebarEntry.PlaylistID == "" && sidebarEntry.RadioPlaylistID == "" {
			continue
		}

		v.InitialSidebarEntries = append(v.InitialSidebarEntries, sidebarEntry.ToSidebarEntry())
	}

	date, wasLive := strings.CutPrefix(output.Date, "Streamed live on ")
	date, isLive := strings.CutPrefix(date, "Started streaming on ")

	v.VideoInfo = FullVideo{
		VideoID:            id,
		Title:              output.Title,
		Description:        output.Description,
		Views:              strings.TrimSuffix(output.Views, " views"),
		IsLive:             output.IsLive || isLive,
		WasLive:            wasLive,
		Date:               date,
		Likes:              output.Likes,
		CommentsCount:      output.CommentsCount,
		Category:           output.Category,
		Username:           output.Username,
		ChannelID:          output.ChannelID,
		NewChannelID:       strings.TrimPrefix(output.RawNewChannelID, "/"),
		ChannelSubscribers: strings.TrimSuffix(output.ChannelSubscribers, " subscribers"),
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
	bs, err = continueInput{VideoID: id}.FillGenericInfo().Construct()
	if err != nil {
		return
	}

	var resp *http.Response
	// using the web key
	resp, err = http.Post("https://youtubei.googleapis.com/youtubei/v1/player?key=AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8", "application/json", bytes.NewReader(bs))
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
