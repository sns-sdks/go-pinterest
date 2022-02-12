package pinterest

import (
	"context"
	"golang.org/x/oauth2"
	"net/http"
)

// Endpoint url for pinterest oauth2
var Endpoint = oauth2.Endpoint{
	AuthURL:   "https://www.pinterest.com/oauth/",
	TokenURL:  "https://api.pinterest.com/v5/oauth/token",
	AuthStyle: oauth2.AuthStyleInHeader,
}

// AuthorizationAPP Pinterest OAuth2 app config
type AuthorizationAPP struct {
	ClientID     string         `json:"client_id"`
	ClientSecret string         `json:"client_secret"`
	RedirectURI  string         `json:"redirect_uri,omitempty"`
	Scope        string         `json:"scope,omitempty"`
	Token        *oauth2.Token  `json:"access_token,omitempty"`
	Config       *oauth2.Config `json:"config,omitempty"`
}

func (app AuthorizationAPP) String() string {
	return Stringify(app)
}

// NewAuthorizationAPP Return app for oauth2 authorization
func NewAuthorizationAPP(app AuthorizationAPP) *AuthorizationAPP {
	app.Config = &oauth2.Config{
		ClientID:     app.ClientID,
		ClientSecret: app.ClientSecret,
		RedirectURL:  app.RedirectURI,
		Scopes:       []string{app.Scope},
		Endpoint:     Endpoint,
	}
	return &app
}

// GetAuthorizationURL Return authorization url for user
func (app *AuthorizationAPP) GetAuthorizationURL() string {
	return app.Config.AuthCodeURL(OAuthState)
}

// GenerateAccessToken Generate user access token for the app
func (app *AuthorizationAPP) GenerateAccessToken(code string) (*oauth2.Token, error) {
	ctx := context.Background()
	token, err := app.Config.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	app.Token = token
	return token, err
}

// GetAuthorizedHttpClient Get user authorized http client
func (app *AuthorizationAPP) GetAuthorizedHttpClient() *http.Client {
	hc := app.Config.Client(context.TODO(), app.Token)
	return hc
}

// GetUserClient get library client with user authorization
func (app *AuthorizationAPP) GetUserClient() *Client {
	hc := app.GetAuthorizedHttpClient()
	return NewUserClint(hc)
}
