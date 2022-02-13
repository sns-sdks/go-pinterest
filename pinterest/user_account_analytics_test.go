package pinterest

import (
	"github.com/jarcoal/httpmock"
)

func (bc *BCSuite) TestGetUserAccountAnalytics() {
	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/user_account/analytics",
		httpmock.NewStringResponder(
			403,
			`{"code":1,"message":"Parameter 'start_date' is required."}`,
		),
	)
	_, err := bc.Pin.UserAccount.GetUserAccountAnalytics(UserAccountAnalyticsOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/user_account/analytics",
		httpmock.NewStringResponder(
			200,
			`{"all":{"summary_metrics":{"IMPRESSION":3,"ENGAGEMENT":1,"ENGAGEMENT_RATE":0.3333333333333333,"SAVE":0,"SAVE_RATE":0.0,"PIN_CLICK":1,"PIN_CLICK_RATE":0.3333333333333333,"OUTBOUND_CLICK":0,"OUTBOUND_CLICK_RATE":0.0},"daily_metrics":[{"date":"2022-02-10","data_status":"READY","metrics":{"IMPRESSION":3,"ENGAGEMENT":1,"ENGAGEMENT_RATE":0.3333333333333333,"SAVE":0,"SAVE_RATE":0.0,"PIN_CLICK":1,"PIN_CLICK_RATE":0.3333333333333333,"OUTBOUND_CLICK":0,"OUTBOUND_CLICK_RATE":0.0}}]}}`,
		),
	)

	analytics, _ := bc.Pin.UserAccount.GetUserAccountAnalytics(UserAccountAnalyticsOpts{StartDate: "2022-02-10", EndDate: "2022-02-10"})
	bc.Equal(*analytics.All.SummaryMetrics.Impression, *Int64(3))
	bc.Equal(*analytics.All.DailyMetrics[0].Date, "2022-02-10")
	bc.Equal(*analytics.All.DailyMetrics[0].Metrics.SaveRate, 0.0)
}
