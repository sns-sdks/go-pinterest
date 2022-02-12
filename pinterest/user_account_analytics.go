package pinterest

/*
	User Account Analytics API
	Refer: https://developers.pinterest.com/docs/api/v5/#operation/user_account/analytics
*/

// Metrics represents the metrics info.
type Metrics struct {
	Engagement        *int64   `json:"ENGAGEMENT"`
	EngagementRate    *float64 `json:"ENGAGEMENT_RATE"`
	ClickThrough      *int64   `json:"CLICKTHROUGH"`
	ClickThroughRate  *float64 `json:"CLICKTHROUGH_RATE"`
	CloseUp           *int64   `json:"CLOSEUP"`
	CloseUpRate       *float64 `json:"CLOSEUP_RATE"`
	Save              *int64   `json:"SAVE"`
	SaveRate          *float64 `json:"SAVE_RATE"`
	Impression        *int64   `json:"IMPRESSION"`
	OutboundClick     *int64   `json:"OUTBOUND_CLICK"`
	OutboundClickRate *float64 `json:"OUTBOUND_CLICK_RATE"`
	PinClick          *int64   `json:"PIN_CLICK"`
	PinClickRate      *float64 `json:"PIN_CLICK_RATE"`
}

func (m Metrics) String() string {
	return Stringify(m)
}

// DailyMetrics represents the metrics info for a date.
type DailyMetrics struct {
	DataStatus *string  `json:"data_status"`
	Date       *string  `json:"date"`
	Metrics    *Metrics `json:"metrics"`
}

func (dm DailyMetrics) String() string {
	return Stringify(dm)
}

// UserAccountAnalyticsMetrics represents the metrics info for days.
type UserAccountAnalyticsMetrics struct {
	DailyMetrics   []*DailyMetrics `json:"daily_metrics"`
	SummaryMetrics *Metrics        `json:"summary_metrics"`
}

func (m UserAccountAnalyticsMetrics) String() string {
	return Stringify(m)
}

// UserAccountAnalytics represents the reponse for the user account analytics.
type UserAccountAnalytics struct {
	All *UserAccountAnalyticsMetrics `json:"all"`
}

func (m UserAccountAnalytics) String() string {
	return Stringify(m)
}

// UserAccountAnalyticsOpts the parameters for the user account analytics.
type UserAccountAnalyticsOpts struct {
	StartDate          string `url:"start_date"`
	EndDate            string `url:"end_date"`
	FromClaimedContent string `url:"from_claimed_content,omitempty"`
	PinFormat          string `url:"pin_format,omitempty"`
	AppTypes           string `url:"app_types,omitempty"`
	MetricTypes        string `url:"metric_types,omitempty"`
	SplitField         string `url:"split_field,omitempty"`
	AdAccountID        string `url:"ad_account_id,omitempty"`
}

// GetUserAccountAnalytics Get analytics for the user account.
func (r *UserAccountResource) GetUserAccountAnalytics(args UserAccountAnalyticsOpts) (*UserAccountAnalytics, *APIError) {
	path := "/user_account/analytics"

	resp := new(UserAccountAnalytics)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
