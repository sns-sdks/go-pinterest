package pinterest

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"
	"testing"
)

func (bc *BCSuite) TestGetUserAccount() {
	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/user_account",
		httpmock.NewStringResponder(
			403,
			`{"code":403,"message":"Not authorized to access the user account."}`,
		),
	)
	_, err := bc.Pin.UserAccount.GetUserAccount("")
	bc.IsType(&APIError{}, err)
	bc.Contains(err.Error(), "403")

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/user_account",
		httpmock.NewStringResponder(
			200,
			`{"account_type":"BUSINESS","profile_image":"https://i.pinimg.com/600x600_R/58/02/68/580268e48c9b3d4c906cfd88aa12cc00.jpg","website_url":"https://example.com/","username":"kun"}`,
		),
	)

	user, _ := bc.Pin.UserAccount.GetUserAccount("")
	bc.Equal(*user.AccountType, "BUSINESS")
	bc.Equal(*user.Username, "kun")
	user, _ = bc.Pin.UserAccount.GetUserAccount("123456")
	bc.Equal(*user.AccountType, "BUSINESS")
	bc.Equal(*user.Username, "kun")
}

func TestBCSuite(t *testing.T) {
	suite.Run(t, new(BCSuite))
}
