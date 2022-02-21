package pinterest

import (
	"github.com/jarcoal/httpmock"
)

func (bc *BCSuite) TestListCampaigns() {
	adAccountID := "12345678"
	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/ad_accounts/"+adAccountID+"/campaigns",
		httpmock.NewStringResponder(
			400,
			`{"code":400,"message":"Invalid ad account campaign parameters."}`,
		),
	)
	_, err := bc.Pin.AdAccount.ListCampaigns(adAccountID, ListCampaignsOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/ad_accounts/"+adAccountID+"/campaigns",
		httpmock.NewStringResponder(
			200,
			`{"items":[{"id":"549755885175","ad_account_id":"549755885175","name":"ACME Tools","status":"ACTIVE","lifetime_spend_cap":1432744744,"daily_spend_cap":1432744744,"order_line_id":"549755885175","tracking_urls":{"impression":["URL1","URL2"],"click":["URL1","URL2"],"engagement":["URL1","URL2"],"buyable_button":["URL1","URL2"],"audience_verification":["URL1","URL2"]},"start_time":1580865126,"end_time":1644023526,"objective_type":"AWARENESS","created_time":1432744744,"updated_time":1432744744,"type":"campaign"}],"bookmark":"string"}`,
		),
	)

	campaigns, _ := bc.Pin.AdAccount.ListCampaigns(adAccountID, ListCampaignsOpts{})
	bc.Equal(*campaigns.Items[0].ID, "549755885175")
	bc.Equal(*campaigns.Items[0].TrackingURLs.Impression[0], "URL1")
}

func (bc *BCSuite) TestGetCampaignAnalytics() {
	adAccountID := "12345678"
	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/ad_accounts/"+adAccountID+"/campaigns/analytics",
		httpmock.NewStringResponder(
			400,
			`{"code":400,"message":"Invalid ad account campaign analytics parameters."}`,
		),
	)
	_, err := bc.Pin.AdAccount.GetCampaignAnalytics(adAccountID, GetCampaignAnalyticsOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/ad_accounts/"+adAccountID+"/campaigns/analytics",
		httpmock.NewStringResponder(
			200,
			`[{"DATE":"2021-04-01","CAMPAIGN_ID":"547602124502","SPEND_IN_DOLLAR":30,"TOTAL_CLICKTHROUGH":216}]`,
		),
	)

	analytics, _ := bc.Pin.AdAccount.GetCampaignAnalytics(adAccountID, GetCampaignAnalyticsOpts{})
	bc.Equal(analytics[0]["DATE"], "2021-04-01")
}
