package pinterest

/*
	Pin API
*/

type PinResource Resource

func newPinResource(cli *Client) *PinResource {
	return &PinResource{Cli: cli}
}

// Pin represents the pin info.
type Pin struct {
	ID             *string     `json:"id"`
	CreatedAt      *string     `json:"created_at"`
	Link           *string     `json:"link"`
	Title          *string     `json:"title"`
	Description    *string     `json:"description"`
	AltText        *string     `json:"alt_text"`
	BoardID        *string     `json:"board_id"`
	BoardSectionID *string     `json:"board_section_id"`
	BoardOwner     *BoardOwner `json:"board_owner"`
	Media          *Media      `json:"media"`
}

func (p Pin) String() string {
	return Stringify(p)
}

// PinsResponse represents the response for list pins
type PinsResponse struct {
	Items    []*Pin  `json:"items"`
	Bookmark *string `json:"bookmark"`
}

func (p PinsResponse) String() string {
	return Stringify(p)
}

// CreatePinMediaSourceOpts represents the parameters for pin media resource
type CreatePinMediaSourceOpts struct {
	SourceType    string `json:"source_type"`
	ContentType   string `json:"content_type,omitempty"`
	Data          string `json:"data,omitempty,omitempty"`
	Url           string `json:"url,omitempty"`
	CoverImageURL string `json:"cover_image_url,omitempty"`
	MediaID       string `json:"media_id,omitempty"`
}

// CreatePinOpts represents the parameters for create a pin
type CreatePinOpts struct {
	Link           string                   `json:"link,omitempty"`
	Title          string                   `json:"title,omitempty"`
	Description    string                   `json:"description,omitempty"`
	AltText        string                   `json:"alt_text,omitempty"`
	BoardID        string                   `json:"board_id"`
	BoardSectionID string                   `json:"board_section_id,omitempty"`
	MediaSource    CreatePinMediaSourceOpts `json:"media_source"`
}

// CreatePin Create a Pin on a board or board section owned by the "operation user_account".
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/pins/get
func (r *PinResource) CreatePin(args CreatePinOpts) (*Pin, *APIError) {
	path := "/pins"

	resp := new(Pin)
	err := r.Cli.DoPost(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// getPinOpts represents the parameters for get pin
type getPinOpts struct {
	AdAccountID string `url:"ad_account_id,omitempty"`
}

// GetPin Get a Pin owned by the "operation user_account" - or on a group board that has been shared with this account.
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/pins/get
func (r *PinResource) GetPin(pinID, adAccountID string) (*Pin, *APIError) {
	path := "/pins/" + pinID

	resp := new(Pin)
	var err *APIError
	if adAccountID != "" {
		params := getPinOpts{AdAccountID: adAccountID}
		err = r.Cli.DoGet(path, params, resp)
	} else {
		err = r.Cli.DoGet(path, nil, resp)
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeletePin Delete a Pins owned by the "operation user_account" - or on a group board that has been shared with this account.
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/pins/delete
func (r *PinResource) DeletePin(pinID string) *APIError {
	path := "/pins/" + pinID
	err := r.Cli.DoDelete(path, nil)
	if err != nil {
		return err
	}
	return nil
}
