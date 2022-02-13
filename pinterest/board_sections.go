package pinterest

// BoardSection represents the board section info
type BoardSection struct {
	ID   *string `json:"id"`
	Name *string `json:"name"`
}

func (b BoardSection) String() string {
	return Stringify(b)
}

// BoardSectionsResponse represents the response for list board sections
type BoardSectionsResponse struct {
	Items    []*BoardSection `json:"items"`
	Bookmark *string         `json:"bookmark"`
}

func (b BoardSectionsResponse) String() string {
	return Stringify(b)
}

// ListBoardSections Get a list of all board sections from a board owned by the "operation user_account" - or a group board that has been shared with this account.
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/board_sections/list
func (r *BoardResource) ListBoardSections(boardID string, args ListOptions) (*BoardSectionsResponse, *APIError) {
	path := "/boards/" + boardID + "/sections"

	resp := new(BoardSectionsResponse)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateBoardSectionOpts represents the parameter for create or update board section.
type CreateBoardSectionOpts struct {
	Name string `json:"name"`
}

// CreateBoardSection Create a board section on a board owned by the "operation user_account" - or on a group board that has been shared with this account.
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/board_sections/create
func (r *BoardResource) CreateBoardSection(boardID string, args CreateBoardSectionOpts) (*BoardSection, *APIError) {
	path := "/boards/" + boardID + "/sections"

	resp := new(BoardSection)
	err := r.Cli.DoPost(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateBoardSection Update a board section on a board owned by the "operation user_account" - or on a group board that has been shared with this account.
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/board_sections/update
func (r *BoardResource) UpdateBoardSection(boardID, sectionID string, args CreateBoardSectionOpts) (*BoardSection, *APIError) {
	path := "/boards/" + boardID + "/sections/" + sectionID

	resp := new(BoardSection)
	err := r.Cli.DoPatch(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteBoardSection Delete a board section on a board owned by the "operation user_account" - or on a group board that has been shared with this account.
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/board_sections/delete
func (r *BoardResource) DeleteBoardSection(boardID, sectionID string) *APIError {
	path := "/boards/" + boardID + "/sections/" + sectionID

	err := r.Cli.DoDelete(path, nil)
	if err != nil {
		return err
	}
	return nil
}

// ListPinsOnBoardSection Get a list of the Pins on a board section of a board owned by the "operation user_account" - or on a group board that has been shared with this account.
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/board_sections/list_pins
func (r *BoardResource) ListPinsOnBoardSection(boardID, sectionID string, args ListOptions) (*PinsResponse, *APIError) {
	path := "/boards/" + boardID + "/sections/" + sectionID + "/pins"

	resp := new(PinsResponse)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
