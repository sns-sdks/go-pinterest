package pinterest

import "github.com/jarcoal/httpmock"

func (bc *BCSuite) TestListBoards() {
	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/boards",
		httpmock.NewStringResponder(
			403,
			`{"code":403,"message":"Not authorized to access the user account."}`,
		),
	)
	_, err := bc.Pin.BoardResource.ListBoards(ListBoardOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/boards",
		httpmock.NewStringResponder(
			200,
			`{"items":[{"owner":{"username":"merleliukun"},"description":"","name":"City","privacy":"PUBLIC","id":"1022106146619699845"},{"owner":{"username":"merleliukun"},"description":"","name":"Food","privacy":"PUBLIC","id":"1022106146619703648"}],"bookmark":null}`,
		),
	)

	boards, _ := bc.Pin.BoardResource.ListBoards(ListBoardOpts{})
	bc.Equal(*boards.Items[0].Privacy, "PUBLIC")
	bc.Nil(boards.Bookmark)
}

func (bc *BCSuite) TestGetBoard() {
	boardID := "1022106146619729163"
	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/boards/"+boardID,
		httpmock.NewStringResponder(
			404,
			`{"code":404,"message":"Board not found."}`,
		),
	)
	_, err := bc.Pin.BoardResource.GetBoard(boardID)
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/boards/"+boardID,
		httpmock.NewStringResponder(
			200,
			`{"description":"","name":"ToBeDelete","owner":{"username":"merleliukun"},"id":"1022106146619729163","privacy":"PUBLIC"}`,
		),
	)

	board, _ := bc.Pin.BoardResource.GetBoard(boardID)
	bc.Equal(*board.Privacy, "PUBLIC")
	bc.Equal(*board.ID, boardID)
}

func (bc *BCSuite) TestCreateBoard() {
	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/boards",
		httpmock.NewStringResponder(
			404,
			`{"code":400,"message":"The board name is invalid or duplicated."}`,
		),
	)
	_, err := bc.Pin.BoardResource.CreateBoard(CreateBoardOpts{Name: "To be delete"})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/boards",
		httpmock.NewStringResponder(
			201,
			`{"description":"","name":"ToBeDelete","owner":{"username":"merleliukun"},"id":"1022106146619729163","privacy":"PUBLIC"}`,
		),
	)

	board, _ := bc.Pin.BoardResource.CreateBoard(CreateBoardOpts{Name: "ToBeDelete"})
	bc.Equal(*board.Name, "ToBeDelete")
	bc.Equal(*board.ID, "1022106146619729163")
}

func (bc *BCSuite) TestUpdateBoard() {
	boardID := "1022106146619729163"
	httpmock.RegisterResponder(
		HttpPatch, Baseurl+"/boards/"+boardID,
		httpmock.NewStringResponder(
			404,
			`{"code":404,"message":"Board not found."}`,
		),
	)
	_, err := bc.Pin.BoardResource.UpdateBoard(boardID, UpdateBoardOpts{Description: "Board will be delete"})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpPatch, Baseurl+"/boards/"+boardID,
		httpmock.NewStringResponder(
			200,
			`{"description":"Board will be delete","name":"ToBeDelete","owner":{"username":"merleliukun"},"id":"1022106146619729163","privacy":"PUBLIC"}`,
		),
	)

	board, _ := bc.Pin.BoardResource.UpdateBoard(boardID, UpdateBoardOpts{Description: "Board will be delete"})
	bc.Equal(*board.Description, "Board will be delete")
	bc.Equal(*board.ID, boardID)
}
func (bc *BCSuite) TestDeleteBoard() {
	boardID := "1022106146619729163"
	httpmock.RegisterResponder(
		HttpDelete, Baseurl+"/boards/"+boardID,
		httpmock.NewStringResponder(
			404,
			`{"code":404,"message":"Board not found."}`,
		),
	)
	err := bc.Pin.BoardResource.DeleteBoard(boardID)
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpDelete, Baseurl+"/boards/"+boardID,
		httpmock.NewStringResponder(
			204,
			``,
		),
	)

	err = bc.Pin.BoardResource.DeleteBoard(boardID)
	bc.Nil(err)
}
