package pinterest

// AdGroup represents the ad group info.
type AdGroup struct {
	ID                         *string              `json:"id"`
	Type                       *string              `json:"type"`
	AdAccountID                *string              `json:"ad_account_id"`
	Name                       *string              `json:"name"`
	Status                     *string              `json:"status"`
	BudgetInMicroCurrency      *int                 `json:"budget_in_micro_currency"`
	BidInMicroCurrency         *int                 `json:"bid_in_micro_currency"`
	BudgetType                 *string              `json:"budget_type"`
	StartTime                  *int                 `json:"start_time"`
	EndTime                    *int                 `json:"end_time"`
	TargetingSpec              *map[string][]string `json:"targeting_spec"`
	LifetimeFrequencyCap       *int                 `json:"lifetime_frequency_cap"`
	TrackingURLs               *TrackingURLs        `json:"tracking_urls"`
	AutoTargetingEnabled       *bool                `json:"auto_targeting_enabled"`
	PlacementGroup             *string              `json:"placement_group"`
	PacingDeliveryType         *string              `json:"pacing_delivery_type"`
	ConversionLearningModeType *string              `json:"conversion_learning_mode_type"`
	SummaryStatus              *string              `json:"summary_status"`
	FeedProfileID              *string              `json:"feed_profile_id"`
	CampaignID                 *string              `json:"campaign_id"`
	BillableEvent              *string              `json:"billable_event"`
	CreatedTime                *int                 `json:"created_time"`
	UpdatedTime                *int                 `json:"updated_time"`
}

func (a AdGroup) String() string {
	return Stringify(a)
}

// ListAdGroupsResponse represents the response for list ad groups.
type ListAdGroupsResponse struct {
	Items    []*AdGroup `json:"items"`
	Bookmark *string    `json:"bookmark"`
}

func (c ListAdGroupsResponse) String() string {
	return Stringify(c)
}

// ListAdGroupsOpts represents the parameters for list ad groups.
type ListAdGroupsOpts struct {
	CampaignIDs               []string `url:"campaign_ids"`
	AdGroupIDs                []string `url:"ad_group_ids"`
	EntityStatuses            []string `url:"entity_statuses"`
	Order                     string   `url:"order"`
	TranslateInterestsToNames bool     `url:"translate_interests_to_names"`
	ListOptions
}

// ListAdGroups Get a list of the ad groups in the specified ad_account_id, filtered by the specified options.
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/ad_groups/list
func (r *AdAccountResource) ListAdGroups(adAccountID string, args ListAdGroupsOpts) (*ListAdGroupsResponse, *APIError) {
	path := "/ad_accounts/" + adAccountID + "/ad_groups"

	resp := new(ListAdGroupsResponse)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetAdGroupAnalyticsOpts represents the parameters for Get ad group analytics.
type GetAdGroupAnalyticsOpts struct {
	StartDate            string   `url:"start_date"`
	EndDate              string   `url:"end_date"`
	AdGroupIDs           []string `url:"ad_group_ids"`
	Columns              []string `url:"columns"`
	Granularity          string   `url:"granularity"`
	ClickWindowDays      int      `url:"click_window_days,omitempty"`
	EngagementWindowDays int      `url:"engagement_window_days,omitempty"`
	ViewWindowDays       int      `url:"view_window_days,omitempty"`
	ConversionReportTime string   `url:"conversion_report_time,omitempty"`
}

// GetAdGroupAnalytics Get analytics for the specified campaigns in the specified ad_account_id, filtered by the specified options.
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/ad_groups/analytics
func (r *AdAccountResource) GetAdGroupAnalytics(adAccountID string, args GetCampaignAnalyticsOpts) ([]map[string]string, *APIError) {
	path := "/ad_accounts/" + adAccountID + "/ad_groups/analytics"

	var resp []map[string]string
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
