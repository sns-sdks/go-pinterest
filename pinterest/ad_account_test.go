package pinterest

import (
	"github.com/jarcoal/httpmock"
)

func (bc *BCSuite) TestListAdAccounts() {
	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/ad_accounts",
		httpmock.NewStringResponder(
			400,
			`{"code":400,"message":"Invalid ad account ads parameters."}`,
		),
	)
	_, err := bc.Pin.AdAccount.ListAdAccounts(ListAdAccountsOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/ad_accounts",
		httpmock.NewStringResponder(
			200,
			`{"items":[{"id":"549763740754","name":"SNS-SDKS","owner":{"username":"merleliukun"},"country":"US","currency":"USD"}],"bookmark":null}`,
		),
	)

	adAccounts, _ := bc.Pin.AdAccount.ListAdAccounts(ListAdAccountsOpts{})
	bc.Equal(*adAccounts.Items[0].ID, "549763740754")
	bc.Equal(*adAccounts.Items[0].Owner.Username, "merleliukun")
}

func (bc *BCSuite) TestGetAdAccountAnalytics() {
	adAccountID := "12345678"
	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/ad_accounts/"+adAccountID+"/analytics",
		httpmock.NewStringResponder(
			400,
			`{"code":400,"message":"Invalid ad account analytics parameters."}`,
		),
	)
	_, err := bc.Pin.AdAccount.GetAdAccountAnalytics(adAccountID, GetAdAccountAnalyticsOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/ad_accounts/"+adAccountID+"/analytics",
		httpmock.NewStringResponder(
			200,
			`[{"DATE":"2021-04-01","AD_ACCOUNT_ID":"547602124502","SPEND_IN_DOLLAR":30,"TOTAL_CLICKTHROUGH":216}]`,
		),
	)

	analytics, _ := bc.Pin.AdAccount.GetAdAccountAnalytics(adAccountID, GetAdAccountAnalyticsOpts{})
	bc.Equal(analytics[0]["DATE"], "2021-04-01")
}
