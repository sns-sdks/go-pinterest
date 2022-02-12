package pinterest

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"
	"golang.org/x/oauth2"
	"net/http"
	"testing"
)

type Auth2Suite struct {
	suite.Suite
	app *AuthorizationAPP
}

func (auth *Auth2Suite) SetupSuite() {
	auth.app = NewAuthorizationAPP(AuthorizationAPP{
		ClientID:     "client id",
		ClientSecret: "client secret",
		RedirectURI:  "https://localhost/",
		Scope:        "boards:read,pins:read",
	})
}

func (auth *Auth2Suite) SetupTest() {
	httpmock.ActivateNonDefault(http.DefaultClient)
}

func (auth *Auth2Suite) TearDownTest() {
	httpmock.DeactivateAndReset()
}

func TestAuth2Suite(t *testing.T) {
	suite.Run(t, new(Auth2Suite))
}

func (auth *Auth2Suite) TestGetAuthorizationURL() {
	authUrl := auth.app.GetAuthorizationURL()
	auth.NotNil(authUrl)
}

func (auth *Auth2Suite) TestGenerateAccessToken() {
	httpmock.RegisterResponder(
		HttpPost, Endpoint.TokenURL,
		httpmock.NewStringResponder(400, `{"code": -1,"message":"error"}`),
	)
	_, err := auth.app.GenerateAccessToken("code")
	auth.Contains(err.Error(), "error")

	httpmock.RegisterResponder(
		HttpPost, Endpoint.TokenURL,
		httpmock.NewStringResponder(
			200,
			`{"response_type":"authorization_code","access_token":"string","token_type":"bearer","expires_in":0,"scope":"string","refresh_token":"string","refresh_token_expires_in":0}`,
		),
	)

	token, _ := auth.app.GenerateAccessToken("code")
	auth.IsType(&oauth2.Token{}, token)

	cli := auth.app.GetUserClient()
	auth.IsType(&Client{}, cli)
}
