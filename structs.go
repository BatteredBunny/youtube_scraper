package youtube_scraper

type initialData struct {
	ResponseContext struct {
		ServiceTrackingParams []struct {
			Service string `json:"service"`
			Params  []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"params"`
		} `json:"serviceTrackingParams"`
		MaxAgeSeconds             int `json:"maxAgeSeconds,omitempty"`
		MainAppWebResponseContext struct {
			LoggedOut     bool   `json:"loggedOut"`
			TrackingParam string `json:"trackingParam"`
		} `json:"mainAppWebResponseContext"`
		WebResponseContextExtensionData struct {
			YtConfigData struct {
				VisitorData           string `json:"visitorData"`
				RootVisualElementType int    `json:"rootVisualElementType"`
			} `json:"ytConfigData"`
			HasDecorated    bool `json:"hasDecorated"`
			WebPrefetchData struct {
				NavigationEndpoints []struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					CommandMetadata     struct {
						WebCommandMetadata struct {
							Url         string `json:"url"`
							WebPageType string `json:"webPageType"`
							RootVe      int    `json:"rootVe"`
						} `json:"webCommandMetadata"`
					} `json:"commandMetadata"`
					WatchEndpoint struct {
						VideoId                              string `json:"videoId"`
						Params                               string `json:"params"`
						PlayerParams                         string `json:"playerParams"`
						WatchEndpointSupportedPrefetchConfig struct {
							PrefetchHintConfig struct {
								PrefetchPriority                            int `json:"prefetchPriority"`
								CountdownUiRelativeSecondsPrefetchCondition int `json:"countdownUiRelativeSecondsPrefetchCondition"`
							} `json:"prefetchHintConfig"`
						} `json:"watchEndpointSupportedPrefetchConfig"`
					} `json:"watchEndpoint"`
				} `json:"navigationEndpoints"`
			} `json:"webPrefetchData,omitempty"`
		} `json:"webResponseContextExtensionData"`
	} `json:"responseContext"`
	Contents struct {
		TwoColumnBrowseResultsRenderer struct {
			Tabs []struct {
				TabRenderer struct {
					Endpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								Url         string `json:"url"`
								WebPageType string `json:"webPageType"`
								RootVe      int    `json:"rootVe"`
								ApiUrl      string `json:"apiUrl"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						BrowseEndpoint struct {
							BrowseId         string `json:"browseId"`
							Params           string `json:"params"`
							CanonicalBaseUrl string `json:"canonicalBaseUrl"`
						} `json:"browseEndpoint"`
					} `json:"endpoint,omitempty"`
					Title          string `json:"title,omitempty"`
					TrackingParams string `json:"trackingParams"`
					Selected       bool   `json:"selected,omitempty"`
					Content        struct {
						RichGridRenderer struct {
							Contents []struct {
								RichItemRenderer struct {
									Content struct {
										VideoRenderer struct {
											VideoId   string `json:"videoId"`
											Thumbnail struct {
												Thumbnails []struct {
													Url    string `json:"url"`
													Width  int    `json:"width"`
													Height int    `json:"height"`
												} `json:"thumbnails"`
											} `json:"thumbnail"`
											Title struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
												Accessibility struct {
													AccessibilityData struct {
														Label string `json:"label"`
													} `json:"accessibilityData"`
												} `json:"accessibility"`
											} `json:"title"`
											DescriptionSnippet struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"descriptionSnippet"`
											PublishedTimeText struct {
												SimpleText string `json:"simpleText"`
											} `json:"publishedTimeText,omitempty"`
											LengthText struct {
												Accessibility struct {
													AccessibilityData struct {
														Label string `json:"label"`
													} `json:"accessibilityData"`
												} `json:"accessibility"`
												SimpleText string `json:"simpleText"`
											} `json:"lengthText,omitempty"`
											ViewCountText struct {
												SimpleText string `json:"simpleText,omitempty"`
												Runs       []struct {
													Text string `json:"text"`
												} `json:"runs,omitempty"`
											} `json:"viewCountText"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														Url         string `json:"url"`
														WebPageType string `json:"webPageType"`
														RootVe      int    `json:"rootVe"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												WatchEndpoint struct {
													VideoId                            string `json:"videoId"`
													WatchEndpointSupportedOnesieConfig struct {
														Html5PlaybackOnesieConfig struct {
															CommonConfig struct {
																Url string `json:"url"`
															} `json:"commonConfig"`
														} `json:"html5PlaybackOnesieConfig"`
													} `json:"watchEndpointSupportedOnesieConfig"`
												} `json:"watchEndpoint"`
											} `json:"navigationEndpoint"`
											OwnerBadges []struct {
												MetadataBadgeRenderer struct {
													Icon struct {
														IconType string `json:"iconType"`
													} `json:"icon"`
													Style             string `json:"style"`
													Tooltip           string `json:"tooltip"`
													TrackingParams    string `json:"trackingParams"`
													AccessibilityData struct {
														Label string `json:"label"`
													} `json:"accessibilityData"`
												} `json:"metadataBadgeRenderer"`
											} `json:"ownerBadges,omitempty"`
											TrackingParams     string `json:"trackingParams"`
											ShowActionMenu     bool   `json:"showActionMenu"`
											ShortViewCountText struct {
												Accessibility struct {
													AccessibilityData struct {
														Label string `json:"label"`
													} `json:"accessibilityData"`
												} `json:"accessibility,omitempty"`
												SimpleText string `json:"simpleText,omitempty"`
												Runs       []struct {
													Text string `json:"text"`
												} `json:"runs,omitempty"`
											} `json:"shortViewCountText"`
											Menu struct {
												MenuRenderer struct {
													Items []struct {
														MenuServiceItemRenderer struct {
															Text struct {
																Runs []struct {
																	Text string `json:"text"`
																} `json:"runs"`
															} `json:"text"`
															Icon struct {
																IconType string `json:"iconType"`
															} `json:"icon"`
															ServiceEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		SendPost bool   `json:"sendPost"`
																		ApiUrl   string `json:"apiUrl,omitempty"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																SignalServiceEndpoint struct {
																	Signal  string `json:"signal"`
																	Actions []struct {
																		ClickTrackingParams  string `json:"clickTrackingParams"`
																		AddToPlaylistCommand struct {
																			OpenMiniplayer      bool   `json:"openMiniplayer"`
																			VideoId             string `json:"videoId"`
																			ListType            string `json:"listType"`
																			OnCreateListCommand struct {
																				ClickTrackingParams string `json:"clickTrackingParams"`
																				CommandMetadata     struct {
																					WebCommandMetadata struct {
																						SendPost bool   `json:"sendPost"`
																						ApiUrl   string `json:"apiUrl"`
																					} `json:"webCommandMetadata"`
																				} `json:"commandMetadata"`
																				CreatePlaylistServiceEndpoint struct {
																					VideoIds []string `json:"videoIds"`
																					Params   string   `json:"params"`
																				} `json:"createPlaylistServiceEndpoint"`
																			} `json:"onCreateListCommand"`
																			VideoIds []string `json:"videoIds"`
																		} `json:"addToPlaylistCommand"`
																	} `json:"actions"`
																} `json:"signalServiceEndpoint,omitempty"`
																ShareEntityServiceEndpoint struct {
																	SerializedShareEntity string `json:"serializedShareEntity"`
																	Commands              []struct {
																		ClickTrackingParams string `json:"clickTrackingParams"`
																		OpenPopupAction     struct {
																			Popup struct {
																				UnifiedSharePanelRenderer struct {
																					TrackingParams     string `json:"trackingParams"`
																					ShowLoadingSpinner bool   `json:"showLoadingSpinner"`
																				} `json:"unifiedSharePanelRenderer"`
																			} `json:"popup"`
																			PopupType string `json:"popupType"`
																			BeReused  bool   `json:"beReused"`
																		} `json:"openPopupAction"`
																	} `json:"commands"`
																} `json:"shareEntityServiceEndpoint,omitempty"`
															} `json:"serviceEndpoint"`
															TrackingParams string `json:"trackingParams"`
															HasSeparator   bool   `json:"hasSeparator,omitempty"`
														} `json:"menuServiceItemRenderer"`
													} `json:"items"`
													TrackingParams string `json:"trackingParams"`
													Accessibility  struct {
														AccessibilityData struct {
															Label string `json:"label"`
														} `json:"accessibilityData"`
													} `json:"accessibility"`
													TargetId string `json:"targetId,omitempty"`
												} `json:"menuRenderer"`
											} `json:"menu"`
											ThumbnailOverlays []struct {
												ThumbnailOverlayTimeStatusRenderer struct {
													Text struct {
														Accessibility struct {
															AccessibilityData struct {
																Label string `json:"label"`
															} `json:"accessibilityData"`
														} `json:"accessibility"`
														SimpleText string `json:"simpleText"`
													} `json:"text"`
													Style string `json:"style"`
												} `json:"thumbnailOverlayTimeStatusRenderer,omitempty"`
												ThumbnailOverlayToggleButtonRenderer struct {
													IsToggled     bool `json:"isToggled,omitempty"`
													UntoggledIcon struct {
														IconType string `json:"iconType"`
													} `json:"untoggledIcon"`
													ToggledIcon struct {
														IconType string `json:"iconType"`
													} `json:"toggledIcon"`
													UntoggledTooltip         string `json:"untoggledTooltip"`
													ToggledTooltip           string `json:"toggledTooltip"`
													UntoggledServiceEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																SendPost bool   `json:"sendPost"`
																ApiUrl   string `json:"apiUrl,omitempty"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														PlaylistEditEndpoint struct {
															PlaylistId string `json:"playlistId"`
															Actions    []struct {
																AddedVideoId string `json:"addedVideoId"`
																Action       string `json:"action"`
															} `json:"actions"`
														} `json:"playlistEditEndpoint,omitempty"`
														SignalServiceEndpoint struct {
															Signal  string `json:"signal"`
															Actions []struct {
																ClickTrackingParams  string `json:"clickTrackingParams"`
																AddToPlaylistCommand struct {
																	OpenMiniplayer      bool   `json:"openMiniplayer"`
																	VideoId             string `json:"videoId"`
																	ListType            string `json:"listType"`
																	OnCreateListCommand struct {
																		ClickTrackingParams string `json:"clickTrackingParams"`
																		CommandMetadata     struct {
																			WebCommandMetadata struct {
																				SendPost bool   `json:"sendPost"`
																				ApiUrl   string `json:"apiUrl"`
																			} `json:"webCommandMetadata"`
																		} `json:"commandMetadata"`
																		CreatePlaylistServiceEndpoint struct {
																			VideoIds []string `json:"videoIds"`
																			Params   string   `json:"params"`
																		} `json:"createPlaylistServiceEndpoint"`
																	} `json:"onCreateListCommand"`
																	VideoIds []string `json:"videoIds"`
																} `json:"addToPlaylistCommand"`
															} `json:"actions"`
														} `json:"signalServiceEndpoint,omitempty"`
													} `json:"untoggledServiceEndpoint"`
													ToggledServiceEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																SendPost bool   `json:"sendPost"`
																ApiUrl   string `json:"apiUrl"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														PlaylistEditEndpoint struct {
															PlaylistId string `json:"playlistId"`
															Actions    []struct {
																Action         string `json:"action"`
																RemovedVideoId string `json:"removedVideoId"`
															} `json:"actions"`
														} `json:"playlistEditEndpoint"`
													} `json:"toggledServiceEndpoint,omitempty"`
													UntoggledAccessibility struct {
														AccessibilityData struct {
															Label string `json:"label"`
														} `json:"accessibilityData"`
													} `json:"untoggledAccessibility"`
													ToggledAccessibility struct {
														AccessibilityData struct {
															Label string `json:"label"`
														} `json:"accessibilityData"`
													} `json:"toggledAccessibility"`
													TrackingParams string `json:"trackingParams"`
												} `json:"thumbnailOverlayToggleButtonRenderer,omitempty"`
												ThumbnailOverlayNowPlayingRenderer struct {
													Text struct {
														Runs []struct {
															Text string `json:"text"`
														} `json:"runs"`
													} `json:"text"`
												} `json:"thumbnailOverlayNowPlayingRenderer,omitempty"`
												ThumbnailOverlayLoadingPreviewRenderer struct {
													Text struct {
														Runs []struct {
															Text string `json:"text"`
														} `json:"runs"`
													} `json:"text"`
												} `json:"thumbnailOverlayLoadingPreviewRenderer,omitempty"`
												ThumbnailOverlayInlineUnplayableRenderer struct {
													Text struct {
														Runs []struct {
															Text string `json:"text"`
														} `json:"runs"`
													} `json:"text"`
													Icon struct {
														IconType string `json:"iconType"`
													} `json:"icon"`
												} `json:"thumbnailOverlayInlineUnplayableRenderer,omitempty"`
											} `json:"thumbnailOverlays"`
											LongBylineText struct {
												Runs []struct {
													Text               string `json:"text"`
													NavigationEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																Url         string `json:"url"`
																WebPageType string `json:"webPageType"`
																RootVe      int    `json:"rootVe"`
																ApiUrl      string `json:"apiUrl"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														BrowseEndpoint struct {
															BrowseId         string `json:"browseId"`
															CanonicalBaseUrl string `json:"canonicalBaseUrl"`
														} `json:"browseEndpoint"`
													} `json:"navigationEndpoint"`
												} `json:"runs"`
											} `json:"longBylineText,omitempty"`
											Badges []struct {
												MetadataBadgeRenderer struct {
													Icon struct {
														IconType string `json:"iconType"`
													} `json:"icon"`
													Style          string `json:"style"`
													Label          string `json:"label"`
													TrackingParams string `json:"trackingParams"`
												} `json:"metadataBadgeRenderer"`
											} `json:"badges,omitempty"`
											OwnerText struct {
												Runs []struct {
													Text               string `json:"text"`
													NavigationEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																Url         string `json:"url"`
																WebPageType string `json:"webPageType"`
																RootVe      int    `json:"rootVe"`
																ApiUrl      string `json:"apiUrl"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														BrowseEndpoint struct {
															BrowseId         string `json:"browseId"`
															CanonicalBaseUrl string `json:"canonicalBaseUrl"`
														} `json:"browseEndpoint"`
													} `json:"navigationEndpoint"`
												} `json:"runs"`
											} `json:"ownerText,omitempty"`
											ShortBylineText struct {
												Runs []struct {
													Text               string `json:"text"`
													NavigationEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																Url         string `json:"url"`
																WebPageType string `json:"webPageType"`
																RootVe      int    `json:"rootVe"`
																ApiUrl      string `json:"apiUrl"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														BrowseEndpoint struct {
															BrowseId         string `json:"browseId"`
															CanonicalBaseUrl string `json:"canonicalBaseUrl"`
														} `json:"browseEndpoint"`
													} `json:"navigationEndpoint"`
												} `json:"runs"`
											} `json:"shortBylineText,omitempty"`
											ChannelThumbnailSupportedRenderers struct {
												ChannelThumbnailWithLinkRenderer struct {
													Thumbnail struct {
														Thumbnails []struct {
															Url    string `json:"url"`
															Width  int    `json:"width"`
															Height int    `json:"height"`
														} `json:"thumbnails"`
													} `json:"thumbnail"`
													NavigationEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																Url         string `json:"url"`
																WebPageType string `json:"webPageType"`
																RootVe      int    `json:"rootVe"`
																ApiUrl      string `json:"apiUrl"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														BrowseEndpoint struct {
															BrowseId         string `json:"browseId"`
															CanonicalBaseUrl string `json:"canonicalBaseUrl"`
														} `json:"browseEndpoint"`
													} `json:"navigationEndpoint"`
													Accessibility struct {
														AccessibilityData struct {
															Label string `json:"label"`
														} `json:"accessibilityData"`
													} `json:"accessibility"`
												} `json:"channelThumbnailWithLinkRenderer"`
											} `json:"channelThumbnailSupportedRenderers,omitempty"`
											InlinePlaybackEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														Url         string `json:"url"`
														WebPageType string `json:"webPageType"`
														RootVe      int    `json:"rootVe"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												WatchEndpoint struct {
													VideoId              string `json:"videoId"`
													PlayerParams         string `json:"playerParams"`
													PlayerExtraUrlParams []struct {
														Key   string `json:"key"`
														Value string `json:"value"`
													} `json:"playerExtraUrlParams"`
													WatchEndpointSupportedOnesieConfig struct {
														Html5PlaybackOnesieConfig struct {
															CommonConfig struct {
																Url string `json:"url"`
															} `json:"commonConfig"`
														} `json:"html5PlaybackOnesieConfig"`
													} `json:"watchEndpointSupportedOnesieConfig"`
												} `json:"watchEndpoint"`
											} `json:"inlinePlaybackEndpoint,omitempty"`
										} `json:"videoRenderer"`
									} `json:"content"`
									TrackingParams string `json:"trackingParams"`
								} `json:"richItemRenderer,omitempty"`
								ContinuationItemRenderer struct {
									Trigger              string `json:"trigger"`
									ContinuationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												SendPost bool   `json:"sendPost"`
												ApiUrl   string `json:"apiUrl"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										ContinuationCommand struct {
											Token   string `json:"token"`
											Request string `json:"request"`
										} `json:"continuationCommand"`
									} `json:"continuationEndpoint"`
									GhostCards struct {
										GhostGridRenderer struct {
											Rows int `json:"rows"`
										} `json:"ghostGridRenderer"`
									} `json:"ghostCards,omitempty"`
								} `json:"continuationItemRenderer,omitempty"`
								RichSectionRenderer struct {
									Content struct {
										RichShelfRenderer struct {
											Title struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"title"`
											Contents []struct {
												RichItemRenderer struct {
													Content struct {
														ReelItemRenderer struct {
															VideoId  string `json:"videoId"`
															Headline struct {
																SimpleText string `json:"simpleText"`
															} `json:"headline"`
															Thumbnail struct {
																Thumbnails []struct {
																	Url    string `json:"url"`
																	Width  int    `json:"width"`
																	Height int    `json:"height"`
																} `json:"thumbnails"`
																IsOriginalAspectRatio bool `json:"isOriginalAspectRatio"`
															} `json:"thumbnail"`
															ViewCountText struct {
																Accessibility struct {
																	AccessibilityData struct {
																		Label string `json:"label"`
																	} `json:"accessibilityData"`
																} `json:"accessibility"`
																SimpleText string `json:"simpleText"`
															} `json:"viewCountText"`
															NavigationEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		Url         string `json:"url"`
																		WebPageType string `json:"webPageType"`
																		RootVe      int    `json:"rootVe"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																ReelWatchEndpoint struct {
																	VideoId      string `json:"videoId"`
																	PlayerParams string `json:"playerParams"`
																	Thumbnail    struct {
																		Thumbnails []struct {
																			Url    string `json:"url"`
																			Width  int    `json:"width"`
																			Height int    `json:"height"`
																		} `json:"thumbnails"`
																		IsOriginalAspectRatio bool `json:"isOriginalAspectRatio"`
																	} `json:"thumbnail"`
																	Overlay struct {
																		ReelPlayerOverlayRenderer struct {
																			Style                     string `json:"style"`
																			TrackingParams            string `json:"trackingParams"`
																			ReelPlayerNavigationModel string `json:"reelPlayerNavigationModel"`
																		} `json:"reelPlayerOverlayRenderer"`
																	} `json:"overlay"`
																	Params           string `json:"params"`
																	SequenceProvider string `json:"sequenceProvider"`
																	SequenceParams   string `json:"sequenceParams"`
																	LoggingContext   struct {
																		VssLoggingContext struct {
																			SerializedContextData string `json:"serializedContextData"`
																		} `json:"vssLoggingContext"`
																		QoeLoggingContext struct {
																			SerializedContextData string `json:"serializedContextData"`
																		} `json:"qoeLoggingContext"`
																	} `json:"loggingContext"`
																	UstreamerConfig string `json:"ustreamerConfig"`
																} `json:"reelWatchEndpoint"`
															} `json:"navigationEndpoint"`
															Menu struct {
																MenuRenderer struct {
																	Items []struct {
																		MenuServiceItemRenderer struct {
																			Text struct {
																				Runs []struct {
																					Text string `json:"text"`
																				} `json:"runs"`
																			} `json:"text"`
																			Icon struct {
																				IconType string `json:"iconType"`
																			} `json:"icon"`
																			ServiceEndpoint struct {
																				ClickTrackingParams string `json:"clickTrackingParams"`
																				CommandMetadata     struct {
																					WebCommandMetadata struct {
																						SendPost bool   `json:"sendPost"`
																						ApiUrl   string `json:"apiUrl"`
																					} `json:"webCommandMetadata"`
																				} `json:"commandMetadata"`
																				ShareEntityServiceEndpoint struct {
																					SerializedShareEntity string `json:"serializedShareEntity"`
																					Commands              []struct {
																						ClickTrackingParams string `json:"clickTrackingParams"`
																						OpenPopupAction     struct {
																							Popup struct {
																								UnifiedSharePanelRenderer struct {
																									TrackingParams     string `json:"trackingParams"`
																									ShowLoadingSpinner bool   `json:"showLoadingSpinner"`
																								} `json:"unifiedSharePanelRenderer"`
																							} `json:"popup"`
																							PopupType string `json:"popupType"`
																							BeReused  bool   `json:"beReused"`
																						} `json:"openPopupAction"`
																					} `json:"commands"`
																				} `json:"shareEntityServiceEndpoint"`
																			} `json:"serviceEndpoint"`
																			TrackingParams string `json:"trackingParams"`
																			Accessibility  struct {
																				AccessibilityData struct {
																					Label string `json:"label"`
																				} `json:"accessibilityData"`
																			} `json:"accessibility"`
																			HasSeparator bool `json:"hasSeparator"`
																		} `json:"menuServiceItemRenderer,omitempty"`
																		MenuNavigationItemRenderer struct {
																			Text struct {
																				Runs []struct {
																					Text string `json:"text"`
																				} `json:"runs"`
																			} `json:"text"`
																			Icon struct {
																				IconType string `json:"iconType"`
																			} `json:"icon"`
																			NavigationEndpoint struct {
																				ClickTrackingParams string `json:"clickTrackingParams"`
																				CommandMetadata     struct {
																					WebCommandMetadata struct {
																						IgnoreNavigation bool `json:"ignoreNavigation"`
																					} `json:"webCommandMetadata"`
																				} `json:"commandMetadata"`
																				UserFeedbackEndpoint struct {
																					AdditionalDatas []struct {
																						UserFeedbackEndpointProductSpecificValueData struct {
																							Key   string `json:"key"`
																							Value string `json:"value"`
																						} `json:"userFeedbackEndpointProductSpecificValueData"`
																					} `json:"additionalDatas"`
																				} `json:"userFeedbackEndpoint"`
																			} `json:"navigationEndpoint"`
																			TrackingParams string `json:"trackingParams"`
																			Accessibility  struct {
																				AccessibilityData struct {
																					Label string `json:"label"`
																				} `json:"accessibilityData"`
																			} `json:"accessibility"`
																		} `json:"menuNavigationItemRenderer,omitempty"`
																	} `json:"items"`
																	TrackingParams string `json:"trackingParams"`
																	Accessibility  struct {
																		AccessibilityData struct {
																			Label string `json:"label"`
																		} `json:"accessibilityData"`
																	} `json:"accessibility"`
																} `json:"menuRenderer"`
															} `json:"menu"`
															TrackingParams string `json:"trackingParams"`
															Accessibility  struct {
																AccessibilityData struct {
																	Label string `json:"label"`
																} `json:"accessibilityData"`
															} `json:"accessibility"`
															Style                  string `json:"style"`
															VideoType              string `json:"videoType"`
															InlinePlaybackEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		Url         string `json:"url"`
																		WebPageType string `json:"webPageType"`
																		RootVe      int    `json:"rootVe"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																WatchEndpoint struct {
																	VideoId              string `json:"videoId"`
																	PlayerParams         string `json:"playerParams"`
																	PlayerExtraUrlParams []struct {
																		Key   string `json:"key"`
																		Value string `json:"value"`
																	} `json:"playerExtraUrlParams"`
																	WatchEndpointSupportedOnesieConfig struct {
																		Html5PlaybackOnesieConfig struct {
																			CommonConfig struct {
																				Url string `json:"url"`
																			} `json:"commonConfig"`
																		} `json:"html5PlaybackOnesieConfig"`
																	} `json:"watchEndpointSupportedOnesieConfig"`
																} `json:"watchEndpoint"`
															} `json:"inlinePlaybackEndpoint"`
															LoggingDirectives struct {
																TrackingParams string `json:"trackingParams"`
																Visibility     struct {
																	Types string `json:"types"`
																} `json:"visibility"`
																EnableDisplayloggerExperiment bool `json:"enableDisplayloggerExperiment"`
															} `json:"loggingDirectives"`
														} `json:"reelItemRenderer,omitempty"`
														VideoRenderer struct {
															VideoId   string `json:"videoId"`
															Thumbnail struct {
																Thumbnails []struct {
																	Url    string `json:"url"`
																	Width  int    `json:"width"`
																	Height int    `json:"height"`
																} `json:"thumbnails"`
															} `json:"thumbnail"`
															Title struct {
																Runs []struct {
																	Text string `json:"text"`
																} `json:"runs"`
																Accessibility struct {
																	AccessibilityData struct {
																		Label string `json:"label"`
																	} `json:"accessibilityData"`
																} `json:"accessibility"`
															} `json:"title"`
															DescriptionSnippet struct {
																Runs []struct {
																	Text string `json:"text"`
																} `json:"runs"`
															} `json:"descriptionSnippet"`
															LongBylineText struct {
																Runs []struct {
																	Text               string `json:"text"`
																	NavigationEndpoint struct {
																		ClickTrackingParams string `json:"clickTrackingParams"`
																		CommandMetadata     struct {
																			WebCommandMetadata struct {
																				Url         string `json:"url"`
																				WebPageType string `json:"webPageType"`
																				RootVe      int    `json:"rootVe"`
																				ApiUrl      string `json:"apiUrl"`
																			} `json:"webCommandMetadata"`
																		} `json:"commandMetadata"`
																		BrowseEndpoint struct {
																			BrowseId         string `json:"browseId"`
																			CanonicalBaseUrl string `json:"canonicalBaseUrl"`
																		} `json:"browseEndpoint"`
																	} `json:"navigationEndpoint"`
																} `json:"runs"`
															} `json:"longBylineText"`
															PublishedTimeText struct {
																SimpleText string `json:"simpleText"`
															} `json:"publishedTimeText"`
															LengthText struct {
																Accessibility struct {
																	AccessibilityData struct {
																		Label string `json:"label"`
																	} `json:"accessibilityData"`
																} `json:"accessibility"`
																SimpleText string `json:"simpleText"`
															} `json:"lengthText"`
															ViewCountText struct {
																SimpleText string `json:"simpleText"`
															} `json:"viewCountText"`
															NavigationEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		Url         string `json:"url"`
																		WebPageType string `json:"webPageType"`
																		RootVe      int    `json:"rootVe"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																WatchEndpoint struct {
																	VideoId                            string `json:"videoId"`
																	WatchEndpointSupportedOnesieConfig struct {
																		Html5PlaybackOnesieConfig struct {
																			CommonConfig struct {
																				Url string `json:"url"`
																			} `json:"commonConfig"`
																		} `json:"html5PlaybackOnesieConfig"`
																	} `json:"watchEndpointSupportedOnesieConfig"`
																} `json:"watchEndpoint"`
															} `json:"navigationEndpoint"`
															OwnerBadges []struct {
																MetadataBadgeRenderer struct {
																	Icon struct {
																		IconType string `json:"iconType"`
																	} `json:"icon"`
																	Style             string `json:"style"`
																	Tooltip           string `json:"tooltip"`
																	TrackingParams    string `json:"trackingParams"`
																	AccessibilityData struct {
																		Label string `json:"label"`
																	} `json:"accessibilityData"`
																} `json:"metadataBadgeRenderer"`
															} `json:"ownerBadges"`
															OwnerText struct {
																Runs []struct {
																	Text               string `json:"text"`
																	NavigationEndpoint struct {
																		ClickTrackingParams string `json:"clickTrackingParams"`
																		CommandMetadata     struct {
																			WebCommandMetadata struct {
																				Url         string `json:"url"`
																				WebPageType string `json:"webPageType"`
																				RootVe      int    `json:"rootVe"`
																				ApiUrl      string `json:"apiUrl"`
																			} `json:"webCommandMetadata"`
																		} `json:"commandMetadata"`
																		BrowseEndpoint struct {
																			BrowseId         string `json:"browseId"`
																			CanonicalBaseUrl string `json:"canonicalBaseUrl"`
																		} `json:"browseEndpoint"`
																	} `json:"navigationEndpoint"`
																} `json:"runs"`
															} `json:"ownerText"`
															ShortBylineText struct {
																Runs []struct {
																	Text               string `json:"text"`
																	NavigationEndpoint struct {
																		ClickTrackingParams string `json:"clickTrackingParams"`
																		CommandMetadata     struct {
																			WebCommandMetadata struct {
																				Url         string `json:"url"`
																				WebPageType string `json:"webPageType"`
																				RootVe      int    `json:"rootVe"`
																				ApiUrl      string `json:"apiUrl"`
																			} `json:"webCommandMetadata"`
																		} `json:"commandMetadata"`
																		BrowseEndpoint struct {
																			BrowseId         string `json:"browseId"`
																			CanonicalBaseUrl string `json:"canonicalBaseUrl"`
																		} `json:"browseEndpoint"`
																	} `json:"navigationEndpoint"`
																} `json:"runs"`
															} `json:"shortBylineText"`
															TrackingParams     string `json:"trackingParams"`
															ShowActionMenu     bool   `json:"showActionMenu"`
															ShortViewCountText struct {
																Accessibility struct {
																	AccessibilityData struct {
																		Label string `json:"label"`
																	} `json:"accessibilityData"`
																} `json:"accessibility"`
																SimpleText string `json:"simpleText"`
															} `json:"shortViewCountText"`
															Menu struct {
																MenuRenderer struct {
																	Items []struct {
																		MenuServiceItemRenderer struct {
																			Text struct {
																				Runs []struct {
																					Text string `json:"text"`
																				} `json:"runs"`
																			} `json:"text"`
																			Icon struct {
																				IconType string `json:"iconType"`
																			} `json:"icon"`
																			ServiceEndpoint struct {
																				ClickTrackingParams string `json:"clickTrackingParams"`
																				CommandMetadata     struct {
																					WebCommandMetadata struct {
																						SendPost bool   `json:"sendPost"`
																						ApiUrl   string `json:"apiUrl,omitempty"`
																					} `json:"webCommandMetadata"`
																				} `json:"commandMetadata"`
																				SignalServiceEndpoint struct {
																					Signal  string `json:"signal"`
																					Actions []struct {
																						ClickTrackingParams  string `json:"clickTrackingParams"`
																						AddToPlaylistCommand struct {
																							OpenMiniplayer      bool   `json:"openMiniplayer"`
																							VideoId             string `json:"videoId"`
																							ListType            string `json:"listType"`
																							OnCreateListCommand struct {
																								ClickTrackingParams string `json:"clickTrackingParams"`
																								CommandMetadata     struct {
																									WebCommandMetadata struct {
																										SendPost bool   `json:"sendPost"`
																										ApiUrl   string `json:"apiUrl"`
																									} `json:"webCommandMetadata"`
																								} `json:"commandMetadata"`
																								CreatePlaylistServiceEndpoint struct {
																									VideoIds []string `json:"videoIds"`
																									Params   string   `json:"params"`
																								} `json:"createPlaylistServiceEndpoint"`
																							} `json:"onCreateListCommand"`
																							VideoIds []string `json:"videoIds"`
																						} `json:"addToPlaylistCommand"`
																					} `json:"actions"`
																				} `json:"signalServiceEndpoint,omitempty"`
																				ShareEntityServiceEndpoint struct {
																					SerializedShareEntity string `json:"serializedShareEntity"`
																					Commands              []struct {
																						ClickTrackingParams string `json:"clickTrackingParams"`
																						OpenPopupAction     struct {
																							Popup struct {
																								UnifiedSharePanelRenderer struct {
																									TrackingParams     string `json:"trackingParams"`
																									ShowLoadingSpinner bool   `json:"showLoadingSpinner"`
																								} `json:"unifiedSharePanelRenderer"`
																							} `json:"popup"`
																							PopupType string `json:"popupType"`
																							BeReused  bool   `json:"beReused"`
																						} `json:"openPopupAction"`
																					} `json:"commands"`
																				} `json:"shareEntityServiceEndpoint,omitempty"`
																			} `json:"serviceEndpoint"`
																			TrackingParams string `json:"trackingParams"`
																			HasSeparator   bool   `json:"hasSeparator,omitempty"`
																		} `json:"menuServiceItemRenderer"`
																	} `json:"items"`
																	TrackingParams string `json:"trackingParams"`
																	Accessibility  struct {
																		AccessibilityData struct {
																			Label string `json:"label"`
																		} `json:"accessibilityData"`
																	} `json:"accessibility"`
																} `json:"menuRenderer"`
															} `json:"menu"`
															ChannelThumbnailSupportedRenderers struct {
																ChannelThumbnailWithLinkRenderer struct {
																	Thumbnail struct {
																		Thumbnails []struct {
																			Url    string `json:"url"`
																			Width  int    `json:"width"`
																			Height int    `json:"height"`
																		} `json:"thumbnails"`
																	} `json:"thumbnail"`
																	NavigationEndpoint struct {
																		ClickTrackingParams string `json:"clickTrackingParams"`
																		CommandMetadata     struct {
																			WebCommandMetadata struct {
																				Url         string `json:"url"`
																				WebPageType string `json:"webPageType"`
																				RootVe      int    `json:"rootVe"`
																				ApiUrl      string `json:"apiUrl"`
																			} `json:"webCommandMetadata"`
																		} `json:"commandMetadata"`
																		BrowseEndpoint struct {
																			BrowseId         string `json:"browseId"`
																			CanonicalBaseUrl string `json:"canonicalBaseUrl"`
																		} `json:"browseEndpoint"`
																	} `json:"navigationEndpoint"`
																	Accessibility struct {
																		AccessibilityData struct {
																			Label string `json:"label"`
																		} `json:"accessibilityData"`
																	} `json:"accessibility"`
																} `json:"channelThumbnailWithLinkRenderer"`
															} `json:"channelThumbnailSupportedRenderers"`
															ThumbnailOverlays []struct {
																ThumbnailOverlayTimeStatusRenderer struct {
																	Text struct {
																		Accessibility struct {
																			AccessibilityData struct {
																				Label string `json:"label"`
																			} `json:"accessibilityData"`
																		} `json:"accessibility"`
																		SimpleText string `json:"simpleText"`
																	} `json:"text"`
																	Style string `json:"style"`
																} `json:"thumbnailOverlayTimeStatusRenderer,omitempty"`
																ThumbnailOverlayToggleButtonRenderer struct {
																	IsToggled     bool `json:"isToggled,omitempty"`
																	UntoggledIcon struct {
																		IconType string `json:"iconType"`
																	} `json:"untoggledIcon"`
																	ToggledIcon struct {
																		IconType string `json:"iconType"`
																	} `json:"toggledIcon"`
																	UntoggledTooltip         string `json:"untoggledTooltip"`
																	ToggledTooltip           string `json:"toggledTooltip"`
																	UntoggledServiceEndpoint struct {
																		ClickTrackingParams string `json:"clickTrackingParams"`
																		CommandMetadata     struct {
																			WebCommandMetadata struct {
																				SendPost bool   `json:"sendPost"`
																				ApiUrl   string `json:"apiUrl,omitempty"`
																			} `json:"webCommandMetadata"`
																		} `json:"commandMetadata"`
																		PlaylistEditEndpoint struct {
																			PlaylistId string `json:"playlistId"`
																			Actions    []struct {
																				AddedVideoId string `json:"addedVideoId"`
																				Action       string `json:"action"`
																			} `json:"actions"`
																		} `json:"playlistEditEndpoint,omitempty"`
																		SignalServiceEndpoint struct {
																			Signal  string `json:"signal"`
																			Actions []struct {
																				ClickTrackingParams  string `json:"clickTrackingParams"`
																				AddToPlaylistCommand struct {
																					OpenMiniplayer      bool   `json:"openMiniplayer"`
																					VideoId             string `json:"videoId"`
																					ListType            string `json:"listType"`
																					OnCreateListCommand struct {
																						ClickTrackingParams string `json:"clickTrackingParams"`
																						CommandMetadata     struct {
																							WebCommandMetadata struct {
																								SendPost bool   `json:"sendPost"`
																								ApiUrl   string `json:"apiUrl"`
																							} `json:"webCommandMetadata"`
																						} `json:"commandMetadata"`
																						CreatePlaylistServiceEndpoint struct {
																							VideoIds []string `json:"videoIds"`
																							Params   string   `json:"params"`
																						} `json:"createPlaylistServiceEndpoint"`
																					} `json:"onCreateListCommand"`
																					VideoIds []string `json:"videoIds"`
																				} `json:"addToPlaylistCommand"`
																			} `json:"actions"`
																		} `json:"signalServiceEndpoint,omitempty"`
																	} `json:"untoggledServiceEndpoint"`
																	ToggledServiceEndpoint struct {
																		ClickTrackingParams string `json:"clickTrackingParams"`
																		CommandMetadata     struct {
																			WebCommandMetadata struct {
																				SendPost bool   `json:"sendPost"`
																				ApiUrl   string `json:"apiUrl"`
																			} `json:"webCommandMetadata"`
																		} `json:"commandMetadata"`
																		PlaylistEditEndpoint struct {
																			PlaylistId string `json:"playlistId"`
																			Actions    []struct {
																				Action         string `json:"action"`
																				RemovedVideoId string `json:"removedVideoId"`
																			} `json:"actions"`
																		} `json:"playlistEditEndpoint"`
																	} `json:"toggledServiceEndpoint,omitempty"`
																	UntoggledAccessibility struct {
																		AccessibilityData struct {
																			Label string `json:"label"`
																		} `json:"accessibilityData"`
																	} `json:"untoggledAccessibility"`
																	ToggledAccessibility struct {
																		AccessibilityData struct {
																			Label string `json:"label"`
																		} `json:"accessibilityData"`
																	} `json:"toggledAccessibility"`
																	TrackingParams string `json:"trackingParams"`
																} `json:"thumbnailOverlayToggleButtonRenderer,omitempty"`
																ThumbnailOverlayNowPlayingRenderer struct {
																	Text struct {
																		Runs []struct {
																			Text string `json:"text"`
																		} `json:"runs"`
																	} `json:"text"`
																} `json:"thumbnailOverlayNowPlayingRenderer,omitempty"`
																ThumbnailOverlayLoadingPreviewRenderer struct {
																	Text struct {
																		Runs []struct {
																			Text string `json:"text"`
																		} `json:"runs"`
																	} `json:"text"`
																} `json:"thumbnailOverlayLoadingPreviewRenderer,omitempty"`
															} `json:"thumbnailOverlays"`
															InlinePlaybackEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		Url         string `json:"url"`
																		WebPageType string `json:"webPageType"`
																		RootVe      int    `json:"rootVe"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																WatchEndpoint struct {
																	VideoId              string `json:"videoId"`
																	PlayerParams         string `json:"playerParams"`
																	PlayerExtraUrlParams []struct {
																		Key   string `json:"key"`
																		Value string `json:"value"`
																	} `json:"playerExtraUrlParams"`
																	WatchEndpointSupportedOnesieConfig struct {
																		Html5PlaybackOnesieConfig struct {
																			CommonConfig struct {
																				Url string `json:"url"`
																			} `json:"commonConfig"`
																		} `json:"html5PlaybackOnesieConfig"`
																	} `json:"watchEndpointSupportedOnesieConfig"`
																} `json:"watchEndpoint"`
															} `json:"inlinePlaybackEndpoint"`
														} `json:"videoRenderer,omitempty"`
													} `json:"content"`
													TrackingParams string `json:"trackingParams"`
												} `json:"richItemRenderer"`
											} `json:"contents"`
											TrackingParams string `json:"trackingParams"`
											ShowMoreButton struct {
												ButtonRenderer struct {
													Style string `json:"style"`
													Size  string `json:"size"`
													Icon  struct {
														IconType string `json:"iconType"`
													} `json:"icon"`
													Accessibility struct {
														Label string `json:"label"`
													} `json:"accessibility"`
													Tooltip        string `json:"tooltip"`
													TrackingParams string `json:"trackingParams"`
												} `json:"buttonRenderer"`
											} `json:"showMoreButton"`
											IsExpanded bool `json:"isExpanded"`
											Icon       struct {
												IconType string `json:"iconType"`
											} `json:"icon,omitempty"`
											IsTopDividerHidden    bool `json:"isTopDividerHidden"`
											IsBottomDividerHidden bool `json:"isBottomDividerHidden"`
											Endpoint              struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														Url         string `json:"url"`
														WebPageType string `json:"webPageType"`
														RootVe      int    `json:"rootVe"`
														ApiUrl      string `json:"apiUrl"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												BrowseEndpoint struct {
													BrowseId string `json:"browseId"`
												} `json:"browseEndpoint"`
											} `json:"endpoint,omitempty"`
										} `json:"richShelfRenderer"`
									} `json:"content"`
									TrackingParams string `json:"trackingParams"`
									FullBleed      bool   `json:"fullBleed"`
								} `json:"richSectionRenderer,omitempty"`
							} `json:"contents"`
							TrackingParams string `json:"trackingParams"`
							Header         struct {
								FeedFilterChipBarRenderer struct {
									Contents []struct {
										ChipCloudChipRenderer struct {
											Text struct {
												SimpleText string `json:"simpleText,omitempty"`
												Runs       []struct {
													Text string `json:"text"`
												} `json:"runs,omitempty"`
											} `json:"text"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														SendPost bool   `json:"sendPost"`
														ApiUrl   string `json:"apiUrl"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												ContinuationCommand struct {
													Token   string `json:"token"`
													Request string `json:"request"`
													Command struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														ShowReloadUiCommand struct {
															TargetId string `json:"targetId"`
														} `json:"showReloadUiCommand"`
													} `json:"command"`
												} `json:"continuationCommand"`
											} `json:"navigationEndpoint,omitempty"`
											TrackingParams string `json:"trackingParams"`
											IsSelected     bool   `json:"isSelected,omitempty"`
											Style          struct {
												StyleType string `json:"styleType"`
											} `json:"style,omitempty"`
											TargetId string `json:"targetId,omitempty"`
										} `json:"chipCloudChipRenderer"`
									} `json:"contents"`
									TrackingParams string `json:"trackingParams"`
									NextButton     struct {
										ButtonRenderer struct {
											Style      string `json:"style"`
											Size       string `json:"size"`
											IsDisabled bool   `json:"isDisabled"`
											Icon       struct {
												IconType string `json:"iconType"`
											} `json:"icon"`
											Tooltip           string `json:"tooltip"`
											TrackingParams    string `json:"trackingParams"`
											AccessibilityData struct {
												AccessibilityData struct {
													Label string `json:"label"`
												} `json:"accessibilityData"`
											} `json:"accessibilityData"`
										} `json:"buttonRenderer"`
									} `json:"nextButton"`
									PreviousButton struct {
										ButtonRenderer struct {
											Style      string `json:"style"`
											Size       string `json:"size"`
											IsDisabled bool   `json:"isDisabled"`
											Icon       struct {
												IconType string `json:"iconType"`
											} `json:"icon"`
											Tooltip           string `json:"tooltip"`
											TrackingParams    string `json:"trackingParams"`
											AccessibilityData struct {
												AccessibilityData struct {
													Label string `json:"label"`
												} `json:"accessibilityData"`
											} `json:"accessibilityData"`
										} `json:"buttonRenderer"`
									} `json:"previousButton"`
									StyleType string `json:"styleType"`
								} `json:"feedFilterChipBarRenderer"`
							} `json:"header"`
							TargetId      string `json:"targetId"`
							Style         string `json:"style,omitempty"`
							ReflowOptions struct {
								MinimumRowsOfVideosAtStart         int `json:"minimumRowsOfVideosAtStart"`
								MinimumRowsOfVideosBetweenSections int `json:"minimumRowsOfVideosBetweenSections"`
							} `json:"reflowOptions,omitempty"`
						} `json:"richGridRenderer"`
					} `json:"content,omitempty"`
					TabIdentifier string `json:"tabIdentifier,omitempty"`
				} `json:"tabRenderer,omitempty"`
				ExpandableTabRenderer struct {
					Endpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								Url         string `json:"url"`
								WebPageType string `json:"webPageType"`
								RootVe      int    `json:"rootVe"`
								ApiUrl      string `json:"apiUrl"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						BrowseEndpoint struct {
							BrowseId         string `json:"browseId"`
							Params           string `json:"params"`
							CanonicalBaseUrl string `json:"canonicalBaseUrl"`
						} `json:"browseEndpoint"`
					} `json:"endpoint"`
					Title    string `json:"title"`
					Selected bool   `json:"selected"`
				} `json:"expandableTabRenderer,omitempty"`
			} `json:"tabs"`
		} `json:"twoColumnBrowseResultsRenderer,omitempty"`
		TwoColumnWatchNextResults struct {
			Results struct {
				Results struct {
					Contents []struct {
						VideoPrimaryInfoRenderer struct {
							Title struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"title"`
							ViewCount struct {
								VideoViewCountRenderer struct {
									ViewCount struct {
										SimpleText string `json:"simpleText"`
									} `json:"viewCount"`
									ShortViewCount struct {
										SimpleText string `json:"simpleText"`
									} `json:"shortViewCount"`
								} `json:"videoViewCountRenderer"`
							} `json:"viewCount"`
							VideoActions struct {
								MenuRenderer struct {
									Items []struct {
										MenuNavigationItemRenderer struct {
											Text struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"text"`
											Icon struct {
												IconType string `json:"iconType"`
											} `json:"icon"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														IgnoreNavigation bool `json:"ignoreNavigation"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												ModalEndpoint struct {
													Modal struct {
														ModalWithTitleAndButtonRenderer struct {
															Title struct {
																Runs []struct {
																	Text string `json:"text"`
																} `json:"runs"`
															} `json:"title"`
															Content struct {
																Runs []struct {
																	Text string `json:"text"`
																} `json:"runs"`
															} `json:"content"`
															Button struct {
																ButtonRenderer struct {
																	Style      string `json:"style"`
																	Size       string `json:"size"`
																	IsDisabled bool   `json:"isDisabled"`
																	Text       struct {
																		SimpleText string `json:"simpleText"`
																	} `json:"text"`
																	NavigationEndpoint struct {
																		ClickTrackingParams string `json:"clickTrackingParams"`
																		CommandMetadata     struct {
																			WebCommandMetadata struct {
																				Url         string `json:"url"`
																				WebPageType string `json:"webPageType"`
																				RootVe      int    `json:"rootVe"`
																			} `json:"webCommandMetadata"`
																		} `json:"commandMetadata"`
																		SignInEndpoint struct {
																			Hack bool `json:"hack"`
																		} `json:"signInEndpoint"`
																	} `json:"navigationEndpoint"`
																	TrackingParams string `json:"trackingParams"`
																} `json:"buttonRenderer"`
															} `json:"button"`
														} `json:"modalWithTitleAndButtonRenderer"`
													} `json:"modal"`
												} `json:"modalEndpoint"`
											} `json:"navigationEndpoint"`
											TrackingParams string `json:"trackingParams"`
										} `json:"menuNavigationItemRenderer,omitempty"`
										MenuServiceItemRenderer struct {
											Text struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"text"`
											Icon struct {
												IconType string `json:"iconType"`
											} `json:"icon"`
											ServiceEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														SendPost bool `json:"sendPost"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												SignalServiceEndpoint struct {
													Signal  string `json:"signal"`
													Actions []struct {
														ClickTrackingParams         string `json:"clickTrackingParams"`
														ShowEngagementPanelEndpoint struct {
															PanelIdentifier string `json:"panelIdentifier"`
														} `json:"showEngagementPanelEndpoint"`
													} `json:"actions"`
												} `json:"signalServiceEndpoint"`
											} `json:"serviceEndpoint"`
											TrackingParams string `json:"trackingParams"`
										} `json:"menuServiceItemRenderer,omitempty"`
									} `json:"items"`
									TrackingParams  string `json:"trackingParams"`
									TopLevelButtons []struct {
										SegmentedLikeDislikeButtonRenderer struct {
											LikeButton struct {
												ToggleButtonRenderer struct {
													Style struct {
														StyleType string `json:"styleType"`
													} `json:"style"`
													IsToggled   bool `json:"isToggled"`
													IsDisabled  bool `json:"isDisabled"`
													DefaultIcon struct {
														IconType string `json:"iconType"`
													} `json:"defaultIcon"`
													DefaultText struct {
														Accessibility struct {
															AccessibilityData struct {
																Label string `json:"label"`
															} `json:"accessibilityData"`
														} `json:"accessibility"`
														SimpleText string `json:"simpleText"`
													} `json:"defaultText"`
													ToggledText struct {
														Accessibility struct {
															AccessibilityData struct {
																Label string `json:"label"`
															} `json:"accessibilityData"`
														} `json:"accessibility"`
														SimpleText string `json:"simpleText"`
													} `json:"toggledText"`
													Accessibility struct {
														Label string `json:"label"`
													} `json:"accessibility"`
													TrackingParams string `json:"trackingParams"`
													DefaultTooltip string `json:"defaultTooltip"`
													ToggledTooltip string `json:"toggledTooltip"`
													ToggledStyle   struct {
														StyleType string `json:"styleType"`
													} `json:"toggledStyle"`
													DefaultNavigationEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																IgnoreNavigation bool `json:"ignoreNavigation"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														ModalEndpoint struct {
															Modal struct {
																ModalWithTitleAndButtonRenderer struct {
																	Title struct {
																		SimpleText string `json:"simpleText"`
																	} `json:"title"`
																	Content struct {
																		SimpleText string `json:"simpleText"`
																	} `json:"content"`
																	Button struct {
																		ButtonRenderer struct {
																			Style      string `json:"style"`
																			Size       string `json:"size"`
																			IsDisabled bool   `json:"isDisabled"`
																			Text       struct {
																				SimpleText string `json:"simpleText"`
																			} `json:"text"`
																			NavigationEndpoint struct {
																				ClickTrackingParams string `json:"clickTrackingParams"`
																				CommandMetadata     struct {
																					WebCommandMetadata struct {
																						Url         string `json:"url"`
																						WebPageType string `json:"webPageType"`
																						RootVe      int    `json:"rootVe"`
																					} `json:"webCommandMetadata"`
																				} `json:"commandMetadata"`
																				SignInEndpoint struct {
																					NextEndpoint struct {
																						ClickTrackingParams string `json:"clickTrackingParams"`
																						CommandMetadata     struct {
																							WebCommandMetadata struct {
																								Url         string `json:"url"`
																								WebPageType string `json:"webPageType"`
																								RootVe      int    `json:"rootVe"`
																							} `json:"webCommandMetadata"`
																						} `json:"commandMetadata"`
																						WatchEndpoint struct {
																							VideoId                            string `json:"videoId"`
																							WatchEndpointSupportedOnesieConfig struct {
																								Html5PlaybackOnesieConfig struct {
																									CommonConfig struct {
																										Url string `json:"url"`
																									} `json:"commonConfig"`
																								} `json:"html5PlaybackOnesieConfig"`
																							} `json:"watchEndpointSupportedOnesieConfig"`
																						} `json:"watchEndpoint"`
																					} `json:"nextEndpoint"`
																					IdamTag string `json:"idamTag"`
																				} `json:"signInEndpoint"`
																			} `json:"navigationEndpoint"`
																			TrackingParams string `json:"trackingParams"`
																		} `json:"buttonRenderer"`
																	} `json:"button"`
																} `json:"modalWithTitleAndButtonRenderer"`
															} `json:"modal"`
														} `json:"modalEndpoint"`
													} `json:"defaultNavigationEndpoint"`
													AccessibilityData struct {
														AccessibilityData struct {
															Label string `json:"label"`
														} `json:"accessibilityData"`
													} `json:"accessibilityData"`
													ToggleButtonSupportedData struct {
														ToggleButtonIdData struct {
															Id string `json:"id"`
														} `json:"toggleButtonIdData"`
													} `json:"toggleButtonSupportedData"`
													TargetId string `json:"targetId"`
												} `json:"toggleButtonRenderer"`
											} `json:"likeButton"`
											DislikeButton struct {
												ToggleButtonRenderer struct {
													Style struct {
														StyleType string `json:"styleType"`
													} `json:"style"`
													IsToggled   bool `json:"isToggled"`
													IsDisabled  bool `json:"isDisabled"`
													DefaultIcon struct {
														IconType string `json:"iconType"`
													} `json:"defaultIcon"`
													Accessibility struct {
														Label string `json:"label"`
													} `json:"accessibility"`
													TrackingParams string `json:"trackingParams"`
													DefaultTooltip string `json:"defaultTooltip"`
													ToggledTooltip string `json:"toggledTooltip"`
													ToggledStyle   struct {
														StyleType string `json:"styleType"`
													} `json:"toggledStyle"`
													DefaultNavigationEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																IgnoreNavigation bool `json:"ignoreNavigation"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														ModalEndpoint struct {
															Modal struct {
																ModalWithTitleAndButtonRenderer struct {
																	Title struct {
																		SimpleText string `json:"simpleText"`
																	} `json:"title"`
																	Content struct {
																		SimpleText string `json:"simpleText"`
																	} `json:"content"`
																	Button struct {
																		ButtonRenderer struct {
																			Style      string `json:"style"`
																			Size       string `json:"size"`
																			IsDisabled bool   `json:"isDisabled"`
																			Text       struct {
																				SimpleText string `json:"simpleText"`
																			} `json:"text"`
																			NavigationEndpoint struct {
																				ClickTrackingParams string `json:"clickTrackingParams"`
																				CommandMetadata     struct {
																					WebCommandMetadata struct {
																						Url         string `json:"url"`
																						WebPageType string `json:"webPageType"`
																						RootVe      int    `json:"rootVe"`
																					} `json:"webCommandMetadata"`
																				} `json:"commandMetadata"`
																				SignInEndpoint struct {
																					NextEndpoint struct {
																						ClickTrackingParams string `json:"clickTrackingParams"`
																						CommandMetadata     struct {
																							WebCommandMetadata struct {
																								Url         string `json:"url"`
																								WebPageType string `json:"webPageType"`
																								RootVe      int    `json:"rootVe"`
																							} `json:"webCommandMetadata"`
																						} `json:"commandMetadata"`
																						WatchEndpoint struct {
																							VideoId                            string `json:"videoId"`
																							WatchEndpointSupportedOnesieConfig struct {
																								Html5PlaybackOnesieConfig struct {
																									CommonConfig struct {
																										Url string `json:"url"`
																									} `json:"commonConfig"`
																								} `json:"html5PlaybackOnesieConfig"`
																							} `json:"watchEndpointSupportedOnesieConfig"`
																						} `json:"watchEndpoint"`
																					} `json:"nextEndpoint"`
																					IdamTag string `json:"idamTag"`
																				} `json:"signInEndpoint"`
																			} `json:"navigationEndpoint"`
																			TrackingParams string `json:"trackingParams"`
																		} `json:"buttonRenderer"`
																	} `json:"button"`
																} `json:"modalWithTitleAndButtonRenderer"`
															} `json:"modal"`
														} `json:"modalEndpoint"`
													} `json:"defaultNavigationEndpoint"`
													AccessibilityData struct {
														AccessibilityData struct {
															Label string `json:"label"`
														} `json:"accessibilityData"`
													} `json:"accessibilityData"`
													ToggleButtonSupportedData struct {
														ToggleButtonIdData struct {
															Id string `json:"id"`
														} `json:"toggleButtonIdData"`
													} `json:"toggleButtonSupportedData"`
													TargetId string `json:"targetId"`
												} `json:"toggleButtonRenderer"`
											} `json:"dislikeButton"`
											LikeCount string `json:"likeCount"`
										} `json:"segmentedLikeDislikeButtonRenderer,omitempty"`
										ButtonRenderer struct {
											Style      string `json:"style"`
											Size       string `json:"size"`
											IsDisabled bool   `json:"isDisabled"`
											Text       struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"text"`
											ServiceEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														SendPost bool   `json:"sendPost"`
														ApiUrl   string `json:"apiUrl"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												ShareEntityServiceEndpoint struct {
													SerializedShareEntity string `json:"serializedShareEntity"`
													Commands              []struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														OpenPopupAction     struct {
															Popup struct {
																UnifiedSharePanelRenderer struct {
																	TrackingParams     string `json:"trackingParams"`
																	ShowLoadingSpinner bool   `json:"showLoadingSpinner"`
																} `json:"unifiedSharePanelRenderer"`
															} `json:"popup"`
															PopupType string `json:"popupType"`
															BeReused  bool   `json:"beReused"`
														} `json:"openPopupAction"`
													} `json:"commands"`
												} `json:"shareEntityServiceEndpoint"`
											} `json:"serviceEndpoint"`
											Icon struct {
												IconType string `json:"iconType"`
											} `json:"icon"`
											Tooltip           string `json:"tooltip"`
											TrackingParams    string `json:"trackingParams"`
											AccessibilityData struct {
												AccessibilityData struct {
													Label string `json:"label"`
												} `json:"accessibilityData"`
											} `json:"accessibilityData"`
										} `json:"buttonRenderer,omitempty"`
									} `json:"topLevelButtons"`
									Accessibility struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"accessibility"`
									FlexibleItems []struct {
										MenuFlexibleItemRenderer struct {
											MenuItem struct {
												MenuServiceItemRenderer struct {
													Text struct {
														Runs []struct {
															Text string `json:"text"`
														} `json:"runs"`
													} `json:"text"`
													Icon struct {
														IconType string `json:"iconType"`
													} `json:"icon"`
													ServiceEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																IgnoreNavigation bool `json:"ignoreNavigation"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														ModalEndpoint struct {
															Modal struct {
																ModalWithTitleAndButtonRenderer struct {
																	Title struct {
																		Runs []struct {
																			Text string `json:"text"`
																		} `json:"runs"`
																	} `json:"title"`
																	Content struct {
																		Runs []struct {
																			Text string `json:"text"`
																		} `json:"runs"`
																	} `json:"content"`
																	Button struct {
																		ButtonRenderer struct {
																			Style      string `json:"style"`
																			Size       string `json:"size"`
																			IsDisabled bool   `json:"isDisabled"`
																			Text       struct {
																				SimpleText string `json:"simpleText"`
																			} `json:"text"`
																			NavigationEndpoint struct {
																				ClickTrackingParams string `json:"clickTrackingParams"`
																				CommandMetadata     struct {
																					WebCommandMetadata struct {
																						Url         string `json:"url"`
																						WebPageType string `json:"webPageType"`
																						RootVe      int    `json:"rootVe"`
																					} `json:"webCommandMetadata"`
																				} `json:"commandMetadata"`
																				SignInEndpoint struct {
																					NextEndpoint struct {
																						ClickTrackingParams string `json:"clickTrackingParams"`
																						CommandMetadata     struct {
																							WebCommandMetadata struct {
																								Url         string `json:"url"`
																								WebPageType string `json:"webPageType"`
																								RootVe      int    `json:"rootVe"`
																							} `json:"webCommandMetadata"`
																						} `json:"commandMetadata"`
																						WatchEndpoint struct {
																							VideoId                            string `json:"videoId"`
																							WatchEndpointSupportedOnesieConfig struct {
																								Html5PlaybackOnesieConfig struct {
																									CommonConfig struct {
																										Url string `json:"url"`
																									} `json:"commonConfig"`
																								} `json:"html5PlaybackOnesieConfig"`
																							} `json:"watchEndpointSupportedOnesieConfig"`
																						} `json:"watchEndpoint"`
																					} `json:"nextEndpoint"`
																					IdamTag string `json:"idamTag"`
																				} `json:"signInEndpoint"`
																			} `json:"navigationEndpoint"`
																			TrackingParams string `json:"trackingParams"`
																		} `json:"buttonRenderer"`
																	} `json:"button"`
																} `json:"modalWithTitleAndButtonRenderer"`
															} `json:"modal"`
														} `json:"modalEndpoint"`
													} `json:"serviceEndpoint"`
													TrackingParams string `json:"trackingParams"`
												} `json:"menuServiceItemRenderer"`
											} `json:"menuItem"`
											TopLevelButton struct {
												ButtonRenderer struct {
													Style      string `json:"style"`
													Size       string `json:"size"`
													IsDisabled bool   `json:"isDisabled"`
													Text       struct {
														Runs []struct {
															Text string `json:"text"`
														} `json:"runs"`
													} `json:"text"`
													Icon struct {
														IconType string `json:"iconType"`
													} `json:"icon"`
													Accessibility struct {
														Label string `json:"label"`
													} `json:"accessibility"`
													Tooltip           string `json:"tooltip"`
													TrackingParams    string `json:"trackingParams"`
													AccessibilityData struct {
														AccessibilityData struct {
															Label string `json:"label"`
														} `json:"accessibilityData"`
													} `json:"accessibilityData"`
													Command struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																IgnoreNavigation bool `json:"ignoreNavigation"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														ModalEndpoint struct {
															Modal struct {
																ModalWithTitleAndButtonRenderer struct {
																	Title struct {
																		Runs []struct {
																			Text string `json:"text"`
																		} `json:"runs"`
																	} `json:"title"`
																	Content struct {
																		Runs []struct {
																			Text string `json:"text"`
																		} `json:"runs"`
																	} `json:"content"`
																	Button struct {
																		ButtonRenderer struct {
																			Style      string `json:"style"`
																			Size       string `json:"size"`
																			IsDisabled bool   `json:"isDisabled"`
																			Text       struct {
																				SimpleText string `json:"simpleText"`
																			} `json:"text"`
																			NavigationEndpoint struct {
																				ClickTrackingParams string `json:"clickTrackingParams"`
																				CommandMetadata     struct {
																					WebCommandMetadata struct {
																						Url         string `json:"url"`
																						WebPageType string `json:"webPageType"`
																						RootVe      int    `json:"rootVe"`
																					} `json:"webCommandMetadata"`
																				} `json:"commandMetadata"`
																				SignInEndpoint struct {
																					NextEndpoint struct {
																						ClickTrackingParams string `json:"clickTrackingParams"`
																						CommandMetadata     struct {
																							WebCommandMetadata struct {
																								Url         string `json:"url"`
																								WebPageType string `json:"webPageType"`
																								RootVe      int    `json:"rootVe"`
																							} `json:"webCommandMetadata"`
																						} `json:"commandMetadata"`
																						WatchEndpoint struct {
																							VideoId                            string `json:"videoId"`
																							WatchEndpointSupportedOnesieConfig struct {
																								Html5PlaybackOnesieConfig struct {
																									CommonConfig struct {
																										Url string `json:"url"`
																									} `json:"commonConfig"`
																								} `json:"html5PlaybackOnesieConfig"`
																							} `json:"watchEndpointSupportedOnesieConfig"`
																						} `json:"watchEndpoint"`
																					} `json:"nextEndpoint"`
																					IdamTag string `json:"idamTag"`
																				} `json:"signInEndpoint"`
																			} `json:"navigationEndpoint"`
																			TrackingParams string `json:"trackingParams"`
																		} `json:"buttonRenderer"`
																	} `json:"button"`
																} `json:"modalWithTitleAndButtonRenderer"`
															} `json:"modal"`
														} `json:"modalEndpoint"`
													} `json:"command"`
												} `json:"buttonRenderer"`
											} `json:"topLevelButton"`
										} `json:"menuFlexibleItemRenderer"`
									} `json:"flexibleItems"`
								} `json:"menuRenderer"`
							} `json:"videoActions"`
							TrackingParams string `json:"trackingParams"`
							SuperTitleLink struct {
								Runs []struct {
									Text               string `json:"text"`
									NavigationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												Url         string `json:"url"`
												WebPageType string `json:"webPageType"`
												RootVe      int    `json:"rootVe"`
												ApiUrl      string `json:"apiUrl"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										BrowseEndpoint struct {
											BrowseId string `json:"browseId"`
										} `json:"browseEndpoint"`
									} `json:"navigationEndpoint"`
									LoggingDirectives struct {
										TrackingParams string `json:"trackingParams"`
										Visibility     struct {
											Types string `json:"types"`
										} `json:"visibility"`
										EnableDisplayloggerExperiment bool `json:"enableDisplayloggerExperiment"`
									} `json:"loggingDirectives"`
								} `json:"runs"`
							} `json:"superTitleLink"`
							DateText struct {
								SimpleText string `json:"simpleText"`
							} `json:"dateText"`
							RelativeDateText struct {
								Accessibility struct {
									AccessibilityData struct {
										Label string `json:"label"`
									} `json:"accessibilityData"`
								} `json:"accessibility"`
								SimpleText string `json:"simpleText"`
							} `json:"relativeDateText"`
						} `json:"videoPrimaryInfoRenderer,omitempty"`
						VideoSecondaryInfoRenderer struct {
							Owner struct {
								VideoOwnerRenderer struct {
									Thumbnail struct {
										Thumbnails []struct {
											Url    string `json:"url"`
											Width  int    `json:"width"`
											Height int    `json:"height"`
										} `json:"thumbnails"`
									} `json:"thumbnail"`
									Title struct {
										Runs []struct {
											Text               string `json:"text"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														Url         string `json:"url"`
														WebPageType string `json:"webPageType"`
														RootVe      int    `json:"rootVe"`
														ApiUrl      string `json:"apiUrl"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												BrowseEndpoint struct {
													BrowseId         string `json:"browseId"`
													CanonicalBaseUrl string `json:"canonicalBaseUrl"`
												} `json:"browseEndpoint"`
											} `json:"navigationEndpoint"`
										} `json:"runs"`
									} `json:"title"`
									SubscriptionButton struct {
										Type string `json:"type"`
									} `json:"subscriptionButton"`
									NavigationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												Url         string `json:"url"`
												WebPageType string `json:"webPageType"`
												RootVe      int    `json:"rootVe"`
												ApiUrl      string `json:"apiUrl"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										BrowseEndpoint struct {
											BrowseId         string `json:"browseId"`
											CanonicalBaseUrl string `json:"canonicalBaseUrl"`
										} `json:"browseEndpoint"`
									} `json:"navigationEndpoint"`
									SubscriberCountText struct {
										Accessibility struct {
											AccessibilityData struct {
												Label string `json:"label"`
											} `json:"accessibilityData"`
										} `json:"accessibility"`
										SimpleText string `json:"simpleText"`
									} `json:"subscriberCountText"`
									TrackingParams string `json:"trackingParams"`
									Badges         []struct {
										MetadataBadgeRenderer struct {
											Icon struct {
												IconType string `json:"iconType"`
											} `json:"icon"`
											Style             string `json:"style"`
											Tooltip           string `json:"tooltip"`
											TrackingParams    string `json:"trackingParams"`
											AccessibilityData struct {
												Label string `json:"label"`
											} `json:"accessibilityData"`
										} `json:"metadataBadgeRenderer"`
									} `json:"badges"`
								} `json:"videoOwnerRenderer"`
							} `json:"owner"`
							SubscribeButton struct {
								SubscribeButtonRenderer struct {
									ButtonText struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"buttonText"`
									Subscribed           bool   `json:"subscribed"`
									Enabled              bool   `json:"enabled"`
									Type                 string `json:"type"`
									ChannelId            string `json:"channelId"`
									ShowPreferences      bool   `json:"showPreferences"`
									SubscribedButtonText struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"subscribedButtonText"`
									UnsubscribedButtonText struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"unsubscribedButtonText"`
									TrackingParams        string `json:"trackingParams"`
									UnsubscribeButtonText struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"unsubscribeButtonText"`
									SubscribeAccessibility struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"subscribeAccessibility"`
									UnsubscribeAccessibility struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"unsubscribeAccessibility"`
									NotificationPreferenceButton struct {
										SubscriptionNotificationToggleButtonRenderer struct {
											States []struct {
												StateId     int `json:"stateId"`
												NextStateId int `json:"nextStateId"`
												State       struct {
													ButtonRenderer struct {
														Style      string `json:"style"`
														Size       string `json:"size"`
														IsDisabled bool   `json:"isDisabled"`
														Icon       struct {
															IconType string `json:"iconType"`
														} `json:"icon"`
														Accessibility struct {
															Label string `json:"label"`
														} `json:"accessibility"`
														TrackingParams    string `json:"trackingParams"`
														AccessibilityData struct {
															AccessibilityData struct {
																Label string `json:"label"`
															} `json:"accessibilityData"`
														} `json:"accessibilityData"`
													} `json:"buttonRenderer"`
												} `json:"state"`
											} `json:"states"`
											CurrentStateId int    `json:"currentStateId"`
											TrackingParams string `json:"trackingParams"`
											Command        struct {
												ClickTrackingParams    string `json:"clickTrackingParams"`
												CommandExecutorCommand struct {
													Commands []struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														OpenPopupAction     struct {
															Popup struct {
																MenuPopupRenderer struct {
																	Items []struct {
																		MenuServiceItemRenderer struct {
																			Text struct {
																				SimpleText string `json:"simpleText,omitempty"`
																				Runs       []struct {
																					Text string `json:"text"`
																				} `json:"runs,omitempty"`
																			} `json:"text"`
																			Icon struct {
																				IconType string `json:"iconType"`
																			} `json:"icon"`
																			ServiceEndpoint struct {
																				ClickTrackingParams string `json:"clickTrackingParams"`
																				CommandMetadata     struct {
																					WebCommandMetadata struct {
																						SendPost bool   `json:"sendPost"`
																						ApiUrl   string `json:"apiUrl,omitempty"`
																					} `json:"webCommandMetadata"`
																				} `json:"commandMetadata"`
																				ModifyChannelNotificationPreferenceEndpoint struct {
																					Params string `json:"params"`
																				} `json:"modifyChannelNotificationPreferenceEndpoint,omitempty"`
																				SignalServiceEndpoint struct {
																					Signal  string `json:"signal"`
																					Actions []struct {
																						ClickTrackingParams string `json:"clickTrackingParams"`
																						OpenPopupAction     struct {
																							Popup struct {
																								ConfirmDialogRenderer struct {
																									TrackingParams string `json:"trackingParams"`
																									DialogMessages []struct {
																										Runs []struct {
																											Text string `json:"text"`
																										} `json:"runs"`
																									} `json:"dialogMessages"`
																									ConfirmButton struct {
																										ButtonRenderer struct {
																											Style      string `json:"style"`
																											Size       string `json:"size"`
																											IsDisabled bool   `json:"isDisabled"`
																											Text       struct {
																												Runs []struct {
																													Text string `json:"text"`
																												} `json:"runs"`
																											} `json:"text"`
																											ServiceEndpoint struct {
																												ClickTrackingParams string `json:"clickTrackingParams"`
																												CommandMetadata     struct {
																													WebCommandMetadata struct {
																														SendPost bool   `json:"sendPost"`
																														ApiUrl   string `json:"apiUrl"`
																													} `json:"webCommandMetadata"`
																												} `json:"commandMetadata"`
																												UnsubscribeEndpoint struct {
																													ChannelIds []string `json:"channelIds"`
																													Params     string   `json:"params"`
																												} `json:"unsubscribeEndpoint"`
																											} `json:"serviceEndpoint"`
																											Accessibility struct {
																												Label string `json:"label"`
																											} `json:"accessibility"`
																											TrackingParams string `json:"trackingParams"`
																										} `json:"buttonRenderer"`
																									} `json:"confirmButton"`
																									CancelButton struct {
																										ButtonRenderer struct {
																											Style      string `json:"style"`
																											Size       string `json:"size"`
																											IsDisabled bool   `json:"isDisabled"`
																											Text       struct {
																												Runs []struct {
																													Text string `json:"text"`
																												} `json:"runs"`
																											} `json:"text"`
																											Accessibility struct {
																												Label string `json:"label"`
																											} `json:"accessibility"`
																											TrackingParams string `json:"trackingParams"`
																										} `json:"buttonRenderer"`
																									} `json:"cancelButton"`
																									PrimaryIsCancel bool `json:"primaryIsCancel"`
																								} `json:"confirmDialogRenderer"`
																							} `json:"popup"`
																							PopupType string `json:"popupType"`
																						} `json:"openPopupAction"`
																					} `json:"actions"`
																				} `json:"signalServiceEndpoint,omitempty"`
																			} `json:"serviceEndpoint"`
																			TrackingParams string `json:"trackingParams"`
																			IsSelected     bool   `json:"isSelected,omitempty"`
																		} `json:"menuServiceItemRenderer"`
																	} `json:"items"`
																} `json:"menuPopupRenderer"`
															} `json:"popup"`
															PopupType string `json:"popupType"`
														} `json:"openPopupAction"`
													} `json:"commands"`
												} `json:"commandExecutorCommand"`
											} `json:"command"`
											TargetId      string `json:"targetId"`
											SecondaryIcon struct {
												IconType string `json:"iconType"`
											} `json:"secondaryIcon"`
										} `json:"subscriptionNotificationToggleButtonRenderer"`
									} `json:"notificationPreferenceButton"`
									TargetId       string `json:"targetId"`
									SignInEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												IgnoreNavigation bool `json:"ignoreNavigation"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										ModalEndpoint struct {
											Modal struct {
												ModalWithTitleAndButtonRenderer struct {
													Title struct {
														SimpleText string `json:"simpleText"`
													} `json:"title"`
													Content struct {
														SimpleText string `json:"simpleText"`
													} `json:"content"`
													Button struct {
														ButtonRenderer struct {
															Style      string `json:"style"`
															Size       string `json:"size"`
															IsDisabled bool   `json:"isDisabled"`
															Text       struct {
																SimpleText string `json:"simpleText"`
															} `json:"text"`
															NavigationEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		Url         string `json:"url"`
																		WebPageType string `json:"webPageType"`
																		RootVe      int    `json:"rootVe"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																SignInEndpoint struct {
																	NextEndpoint struct {
																		ClickTrackingParams string `json:"clickTrackingParams"`
																		CommandMetadata     struct {
																			WebCommandMetadata struct {
																				Url         string `json:"url"`
																				WebPageType string `json:"webPageType"`
																				RootVe      int    `json:"rootVe"`
																			} `json:"webCommandMetadata"`
																		} `json:"commandMetadata"`
																		WatchEndpoint struct {
																			VideoId                            string `json:"videoId"`
																			WatchEndpointSupportedOnesieConfig struct {
																				Html5PlaybackOnesieConfig struct {
																					CommonConfig struct {
																						Url string `json:"url"`
																					} `json:"commonConfig"`
																				} `json:"html5PlaybackOnesieConfig"`
																			} `json:"watchEndpointSupportedOnesieConfig"`
																		} `json:"watchEndpoint"`
																	} `json:"nextEndpoint"`
																	ContinueAction string `json:"continueAction"`
																	IdamTag        string `json:"idamTag"`
																} `json:"signInEndpoint"`
															} `json:"navigationEndpoint"`
															TrackingParams string `json:"trackingParams"`
														} `json:"buttonRenderer"`
													} `json:"button"`
												} `json:"modalWithTitleAndButtonRenderer"`
											} `json:"modal"`
										} `json:"modalEndpoint"`
									} `json:"signInEndpoint"`
									SubscribedEntityKey  string `json:"subscribedEntityKey"`
									OnSubscribeEndpoints []struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												SendPost bool   `json:"sendPost"`
												ApiUrl   string `json:"apiUrl"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										SubscribeEndpoint struct {
											ChannelIds []string `json:"channelIds"`
											Params     string   `json:"params"`
										} `json:"subscribeEndpoint"`
									} `json:"onSubscribeEndpoints"`
									OnUnsubscribeEndpoints []struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												SendPost bool `json:"sendPost"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										SignalServiceEndpoint struct {
											Signal  string `json:"signal"`
											Actions []struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												OpenPopupAction     struct {
													Popup struct {
														ConfirmDialogRenderer struct {
															TrackingParams string `json:"trackingParams"`
															DialogMessages []struct {
																Runs []struct {
																	Text string `json:"text"`
																} `json:"runs"`
															} `json:"dialogMessages"`
															ConfirmButton struct {
																ButtonRenderer struct {
																	Style      string `json:"style"`
																	Size       string `json:"size"`
																	IsDisabled bool   `json:"isDisabled"`
																	Text       struct {
																		Runs []struct {
																			Text string `json:"text"`
																		} `json:"runs"`
																	} `json:"text"`
																	ServiceEndpoint struct {
																		ClickTrackingParams string `json:"clickTrackingParams"`
																		CommandMetadata     struct {
																			WebCommandMetadata struct {
																				SendPost bool   `json:"sendPost"`
																				ApiUrl   string `json:"apiUrl"`
																			} `json:"webCommandMetadata"`
																		} `json:"commandMetadata"`
																		UnsubscribeEndpoint struct {
																			ChannelIds []string `json:"channelIds"`
																			Params     string   `json:"params"`
																		} `json:"unsubscribeEndpoint"`
																	} `json:"serviceEndpoint"`
																	Accessibility struct {
																		Label string `json:"label"`
																	} `json:"accessibility"`
																	TrackingParams string `json:"trackingParams"`
																} `json:"buttonRenderer"`
															} `json:"confirmButton"`
															CancelButton struct {
																ButtonRenderer struct {
																	Style      string `json:"style"`
																	Size       string `json:"size"`
																	IsDisabled bool   `json:"isDisabled"`
																	Text       struct {
																		Runs []struct {
																			Text string `json:"text"`
																		} `json:"runs"`
																	} `json:"text"`
																	Accessibility struct {
																		Label string `json:"label"`
																	} `json:"accessibility"`
																	TrackingParams string `json:"trackingParams"`
																} `json:"buttonRenderer"`
															} `json:"cancelButton"`
															PrimaryIsCancel bool `json:"primaryIsCancel"`
														} `json:"confirmDialogRenderer"`
													} `json:"popup"`
													PopupType string `json:"popupType"`
												} `json:"openPopupAction"`
											} `json:"actions"`
										} `json:"signalServiceEndpoint"`
									} `json:"onUnsubscribeEndpoints"`
								} `json:"subscribeButtonRenderer"`
							} `json:"subscribeButton"`
							MetadataRowContainer struct {
								MetadataRowContainerRenderer struct {
									Rows []struct {
										RichMetadataRowRenderer struct {
											Contents []struct {
												RichMetadataRenderer struct {
													Style     string `json:"style"`
													Thumbnail struct {
														Thumbnails []struct {
															Url    string `json:"url"`
															Width  int    `json:"width"`
															Height int    `json:"height"`
														} `json:"thumbnails"`
													} `json:"thumbnail"`
													Title struct {
														SimpleText string `json:"simpleText,omitempty"`
														Runs       []struct {
															Text string `json:"text"`
														} `json:"runs,omitempty"`
													} `json:"title"`
													Subtitle struct {
														SimpleText string `json:"simpleText"`
													} `json:"subtitle,omitempty"`
													CallToAction struct {
														Runs []struct {
															Text string `json:"text"`
														} `json:"runs"`
													} `json:"callToAction"`
													CallToActionIcon struct {
														IconType string `json:"iconType"`
													} `json:"callToActionIcon"`
													Endpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																Url         string `json:"url"`
																WebPageType string `json:"webPageType"`
																RootVe      int    `json:"rootVe"`
																ApiUrl      string `json:"apiUrl"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														BrowseEndpoint struct {
															BrowseId string `json:"browseId"`
														} `json:"browseEndpoint"`
													} `json:"endpoint"`
													TrackingParams string `json:"trackingParams"`
												} `json:"richMetadataRenderer"`
											} `json:"contents"`
											TrackingParams string `json:"trackingParams"`
										} `json:"richMetadataRowRenderer"`
									} `json:"rows"`
									CollapsedItemCount int    `json:"collapsedItemCount"`
									TrackingParams     string `json:"trackingParams"`
								} `json:"metadataRowContainerRenderer"`
							} `json:"metadataRowContainer"`
							ShowMoreText struct {
								SimpleText string `json:"simpleText"`
							} `json:"showMoreText"`
							ShowLessText struct {
								SimpleText string `json:"simpleText"`
							} `json:"showLessText"`
							TrackingParams            string `json:"trackingParams"`
							DefaultExpanded           bool   `json:"defaultExpanded"`
							DescriptionCollapsedLines int    `json:"descriptionCollapsedLines"`
							ShowMoreCommand           struct {
								ClickTrackingParams    string `json:"clickTrackingParams"`
								CommandExecutorCommand struct {
									Commands []struct {
										ClickTrackingParams                   string `json:"clickTrackingParams"`
										ChangeEngagementPanelVisibilityAction struct {
											TargetId   string `json:"targetId"`
											Visibility string `json:"visibility"`
										} `json:"changeEngagementPanelVisibilityAction,omitempty"`
										ScrollToEngagementPanelCommand struct {
											TargetId string `json:"targetId"`
										} `json:"scrollToEngagementPanelCommand,omitempty"`
									} `json:"commands"`
								} `json:"commandExecutorCommand"`
							} `json:"showMoreCommand"`
							ShowLessCommand struct {
								ClickTrackingParams                   string `json:"clickTrackingParams"`
								ChangeEngagementPanelVisibilityAction struct {
									TargetId   string `json:"targetId"`
									Visibility string `json:"visibility"`
								} `json:"changeEngagementPanelVisibilityAction"`
							} `json:"showLessCommand"`
							AttributedDescription struct {
								Content     string `json:"content"`
								CommandRuns []struct {
									StartIndex int `json:"startIndex"`
									Length     int `json:"length"`
									OnTap      struct {
										InnertubeCommand struct {
											ClickTrackingParams string `json:"clickTrackingParams"`
											CommandMetadata     struct {
												WebCommandMetadata struct {
													Url         string `json:"url"`
													WebPageType string `json:"webPageType"`
													RootVe      int    `json:"rootVe"`
													ApiUrl      string `json:"apiUrl,omitempty"`
												} `json:"webCommandMetadata"`
											} `json:"commandMetadata"`
											UrlEndpoint struct {
												Url      string `json:"url"`
												Target   string `json:"target,omitempty"`
												Nofollow bool   `json:"nofollow"`
											} `json:"urlEndpoint,omitempty"`
											BrowseEndpoint struct {
												BrowseId         string `json:"browseId"`
												CanonicalBaseUrl string `json:"canonicalBaseUrl"`
											} `json:"browseEndpoint,omitempty"`
										} `json:"innertubeCommand"`
									} `json:"onTap"`
									LoggingDirectives struct {
										TrackingParams                string `json:"trackingParams"`
										EnableDisplayloggerExperiment bool   `json:"enableDisplayloggerExperiment"`
									} `json:"loggingDirectives,omitempty"`
								} `json:"commandRuns"`
								StyleRuns []struct {
									StartIndex         int `json:"startIndex"`
									Length             int `json:"length"`
									StyleRunExtensions struct {
										StyleRunColorMapExtension struct {
											ColorMap []struct {
												Key   string `json:"key"`
												Value int64  `json:"value"`
											} `json:"colorMap"`
										} `json:"styleRunColorMapExtension"`
									} `json:"styleRunExtensions"`
								} `json:"styleRuns"`
								AttachmentRuns []struct {
									StartIndex int `json:"startIndex"`
									Length     int `json:"length"`
									Element    struct {
										Type struct {
											ImageType struct {
												Image struct {
													Sources []struct {
														Url string `json:"url"`
													} `json:"sources"`
												} `json:"image"`
											} `json:"imageType"`
										} `json:"type"`
										Properties struct {
											LayoutProperties struct {
												Height struct {
													Value int    `json:"value"`
													Unit  string `json:"unit"`
												} `json:"height"`
												Width struct {
													Value int    `json:"value"`
													Unit  string `json:"unit"`
												} `json:"width"`
											} `json:"layoutProperties"`
										} `json:"properties"`
									} `json:"element"`
									Alignment string `json:"alignment"`
								} `json:"attachmentRuns"`
								DecorationRuns []struct {
									TextDecorator struct {
										HighlightTextDecorator struct {
											StartIndex                       int `json:"startIndex"`
											Length                           int `json:"length"`
											BackgroundCornerRadius           int `json:"backgroundCornerRadius"`
											HighlightTextDecoratorExtensions struct {
												HighlightTextDecoratorColorMapExtension struct {
													ColorMap []struct {
														Key   string `json:"key"`
														Value int    `json:"value"`
													} `json:"colorMap"`
												} `json:"highlightTextDecoratorColorMapExtension"`
											} `json:"highlightTextDecoratorExtensions"`
										} `json:"highlightTextDecorator"`
									} `json:"textDecorator"`
								} `json:"decorationRuns"`
							} `json:"attributedDescription"`
						} `json:"videoSecondaryInfoRenderer,omitempty"`
						ItemSectionRenderer struct {
							Contents []struct {
								CommentsEntryPointHeaderRenderer struct {
									HeaderText struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"headerText"`
									OnTap struct {
										ClickTrackingParams    string `json:"clickTrackingParams"`
										CommandExecutorCommand struct {
											Commands []struct {
												ClickTrackingParams                   string `json:"clickTrackingParams"`
												ChangeEngagementPanelVisibilityAction struct {
													TargetId   string `json:"targetId"`
													Visibility string `json:"visibility"`
												} `json:"changeEngagementPanelVisibilityAction,omitempty"`
												ScrollToEngagementPanelCommand struct {
													TargetId string `json:"targetId"`
												} `json:"scrollToEngagementPanelCommand,omitempty"`
											} `json:"commands"`
										} `json:"commandExecutorCommand"`
									} `json:"onTap"`
									TrackingParams string `json:"trackingParams"`
									CommentCount   struct {
										SimpleText string `json:"simpleText"`
									} `json:"commentCount"`
									ContentRenderer struct {
										CommentsEntryPointTeaserRenderer struct {
											TeaserAvatar struct {
												Thumbnails []struct {
													Url    string `json:"url"`
													Width  int    `json:"width"`
													Height int    `json:"height"`
												} `json:"thumbnails"`
												Accessibility struct {
													AccessibilityData struct {
														Label string `json:"label"`
													} `json:"accessibilityData"`
												} `json:"accessibility"`
											} `json:"teaserAvatar"`
											TeaserContent struct {
												SimpleText string `json:"simpleText"`
											} `json:"teaserContent"`
											TrackingParams string `json:"trackingParams"`
										} `json:"commentsEntryPointTeaserRenderer"`
									} `json:"contentRenderer"`
									TargetId string `json:"targetId"`
								} `json:"commentsEntryPointHeaderRenderer,omitempty"`
								ContinuationItemRenderer struct {
									Trigger              string `json:"trigger"`
									ContinuationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												SendPost bool   `json:"sendPost"`
												ApiUrl   string `json:"apiUrl"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										ContinuationCommand struct {
											Token   string `json:"token"`
											Request string `json:"request"`
										} `json:"continuationCommand"`
									} `json:"continuationEndpoint"`
								} `json:"continuationItemRenderer,omitempty"`
							} `json:"contents"`
							TrackingParams    string `json:"trackingParams"`
							SectionIdentifier string `json:"sectionIdentifier"`
							TargetId          string `json:"targetId,omitempty"`
						} `json:"itemSectionRenderer,omitempty"`
					} `json:"contents"`
					TrackingParams string `json:"trackingParams"`
				} `json:"results"`
			} `json:"results"`
			SecondaryResults struct {
				SecondaryResults struct {
					Results []struct {
						CompactVideoRenderer struct {
							VideoId   string `json:"videoId"`
							Thumbnail struct {
								Thumbnails []struct {
									Url    string `json:"url"`
									Width  int    `json:"width"`
									Height int    `json:"height"`
								} `json:"thumbnails"`
							} `json:"thumbnail"`
							Title struct {
								Accessibility struct {
									AccessibilityData struct {
										Label string `json:"label"`
									} `json:"accessibilityData"`
								} `json:"accessibility"`
								SimpleText string `json:"simpleText"`
							} `json:"title"`
							LongBylineText struct {
								Runs []struct {
									Text               string `json:"text"`
									NavigationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												Url         string `json:"url"`
												WebPageType string `json:"webPageType"`
												RootVe      int    `json:"rootVe"`
												ApiUrl      string `json:"apiUrl"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										BrowseEndpoint struct {
											BrowseId         string `json:"browseId"`
											CanonicalBaseUrl string `json:"canonicalBaseUrl"`
										} `json:"browseEndpoint"`
									} `json:"navigationEndpoint"`
								} `json:"runs"`
							} `json:"longBylineText"`
							PublishedTimeText struct {
								SimpleText string `json:"simpleText"`
							} `json:"publishedTimeText"`
							ViewCountText struct {
								SimpleText string `json:"simpleText"`
							} `json:"viewCountText"`
							LengthText struct {
								Accessibility struct {
									AccessibilityData struct {
										Label string `json:"label"`
									} `json:"accessibilityData"`
								} `json:"accessibility"`
								SimpleText string `json:"simpleText"`
							} `json:"lengthText"`
							NavigationEndpoint struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										Url         string `json:"url"`
										WebPageType string `json:"webPageType"`
										RootVe      int    `json:"rootVe"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								WatchEndpoint struct {
									VideoId                            string `json:"videoId"`
									Nofollow                           bool   `json:"nofollow"`
									WatchEndpointSupportedOnesieConfig struct {
										Html5PlaybackOnesieConfig struct {
											CommonConfig struct {
												Url string `json:"url"`
											} `json:"commonConfig"`
										} `json:"html5PlaybackOnesieConfig"`
									} `json:"watchEndpointSupportedOnesieConfig"`
								} `json:"watchEndpoint"`
							} `json:"navigationEndpoint"`
							ShortBylineText struct {
								Runs []struct {
									Text               string `json:"text"`
									NavigationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												Url         string `json:"url"`
												WebPageType string `json:"webPageType"`
												RootVe      int    `json:"rootVe"`
												ApiUrl      string `json:"apiUrl"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										BrowseEndpoint struct {
											BrowseId         string `json:"browseId"`
											CanonicalBaseUrl string `json:"canonicalBaseUrl"`
										} `json:"browseEndpoint"`
									} `json:"navigationEndpoint"`
								} `json:"runs"`
							} `json:"shortBylineText"`
							ChannelThumbnail struct {
								Thumbnails []struct {
									Url    string `json:"url"`
									Width  int    `json:"width"`
									Height int    `json:"height"`
								} `json:"thumbnails"`
							} `json:"channelThumbnail"`
							OwnerBadges []struct {
								MetadataBadgeRenderer struct {
									Icon struct {
										IconType string `json:"iconType"`
									} `json:"icon"`
									Style             string `json:"style"`
									Tooltip           string `json:"tooltip"`
									TrackingParams    string `json:"trackingParams"`
									AccessibilityData struct {
										Label string `json:"label"`
									} `json:"accessibilityData"`
								} `json:"metadataBadgeRenderer"`
							} `json:"ownerBadges"`
							TrackingParams     string `json:"trackingParams"`
							ShortViewCountText struct {
								Accessibility struct {
									AccessibilityData struct {
										Label string `json:"label"`
									} `json:"accessibilityData"`
								} `json:"accessibility"`
								SimpleText string `json:"simpleText"`
							} `json:"shortViewCountText"`
							Menu struct {
								MenuRenderer struct {
									Items []struct {
										MenuServiceItemRenderer struct {
											Text struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"text"`
											Icon struct {
												IconType string `json:"iconType"`
											} `json:"icon"`
											ServiceEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														SendPost bool   `json:"sendPost"`
														ApiUrl   string `json:"apiUrl,omitempty"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												SignalServiceEndpoint struct {
													Signal  string `json:"signal"`
													Actions []struct {
														ClickTrackingParams  string `json:"clickTrackingParams"`
														AddToPlaylistCommand struct {
															OpenMiniplayer      bool   `json:"openMiniplayer"`
															OpenListPanel       bool   `json:"openListPanel"`
															VideoId             string `json:"videoId"`
															ListType            string `json:"listType"`
															OnCreateListCommand struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		SendPost bool   `json:"sendPost"`
																		ApiUrl   string `json:"apiUrl"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																CreatePlaylistServiceEndpoint struct {
																	VideoIds []string `json:"videoIds"`
																	Params   string   `json:"params"`
																} `json:"createPlaylistServiceEndpoint"`
															} `json:"onCreateListCommand"`
															VideoIds []string `json:"videoIds"`
														} `json:"addToPlaylistCommand,omitempty"`
														OpenPopupAction struct {
															Popup struct {
																NotificationActionRenderer struct {
																	ResponseText struct {
																		SimpleText string `json:"simpleText"`
																	} `json:"responseText"`
																	TrackingParams string `json:"trackingParams"`
																} `json:"notificationActionRenderer"`
															} `json:"popup"`
															PopupType string `json:"popupType"`
														} `json:"openPopupAction,omitempty"`
													} `json:"actions"`
												} `json:"signalServiceEndpoint,omitempty"`
												ShareEntityServiceEndpoint struct {
													SerializedShareEntity string `json:"serializedShareEntity"`
													Commands              []struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														OpenPopupAction     struct {
															Popup struct {
																UnifiedSharePanelRenderer struct {
																	TrackingParams     string `json:"trackingParams"`
																	ShowLoadingSpinner bool   `json:"showLoadingSpinner"`
																} `json:"unifiedSharePanelRenderer"`
															} `json:"popup"`
															PopupType string `json:"popupType"`
															BeReused  bool   `json:"beReused"`
														} `json:"openPopupAction"`
													} `json:"commands"`
												} `json:"shareEntityServiceEndpoint,omitempty"`
											} `json:"serviceEndpoint"`
											TrackingParams string `json:"trackingParams"`
											HasSeparator   bool   `json:"hasSeparator,omitempty"`
										} `json:"menuServiceItemRenderer"`
									} `json:"items"`
									TrackingParams string `json:"trackingParams"`
									Accessibility  struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"accessibility"`
									TargetId string `json:"targetId,omitempty"`
								} `json:"menuRenderer"`
							} `json:"menu"`
							ThumbnailOverlays []struct {
								ThumbnailOverlayTimeStatusRenderer struct {
									Text struct {
										Accessibility struct {
											AccessibilityData struct {
												Label string `json:"label"`
											} `json:"accessibilityData"`
										} `json:"accessibility"`
										SimpleText string `json:"simpleText"`
									} `json:"text"`
									Style string `json:"style"`
								} `json:"thumbnailOverlayTimeStatusRenderer,omitempty"`
								ThumbnailOverlayToggleButtonRenderer struct {
									IsToggled     bool `json:"isToggled,omitempty"`
									UntoggledIcon struct {
										IconType string `json:"iconType"`
									} `json:"untoggledIcon"`
									ToggledIcon struct {
										IconType string `json:"iconType"`
									} `json:"toggledIcon"`
									UntoggledTooltip         string `json:"untoggledTooltip"`
									ToggledTooltip           string `json:"toggledTooltip"`
									UntoggledServiceEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												SendPost bool   `json:"sendPost"`
												ApiUrl   string `json:"apiUrl,omitempty"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										PlaylistEditEndpoint struct {
											PlaylistId string `json:"playlistId"`
											Actions    []struct {
												AddedVideoId string `json:"addedVideoId"`
												Action       string `json:"action"`
											} `json:"actions"`
										} `json:"playlistEditEndpoint,omitempty"`
										SignalServiceEndpoint struct {
											Signal  string `json:"signal"`
											Actions []struct {
												ClickTrackingParams  string `json:"clickTrackingParams"`
												AddToPlaylistCommand struct {
													OpenMiniplayer      bool   `json:"openMiniplayer"`
													OpenListPanel       bool   `json:"openListPanel"`
													VideoId             string `json:"videoId"`
													ListType            string `json:"listType"`
													OnCreateListCommand struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																SendPost bool   `json:"sendPost"`
																ApiUrl   string `json:"apiUrl"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														CreatePlaylistServiceEndpoint struct {
															VideoIds []string `json:"videoIds"`
															Params   string   `json:"params"`
														} `json:"createPlaylistServiceEndpoint"`
													} `json:"onCreateListCommand"`
													VideoIds []string `json:"videoIds"`
												} `json:"addToPlaylistCommand"`
											} `json:"actions"`
										} `json:"signalServiceEndpoint,omitempty"`
									} `json:"untoggledServiceEndpoint"`
									ToggledServiceEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												SendPost bool   `json:"sendPost"`
												ApiUrl   string `json:"apiUrl"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										PlaylistEditEndpoint struct {
											PlaylistId string `json:"playlistId"`
											Actions    []struct {
												Action         string `json:"action"`
												RemovedVideoId string `json:"removedVideoId"`
											} `json:"actions"`
										} `json:"playlistEditEndpoint"`
									} `json:"toggledServiceEndpoint,omitempty"`
									UntoggledAccessibility struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"untoggledAccessibility"`
									ToggledAccessibility struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"toggledAccessibility"`
									TrackingParams string `json:"trackingParams"`
								} `json:"thumbnailOverlayToggleButtonRenderer,omitempty"`
								ThumbnailOverlayNowPlayingRenderer struct {
									Text struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"text"`
								} `json:"thumbnailOverlayNowPlayingRenderer,omitempty"`
							} `json:"thumbnailOverlays"`
							Accessibility struct {
								AccessibilityData struct {
									Label string `json:"label"`
								} `json:"accessibilityData"`
							} `json:"accessibility"`
							Badges []struct {
								MetadataBadgeRenderer struct {
									Style          string `json:"style"`
									Label          string `json:"label"`
									TrackingParams string `json:"trackingParams"`
								} `json:"metadataBadgeRenderer"`
							} `json:"badges,omitempty"`
						} `json:"compactVideoRenderer,omitempty"`
						ContinuationItemRenderer struct {
							Trigger              string `json:"trigger"`
							ContinuationEndpoint struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										SendPost bool   `json:"sendPost"`
										ApiUrl   string `json:"apiUrl"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								ContinuationCommand struct {
									Token   string `json:"token"`
									Request string `json:"request"`
								} `json:"continuationCommand"`
							} `json:"continuationEndpoint"`
							Button struct {
								ButtonRenderer struct {
									Style      string `json:"style"`
									Size       string `json:"size"`
									IsDisabled bool   `json:"isDisabled"`
									Text       struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"text"`
									TrackingParams string `json:"trackingParams"`
									Command        struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												SendPost bool   `json:"sendPost"`
												ApiUrl   string `json:"apiUrl"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										ContinuationCommand struct {
											Token   string `json:"token"`
											Request string `json:"request"`
										} `json:"continuationCommand"`
									} `json:"command"`
								} `json:"buttonRenderer"`
							} `json:"button"`
						} `json:"continuationItemRenderer,omitempty"`
					} `json:"results"`
					TrackingParams string `json:"trackingParams"`
					TargetId       string `json:"targetId"`
				} `json:"secondaryResults"`
			} `json:"secondaryResults"`
			Autoplay struct {
				Autoplay struct {
					Sets []struct {
						Mode          string `json:"mode"`
						AutoplayVideo struct {
							ClickTrackingParams string `json:"clickTrackingParams"`
							CommandMetadata     struct {
								WebCommandMetadata struct {
									Url         string `json:"url"`
									WebPageType string `json:"webPageType"`
									RootVe      int    `json:"rootVe"`
								} `json:"webCommandMetadata"`
							} `json:"commandMetadata"`
							WatchEndpoint struct {
								VideoId                              string `json:"videoId"`
								Params                               string `json:"params"`
								PlayerParams                         string `json:"playerParams"`
								WatchEndpointSupportedPrefetchConfig struct {
									PrefetchHintConfig struct {
										PrefetchPriority                            int `json:"prefetchPriority"`
										CountdownUiRelativeSecondsPrefetchCondition int `json:"countdownUiRelativeSecondsPrefetchCondition"`
									} `json:"prefetchHintConfig"`
								} `json:"watchEndpointSupportedPrefetchConfig"`
							} `json:"watchEndpoint"`
						} `json:"autoplayVideo"`
					} `json:"sets"`
					CountDownSecs  int    `json:"countDownSecs"`
					TrackingParams string `json:"trackingParams"`
				} `json:"autoplay"`
			} `json:"autoplay"`
		} `json:"twoColumnWatchNextResults,omitempty"`
	} `json:"contents"`
	Header struct {
		C4TabbedHeaderRenderer struct {
			ChannelId          string `json:"channelId"`
			Title              string `json:"title"`
			NavigationEndpoint struct {
				ClickTrackingParams string `json:"clickTrackingParams"`
				CommandMetadata     struct {
					WebCommandMetadata struct {
						Url         string `json:"url"`
						WebPageType string `json:"webPageType"`
						RootVe      int    `json:"rootVe"`
						ApiUrl      string `json:"apiUrl"`
					} `json:"webCommandMetadata"`
				} `json:"commandMetadata"`
				BrowseEndpoint struct {
					BrowseId         string `json:"browseId"`
					CanonicalBaseUrl string `json:"canonicalBaseUrl"`
				} `json:"browseEndpoint"`
			} `json:"navigationEndpoint"`
			Avatar struct {
				Thumbnails []struct {
					Url    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"thumbnails"`
			} `json:"avatar"`
			Banner struct {
				Thumbnails []struct {
					Url    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"thumbnails"`
			} `json:"banner"`
			Badges []struct {
				MetadataBadgeRenderer struct {
					Icon struct {
						IconType string `json:"iconType"`
					} `json:"icon"`
					Style             string `json:"style"`
					Tooltip           string `json:"tooltip"`
					TrackingParams    string `json:"trackingParams"`
					AccessibilityData struct {
						Label string `json:"label"`
					} `json:"accessibilityData"`
				} `json:"metadataBadgeRenderer"`
			} `json:"badges"`
			HeaderLinks struct {
				ChannelHeaderLinksRenderer struct {
					PrimaryLinks []struct {
						NavigationEndpoint struct {
							ClickTrackingParams string `json:"clickTrackingParams"`
							CommandMetadata     struct {
								WebCommandMetadata struct {
									Url         string `json:"url"`
									WebPageType string `json:"webPageType"`
									RootVe      int    `json:"rootVe"`
								} `json:"webCommandMetadata"`
							} `json:"commandMetadata"`
							UrlEndpoint struct {
								Url      string `json:"url"`
								Target   string `json:"target"`
								Nofollow bool   `json:"nofollow"`
							} `json:"urlEndpoint"`
						} `json:"navigationEndpoint"`
						Icon struct {
							Thumbnails []struct {
								Url string `json:"url"`
							} `json:"thumbnails"`
						} `json:"icon"`
						Title struct {
							SimpleText string `json:"simpleText"`
						} `json:"title"`
					} `json:"primaryLinks"`
					SecondaryLinks []struct {
						NavigationEndpoint struct {
							ClickTrackingParams string `json:"clickTrackingParams"`
							CommandMetadata     struct {
								WebCommandMetadata struct {
									Url         string `json:"url"`
									WebPageType string `json:"webPageType"`
									RootVe      int    `json:"rootVe"`
								} `json:"webCommandMetadata"`
							} `json:"commandMetadata"`
							UrlEndpoint struct {
								Url      string `json:"url"`
								Target   string `json:"target"`
								Nofollow bool   `json:"nofollow"`
							} `json:"urlEndpoint"`
						} `json:"navigationEndpoint"`
						Icon struct {
							Thumbnails []struct {
								Url string `json:"url"`
							} `json:"thumbnails"`
						} `json:"icon"`
						Title struct {
							SimpleText string `json:"simpleText"`
						} `json:"title"`
					} `json:"secondaryLinks"`
				} `json:"channelHeaderLinksRenderer"`
			} `json:"headerLinks"`
			SubscribeButton struct {
				ButtonRenderer struct {
					Style      string `json:"style"`
					Size       string `json:"size"`
					IsDisabled bool   `json:"isDisabled"`
					Text       struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"text"`
					NavigationEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								IgnoreNavigation bool `json:"ignoreNavigation"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						ModalEndpoint struct {
							Modal struct {
								ModalWithTitleAndButtonRenderer struct {
									Title struct {
										SimpleText string `json:"simpleText"`
									} `json:"title"`
									Content struct {
										SimpleText string `json:"simpleText"`
									} `json:"content"`
									Button struct {
										ButtonRenderer struct {
											Style      string `json:"style"`
											Size       string `json:"size"`
											IsDisabled bool   `json:"isDisabled"`
											Text       struct {
												SimpleText string `json:"simpleText"`
											} `json:"text"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														Url         string `json:"url"`
														WebPageType string `json:"webPageType"`
														RootVe      int    `json:"rootVe"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												SignInEndpoint struct {
													NextEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																Url         string `json:"url"`
																WebPageType string `json:"webPageType"`
																RootVe      int    `json:"rootVe"`
																ApiUrl      string `json:"apiUrl"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														BrowseEndpoint struct {
															BrowseId         string `json:"browseId"`
															Params           string `json:"params"`
															CanonicalBaseUrl string `json:"canonicalBaseUrl"`
														} `json:"browseEndpoint"`
													} `json:"nextEndpoint"`
													ContinueAction string `json:"continueAction"`
													IdamTag        string `json:"idamTag"`
												} `json:"signInEndpoint"`
											} `json:"navigationEndpoint"`
											TrackingParams string `json:"trackingParams"`
										} `json:"buttonRenderer"`
									} `json:"button"`
								} `json:"modalWithTitleAndButtonRenderer"`
							} `json:"modal"`
						} `json:"modalEndpoint"`
					} `json:"navigationEndpoint"`
					TrackingParams string `json:"trackingParams"`
				} `json:"buttonRenderer"`
			} `json:"subscribeButton"`
			SubscriberCountText struct {
				Accessibility struct {
					AccessibilityData struct {
						Label string `json:"label"`
					} `json:"accessibilityData"`
				} `json:"accessibility"`
				SimpleText string `json:"simpleText"`
			} `json:"subscriberCountText"`
			TvBanner struct {
				Thumbnails []struct {
					Url    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"thumbnails"`
			} `json:"tvBanner"`
			MobileBanner struct {
				Thumbnails []struct {
					Url    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"thumbnails"`
			} `json:"mobileBanner"`
			TrackingParams    string `json:"trackingParams"`
			ChannelHandleText struct {
				Runs []struct {
					Text string `json:"text"`
				} `json:"runs"`
			} `json:"channelHandleText"`
			Style           string `json:"style"`
			VideosCountText struct {
				Runs []struct {
					Text string `json:"text"`
				} `json:"runs"`
			} `json:"videosCountText"`
			Tagline struct {
				ChannelTaglineRenderer struct {
					Content      string `json:"content"`
					MaxLines     int    `json:"maxLines"`
					MoreLabel    string `json:"moreLabel"`
					MoreEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								Url         string `json:"url"`
								WebPageType string `json:"webPageType"`
								RootVe      int    `json:"rootVe"`
								ApiUrl      string `json:"apiUrl"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						BrowseEndpoint struct {
							BrowseId         string `json:"browseId"`
							Params           string `json:"params"`
							CanonicalBaseUrl string `json:"canonicalBaseUrl"`
						} `json:"browseEndpoint"`
					} `json:"moreEndpoint"`
					MoreIcon struct {
						IconType string `json:"iconType"`
					} `json:"moreIcon"`
				} `json:"channelTaglineRenderer"`
			} `json:"tagline"`
		} `json:"c4TabbedHeaderRenderer,omitempty"`
		FeedTabbedHeaderRenderer struct {
			Title struct {
				Runs []struct {
					Text string `json:"text"`
				} `json:"runs"`
			} `json:"title"`
		} `json:"feedTabbedHeaderRenderer,omitempty"`
	} `json:"header,omitempty"`
	Metadata struct {
		ChannelMetadataRenderer struct {
			Title       string   `json:"title"`
			Description string   `json:"description"`
			RssUrl      string   `json:"rssUrl"`
			ExternalId  string   `json:"externalId"`
			Keywords    string   `json:"keywords"`
			OwnerUrls   []string `json:"ownerUrls"`
			Avatar      struct {
				Thumbnails []struct {
					Url    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"thumbnails"`
			} `json:"avatar"`
			ChannelUrl             string   `json:"channelUrl"`
			IsFamilySafe           bool     `json:"isFamilySafe"`
			FacebookProfileId      string   `json:"facebookProfileId"`
			AvailableCountryCodes  []string `json:"availableCountryCodes"`
			AndroidDeepLink        string   `json:"androidDeepLink"`
			AndroidAppindexingLink string   `json:"androidAppindexingLink"`
			IosAppindexingLink     string   `json:"iosAppindexingLink"`
			VanityChannelUrl       string   `json:"vanityChannelUrl"`
		} `json:"channelMetadataRenderer"`
	} `json:"metadata,omitempty"`
	TrackingParams string `json:"trackingParams"`
	Topbar         struct {
		DesktopTopbarRenderer struct {
			Logo struct {
				TopbarLogoRenderer struct {
					IconImage struct {
						IconType string `json:"iconType"`
					} `json:"iconImage"`
					TooltipText struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"tooltipText"`
					Endpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								Url         string `json:"url"`
								WebPageType string `json:"webPageType"`
								RootVe      int    `json:"rootVe"`
								ApiUrl      string `json:"apiUrl"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						BrowseEndpoint struct {
							BrowseId string `json:"browseId"`
						} `json:"browseEndpoint"`
					} `json:"endpoint"`
					TrackingParams    string `json:"trackingParams"`
					OverrideEntityKey string `json:"overrideEntityKey"`
				} `json:"topbarLogoRenderer"`
			} `json:"logo"`
			Searchbox struct {
				FusionSearchboxRenderer struct {
					Icon struct {
						IconType string `json:"iconType"`
					} `json:"icon"`
					PlaceholderText struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"placeholderText"`
					Config struct {
						WebSearchboxConfig struct {
							RequestLanguage     string `json:"requestLanguage"`
							RequestDomain       string `json:"requestDomain"`
							HasOnscreenKeyboard bool   `json:"hasOnscreenKeyboard"`
							FocusSearchbox      bool   `json:"focusSearchbox"`
						} `json:"webSearchboxConfig"`
					} `json:"config"`
					TrackingParams string `json:"trackingParams"`
					SearchEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								Url         string `json:"url"`
								WebPageType string `json:"webPageType"`
								RootVe      int    `json:"rootVe"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SearchEndpoint struct {
							Query string `json:"query"`
						} `json:"searchEndpoint"`
					} `json:"searchEndpoint"`
					ClearButton struct {
						ButtonRenderer struct {
							Style      string `json:"style"`
							Size       string `json:"size"`
							IsDisabled bool   `json:"isDisabled"`
							Icon       struct {
								IconType string `json:"iconType"`
							} `json:"icon"`
							TrackingParams    string `json:"trackingParams"`
							AccessibilityData struct {
								AccessibilityData struct {
									Label string `json:"label"`
								} `json:"accessibilityData"`
							} `json:"accessibilityData"`
						} `json:"buttonRenderer"`
					} `json:"clearButton"`
				} `json:"fusionSearchboxRenderer"`
			} `json:"searchbox"`
			TrackingParams string `json:"trackingParams"`
			CountryCode    string `json:"countryCode"`
			TopbarButtons  []struct {
				TopbarMenuButtonRenderer struct {
					Icon struct {
						IconType string `json:"iconType"`
					} `json:"icon"`
					MenuRequest struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool   `json:"sendPost"`
								ApiUrl   string `json:"apiUrl"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SignalServiceEndpoint struct {
							Signal  string `json:"signal"`
							Actions []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								OpenPopupAction     struct {
									Popup struct {
										MultiPageMenuRenderer struct {
											TrackingParams     string `json:"trackingParams"`
											Style              string `json:"style"`
											ShowLoadingSpinner bool   `json:"showLoadingSpinner"`
										} `json:"multiPageMenuRenderer"`
									} `json:"popup"`
									PopupType string `json:"popupType"`
									BeReused  bool   `json:"beReused"`
								} `json:"openPopupAction"`
							} `json:"actions"`
						} `json:"signalServiceEndpoint"`
					} `json:"menuRequest"`
					TrackingParams string `json:"trackingParams"`
					Accessibility  struct {
						AccessibilityData struct {
							Label string `json:"label"`
						} `json:"accessibilityData"`
					} `json:"accessibility"`
					Tooltip string `json:"tooltip"`
					Style   string `json:"style"`
				} `json:"topbarMenuButtonRenderer,omitempty"`
				ButtonRenderer struct {
					Style string `json:"style"`
					Size  string `json:"size"`
					Text  struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"text"`
					Icon struct {
						IconType string `json:"iconType"`
					} `json:"icon"`
					NavigationEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								Url         string `json:"url"`
								WebPageType string `json:"webPageType"`
								RootVe      int    `json:"rootVe"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SignInEndpoint struct {
							IdamTag string `json:"idamTag"`
						} `json:"signInEndpoint"`
					} `json:"navigationEndpoint"`
					TrackingParams string `json:"trackingParams"`
					TargetId       string `json:"targetId"`
				} `json:"buttonRenderer,omitempty"`
			} `json:"topbarButtons"`
			HotkeyDialog struct {
				HotkeyDialogRenderer struct {
					Title struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"title"`
					Sections []struct {
						HotkeyDialogSectionRenderer struct {
							Title struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"title"`
							Options []struct {
								HotkeyDialogSectionOptionRenderer struct {
									Label struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"label"`
									Hotkey                   string `json:"hotkey"`
									HotkeyAccessibilityLabel struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"hotkeyAccessibilityLabel,omitempty"`
								} `json:"hotkeyDialogSectionOptionRenderer"`
							} `json:"options"`
						} `json:"hotkeyDialogSectionRenderer"`
					} `json:"sections"`
					DismissButton struct {
						ButtonRenderer struct {
							Style      string `json:"style"`
							Size       string `json:"size"`
							IsDisabled bool   `json:"isDisabled"`
							Text       struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"text"`
							TrackingParams string `json:"trackingParams"`
						} `json:"buttonRenderer"`
					} `json:"dismissButton"`
					TrackingParams string `json:"trackingParams"`
				} `json:"hotkeyDialogRenderer"`
			} `json:"hotkeyDialog"`
			BackButton struct {
				ButtonRenderer struct {
					TrackingParams string `json:"trackingParams"`
					Command        struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool `json:"sendPost"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SignalServiceEndpoint struct {
							Signal  string `json:"signal"`
							Actions []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								SignalAction        struct {
									Signal string `json:"signal"`
								} `json:"signalAction"`
							} `json:"actions"`
						} `json:"signalServiceEndpoint"`
					} `json:"command"`
				} `json:"buttonRenderer"`
			} `json:"backButton"`
			ForwardButton struct {
				ButtonRenderer struct {
					TrackingParams string `json:"trackingParams"`
					Command        struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool `json:"sendPost"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SignalServiceEndpoint struct {
							Signal  string `json:"signal"`
							Actions []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								SignalAction        struct {
									Signal string `json:"signal"`
								} `json:"signalAction"`
							} `json:"actions"`
						} `json:"signalServiceEndpoint"`
					} `json:"command"`
				} `json:"buttonRenderer"`
			} `json:"forwardButton"`
			A11YSkipNavigationButton struct {
				ButtonRenderer struct {
					Style      string `json:"style"`
					Size       string `json:"size"`
					IsDisabled bool   `json:"isDisabled"`
					Text       struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"text"`
					TrackingParams string `json:"trackingParams"`
					Command        struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool `json:"sendPost"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SignalServiceEndpoint struct {
							Signal  string `json:"signal"`
							Actions []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								SignalAction        struct {
									Signal string `json:"signal"`
								} `json:"signalAction"`
							} `json:"actions"`
						} `json:"signalServiceEndpoint"`
					} `json:"command"`
				} `json:"buttonRenderer"`
			} `json:"a11ySkipNavigationButton"`
			Interstitial struct {
				ConsentBumpV2Renderer struct {
					InterstitialLogoAside struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"interstitialLogoAside"`
					LanguagePickerButton struct {
						ButtonRenderer struct {
							Style      string `json:"style"`
							Size       string `json:"size"`
							IsDisabled bool   `json:"isDisabled"`
							Text       struct {
								SimpleText string `json:"simpleText"`
							} `json:"text"`
							Icon struct {
								IconType string `json:"iconType"`
							} `json:"icon"`
							Accessibility struct {
								Label string `json:"label"`
							} `json:"accessibility"`
							TrackingParams string `json:"trackingParams"`
						} `json:"buttonRenderer"`
					} `json:"languagePickerButton"`
					InterstitialTitle struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"interstitialTitle"`
					CustomizeButton struct {
						ButtonRenderer struct {
							Style      string `json:"style"`
							Size       string `json:"size"`
							IsDisabled bool   `json:"isDisabled"`
							Text       struct {
								SimpleText string `json:"simpleText"`
							} `json:"text"`
							TrackingParams string `json:"trackingParams"`
							Command        struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										Url         string `json:"url"`
										WebPageType string `json:"webPageType"`
										RootVe      int    `json:"rootVe"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								UrlEndpoint struct {
									Url string `json:"url"`
								} `json:"urlEndpoint"`
							} `json:"command"`
						} `json:"buttonRenderer"`
					} `json:"customizeButton"`
					AgreeButton struct {
						ButtonRenderer struct {
							Style      string `json:"style"`
							Size       string `json:"size"`
							IsDisabled bool   `json:"isDisabled"`
							Text       struct {
								SimpleText string `json:"simpleText"`
							} `json:"text"`
							Accessibility struct {
								Label string `json:"label"`
							} `json:"accessibility"`
							TrackingParams string `json:"trackingParams"`
							Command        struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								SaveConsentAction   struct {
									SocsCookie        string `json:"socsCookie"`
									SavePreferenceUrl string `json:"savePreferenceUrl"`
								} `json:"saveConsentAction"`
							} `json:"command"`
						} `json:"buttonRenderer"`
					} `json:"agreeButton"`
					PrivacyLink struct {
						Runs []struct {
							Text               string `json:"text"`
							NavigationEndpoint struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										Url         string `json:"url"`
										WebPageType string `json:"webPageType"`
										RootVe      int    `json:"rootVe"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								UrlEndpoint struct {
									Url string `json:"url"`
								} `json:"urlEndpoint"`
							} `json:"navigationEndpoint"`
						} `json:"runs"`
					} `json:"privacyLink"`
					TermsLink struct {
						Runs []struct {
							Text               string `json:"text"`
							NavigationEndpoint struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										Url         string `json:"url"`
										WebPageType string `json:"webPageType"`
										RootVe      int    `json:"rootVe"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								UrlEndpoint struct {
									Url string `json:"url"`
								} `json:"urlEndpoint"`
							} `json:"navigationEndpoint"`
						} `json:"runs"`
					} `json:"termsLink"`
					TrackingParams string `json:"trackingParams"`
					SignInButton   struct {
						ButtonRenderer struct {
							Style      string `json:"style"`
							Size       string `json:"size"`
							IsDisabled bool   `json:"isDisabled"`
							Text       struct {
								SimpleText string `json:"simpleText"`
							} `json:"text"`
							Icon struct {
								IconType string `json:"iconType"`
							} `json:"icon"`
							Tooltip        string `json:"tooltip"`
							TrackingParams string `json:"trackingParams"`
							Command        struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										Url         string `json:"url"`
										WebPageType string `json:"webPageType"`
										RootVe      int    `json:"rootVe"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								SignInEndpoint struct {
									Hack bool `json:"hack"`
								} `json:"signInEndpoint"`
							} `json:"command"`
						} `json:"buttonRenderer"`
					} `json:"signInButton"`
					LanguageList struct {
						DropdownRenderer struct {
							Entries []struct {
								DropdownItemRenderer struct {
									Label struct {
										SimpleText string `json:"simpleText"`
									} `json:"label"`
									IsSelected      bool   `json:"isSelected"`
									StringValue     string `json:"stringValue"`
									OnSelectCommand struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												SendPost bool `json:"sendPost"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										SignalServiceEndpoint struct {
											Signal  string `json:"signal"`
											Actions []struct {
												ClickTrackingParams   string `json:"clickTrackingParams"`
												SelectLanguageCommand struct {
													Hl string `json:"hl"`
												} `json:"selectLanguageCommand"`
											} `json:"actions"`
										} `json:"signalServiceEndpoint"`
									} `json:"onSelectCommand"`
								} `json:"dropdownItemRenderer"`
							} `json:"entries"`
							Accessibility struct {
								Label string `json:"label"`
							} `json:"accessibility"`
						} `json:"dropdownRenderer"`
					} `json:"languageList"`
					ReadMoreButton struct {
						ButtonRenderer struct {
							Style      string `json:"style"`
							Size       string `json:"size"`
							IsDisabled bool   `json:"isDisabled"`
							Text       struct {
								SimpleText string `json:"simpleText"`
							} `json:"text"`
							Icon struct {
								IconType string `json:"iconType"`
							} `json:"icon"`
							TrackingParams string `json:"trackingParams"`
							IconPosition   string `json:"iconPosition"`
						} `json:"buttonRenderer"`
					} `json:"readMoreButton"`
					DisableP13NButton struct {
						ButtonRenderer struct {
							Style      string `json:"style"`
							Size       string `json:"size"`
							IsDisabled bool   `json:"isDisabled"`
							Text       struct {
								SimpleText string `json:"simpleText"`
							} `json:"text"`
							TrackingParams    string `json:"trackingParams"`
							AccessibilityData struct {
								AccessibilityData struct {
									Label string `json:"label"`
								} `json:"accessibilityData"`
							} `json:"accessibilityData"`
							Command struct {
								ClickTrackingParams          string `json:"clickTrackingParams"`
								DisablePersonalizationAction struct {
									SocsCookie        string `json:"socsCookie"`
									SavePreferenceUrl string `json:"savePreferenceUrl"`
								} `json:"disablePersonalizationAction"`
							} `json:"command"`
						} `json:"buttonRenderer"`
					} `json:"disableP13nButton"`
					LoadingMessage struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"loadingMessage"`
					ErrorMessage struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"errorMessage"`
					EomV1Text struct {
						EssentialCookieMsg struct {
							Begin struct {
								Runs []struct {
									Text               string `json:"text"`
									NavigationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												Url         string `json:"url"`
												WebPageType string `json:"webPageType"`
												RootVe      int    `json:"rootVe"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										UrlEndpoint struct {
											Url string `json:"url"`
										} `json:"urlEndpoint"`
									} `json:"navigationEndpoint,omitempty"`
								} `json:"runs"`
							} `json:"begin"`
							Items []struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"items"`
						} `json:"essentialCookieMsg"`
						NonEssentialCookieMsg struct {
							Begin struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"begin"`
							Items []struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"items"`
						} `json:"nonEssentialCookieMsg"`
						IfReject struct {
							Runs []struct {
								Text string `json:"text"`
							} `json:"runs"`
						} `json:"ifReject"`
						Personalization struct {
							Runs []struct {
								Text string `json:"text"`
							} `json:"runs"`
						} `json:"personalization"`
						MoreOptions struct {
							Runs []struct {
								Text string `json:"text"`
							} `json:"runs"`
						} `json:"moreOptions"`
					} `json:"eomV1Text"`
				} `json:"consentBumpV2Renderer"`
			} `json:"interstitial,omitempty"`
		} `json:"desktopTopbarRenderer"`
	} `json:"topbar"`
	Microformat struct {
		MicroformatDataRenderer struct {
			UrlCanonical string `json:"urlCanonical"`
			Title        string `json:"title"`
			Description  string `json:"description"`
			Thumbnail    struct {
				Thumbnails []struct {
					Url    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"thumbnails"`
			} `json:"thumbnail"`
			SiteName           string   `json:"siteName"`
			AppName            string   `json:"appName"`
			AndroidPackage     string   `json:"androidPackage"`
			IosAppStoreId      string   `json:"iosAppStoreId"`
			IosAppArguments    string   `json:"iosAppArguments"`
			OgType             string   `json:"ogType"`
			UrlApplinksWeb     string   `json:"urlApplinksWeb"`
			UrlApplinksIos     string   `json:"urlApplinksIos"`
			UrlApplinksAndroid string   `json:"urlApplinksAndroid"`
			UrlTwitterIos      string   `json:"urlTwitterIos"`
			UrlTwitterAndroid  string   `json:"urlTwitterAndroid"`
			TwitterCardType    string   `json:"twitterCardType"`
			TwitterSiteHandle  string   `json:"twitterSiteHandle"`
			SchemaDotOrgType   string   `json:"schemaDotOrgType"`
			Noindex            bool     `json:"noindex"`
			Unlisted           bool     `json:"unlisted"`
			FamilySafe         bool     `json:"familySafe"`
			AvailableCountries []string `json:"availableCountries"`
			LinkAlternates     []struct {
				HrefUrl string `json:"hrefUrl"`
			} `json:"linkAlternates"`
		} `json:"microformatDataRenderer"`
	} `json:"microformat,omitempty"`
	FrameworkUpdates struct {
		EntityBatchUpdate struct {
			Mutations []struct {
				EntityKey string `json:"entityKey"`
				Type      string `json:"type"`
				Options   struct {
					PersistenceOption string `json:"persistenceOption"`
				} `json:"options,omitempty"`
				Payload struct {
					SubscriptionStateEntity struct {
						Key        string `json:"key"`
						Subscribed bool   `json:"subscribed"`
					} `json:"subscriptionStateEntity,omitempty"`
					TranscriptTrackSelectionEntity struct {
						Key                string `json:"key"`
						SelectedTrackIndex int    `json:"selectedTrackIndex"`
						SerializedParams   string `json:"serializedParams"`
					} `json:"transcriptTrackSelectionEntity,omitempty"`
					TranscriptSearchBoxStateEntity struct {
						Key      string `json:"key"`
						IsHidden bool   `json:"isHidden"`
					} `json:"transcriptSearchBoxStateEntity,omitempty"`
				} `json:"payload,omitempty"`
			} `json:"mutations"`
			Timestamp struct {
				Seconds string `json:"seconds"`
				Nanos   int    `json:"nanos"`
			} `json:"timestamp"`
		} `json:"entityBatchUpdate"`
		ElementUpdate struct {
			Updates []struct {
				TemplateUpdate struct {
					Identifier               string   `json:"identifier"`
					SerializedTemplateConfig string   `json:"serializedTemplateConfig"`
					Dependencies             []string `json:"dependencies,omitempty"`
				} `json:"templateUpdate,omitempty"`
				ResourceStatusInResponseCheck struct {
					ResourceStatuses []struct {
						Identifier string `json:"identifier"`
						Status     string `json:"status"`
					} `json:"resourceStatuses"`
					ServerBuildLabel string `json:"serverBuildLabel"`
				} `json:"resourceStatusInResponseCheck,omitempty"`
			} `json:"updates"`
		} `json:"elementUpdate,omitempty"`
	} `json:"frameworkUpdates,omitempty"`
	CurrentVideoEndpoint struct {
		ClickTrackingParams string `json:"clickTrackingParams"`
		CommandMetadata     struct {
			WebCommandMetadata struct {
				Url         string `json:"url"`
				WebPageType string `json:"webPageType"`
				RootVe      int    `json:"rootVe"`
			} `json:"webCommandMetadata"`
		} `json:"commandMetadata"`
		WatchEndpoint struct {
			VideoId                            string `json:"videoId"`
			WatchEndpointSupportedOnesieConfig struct {
				Html5PlaybackOnesieConfig struct {
					CommonConfig struct {
						Url string `json:"url"`
					} `json:"commonConfig"`
				} `json:"html5PlaybackOnesieConfig"`
			} `json:"watchEndpointSupportedOnesieConfig"`
		} `json:"watchEndpoint"`
	} `json:"currentVideoEndpoint,omitempty"`
	PlayerOverlays struct {
		PlayerOverlayRenderer struct {
			EndScreen struct {
				WatchNextEndScreenRenderer struct {
					Results []struct {
						EndScreenVideoRenderer struct {
							VideoId   string `json:"videoId"`
							Thumbnail struct {
								Thumbnails []struct {
									Url    string `json:"url"`
									Width  int    `json:"width"`
									Height int    `json:"height"`
								} `json:"thumbnails"`
							} `json:"thumbnail"`
							Title struct {
								Accessibility struct {
									AccessibilityData struct {
										Label string `json:"label"`
									} `json:"accessibilityData"`
								} `json:"accessibility"`
								SimpleText string `json:"simpleText"`
							} `json:"title"`
							ShortBylineText struct {
								Runs []struct {
									Text               string `json:"text"`
									NavigationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												Url         string `json:"url"`
												WebPageType string `json:"webPageType"`
												RootVe      int    `json:"rootVe"`
												ApiUrl      string `json:"apiUrl"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										BrowseEndpoint struct {
											BrowseId         string `json:"browseId"`
											CanonicalBaseUrl string `json:"canonicalBaseUrl"`
										} `json:"browseEndpoint"`
									} `json:"navigationEndpoint"`
								} `json:"runs"`
							} `json:"shortBylineText"`
							LengthText struct {
								Accessibility struct {
									AccessibilityData struct {
										Label string `json:"label"`
									} `json:"accessibilityData"`
								} `json:"accessibility"`
								SimpleText string `json:"simpleText"`
							} `json:"lengthText"`
							LengthInSeconds    int `json:"lengthInSeconds"`
							NavigationEndpoint struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										Url         string `json:"url"`
										WebPageType string `json:"webPageType"`
										RootVe      int    `json:"rootVe"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								WatchEndpoint struct {
									VideoId                            string `json:"videoId"`
									WatchEndpointSupportedOnesieConfig struct {
										Html5PlaybackOnesieConfig struct {
											CommonConfig struct {
												Url string `json:"url"`
											} `json:"commonConfig"`
										} `json:"html5PlaybackOnesieConfig"`
									} `json:"watchEndpointSupportedOnesieConfig"`
								} `json:"watchEndpoint"`
							} `json:"navigationEndpoint"`
							TrackingParams     string `json:"trackingParams"`
							ShortViewCountText struct {
								Accessibility struct {
									AccessibilityData struct {
										Label string `json:"label"`
									} `json:"accessibilityData"`
								} `json:"accessibility"`
								SimpleText string `json:"simpleText"`
							} `json:"shortViewCountText"`
							PublishedTimeText struct {
								SimpleText string `json:"simpleText"`
							} `json:"publishedTimeText"`
							ThumbnailOverlays []struct {
								ThumbnailOverlayTimeStatusRenderer struct {
									Text struct {
										Accessibility struct {
											AccessibilityData struct {
												Label string `json:"label"`
											} `json:"accessibilityData"`
										} `json:"accessibility"`
										SimpleText string `json:"simpleText"`
									} `json:"text"`
									Style string `json:"style"`
								} `json:"thumbnailOverlayTimeStatusRenderer,omitempty"`
								ThumbnailOverlayNowPlayingRenderer struct {
									Text struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"text"`
								} `json:"thumbnailOverlayNowPlayingRenderer,omitempty"`
							} `json:"thumbnailOverlays"`
						} `json:"endScreenVideoRenderer"`
					} `json:"results"`
					Title struct {
						SimpleText string `json:"simpleText"`
					} `json:"title"`
					TrackingParams string `json:"trackingParams"`
				} `json:"watchNextEndScreenRenderer"`
			} `json:"endScreen"`
			Autoplay struct {
				PlayerOverlayAutoplayRenderer struct {
					Title struct {
						SimpleText string `json:"simpleText"`
					} `json:"title"`
					VideoTitle struct {
						Accessibility struct {
							AccessibilityData struct {
								Label string `json:"label"`
							} `json:"accessibilityData"`
						} `json:"accessibility"`
						SimpleText string `json:"simpleText"`
					} `json:"videoTitle"`
					Byline struct {
						Runs []struct {
							Text               string `json:"text"`
							NavigationEndpoint struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										Url         string `json:"url"`
										WebPageType string `json:"webPageType"`
										RootVe      int    `json:"rootVe"`
										ApiUrl      string `json:"apiUrl"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								BrowseEndpoint struct {
									BrowseId         string `json:"browseId"`
									CanonicalBaseUrl string `json:"canonicalBaseUrl"`
								} `json:"browseEndpoint"`
							} `json:"navigationEndpoint"`
						} `json:"runs"`
					} `json:"byline"`
					PauseText struct {
						SimpleText string `json:"simpleText"`
					} `json:"pauseText"`
					Background struct {
						Thumbnails []struct {
							Url    string `json:"url"`
							Width  int    `json:"width"`
							Height int    `json:"height"`
						} `json:"thumbnails"`
					} `json:"background"`
					CountDownSecs int `json:"countDownSecs"`
					CancelButton  struct {
						ButtonRenderer struct {
							Style      string `json:"style"`
							Size       string `json:"size"`
							IsDisabled bool   `json:"isDisabled"`
							Text       struct {
								SimpleText string `json:"simpleText"`
							} `json:"text"`
							Accessibility struct {
								Label string `json:"label"`
							} `json:"accessibility"`
							TrackingParams    string `json:"trackingParams"`
							AccessibilityData struct {
								AccessibilityData struct {
									Label string `json:"label"`
								} `json:"accessibilityData"`
							} `json:"accessibilityData"`
							Command struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										SendPost bool   `json:"sendPost"`
										ApiUrl   string `json:"apiUrl"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								GetSurveyCommand struct {
									Endpoint struct {
										Watch struct {
											Hack bool `json:"hack"`
										} `json:"watch"`
									} `json:"endpoint"`
									Action string `json:"action"`
								} `json:"getSurveyCommand"`
							} `json:"command"`
						} `json:"buttonRenderer"`
					} `json:"cancelButton"`
					NextButton struct {
						ButtonRenderer struct {
							Style              string `json:"style"`
							Size               string `json:"size"`
							IsDisabled         bool   `json:"isDisabled"`
							NavigationEndpoint struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										Url         string `json:"url"`
										WebPageType string `json:"webPageType"`
										RootVe      int    `json:"rootVe"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								WatchEndpoint struct {
									VideoId                            string `json:"videoId"`
									WatchEndpointSupportedOnesieConfig struct {
										Html5PlaybackOnesieConfig struct {
											CommonConfig struct {
												Url string `json:"url"`
											} `json:"commonConfig"`
										} `json:"html5PlaybackOnesieConfig"`
									} `json:"watchEndpointSupportedOnesieConfig"`
								} `json:"watchEndpoint"`
							} `json:"navigationEndpoint"`
							Accessibility struct {
								Label string `json:"label"`
							} `json:"accessibility"`
							TrackingParams    string `json:"trackingParams"`
							AccessibilityData struct {
								AccessibilityData struct {
									Label string `json:"label"`
								} `json:"accessibilityData"`
							} `json:"accessibilityData"`
						} `json:"buttonRenderer"`
					} `json:"nextButton"`
					TrackingParams string `json:"trackingParams"`
					CloseButton    struct {
						ButtonRenderer struct {
							Style      string `json:"style"`
							Size       string `json:"size"`
							IsDisabled bool   `json:"isDisabled"`
							Icon       struct {
								IconType string `json:"iconType"`
							} `json:"icon"`
							Accessibility struct {
								Label string `json:"label"`
							} `json:"accessibility"`
							TrackingParams string `json:"trackingParams"`
						} `json:"buttonRenderer"`
					} `json:"closeButton"`
					ThumbnailOverlays []struct {
						ThumbnailOverlayTimeStatusRenderer struct {
							Text struct {
								Accessibility struct {
									AccessibilityData struct {
										Label string `json:"label"`
									} `json:"accessibilityData"`
								} `json:"accessibility"`
								SimpleText string `json:"simpleText"`
							} `json:"text"`
							Style string `json:"style"`
						} `json:"thumbnailOverlayTimeStatusRenderer"`
					} `json:"thumbnailOverlays"`
					PreferImmediateRedirect bool   `json:"preferImmediateRedirect"`
					VideoId                 string `json:"videoId"`
					PublishedTimeText       struct {
						SimpleText string `json:"simpleText"`
					} `json:"publishedTimeText"`
					WebShowNewAutonavCountdown   bool `json:"webShowNewAutonavCountdown"`
					WebShowBigThumbnailEndscreen bool `json:"webShowBigThumbnailEndscreen"`
					ShortViewCountText           struct {
						Accessibility struct {
							AccessibilityData struct {
								Label string `json:"label"`
							} `json:"accessibilityData"`
						} `json:"accessibility"`
						SimpleText string `json:"simpleText"`
					} `json:"shortViewCountText"`
					CountDownSecsForFullscreen int `json:"countDownSecsForFullscreen"`
				} `json:"playerOverlayAutoplayRenderer"`
			} `json:"autoplay"`
			ShareButton struct {
				ButtonRenderer struct {
					Style      string `json:"style"`
					Size       string `json:"size"`
					IsDisabled bool   `json:"isDisabled"`
					Icon       struct {
						IconType string `json:"iconType"`
					} `json:"icon"`
					NavigationEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool   `json:"sendPost"`
								ApiUrl   string `json:"apiUrl"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						ShareEntityServiceEndpoint struct {
							SerializedShareEntity string `json:"serializedShareEntity"`
							Commands              []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								OpenPopupAction     struct {
									Popup struct {
										UnifiedSharePanelRenderer struct {
											TrackingParams     string `json:"trackingParams"`
											ShowLoadingSpinner bool   `json:"showLoadingSpinner"`
										} `json:"unifiedSharePanelRenderer"`
									} `json:"popup"`
									PopupType string `json:"popupType"`
									BeReused  bool   `json:"beReused"`
								} `json:"openPopupAction"`
							} `json:"commands"`
						} `json:"shareEntityServiceEndpoint"`
					} `json:"navigationEndpoint"`
					Tooltip        string `json:"tooltip"`
					TrackingParams string `json:"trackingParams"`
				} `json:"buttonRenderer"`
			} `json:"shareButton"`
			AddToMenu struct {
				MenuRenderer struct {
					TrackingParams string `json:"trackingParams"`
				} `json:"menuRenderer"`
			} `json:"addToMenu"`
			VideoDetails struct {
				PlayerOverlayVideoDetailsRenderer struct {
					Title struct {
						SimpleText string `json:"simpleText"`
					} `json:"title"`
					Subtitle struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"subtitle"`
				} `json:"playerOverlayVideoDetailsRenderer"`
			} `json:"videoDetails"`
			AutonavToggle struct {
				AutoplaySwitchButtonRenderer struct {
					OnEnabledCommand struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool   `json:"sendPost"`
								ApiUrl   string `json:"apiUrl"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SetSettingEndpoint struct {
							SettingItemId          string `json:"settingItemId"`
							BoolValue              bool   `json:"boolValue"`
							SettingItemIdForClient string `json:"settingItemIdForClient"`
						} `json:"setSettingEndpoint"`
					} `json:"onEnabledCommand"`
					OnDisabledCommand struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool   `json:"sendPost"`
								ApiUrl   string `json:"apiUrl"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SetSettingEndpoint struct {
							SettingItemId          string `json:"settingItemId"`
							BoolValue              bool   `json:"boolValue"`
							SettingItemIdForClient string `json:"settingItemIdForClient"`
						} `json:"setSettingEndpoint"`
					} `json:"onDisabledCommand"`
					EnabledAccessibilityData struct {
						AccessibilityData struct {
							Label string `json:"label"`
						} `json:"accessibilityData"`
					} `json:"enabledAccessibilityData"`
					DisabledAccessibilityData struct {
						AccessibilityData struct {
							Label string `json:"label"`
						} `json:"accessibilityData"`
					} `json:"disabledAccessibilityData"`
					TrackingParams string `json:"trackingParams"`
					Enabled        bool   `json:"enabled"`
				} `json:"autoplaySwitchButtonRenderer"`
			} `json:"autonavToggle"`
			DecoratedPlayerBarRenderer struct {
				DecoratedPlayerBarRenderer struct {
					PlayerBar struct {
						MultiMarkersPlayerBarRenderer struct {
							VisibleOnLoad struct {
								Key string `json:"key"`
							} `json:"visibleOnLoad"`
							MarkersMap []struct {
								Key   string `json:"key"`
								Value struct {
									TrackingParams string `json:"trackingParams"`
									Heatmap        struct {
										HeatmapRenderer struct {
											MaxHeightDp                     int `json:"maxHeightDp"`
											MinHeightDp                     int `json:"minHeightDp"`
											ShowHideAnimationDurationMillis int `json:"showHideAnimationDurationMillis"`
											HeatMarkers                     []struct {
												HeatMarkerRenderer struct {
													TimeRangeStartMillis               int     `json:"timeRangeStartMillis"`
													MarkerDurationMillis               int     `json:"markerDurationMillis"`
													HeatMarkerIntensityScoreNormalized float64 `json:"heatMarkerIntensityScoreNormalized"`
												} `json:"heatMarkerRenderer"`
											} `json:"heatMarkers"`
											HeatMarkersDecorations []struct {
												TimedMarkerDecorationRenderer struct {
													VisibleTimeRangeStartMillis int `json:"visibleTimeRangeStartMillis"`
													VisibleTimeRangeEndMillis   int `json:"visibleTimeRangeEndMillis"`
													DecorationTimeMillis        int `json:"decorationTimeMillis"`
													Label                       struct {
														Runs []struct {
															Text string `json:"text"`
														} `json:"runs"`
													} `json:"label"`
													Icon           string `json:"icon"`
													TrackingParams string `json:"trackingParams"`
												} `json:"timedMarkerDecorationRenderer"`
											} `json:"heatMarkersDecorations"`
										} `json:"heatmapRenderer"`
									} `json:"heatmap"`
								} `json:"value"`
							} `json:"markersMap"`
						} `json:"multiMarkersPlayerBarRenderer"`
					} `json:"playerBar"`
				} `json:"decoratedPlayerBarRenderer"`
			} `json:"decoratedPlayerBarRenderer"`
		} `json:"playerOverlayRenderer"`
	} `json:"playerOverlays,omitempty"`
	OnResponseReceivedEndpoints []struct {
		ClickTrackingParams string `json:"clickTrackingParams"`
		CommandMetadata     struct {
			WebCommandMetadata struct {
				SendPost bool `json:"sendPost"`
			} `json:"webCommandMetadata"`
		} `json:"commandMetadata,omitempty"`
		SignalServiceEndpoint struct {
			Signal  string `json:"signal"`
			Actions []struct {
				ClickTrackingParams string `json:"clickTrackingParams"`
				SignalAction        struct {
					Signal string `json:"signal"`
				} `json:"signalAction"`
			} `json:"actions"`
		} `json:"signalServiceEndpoint,omitempty"`
		ChangeKeyedMarkersVisibilityCommand struct {
			IsVisible bool   `json:"isVisible"`
			Key       string `json:"key"`
		} `json:"changeKeyedMarkersVisibilityCommand,omitempty"`
		LoadMarkersCommand struct {
			EntityKeys []string `json:"entityKeys"`
		} `json:"loadMarkersCommand,omitempty"`
	} `json:"onResponseReceivedEndpoints,omitempty"`
	EngagementPanels []struct {
		EngagementPanelSectionListRenderer struct {
			Content struct {
				AdsEngagementPanelContentRenderer struct {
					Hack bool `json:"hack"`
				} `json:"adsEngagementPanelContentRenderer,omitempty"`
				StructuredDescriptionContentRenderer struct {
					Items []struct {
						VideoDescriptionHeaderRenderer struct {
							Title struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"title"`
							Channel struct {
								SimpleText string `json:"simpleText"`
							} `json:"channel"`
							Views struct {
								SimpleText string `json:"simpleText"`
							} `json:"views"`
							PublishDate struct {
								SimpleText string `json:"simpleText"`
							} `json:"publishDate"`
							Factoid []struct {
								FactoidRenderer struct {
									Value struct {
										SimpleText string `json:"simpleText"`
									} `json:"value"`
									Label struct {
										SimpleText string `json:"simpleText"`
									} `json:"label"`
									AccessibilityText string `json:"accessibilityText"`
								} `json:"factoidRenderer"`
							} `json:"factoid"`
							ChannelNavigationEndpoint struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										Url         string `json:"url"`
										WebPageType string `json:"webPageType"`
										RootVe      int    `json:"rootVe"`
										ApiUrl      string `json:"apiUrl"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								BrowseEndpoint struct {
									BrowseId         string `json:"browseId"`
									CanonicalBaseUrl string `json:"canonicalBaseUrl"`
								} `json:"browseEndpoint"`
							} `json:"channelNavigationEndpoint"`
							ChannelThumbnail struct {
								Thumbnails []struct {
									Url string `json:"url"`
								} `json:"thumbnails"`
							} `json:"channelThumbnail"`
						} `json:"videoDescriptionHeaderRenderer,omitempty"`
						ExpandableVideoDescriptionBodyRenderer struct {
							ShowMoreText struct {
								SimpleText string `json:"simpleText"`
							} `json:"showMoreText"`
							ShowLessText struct {
								SimpleText string `json:"simpleText"`
							} `json:"showLessText"`
							AttributedDescriptionBodyText struct {
								Content     string `json:"content"`
								CommandRuns []struct {
									StartIndex int `json:"startIndex"`
									Length     int `json:"length"`
									OnTap      struct {
										InnertubeCommand struct {
											ClickTrackingParams string `json:"clickTrackingParams"`
											CommandMetadata     struct {
												WebCommandMetadata struct {
													Url         string `json:"url"`
													WebPageType string `json:"webPageType"`
													RootVe      int    `json:"rootVe"`
													ApiUrl      string `json:"apiUrl,omitempty"`
												} `json:"webCommandMetadata"`
											} `json:"commandMetadata"`
											UrlEndpoint struct {
												Url      string `json:"url"`
												Target   string `json:"target,omitempty"`
												Nofollow bool   `json:"nofollow"`
											} `json:"urlEndpoint,omitempty"`
											BrowseEndpoint struct {
												BrowseId         string `json:"browseId"`
												CanonicalBaseUrl string `json:"canonicalBaseUrl"`
											} `json:"browseEndpoint,omitempty"`
										} `json:"innertubeCommand"`
									} `json:"onTap"`
									LoggingDirectives struct {
										TrackingParams                string `json:"trackingParams"`
										EnableDisplayloggerExperiment bool   `json:"enableDisplayloggerExperiment"`
									} `json:"loggingDirectives,omitempty"`
								} `json:"commandRuns"`
								StyleRuns []struct {
									StartIndex         int `json:"startIndex"`
									Length             int `json:"length"`
									StyleRunExtensions struct {
										StyleRunColorMapExtension struct {
											ColorMap []struct {
												Key   string `json:"key"`
												Value int64  `json:"value"`
											} `json:"colorMap"`
										} `json:"styleRunColorMapExtension"`
									} `json:"styleRunExtensions"`
								} `json:"styleRuns"`
								AttachmentRuns []struct {
									StartIndex int `json:"startIndex"`
									Length     int `json:"length"`
									Element    struct {
										Type struct {
											ImageType struct {
												Image struct {
													Sources []struct {
														Url string `json:"url"`
													} `json:"sources"`
												} `json:"image"`
											} `json:"imageType"`
										} `json:"type"`
										Properties struct {
											LayoutProperties struct {
												Height struct {
													Value int    `json:"value"`
													Unit  string `json:"unit"`
												} `json:"height"`
												Width struct {
													Value int    `json:"value"`
													Unit  string `json:"unit"`
												} `json:"width"`
											} `json:"layoutProperties"`
										} `json:"properties"`
									} `json:"element"`
									Alignment string `json:"alignment"`
								} `json:"attachmentRuns"`
								DecorationRuns []struct {
									TextDecorator struct {
										HighlightTextDecorator struct {
											StartIndex                       int `json:"startIndex"`
											Length                           int `json:"length"`
											BackgroundCornerRadius           int `json:"backgroundCornerRadius"`
											HighlightTextDecoratorExtensions struct {
												HighlightTextDecoratorColorMapExtension struct {
													ColorMap []struct {
														Key   string `json:"key"`
														Value int    `json:"value"`
													} `json:"colorMap"`
												} `json:"highlightTextDecoratorColorMapExtension"`
											} `json:"highlightTextDecoratorExtensions"`
										} `json:"highlightTextDecorator"`
									} `json:"textDecorator"`
								} `json:"decorationRuns"`
							} `json:"attributedDescriptionBodyText"`
						} `json:"expandableVideoDescriptionBodyRenderer,omitempty"`
						VideoDescriptionInfocardsSectionRenderer struct {
							SectionTitle struct {
								SimpleText string `json:"simpleText"`
							} `json:"sectionTitle"`
							CreatorVideosButton struct {
								ButtonRenderer struct {
									Style      string `json:"style"`
									Size       string `json:"size"`
									IsDisabled bool   `json:"isDisabled"`
									Text       struct {
										SimpleText string `json:"simpleText"`
									} `json:"text"`
									Icon struct {
										IconType string `json:"iconType"`
									} `json:"icon"`
									TrackingParams string `json:"trackingParams"`
									Command        struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												Url         string `json:"url"`
												WebPageType string `json:"webPageType"`
												RootVe      int    `json:"rootVe"`
												ApiUrl      string `json:"apiUrl"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										BrowseEndpoint struct {
											BrowseId string `json:"browseId"`
											Params   string `json:"params"`
										} `json:"browseEndpoint"`
									} `json:"command"`
								} `json:"buttonRenderer"`
							} `json:"creatorVideosButton"`
							CreatorAboutButton struct {
								ButtonRenderer struct {
									Style      string `json:"style"`
									Size       string `json:"size"`
									IsDisabled bool   `json:"isDisabled"`
									Text       struct {
										SimpleText string `json:"simpleText"`
									} `json:"text"`
									Icon struct {
										IconType string `json:"iconType"`
									} `json:"icon"`
									TrackingParams string `json:"trackingParams"`
									Command        struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												Url         string `json:"url"`
												WebPageType string `json:"webPageType"`
												RootVe      int    `json:"rootVe"`
												ApiUrl      string `json:"apiUrl"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										BrowseEndpoint struct {
											BrowseId string `json:"browseId"`
											Params   string `json:"params"`
										} `json:"browseEndpoint"`
									} `json:"command"`
								} `json:"buttonRenderer"`
							} `json:"creatorAboutButton"`
							SectionSubtitle struct {
								Accessibility struct {
									AccessibilityData struct {
										Label string `json:"label"`
									} `json:"accessibilityData"`
								} `json:"accessibility"`
								SimpleText string `json:"simpleText"`
							} `json:"sectionSubtitle"`
							ChannelAvatar struct {
								Thumbnails []struct {
									Url string `json:"url"`
								} `json:"thumbnails"`
							} `json:"channelAvatar"`
							ChannelEndpoint struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										Url         string `json:"url"`
										WebPageType string `json:"webPageType"`
										RootVe      int    `json:"rootVe"`
										ApiUrl      string `json:"apiUrl"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								BrowseEndpoint struct {
									BrowseId         string `json:"browseId"`
									CanonicalBaseUrl string `json:"canonicalBaseUrl"`
								} `json:"browseEndpoint"`
							} `json:"channelEndpoint"`
							TrackingParams string `json:"trackingParams"`
						} `json:"videoDescriptionInfocardsSectionRenderer,omitempty"`
					} `json:"items"`
				} `json:"structuredDescriptionContentRenderer,omitempty"`
				SectionListRenderer struct {
					Contents []struct {
						ItemSectionRenderer struct {
							Contents []struct {
								ContinuationItemRenderer struct {
									Trigger              string `json:"trigger"`
									ContinuationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												SendPost bool   `json:"sendPost"`
												ApiUrl   string `json:"apiUrl"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										ContinuationCommand struct {
											Token   string `json:"token"`
											Request string `json:"request"`
										} `json:"continuationCommand"`
									} `json:"continuationEndpoint"`
								} `json:"continuationItemRenderer"`
							} `json:"contents"`
							TrackingParams    string `json:"trackingParams"`
							SectionIdentifier string `json:"sectionIdentifier"`
							TargetId          string `json:"targetId"`
						} `json:"itemSectionRenderer"`
					} `json:"contents"`
					TrackingParams string `json:"trackingParams"`
				} `json:"sectionListRenderer,omitempty"`
				ContinuationItemRenderer struct {
					Trigger              string `json:"trigger"`
					ContinuationEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool   `json:"sendPost"`
								ApiUrl   string `json:"apiUrl"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						GetTranscriptEndpoint struct {
							Params string `json:"params"`
						} `json:"getTranscriptEndpoint"`
					} `json:"continuationEndpoint"`
				} `json:"continuationItemRenderer,omitempty"`
			} `json:"content"`
			TargetId          string `json:"targetId"`
			Visibility        string `json:"visibility"`
			LoggingDirectives struct {
				TrackingParams string `json:"trackingParams"`
				Visibility     struct {
					Types string `json:"types"`
				} `json:"visibility"`
				EnableDisplayloggerExperiment bool `json:"enableDisplayloggerExperiment"`
			} `json:"loggingDirectives"`
			PanelIdentifier string `json:"panelIdentifier,omitempty"`
			Header          struct {
				EngagementPanelTitleHeaderRenderer struct {
					Title struct {
						SimpleText string `json:"simpleText,omitempty"`
						Runs       []struct {
							Text string `json:"text"`
						} `json:"runs,omitempty"`
					} `json:"title"`
					VisibilityButton struct {
						ButtonRenderer struct {
							Icon struct {
								IconType string `json:"iconType"`
							} `json:"icon"`
							TrackingParams    string `json:"trackingParams"`
							AccessibilityData struct {
								AccessibilityData struct {
									Label string `json:"label"`
								} `json:"accessibilityData"`
							} `json:"accessibilityData"`
							Command struct {
								ClickTrackingParams    string `json:"clickTrackingParams"`
								CommandExecutorCommand struct {
									Commands []struct {
										ClickTrackingParams                   string `json:"clickTrackingParams"`
										ChangeEngagementPanelVisibilityAction struct {
											TargetId   string `json:"targetId"`
											Visibility string `json:"visibility"`
										} `json:"changeEngagementPanelVisibilityAction,omitempty"`
										UpdateToggleButtonStateCommand struct {
											Toggled  bool   `json:"toggled"`
											ButtonId string `json:"buttonId"`
										} `json:"updateToggleButtonStateCommand,omitempty"`
									} `json:"commands"`
								} `json:"commandExecutorCommand,omitempty"`
								ChangeEngagementPanelVisibilityAction struct {
									TargetId   string `json:"targetId"`
									Visibility string `json:"visibility"`
								} `json:"changeEngagementPanelVisibilityAction,omitempty"`
							} `json:"command"`
							Style         string `json:"style,omitempty"`
							Size          string `json:"size,omitempty"`
							Accessibility struct {
								Label string `json:"label"`
							} `json:"accessibility,omitempty"`
						} `json:"buttonRenderer"`
					} `json:"visibilityButton"`
					TrackingParams string `json:"trackingParams"`
					ContextualInfo struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"contextualInfo,omitempty"`
					Menu struct {
						SortFilterSubMenuRenderer struct {
							SubMenuItems []struct {
								Title           string `json:"title"`
								Selected        bool   `json:"selected"`
								ServiceEndpoint struct {
									ClickTrackingParams string `json:"clickTrackingParams"`
									CommandMetadata     struct {
										WebCommandMetadata struct {
											SendPost bool   `json:"sendPost"`
											ApiUrl   string `json:"apiUrl"`
										} `json:"webCommandMetadata"`
									} `json:"commandMetadata"`
									ContinuationCommand struct {
										Token   string `json:"token"`
										Request string `json:"request"`
										Command struct {
											ClickTrackingParams string `json:"clickTrackingParams"`
											ShowReloadUiCommand struct {
												TargetId string `json:"targetId"`
											} `json:"showReloadUiCommand"`
										} `json:"command"`
									} `json:"continuationCommand"`
								} `json:"serviceEndpoint"`
								TrackingParams string `json:"trackingParams"`
							} `json:"subMenuItems"`
							Icon struct {
								IconType string `json:"iconType"`
							} `json:"icon"`
							Accessibility struct {
								AccessibilityData struct {
									Label string `json:"label"`
								} `json:"accessibilityData"`
							} `json:"accessibility"`
							TrackingParams string `json:"trackingParams"`
						} `json:"sortFilterSubMenuRenderer,omitempty"`
						MenuRenderer struct {
							Items []struct {
								MenuServiceItemRenderer struct {
									Text struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"text"`
									ServiceEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												SendPost bool `json:"sendPost"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										SignalServiceEndpoint struct {
											Signal  string `json:"signal"`
											Actions []struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												SignalAction        struct {
													Signal string `json:"signal"`
												} `json:"signalAction"`
											} `json:"actions"`
										} `json:"signalServiceEndpoint"`
									} `json:"serviceEndpoint"`
									TrackingParams string `json:"trackingParams"`
								} `json:"menuServiceItemRenderer"`
							} `json:"items"`
							TrackingParams string `json:"trackingParams"`
							Accessibility  struct {
								AccessibilityData struct {
									Label string `json:"label"`
								} `json:"accessibilityData"`
							} `json:"accessibility"`
						} `json:"menuRenderer,omitempty"`
					} `json:"menu,omitempty"`
				} `json:"engagementPanelTitleHeaderRenderer"`
			} `json:"header,omitempty"`
			VeType         int `json:"veType,omitempty"`
			OnShowCommands []struct {
				ClickTrackingParams            string `json:"clickTrackingParams"`
				ScrollToEngagementPanelCommand struct {
					TargetId string `json:"targetId"`
				} `json:"scrollToEngagementPanelCommand"`
			} `json:"onShowCommands,omitempty"`
		} `json:"engagementPanelSectionListRenderer"`
	} `json:"engagementPanels,omitempty"`
	PageVisualEffects []struct {
		CinematicContainerRenderer struct {
			GradientColorConfig []struct {
				DarkThemeColor int64 `json:"darkThemeColor"`
				StartLocation  int   `json:"startLocation,omitempty"`
			} `json:"gradientColorConfig"`
			PresentationStyle string `json:"presentationStyle"`
			Config            struct {
				LightThemeBackgroundColor int64 `json:"lightThemeBackgroundColor"`
				DarkThemeBackgroundColor  int64 `json:"darkThemeBackgroundColor"`
				AnimationConfig           struct {
					MinImageUpdateIntervalMs int `json:"minImageUpdateIntervalMs"`
					CrossfadeDurationMs      int `json:"crossfadeDurationMs"`
					CrossfadeStartOffset     int `json:"crossfadeStartOffset"`
					MaxFrameRate             int `json:"maxFrameRate"`
				} `json:"animationConfig"`
				ColorSourceSizeMultiplier         float64 `json:"colorSourceSizeMultiplier"`
				ApplyClientImageBlur              bool    `json:"applyClientImageBlur"`
				BottomColorSourceHeightMultiplier float64 `json:"bottomColorSourceHeightMultiplier"`
				MaxBottomColorSourceHeight        int     `json:"maxBottomColorSourceHeight"`
				ColorSourceWidthMultiplier        float64 `json:"colorSourceWidthMultiplier"`
				ColorSourceHeightMultiplier       int     `json:"colorSourceHeightMultiplier"`
				BlurStrength                      int     `json:"blurStrength"`
				EnableInLightTheme                bool    `json:"enableInLightTheme"`
			} `json:"config"`
		} `json:"cinematicContainerRenderer"`
	} `json:"pageVisualEffects,omitempty"`
}

