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

func (bc *BCSuite) TestListPinsOnBoard() {
	boardID := "1022106146619729163"
	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/boards/"+boardID+"/pins",
		httpmock.NewStringResponder(
			404,
			`{"code":404,"message":"Board not found."}`,
		),
	)
	_, err := bc.Pin.BoardResource.ListPinsOnBoard(boardID, ListOptions{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/boards/"+boardID+"/pins",
		httpmock.NewStringResponder(
			200,
			`{"items":[{"title":"","board_id":"1022106146619699845","id":"1022106077902810180","media":{"media_type":"image","images":{"150x150":{"width":150,"height":150,"url":"https://i.pinimg.com/150x150/39/90/d9/3990d935052091b45865fb001609b97e.jpg"},"400x300":{"width":400,"height":300,"url":"https://i.pinimg.com/400x300/39/90/d9/3990d935052091b45865fb001609b97e.jpg"},"600x":{"width":600,"height":893,"url":"https://i.pinimg.com/600x/39/90/d9/3990d935052091b45865fb001609b97e.jpg"},"1200x":{"width":1200,"height":1786,"url":"https://i.pinimg.com/1200x/39/90/d9/3990d935052091b45865fb001609b97e.jpg"},"originals":{"width":1920,"height":2858,"url":"https://i.pinimg.com/originals/39/90/d9/3990d935052091b45865fb001609b97e.jpg"}}},"board_section_id":null,"description":" ","board_owner":{"username":"merleliukun"},"alt_text":null,"link":null,"created_at":"2022-02-14T02:54:38"}],"bookmark":null}`,
		),
	)

	boards, _ := bc.Pin.BoardResource.ListPinsOnBoard(boardID, ListOptions{})
	bc.Equal(*boards.Items[0].ID, "1022106077902810180")
	bc.Nil(boards.Bookmark)
}
