package pinterest

import (
	"github.com/go-resty/resty/v2"
	"net/http"
)

const (
	Baseurl    = "https://api.pinterest.com/v5"
	HttpGet    = resty.MethodGet
	HttpPost   = resty.MethodPost
	HttpPatch  = resty.MethodPatch
	HttpDelete = resty.MethodDelete

	OAuthState = "go-pinterest"
)

type Client struct {
	Cli *resty.Client
	// API Resource
	UserAccount *UserAccountResource
	Board       *BoardResource
	Pin         *PinResource
	Media       *MediaResource
}

type Resource struct {
	Cli *Client
}

func NewClient(client *resty.Client) *Client {
	c := &Client{Cli: client}

	// Register data resource
	c.UserAccount = newUserAccountResource(c)
	c.Board = newBoardResource(c)
	c.Pin = newPinResource(c)
	c.Media = newMediaResource(c)
	return c
}

func NewBearerClient(bearerToken string) *Client {
	rCli := resty.New()
	rCli.SetAuthToken(bearerToken)

	return NewClient(rCli)
}

func NewUserClint(hc *http.Client) *Client {
	rCli := resty.NewWithClient(hc)
	return NewClient(rCli)
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }
