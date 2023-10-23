package youtube_scraper

type accountScrapeInitial struct {
	//ResponseContext struct {
	//	ServiceTrackingParams []struct {
	//		Service string `json:"service"`
	//		Params  []struct {
	//			Key   string `json:"key"`
	//			Value string `json:"value"`
	//		} `json:"params"`
	//	} `json:"serviceTrackingParams"`
	//	MaxAgeSeconds             int `json:"maxAgeSeconds"`
	//	MainAppWebResponseContext struct {
	//		LoggedOut     bool   `json:"loggedOut"`
	//		TrackingParam string `json:"trackingParam"`
	//	} `json:"mainAppWebResponseContext"`
	//	WebResponseContextExtensionData struct {
	//		YtConfigData struct {
	//			VisitorData           string `json:"visitorData"`
	//			RootVisualElementType int    `json:"rootVisualElementType"`
	//		} `json:"ytConfigData"`
	//		HasDecorated bool `json:"hasDecorated"`
	//	} `json:"webResponseContextExtensionData"`
	//} `json:"responseContext"`
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
					} `json:"endpoint"`
					Title          string `json:"title"`
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
											//NavigationEndpoint struct {
											//	ClickTrackingParams string `json:"clickTrackingParams"`
											//	CommandMetadata     struct {
											//		WebCommandMetadata struct {
											//			Url         string `json:"url"`
											//			WebPageType string `json:"webPageType"`
											//			RootVe      int    `json:"rootVe"`
											//		} `json:"webCommandMetadata"`
											//	} `json:"commandMetadata"`
											//	WatchEndpoint struct {
											//		VideoId                            string `json:"videoId"`
											//		WatchEndpointSupportedOnesieConfig struct {
											//			Html5PlaybackOnesieConfig struct {
											//				CommonConfig struct {
											//					Url string `json:"url"`
											//				} `json:"commonConfig"`
											//			} `json:"html5PlaybackOnesieConfig"`
											//		} `json:"watchEndpointSupportedOnesieConfig"`
											//	} `json:"watchEndpoint"`
											//} `json:"navigationEndpoint"`
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
											} `json:"thumbnailOverlays"`
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
								} `json:"continuationItemRenderer,omitempty"`
							} `json:"contents"`
							TrackingParams string `json:"trackingParams"`
							Header         struct {
								FeedFilterChipBarRenderer struct {
									Contents []struct {
										ChipCloudChipRenderer struct {
											Text struct {
												SimpleText string `json:"simpleText"`
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
											} `json:"navigationEndpoint"`
											TrackingParams string `json:"trackingParams"`
											IsSelected     bool   `json:"isSelected"`
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
							TargetId string `json:"targetId"`
							Style    string `json:"style"`
						} `json:"richGridRenderer"`
					} `json:"content,omitempty"`
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
		} `json:"twoColumnBrowseResultsRenderer"`
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
		} `json:"c4TabbedHeaderRenderer"`
	} `json:"header"`
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
			AvailableCountryCodes  []string `json:"availableCountryCodes"`
			AndroidDeepLink        string   `json:"androidDeepLink"`
			AndroidAppindexingLink string   `json:"androidAppindexingLink"`
			IosAppindexingLink     string   `json:"iosAppindexingLink"`
			VanityChannelUrl       string   `json:"vanityChannelUrl"`
		} `json:"channelMetadataRenderer"`
	} `json:"metadata"`
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
			Tags               []string `json:"tags"`
			AvailableCountries []string `json:"availableCountries"`
			LinkAlternates     []struct {
				HrefUrl string `json:"hrefUrl"`
			} `json:"linkAlternates"`
		} `json:"microformatDataRenderer"`
	} `json:"microformat"`
}
type accountScrapeContinueInput struct {
	Context struct {
		Client struct {
			Hl string `json:"hl"` // language you want the data in, for english "en"
			//Gl string `json:"gl"`
			//RemoteHost string `json:"remoteHost"`
			//DeviceMake    string `json:"deviceMake"`
			//DeviceModel   string `json:"deviceModel"`
			//VisitorData string `json:"visitorData"`
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
	Continuation string `json:"continuation"`
}
type accountScrapeContinueOutput struct {
	ResponseContext struct {
		VisitorData           string `json:"visitorData"`
		ServiceTrackingParams []struct {
			Service string `json:"service"`
			Params  []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"params"`
		} `json:"serviceTrackingParams"`
		MaxAgeSeconds             int `json:"maxAgeSeconds"`
		MainAppWebResponseContext struct {
			LoggedOut     bool   `json:"loggedOut"`
			TrackingParam string `json:"trackingParam"`
		} `json:"mainAppWebResponseContext"`
		WebResponseContextExtensionData struct {
			HasDecorated bool `json:"hasDecorated"`
		} `json:"webResponseContextExtensionData"`
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
							} `json:"descriptionSnippet"`
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
							} `json:"thumbnailOverlays"`
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
				} `json:"continuationItemRenderer,omitempty"`
			} `json:"continuationItems"`
			TargetId string `json:"targetId"`
		} `json:"appendContinuationItemsAction"`
	} `json:"onResponseReceivedActions"`
}
