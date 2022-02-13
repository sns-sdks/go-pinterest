package pinterest

/*
	Boards API
*/

type BoardResource Resource

func newBoardResource(cli *Client) *BoardResource {
	return &BoardResource{Cli: cli}
}

// BoardOwner represents the owner for board
type BoardOwner struct {
	Username *string `json:"username"`
}

func (b BoardOwner) String() string {
	return Stringify(b)
}

// Board represents the board info
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/boards/get
type Board struct {
	ID          *string     `json:"id"`
	Name        *string     `json:"name"`
	Description *string     `json:"description"`
	Owner       *BoardOwner `json:"owner"`
	Privacy     *string     `json:"privacy"`
}

func (b Board) String() string {
	return Stringify(b)
}

// BoardsResponse represents the response for list boards
type BoardsResponse struct {
	Items    []*Board `json:"items"`
	Bookmark *string  `json:"bookmark"`
}

func (b BoardsResponse) String() string {
	return Stringify(b)
}

// ListBoardOpts represents the parameters for list boards
type ListBoardOpts struct {
	Bookmark string `url:"bookmark,omitempty"`
	PageSize int    `url:"page_size,omitempty"`
	Privacy  string `url:"privacy,omitempty"`
}

// ListBoards Get a list of the boards owned by the "operation user_account" + group boards where this account is a collaborator
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/boards/list
func (r *BoardResource) ListBoards(args ListBoardOpts) (*BoardsResponse, *APIError) {
	path := "/boards"

	resp := new(BoardsResponse)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetBoard Get a board owned by the operation user_account - or a group board that has been shared with this account.
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/boards/get
func (r *BoardResource) GetBoard(boardID string) (*Board, *APIError) {
	path := "/boards/" + boardID

	resp := new(Board)
	err := r.Cli.DoGet(path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateBoardOpts represents the parameters for create a board
type CreateBoardOpts struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Privacy     string `json:"privacy,omitempty"`
}

// CreateBoard Create a board owned by the "operation user_account".
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/boards/create
func (r *BoardResource) CreateBoard(args CreateBoardOpts) (*Board, *APIError) {
	path := "/boards"
	resp := new(Board)
	err := r.Cli.DoPost(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateBoardOpts represents the parameters for update board
type UpdateBoardOpts struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Privacy     string `json:"privacy,omitempty"`
}

// UpdateBoard Update a board owned by the "operating user_account".
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/boards/update
func (r *BoardResource) UpdateBoard(boardID string, args UpdateBoardOpts) (*Board, *APIError) {
	path := "/boards/" + boardID
	resp := new(Board)
	err := r.Cli.DoPatch(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteBoard Delete a board owned by the "operation user_account".
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/boards/delete
func (r *BoardResource) DeleteBoard(boardID string) *APIError {
	path := "/boards/" + boardID

	err := r.Cli.DoDelete(path, nil)
	if err != nil {
		return err
	}
	return nil
}
