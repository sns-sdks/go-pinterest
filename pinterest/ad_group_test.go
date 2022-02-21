package pinterest

import (
	"github.com/jarcoal/httpmock"
)

func (bc *BCSuite) TestListAdGroups() {
	adAccountID := "12345678"
	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/ad_accounts/"+adAccountID+"/ad_groups",
		httpmock.NewStringResponder(
			400,
			`{"code":400,"message":"Invalid ad account group parameters."}`,
		),
	)
	_, err := bc.Pin.AdAccount.ListAdGroups(adAccountID, ListAdGroupsOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/ad_accounts/"+adAccountID+"/ad_groups",
		httpmock.NewStringResponder(
			200,
			`{"items":[{"name":"Ad Group For Pin: 687195905986","status":"ACTIVE","budget_in_micro_currency":5000000,"bid_in_micro_currency":5000000,"budget_type":"DAILY","start_time":5686848000,"end_time":5705424000,"targeting_spec":{"property1":["string"],"property2":["string"]},"lifetime_frequency_cap":100,"tracking_urls":{"impression":["URL1","URL2"],"click":["URL1","URL2"],"engagement":["URL1","URL2"],"buyable_button":["URL1","URL2"],"audience_verification":["URL1","URL2"]},"auto_targeting_enabled":true,"placement_group":"ALL","pacing_delivery_type":"STANDARD","conversion_learning_mode_type":"ACTIVE","summary_status":"RUNNING","feed_profile_id":"626736533506","campaign_id":"626736533506","billable_event":"CLICKTHROUGH","id":"2680060704746","type":"string","ad_account_id":"549755885175","created_time":1476477189,"updated_time":1476477189}],"bookmark":"string"}`,
		),
	)

	adGroups, _ := bc.Pin.AdAccount.ListAdGroups(adAccountID, ListAdGroupsOpts{})
	bc.Equal(*adGroups.Items[0].ID, "2680060704746")
	bc.Equal(*adGroups.Items[0].TrackingURLs.Impression[0], "URL1")
}

func (bc *BCSuite) TestGetAdGroupAnalytics() {
	adAccountID := "12345678"
	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/ad_accounts/"+adAccountID+"/ad_groups/analytics",
		httpmock.NewStringResponder(
			400,
			`{"code":400,"message":"Invalid ad account group analytics parameters."}`,
		),
	)
	_, err := bc.Pin.AdAccount.GetAdGroupAnalytics(adAccountID, GetAdGroupAnalyticsOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/ad_accounts/"+adAccountID+"/ad_groups/analytics",
		httpmock.NewStringResponder(
			200,
			`[{"DATE":"2021-04-01","AD_ID":"547602124502","SPEND_IN_DOLLAR":30,"TOTAL_CLICKTHROUGH":216}]`,
		),
	)

	analytics, _ := bc.Pin.AdAccount.GetAdGroupAnalytics(adAccountID, GetAdGroupAnalyticsOpts{})
	bc.Equal(analytics[0]["DATE"], "2021-04-01")
}
