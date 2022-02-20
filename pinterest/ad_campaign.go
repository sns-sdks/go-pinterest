package pinterest

// TrackingURLs represents the Third-party tracking URLs.
type TrackingURLs struct {
	Impression           []*string `json:"impression"`
	Click                []*string `json:"click"`
	Engagement           []*string `json:"engagement"`
	BuyableButton        []*string `json:"buyable_button"`
	AudienceVerification []*string `json:"audience_verification"`
}

func (t TrackingURLs) String() string {
	return Stringify(t)
}

// Campaign represents the campaign info.
type Campaign struct {
	ID               *string       `json:"id"`
	Type             *string       `json:"type"`
	AdAccountID      *string       `json:"ad_account_id"`
	Name             *string       `json:"name"`
	Status           *string       `json:"status"`
	LifetimeSpendCap *int          `json:"lifetime_spend_cap"`
	DailySpendCap    *int          `json:"daily_spend_cap"`
	OrderLineID      *string       `json:"order_line_id"`
	TrackingURLs     *TrackingURLs `json:"tracking_urls"`
	StartTime        *int          `json:"start_time"`
	EndTime          *int          `json:"end_time"`
	ObjectiveType    *string       `json:"objective_type"`
	CreatedTime      *int          `json:"created_time"`
	UpdatedTime      *int          `json:"updated_time"`
}

func (c Campaign) String() string {
	return Stringify(c)
}

// ListCampaignsResponse represents the response for list campaigns.
type ListCampaignsResponse struct {
	Items    []*Campaign `json:"items"`
	Bookmark *string     `json:"bookmark"`
}

func (c ListCampaignsResponse) String() string {
	return Stringify(c)
}

// ListCampaignsOpts represents the parameters for list campaigns.
type ListCampaignsOpts struct {
	CampaignIDs    []string `url:"campaign_ids"`
	EntityStatuses []string `url:"entity_statuses"`
	Order          string   `url:"order"`
	ListOptions
}

// ListCampaigns Get a list of the campaigns in the specified ad_account_id, filtered by the specified options.
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/campaigns/list
func (r *AdAccountResource) ListCampaigns(adAccountID string, args ListCampaignsOpts) (*ListCampaignsResponse, *APIError) {
	path := "/ad_accounts/" + adAccountID + "/campaigns"

	resp := new(ListCampaignsResponse)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetCampaignAnalyticsOpts represents the parameters for Get campaign analytics.
type GetCampaignAnalyticsOpts struct {
	StartDate            string   `url:"start_date"`
	EndDate              string   `url:"end_date"`
	CampaignIDs          []string `url:"campaign_ids"`
	Columns              []string `url:"columns"`
	Granularity          string   `url:"granularity"`
	ClickWindowDays      int      `url:"click_window_days,omitempty"`
	EngagementWindowDays int      `url:"engagement_window_days,omitempty"`
	ViewWindowDays       int      `url:"view_window_days,omitempty"`
	ConversionReportTime string   `url:"conversion_report_time,omitempty"`
}

// GetCampaignAnalytics Get analytics for the specified campaigns in the specified ad_account_id, filtered by the specified options.
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/campaigns/analytics
func (r *AdAccountResource) GetCampaignAnalytics(adAccountID string, args GetCampaignAnalyticsOpts) ([]map[string]string, *APIError) {
	path := "/ad_accounts/" + adAccountID + "/campaigns/analytics"

	var resp []map[string]string
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
