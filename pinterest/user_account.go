package pinterest

/*
	Users API
*/

type UserAccountResource Resource

func newUserAccountResource(cli *Client) *UserAccountResource {
	return &UserAccountResource{
		Cli: cli,
	}
}

// UserAccount represent a Pinterest user account
// Refer: https://developers.pinterest.com/docs/api/v5/#tag/user_account
type UserAccount struct {
	AccountType  *string `json:"account_type"`
	ProfileImage *string `json:"profile_image"`
	WebsiteURL   *string `json:"website_url"`
	Username     *string `json:"username"`
}

func (u UserAccount) String() string {
	return Stringify(u)
}

type userAccountOpts struct {
	AdAccountID string `url:"ad_account_id"`
}

// GetUserAccount Get account information for the "operation user_account"
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/user_account/get
func (r *UserAccountResource) GetUserAccount(AdAccountID string) (*UserAccount, *APIError) {
	path := "/user_account"

	resp := new(UserAccount)
	var err *APIError
	if AdAccountID != "" {
		params := userAccountOpts{AdAccountID: AdAccountID}
		err = r.Cli.DoGet(path, params, resp)
	} else {
		err = r.Cli.DoGet(path, nil, resp)
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}
