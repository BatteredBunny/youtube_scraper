package scraper

type continueInput struct {
	Context struct {
		Client struct {
			Hl string `json:"hl"`           // language you want the data in, for english "en"
			Gl string `json:"gl,omitempty"` // data region
			//RemoteHost string `json:"remoteHost"`
			//DeviceMake    string `json:"deviceMake"`
			//DeviceModel   string `json:"deviceModel"`
			VisitorData string `json:"visitorData,omitempty"`
			//UserAgent     string `json:"userAgent"`
			ClientName    string `json:"clientName"`
			ClientVersion string `json:"clientVersion"`
			//OsName        string `json:"osName"`
			//OsVersion     string `json:"osVersion"`
			//OriginalUrl   string `json:"originalUrl"`
			//ScreenPixelDensity int    `json:"screenPixelDensity"`
			//Platform           string `json:"platform"`
			//ClientFormFactor   string `json:"clientFormFactor"`
			//ConfigInfo         struct {
			//	AppInstallData string `json:"appInstallData"`
			//} `json:"configInfo"`
			//ScreenDensityFloat int    `json:"screenDensityFloat"`
			//UserInterfaceTheme string `json:"userInterfaceTheme"`
			//TimeZone           string `json:"timeZone"`
			//BrowserName        string `json:"browserName"`
			//BrowserVersion     string `json:"browserVersion"`
			//AcceptHeader       string `json:"acceptHeader"`
			//DeviceExperimentId string `json:"deviceExperimentId"`
			//ScreenWidthPoints  int    `json:"screenWidthPoints"`
			//ScreenHeightPoints int    `json:"screenHeightPoints"`
			//UtcOffsetMinutes   int    `json:"utcOffsetMinutes"`
			//MainAppWebInfo     struct {
			//	GraftUrl                  string `json:"graftUrl"`
			//	PwaInstallabilityStatus   string `json:"pwaInstallabilityStatus"`
			//	WebDisplayMode            string `json:"webDisplayMode"`
			//	IsWebNativeShareAvailable bool   `json:"isWebNativeShareAvailable"`
			//} `json:"mainAppWebInfo"`
		} `json:"client"`
		//User struct {
		//	LockedSafetyMode bool `json:"lockedSafetyMode"`
		//} `json:"user"`
		//Request struct {
		//	UseSsl bool `json:"useSsl"`
		//	InternalExperimentFlags []interface{} `json:"internalExperimentFlags"`
		//	ConsistencyTokenJars    []interface{} `json:"consistencyTokenJars"`
		//} `json:"request"`
		//ClickTracking struct {
		//	ClickTrackingParams string `json:"clickTrackingParams"`
		//} `json:"clickTracking"`
		//AdSignalsInfo struct {
		//	Params []struct {
		//		Key   string `json:"key"`
		//		Value string `json:"value"`
		//	} `json:"params"`
		//} `json:"adSignalsInfo"`
	} `json:"context"`
	VideoID             string `json:"videoId"`
	Continuation        string `json:"continuation"`
	BrowseId            string `json:"browseId,omitempty"`
	InlineSettingStatus string `json:"inlineSettingStatus,omitempty"`
}
