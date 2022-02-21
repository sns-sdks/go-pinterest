package pinterest

import (
	"github.com/jarcoal/httpmock"
)

func (bc *BCSuite) TestListAds() {
	adAccountID := "12345678"
	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/ad_accounts/"+adAccountID+"/ads",
		httpmock.NewStringResponder(
			400,
			`{"code":400,"message":"Invalid ad account ads parameters."}`,
		),
	)
	_, err := bc.Pin.AdAccount.ListAds(adAccountID, ListAdsOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/ad_accounts/"+adAccountID+"/ads",
		httpmock.NewStringResponder(
			200,
			`{"items":[{"ad_group_id":"2680059592705","android_deep_link":"string","carousel_android_deep_links":["string"],"carousel_destination_urls":["string"],"carousel_ios_deep_links":["string"],"click_tracking_url":"string","creative_type":"REGULAR","destination_url":"string","ios_deep_link":"string","is_pin_deleted":false,"is_removable":false,"name":"string","pin_id":"394205773611545468","status":"ACTIVE","tracking_urls":{"impression":["URL1","URL2"],"click":["URL1","URL2"],"engagement":["URL1","URL2"],"buyable_button":["URL1","URL2"],"audience_verification":["URL1","URL2"]},"view_tracking_url":"string","ad_account_id":"549755885175","campaign_id":"626735565838","collection_items_destination_url_template":"string","created_time":1451431341,"id":"687195134316","rejected_reasons":["HASHTAGS"],"rejection_labels":["string"],"review_status":"PENDING","type":"pinpromotion","updated_time":1451431341,"summary_status":"APPROVED"}],"bookmark":null}`,
		),
	)

	ads, _ := bc.Pin.AdAccount.ListAds(adAccountID, ListAdsOpts{})
	bc.Equal(*ads.Items[0].ID, "687195134316")
	bc.Equal(*ads.Items[0].TrackingURLs.Impression[0], "URL1")
	bc.Nil(ads.Bookmark)
}

func (bc *BCSuite) TestGetAdAnalytics() {
	adAccountID := "12345678"
	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/ad_accounts/"+adAccountID+"/ads/analytics",
		httpmock.NewStringResponder(
			400,
			`{"code":400,"message":"Invalid ad account ads analytics parameters."}`,
		),
	)
	_, err := bc.Pin.AdAccount.GetAdAnalytics(adAccountID, GetAdAnalyticsOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/ad_accounts/"+adAccountID+"/ads/analytics",
		httpmock.NewStringResponder(
			200,
			`[{"DATE":"2021-04-01","AD_ID":"547602124502","SPEND_IN_DOLLAR":30,"TOTAL_CLICKTHROUGH":216}]`,
		),
	)

	analytics, _ := bc.Pin.AdAccount.GetAdAnalytics(adAccountID, GetAdAnalyticsOpts{})
	bc.Equal(analytics[0]["DATE"], "2021-04-01")
}

func (bc *BCSuite) TestGetProductGroupAnalytics() {
	adAccountID := "12345678"
	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/ad_accounts/"+adAccountID+"/product_groups/analytics",
		httpmock.NewStringResponder(
			400,
			`{"code":400,"message":"Invalid ad account ads analytics parameters."}`,
		),
	)
	_, err := bc.Pin.AdAccount.GetProductGroupAnalytics(adAccountID, GetProductGroupAnalyticsOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/ad_accounts/"+adAccountID+"/product_groups/analytics",
		httpmock.NewStringResponder(
			200,
			`[{"DATE":"2021-04-01","AD_ID":"547602124502","SPEND_IN_DOLLAR":30,"TOTAL_CLICKTHROUGH":216}]`,
		),
	)

	analytics, _ := bc.Pin.AdAccount.GetProductGroupAnalytics(adAccountID, GetProductGroupAnalyticsOpts{})
	bc.Equal(analytics[0]["DATE"], "2021-04-01")
}