type continueInput struct {
	Context struct {
		Client struct {
			Hl string `json:"hl"` // language you want the data in, for english "en"
			//Gl string `json:"gl"`
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
	Continuation        string `json:"continuation"`
	BrowseId            string `json:"browseId,omitempty"`
	InlineSettingStatus string `json:"inlineSettingStatus,omitempty"`
}

type continueOutput struct {
	ResponseContext struct {
		ServiceTrackingParams []struct {
			Service string `json:"service"`
			Params  []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"params"`
		} `json:"serviceTrackingParams"`
		MaxAgeSeconds             int `json:"maxAgeSeconds"`
		MainAppWebResponseContext struct {
			DatasyncId    string `json:"datasyncId,omitempty"`
			LoggedOut     bool   `json:"loggedOut"`
			TrackingParam string `json:"trackingParam"`
		} `json:"mainAppWebResponseContext"`
		WebResponseContextExtensionData struct {
			HasDecorated bool `json:"hasDecorated"`
		} `json:"webResponseContextExtensionData"`
		VisitorData string `json:"visitorData,omitempty"`
	} `json:"responseContext"`
	TrackingParams            string `json:"trackingParams"`
	OnResponseReceivedActions []struct {
		ClickTrackingParams           string `json:"clickTrackingParams"`
		AppendContinuationItemsAction struct {
			ContinuationItems []struct {
				RichItemRenderer struct {
					Content struct {
						VideoRenderer struct {
							VideoId   string `json:"videoId"`
							Thumbnail struct {
								Thumbnails []struct {
									Url    string `json:"url"`
									Width  int    `json:"width"`
									Height int    `json:"height"`
								} `json:"thumbnails"`
							} `json:"thumbnail"`
							Title struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
								Accessibility struct {
									AccessibilityData struct {
										Label string `json:"label"`
									} `json:"accessibilityData"`
								} `json:"accessibility"`
							} `json:"title"`
							DescriptionSnippet struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"descriptionSnippet,omitempty"`
							LongBylineText struct {
								Runs []struct {
									Text               string `json:"text"`
									NavigationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												Url         string `json:"url"`
												WebPageType string `json:"webPageType"`
												RootVe      int    `json:"rootVe"`
												ApiUrl      string `json:"apiUrl"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										BrowseEndpoint struct {
											BrowseId         string `json:"browseId"`
											CanonicalBaseUrl string `json:"canonicalBaseUrl"`
										} `json:"browseEndpoint"`
									} `json:"navigationEndpoint"`
								} `json:"runs"`
							} `json:"longBylineText,omitempty"`
							PublishedTimeText struct {
								SimpleText string `json:"simpleText"`
							} `json:"publishedTimeText"`
							LengthText struct {
								Accessibility struct {
									AccessibilityData struct {
										Label string `json:"label"`
									} `json:"accessibilityData"`
								} `json:"accessibility"`
								SimpleText string `json:"simpleText"`
							} `json:"lengthText"`
							ViewCountText struct {
								SimpleText string `json:"simpleText"`
							} `json:"viewCountText"`
							NavigationEndpoint struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										Url         string `json:"url"`
										WebPageType string `json:"webPageType"`
										RootVe      int    `json:"rootVe"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								WatchEndpoint struct {
									VideoId                            string `json:"videoId"`
									WatchEndpointSupportedOnesieConfig struct {
										Html5PlaybackOnesieConfig struct {
											CommonConfig struct {
												Url string `json:"url"`
											} `json:"commonConfig"`
										} `json:"html5PlaybackOnesieConfig"`
									} `json:"watchEndpointSupportedOnesieConfig"`
								} `json:"watchEndpoint"`
							} `json:"navigationEndpoint"`
							OwnerText struct {
								Runs []struct {
									Text               string `json:"text"`
									NavigationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												Url         string `json:"url"`
												WebPageType string `json:"webPageType"`
												RootVe      int    `json:"rootVe"`
												ApiUrl      string `json:"apiUrl"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										BrowseEndpoint struct {
											BrowseId         string `json:"browseId"`
											CanonicalBaseUrl string `json:"canonicalBaseUrl"`
										} `json:"browseEndpoint"`
									} `json:"navigationEndpoint"`
								} `json:"runs"`
							} `json:"ownerText,omitempty"`
							ShortBylineText struct {
								Runs []struct {
									Text               string `json:"text"`
									NavigationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												Url         string `json:"url"`
												WebPageType string `json:"webPageType"`
												RootVe      int    `json:"rootVe"`
												ApiUrl      string `json:"apiUrl"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										BrowseEndpoint struct {
											BrowseId         string `json:"browseId"`
											CanonicalBaseUrl string `json:"canonicalBaseUrl"`
										} `json:"browseEndpoint"`
									} `json:"navigationEndpoint"`
								} `json:"runs"`
							} `json:"shortBylineText,omitempty"`
							TrackingParams     string `json:"trackingParams"`
							ShowActionMenu     bool   `json:"showActionMenu"`
							ShortViewCountText struct {
								Accessibility struct {
									AccessibilityData struct {
										Label string `json:"label"`
									} `json:"accessibilityData"`
								} `json:"accessibility"`
								SimpleText string `json:"simpleText"`
							} `json:"shortViewCountText"`
							Menu struct {
								MenuRenderer struct {
									Items []struct {
										MenuServiceItemRenderer struct {
											Text struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"text"`
											Icon struct {
												IconType string `json:"iconType"`
											} `json:"icon"`
											ServiceEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														SendPost bool   `json:"sendPost"`
														ApiUrl   string `json:"apiUrl,omitempty"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												SignalServiceEndpoint struct {
													Signal  string `json:"signal"`
													Actions []struct {
														ClickTrackingParams  string `json:"clickTrackingParams"`
														AddToPlaylistCommand struct {
															OpenMiniplayer      bool   `json:"openMiniplayer"`
															VideoId             string `json:"videoId"`
															ListType            string `json:"listType"`
															OnCreateListCommand struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		SendPost bool   `json:"sendPost"`
																		ApiUrl   string `json:"apiUrl"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																CreatePlaylistServiceEndpoint struct {
																	VideoIds []string `json:"videoIds"`
																	Params   string   `json:"params"`
																} `json:"createPlaylistServiceEndpoint"`
															} `json:"onCreateListCommand"`
															VideoIds []string `json:"videoIds"`
														} `json:"addToPlaylistCommand"`
													} `json:"actions"`
												} `json:"signalServiceEndpoint,omitempty"`
												PlaylistEditEndpoint struct {
													PlaylistId string `json:"playlistId"`
													Actions    []struct {
														AddedVideoId string `json:"addedVideoId"`
														Action       string `json:"action"`
													} `json:"actions"`
												} `json:"playlistEditEndpoint,omitempty"`
												AddToPlaylistServiceEndpoint struct {
													VideoId string `json:"videoId"`
												} `json:"addToPlaylistServiceEndpoint,omitempty"`
												ShareEntityServiceEndpoint struct {
													SerializedShareEntity string `json:"serializedShareEntity"`
													Commands              []struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														OpenPopupAction     struct {
															Popup struct {
																UnifiedSharePanelRenderer struct {
																	TrackingParams     string `json:"trackingParams"`
																	ShowLoadingSpinner bool   `json:"showLoadingSpinner"`
																} `json:"unifiedSharePanelRenderer"`
															} `json:"popup"`
															PopupType string `json:"popupType"`
															BeReused  bool   `json:"beReused"`
														} `json:"openPopupAction"`
													} `json:"commands"`
												} `json:"shareEntityServiceEndpoint,omitempty"`
												FeedbackEndpoint struct {
													FeedbackToken string `json:"feedbackToken"`
													UiActions     struct {
														HideEnclosingContainer bool `json:"hideEnclosingContainer"`
													} `json:"uiActions"`
													Actions []struct {
														ClickTrackingParams    string `json:"clickTrackingParams"`
														ReplaceEnclosingAction struct {
															Item struct {
																NotificationMultiActionRenderer struct {
																	ResponseText struct {
																		Accessibility struct {
																			AccessibilityData struct {
																				Label string `json:"label"`
																			} `json:"accessibilityData"`
																		} `json:"accessibility"`
																		SimpleText string `json:"simpleText,omitempty"`
																		Runs       []struct {
																			Text string `json:"text"`
																		} `json:"runs,omitempty"`
																	} `json:"responseText"`
																	Buttons []struct {
																		ButtonRenderer struct {
																			Style string `json:"style"`
																			Text  struct {
																				SimpleText string `json:"simpleText,omitempty"`
																				Runs       []struct {
																					Text string `json:"text"`
																				} `json:"runs,omitempty"`
																			} `json:"text"`
																			ServiceEndpoint struct {
																				ClickTrackingParams string `json:"clickTrackingParams"`
																				CommandMetadata     struct {
																					WebCommandMetadata struct {
																						SendPost bool   `json:"sendPost"`
																						ApiUrl   string `json:"apiUrl,omitempty"`
																					} `json:"webCommandMetadata"`
																				} `json:"commandMetadata"`
																				UndoFeedbackEndpoint struct {
																					UndoToken string `json:"undoToken"`
																					Actions   []struct {
																						ClickTrackingParams string `json:"clickTrackingParams"`
																						UndoFeedbackAction  struct {
																							Hack bool `json:"hack"`
																						} `json:"undoFeedbackAction"`
																					} `json:"actions"`
																				} `json:"undoFeedbackEndpoint,omitempty"`
																				SignalServiceEndpoint struct {
																					Signal  string `json:"signal"`
																					Actions []struct {
																						ClickTrackingParams string `json:"clickTrackingParams"`
																						SignalAction        struct {
																							Signal string `json:"signal"`
																						} `json:"signalAction"`
																					} `json:"actions"`
																				} `json:"signalServiceEndpoint,omitempty"`
																			} `json:"serviceEndpoint,omitempty"`
																			TrackingParams string `json:"trackingParams"`
																			Command        struct {
																				ClickTrackingParams string `json:"clickTrackingParams"`
																				CommandMetadata     struct {
																					WebCommandMetadata struct {
																						Url         string `json:"url"`
																						WebPageType string `json:"webPageType"`
																						RootVe      int    `json:"rootVe"`
																					} `json:"webCommandMetadata"`
																				} `json:"commandMetadata"`
																				UrlEndpoint struct {
																					Url    string `json:"url"`
																					Target string `json:"target"`
																				} `json:"urlEndpoint"`
																			} `json:"command,omitempty"`
																		} `json:"buttonRenderer"`
																	} `json:"buttons"`
																	TrackingParams     string `json:"trackingParams"`
																	DismissalViewStyle string `json:"dismissalViewStyle"`
																} `json:"notificationMultiActionRenderer"`
															} `json:"item"`
														} `json:"replaceEnclosingAction"`
													} `json:"actions"`
												} `json:"feedbackEndpoint,omitempty"`
												GetReportFormEndpoint struct {
													Params string `json:"params"`
												} `json:"getReportFormEndpoint,omitempty"`
											} `json:"serviceEndpoint"`
											TrackingParams string `json:"trackingParams"`
											HasSeparator   bool   `json:"hasSeparator,omitempty"`
										} `json:"menuServiceItemRenderer,omitempty"`
										MenuServiceItemDownloadRenderer struct {
											ServiceEndpoint struct {
												ClickTrackingParams  string `json:"clickTrackingParams"`
												OfflineVideoEndpoint struct {
													VideoId      string `json:"videoId"`
													OnAddCommand struct {
														ClickTrackingParams      string `json:"clickTrackingParams"`
														GetDownloadActionCommand struct {
															VideoId string `json:"videoId"`
															Params  string `json:"params"`
														} `json:"getDownloadActionCommand"`
													} `json:"onAddCommand"`
												} `json:"offlineVideoEndpoint"`
											} `json:"serviceEndpoint"`
											TrackingParams string `json:"trackingParams"`
										} `json:"menuServiceItemDownloadRenderer,omitempty"`
									} `json:"items"`
									TrackingParams string `json:"trackingParams"`
									Accessibility  struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"accessibility"`
								} `json:"menuRenderer"`
							} `json:"menu"`
							ChannelThumbnailSupportedRenderers struct {
								ChannelThumbnailWithLinkRenderer struct {
									Thumbnail struct {
										Thumbnails []struct {
											Url    string `json:"url"`
											Width  int    `json:"width"`
											Height int    `json:"height"`
										} `json:"thumbnails"`
									} `json:"thumbnail"`
									NavigationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												Url         string `json:"url"`
												WebPageType string `json:"webPageType"`
												RootVe      int    `json:"rootVe"`
												ApiUrl      string `json:"apiUrl"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										BrowseEndpoint struct {
											BrowseId         string `json:"browseId"`
											CanonicalBaseUrl string `json:"canonicalBaseUrl"`
										} `json:"browseEndpoint"`
									} `json:"navigationEndpoint"`
									Accessibility struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"accessibility"`
								} `json:"channelThumbnailWithLinkRenderer"`
							} `json:"channelThumbnailSupportedRenderers,omitempty"`
							ThumbnailOverlays []struct {
								ThumbnailOverlayTimeStatusRenderer struct {
									Text struct {
										Accessibility struct {
											AccessibilityData struct {
												Label string `json:"label"`
											} `json:"accessibilityData"`
										} `json:"accessibility"`
										SimpleText string `json:"simpleText"`
									} `json:"text"`
									Style string `json:"style"`
								} `json:"thumbnailOverlayTimeStatusRenderer,omitempty"`
								ThumbnailOverlayToggleButtonRenderer struct {
									IsToggled     bool `json:"isToggled,omitempty"`
									UntoggledIcon struct {
										IconType string `json:"iconType"`
									} `json:"untoggledIcon"`
									ToggledIcon struct {
										IconType string `json:"iconType"`
									} `json:"toggledIcon"`
									UntoggledTooltip         string `json:"untoggledTooltip"`
									ToggledTooltip           string `json:"toggledTooltip"`
									UntoggledServiceEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												SendPost bool   `json:"sendPost"`
												ApiUrl   string `json:"apiUrl,omitempty"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										PlaylistEditEndpoint struct {
											PlaylistId string `json:"playlistId"`
											Actions    []struct {
												AddedVideoId string `json:"addedVideoId"`
												Action       string `json:"action"`
											} `json:"actions"`
										} `json:"playlistEditEndpoint,omitempty"`
										SignalServiceEndpoint struct {
											Signal  string `json:"signal"`
											Actions []struct {
												ClickTrackingParams  string `json:"clickTrackingParams"`
												AddToPlaylistCommand struct {
													OpenMiniplayer      bool   `json:"openMiniplayer"`
													VideoId             string `json:"videoId"`
													ListType            string `json:"listType"`
													OnCreateListCommand struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																SendPost bool   `json:"sendPost"`
																ApiUrl   string `json:"apiUrl"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														CreatePlaylistServiceEndpoint struct {
															VideoIds []string `json:"videoIds"`
															Params   string   `json:"params"`
														} `json:"createPlaylistServiceEndpoint"`
													} `json:"onCreateListCommand"`
													VideoIds []string `json:"videoIds"`
												} `json:"addToPlaylistCommand"`
											} `json:"actions"`
										} `json:"signalServiceEndpoint,omitempty"`
									} `json:"untoggledServiceEndpoint"`
									ToggledServiceEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												SendPost bool   `json:"sendPost"`
												ApiUrl   string `json:"apiUrl"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										PlaylistEditEndpoint struct {
											PlaylistId string `json:"playlistId"`
											Actions    []struct {
												Action         string `json:"action"`
												RemovedVideoId string `json:"removedVideoId"`
											} `json:"actions"`
										} `json:"playlistEditEndpoint"`
									} `json:"toggledServiceEndpoint,omitempty"`
									UntoggledAccessibility struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"untoggledAccessibility"`
									ToggledAccessibility struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"toggledAccessibility"`
									TrackingParams string `json:"trackingParams"`
								} `json:"thumbnailOverlayToggleButtonRenderer,omitempty"`
								ThumbnailOverlayNowPlayingRenderer struct {
									Text struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"text"`
								} `json:"thumbnailOverlayNowPlayingRenderer,omitempty"`
								ThumbnailOverlayLoadingPreviewRenderer struct {
									Text struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"text"`
								} `json:"thumbnailOverlayLoadingPreviewRenderer,omitempty"`
							} `json:"thumbnailOverlays"`
							RichThumbnail struct {
								MovingThumbnailRenderer struct {
									MovingThumbnailDetails struct {
										Thumbnails []struct {
											Url    string `json:"url"`
											Width  int    `json:"width"`
											Height int    `json:"height"`
										} `json:"thumbnails"`
										LogAsMovingThumbnail bool `json:"logAsMovingThumbnail"`
									} `json:"movingThumbnailDetails"`
									EnableHoveredLogging bool `json:"enableHoveredLogging"`
									EnableOverlay        bool `json:"enableOverlay"`
								} `json:"movingThumbnailRenderer"`
							} `json:"richThumbnail,omitempty"`
							InlinePlaybackEndpoint struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										Url         string `json:"url"`
										WebPageType string `json:"webPageType"`
										RootVe      int    `json:"rootVe"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								WatchEndpoint struct {
									VideoId              string `json:"videoId"`
									PlayerParams         string `json:"playerParams"`
									PlayerExtraUrlParams []struct {
										Key   string `json:"key"`
										Value string `json:"value"`
									} `json:"playerExtraUrlParams"`
									WatchEndpointSupportedOnesieConfig struct {
										Html5PlaybackOnesieConfig struct {
											CommonConfig struct {
												Url string `json:"url"`
											} `json:"commonConfig"`
										} `json:"html5PlaybackOnesieConfig"`
									} `json:"watchEndpointSupportedOnesieConfig"`
								} `json:"watchEndpoint"`
							} `json:"inlinePlaybackEndpoint,omitempty"`
							OwnerBadges []struct {
								MetadataBadgeRenderer struct {
									Icon struct {
										IconType string `json:"iconType"`
									} `json:"icon"`
									Style             string `json:"style"`
									Tooltip           string `json:"tooltip"`
									TrackingParams    string `json:"trackingParams"`
									AccessibilityData struct {
										Label string `json:"label"`
									} `json:"accessibilityData"`
								} `json:"metadataBadgeRenderer"`
							} `json:"ownerBadges,omitempty"`
						} `json:"videoRenderer"`
					} `json:"content"`
					TrackingParams string `json:"trackingParams"`
				} `json:"richItemRenderer,omitempty"`
				ContinuationItemRenderer struct {
					Trigger              string `json:"trigger"`
					ContinuationEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool   `json:"sendPost"`
								ApiUrl   string `json:"apiUrl"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						ContinuationCommand struct {
							Token   string `json:"token"`
							Request string `json:"request"`
						} `json:"continuationCommand"`
					} `json:"continuationEndpoint"`
					GhostCards struct {
						GhostGridRenderer struct {
							Rows int `json:"rows"`
						} `json:"ghostGridRenderer"`
					} `json:"ghostCards,omitempty"`
				} `json:"continuationItemRenderer,omitempty"`
			} `json:"continuationItems"`
			TargetId string `json:"targetId"`
		} `json:"appendContinuationItemsAction"`
	} `json:"onResponseReceivedActions"`
}
