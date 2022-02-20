package pinterest

// Ad represents the ad info.
type Ad struct {
	ID                                    *string       `json:"id"`
	Type                                  *string       `json:"type"`
	AdAccountID                           *string       `json:"ad_account_id"`
	AdGroupID                             *string       `json:"ad_group_id"`
	CampaignID                            *string       `json:"campaign_id"`
	PinID                                 *string       `json:"pin_id"`
	Name                                  *string       `json:"name"`
	Status                                *string       `json:"status"`
	AndroidDeepLink                       *string       `json:"android_deep_link"`
	IOSDeepLink                           *string       `json:"ios_deep_link"`
	CarouselAndroidDeepLinks              []*string     `json:"carousel_android_deep_links"`
	CarouselDestinationURLs               []*string     `json:"carousel_destination_urls"`
	CarouselIOSDeepLinks                  []*string     `json:"carousel_ios_deep_links"`
	ClickTrackingURL                      *string       `json:"click_tracking_url"`
	CreativeType                          *string       `json:"creative_type"`
	DestinationURL                        *string       `json:"destination_url"`
	IsPinDeleted                          *bool         `json:"is_pin_deleted"`
	IsRemovable                           *bool         `json:"is_removable"`
	TrackingURLs                          *TrackingURLs `json:"tracking_urls"`
	ViewTrackingURL                       *string       `json:"view_tracking_url"`
	CollectionItemsDestinationURLTemplate *string       `json:"collection_items_destination_url_template"`
	CreatedTime                           *int          `json:"created_time"`
	UpdatedTime                           *int          `json:"updated_time"`
	RejectedReasons                       []*string     `json:"rejected_reasons"`
	RejectionLabels                       []*string     `json:"rejection_labels"`
	ReviewStatus                          *string       `json:"review_status"`
	SummaryStatus                         *string       `json:"summary_status"`
}

func (a Ad) String() string {
	return Stringify(a)
}

// ListAdsResponse represents the response for list ads.
type ListAdsResponse struct {
	Items    []*Ad   `json:"items"`
	Bookmark *string `json:"bookmark"`
}

func (a ListAdsResponse) String() string {
	return Stringify(a)
}

// ListAdsOpts represents the parameters for list ads.
type ListAdsOpts struct {
	CampaignIDs               []string `url:"campaign_ids"`
	AdGroupIDs                []string `url:"ad_group_ids"`
	AdIDs                     []string `url:"ad_ids"`
	EntityStatuses            []string `url:"entity_statuses"`
	Order                     string   `url:"order"`
	TranslateInterestsToNames bool     `url:"translate_interests_to_names"`
	ListOptions
}

// ListAds Get a list of the ads in the specified ad_account_id, filtered by the specified options.
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/ads/list
func (r *AdAccountResource) ListAds(adAccountID string, args ListAdsOpts) (*ListAdsResponse, *APIError) {
	path := "/ad_accounts/" + adAccountID + "/ads"

	resp := new(ListAdsResponse)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetAdAnalyticsOpts represents the parameters for Get ad analytics.
type GetAdAnalyticsOpts struct {
	StartDate            string   `url:"start_date"`
	EndDate              string   `url:"end_date"`
	AdIDs                []string `url:"ad_ids"`
	Columns              []string `url:"columns"`
	Granularity          string   `url:"granularity"`
	ClickWindowDays      int      `url:"click_window_days,omitempty"`
	EngagementWindowDays int      `url:"engagement_window_days,omitempty"`
	ViewWindowDays       int      `url:"view_window_days,omitempty"`
	ConversionReportTime string   `url:"conversion_report_time,omitempty"`
}

// GetAdAnalytics Get analytics for the specified ads in the specified ad_account_id, filtered by the specified options.
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/ads/analytics
func (r *AdAccountResource) GetAdAnalytics(adAccountID string, args GetAdAnalyticsOpts) ([]map[string]string, *APIError) {
	path := "/ad_accounts/" + adAccountID + "/ads/analytics"

	var resp []map[string]string
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetProductGroupAnalyticsOpts represents the parameters for Get product group analytics.
type GetProductGroupAnalyticsOpts struct {
	StartDate            string   `url:"start_date"`
	EndDate              string   `url:"end_date"`
	ProductGroupIDs      []string `url:"product_group_ids"`
	Columns              []string `url:"columns"`
	Granularity          string   `url:"granularity"`
	ClickWindowDays      int      `url:"click_window_days,omitempty"`
	EngagementWindowDays int      `url:"engagement_window_days,omitempty"`
	ViewWindowDays       int      `url:"view_window_days,omitempty"`
	ConversionReportTime string   `url:"conversion_report_time,omitempty"`
}

// GetProductGroupAnalytics Get analytics for the specified product groups in the specified ad_account_id, filtered by the specified options.
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/product_groups/analytics
func (r *AdAccountResource) GetProductGroupAnalytics(adAccountID string, args GetAdAnalyticsOpts) ([]map[string]string, *APIError) {
	path := "/ad_accounts/" + adAccountID + "/product_groups/analytics"

	var resp []map[string]string
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
