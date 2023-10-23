package scraper

// Youtube api input json
type ContinueInput struct {
	Context struct {
		Client struct {
			Hl            string `json:"hl"`           // language you want the data in, for english "en"
			Gl            string `json:"gl,omitempty"` // data region
			VisitorData   string `json:"visitorData,omitempty"`
			ClientName    string `json:"clientName"`
			ClientVersion string `json:"clientVersion"`
		} `json:"client"`
	} `json:"context"`
	VideoID             string `json:"videoId,omitempty"`
	Continuation        string `json:"continuation"`
	BrowseId            string `json:"browseId,omitempty"`
	InlineSettingStatus string `json:"inlineSettingStatus,omitempty"`
}
