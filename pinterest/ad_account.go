package pinterest

/*
	Ad Accounts API
*/

type AdAccountResource Resource

func newAdAccountResource(cli *Client) *AdAccountResource {
	return &AdAccountResource{Cli: cli}
}

type AdAccountOwner struct {
	Username *string `json:"username"`
}

// AdAccount represents the ad account info.
type AdAccount struct {
	ID       *string         `json:"id"`
	Name     *string         `json:"name"`
	Owner    *AdAccountOwner `json:"owner"`
	Country  *string         `json:"country"`
	Currency *string         `json:"currency"`
}

func (a AdAccount) String() string {
	return Stringify(a)
}

// AdAccountsResponse represents the response for list ad accounts.
type AdAccountsResponse struct {
	Items    []*AdAccount `json:"items"`
	Bookmark *string      `json:"bookmark"`
}

func (ar AdAccountsResponse) String() string {
	return Stringify(ar)
}

// ListAdAccountsOptions represents the parameters for list ad accounts.
type ListAdAccountsOptions struct {
	IncludeSharedAccounts bool `url:"include_shared_accounts,omitempty"`
	ListOptions
}

// ListAdAccounts Get a list of the ad_accounts that the "operation user_account" has access to.
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/ad_accounts/list
func (r *AdAccountResource) ListAdAccounts(args ListAdAccountsOptions) (*AdAccountsResponse, *APIError) {
	path := "/ad_accounts"

	resp := new(AdAccountsResponse)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetAdAccountAnalyticsOpts represents the parameters for Get ad account analytics.
type GetAdAccountAnalyticsOpts struct {
	StartDate            string   `url:"start_date"`
	EndDate              string   `url:"end_date"`
	Columns              []string `url:"columns"`
	Granularity          string   `url:"granularity"`
	ClickWindowDays      int      `url:"click_window_days,omitempty"`
	EngagementWindowDays int      `url:"engagement_window_days,omitempty"`
	ViewWindowDays       int      `url:"view_window_days,omitempty"`
	ConversionReportTime string   `url:"conversion_report_time,omitempty"`
}

// GetAdAccountAnalytics Get analytics for the specified ad_account_id, filtered by the specified options.
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/ad_account/analytics
func (r *AdAccountResource) GetAdAccountAnalytics(adAccountID string, args GetAdAccountAnalyticsOpts) ([]map[string]string, *APIError) {
	path := "/ad_accounts/" + adAccountID + "/analytics"

	var resp []map[string]string
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
