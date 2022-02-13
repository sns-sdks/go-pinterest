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
