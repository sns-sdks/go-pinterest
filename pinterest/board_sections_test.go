package pinterest

import (
	"github.com/jarcoal/httpmock"
)

func (bc *BCSuite) TestListBoardSections() {
	boardID := "1022106146619699845"
	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/boards/"+boardID+"/sections",
		httpmock.NewStringResponder(
			403,
			`{"code":403,"message":"Not authorized to access the board."}`,
		),
	)
	_, err := bc.Pin.BoardResource.ListBoardSections(boardID, ListOptions{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/boards/"+boardID+"/sections",
		httpmock.NewStringResponder(
			200,
			`{"items":[{"id":"5215150022519213435","name":"Night"},{"id":"5215148235938652082","name":"Big"}],"bookmark":null}`,
		),
	)

	bs, _ := bc.Pin.BoardResource.ListBoardSections(boardID, ListOptions{})
	bc.Equal(*bs.Items[0].Name, "Night")
	bc.Equal(*bs.Items[0].ID, "5215150022519213435")
	bc.Nil(bs.Bookmark)
}

func (bc *BCSuite) TestCreateBoardSection() {
	boardID := "1022106146619699845"
	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/boards/"+boardID+"/sections",
		httpmock.NewStringResponder(
			400,
			`{"code":400,"message":"Invalid board section parameters."}`,
		),
	)
	_, err := bc.Pin.BoardResource.CreateBoardSection(boardID, CreateBoardSectionOpts{Name: "Day"})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/boards/"+boardID+"/sections",
		httpmock.NewStringResponder(
			201,
			`{"id":"5215175925383086784","name":"Day"}`,
		),
	)

	bs, _ := bc.Pin.BoardResource.CreateBoardSection(boardID, CreateBoardSectionOpts{Name: "Day"})
	bc.Equal(*bs.Name, "Day")
}

func (bc *BCSuite) TestUpdateBoardSection() {
	boardID := "1022106146619699845"
	sectionID := "5215175925383086784"
	httpmock.RegisterResponder(
		HttpPatch, Baseurl+"/boards/"+boardID+"/sections/"+sectionID,
		httpmock.NewStringResponder(
			400,
			`{"code":400,"message":"Invalid board section parameters."}`,
		),
	)
	_, err := bc.Pin.BoardResource.UpdateBoardSection(boardID, sectionID, CreateBoardSectionOpts{Name: "DayModify"})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpPatch, Baseurl+"/boards/"+boardID+"/sections/"+sectionID,
		httpmock.NewStringResponder(
			201,
			`{"id":"5215175925383086784","name":"DayModify"}`,
		),
	)

	bs, _ := bc.Pin.BoardResource.UpdateBoardSection(boardID, sectionID, CreateBoardSectionOpts{Name: "DayModify"})
	bc.Equal(*bs.Name, "DayModify")
}

func (bc *BCSuite) TestDeleteBoardSection() {
	boardID := "1022106146619699845"
	sectionID := "5215175925383086784"
	httpmock.RegisterResponder(
		HttpDelete, Baseurl+"/boards/"+boardID+"/sections/"+sectionID,
		httpmock.NewStringResponder(
			403,
			`{"code":403,"message":"Not authorized to delete board section."}`,
		),
	)
	err := bc.Pin.BoardResource.DeleteBoardSection(boardID, sectionID)
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpDelete, Baseurl+"/boards/"+boardID+"/sections/"+sectionID,
		httpmock.NewStringResponder(
			204,
			``,
		),
	)

	err = bc.Pin.BoardResource.DeleteBoardSection(boardID, sectionID)
	bc.Nil(err)
}

func (bc *BCSuite) TestListPinsOnBoardSection() {
	boardID := "1022106146619729163"
	sectionID := "5215175925383086784"
	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/boards/"+boardID+"/sections/"+sectionID+"/pins",
		httpmock.NewStringResponder(
			404,
			`{"code":404,"message":"Board not found."}`,
		),
	)
	_, err := bc.Pin.BoardResource.ListPinsOnBoardSection(boardID, sectionID, ListOptions{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/boards/"+boardID+"/sections/"+sectionID+"/pins",
		httpmock.NewStringResponder(
			200,
			`{"items":[{"link":null,"media":{"media_type":"image","images":{"150x150":{"width":150,"height":150,"url":"https://i.pinimg.com/150x150/3e/36/63/3e3663eba0b3e150204fdfc2df7809a0.jpg"},"400x300":{"width":400,"height":300,"url":"https://i.pinimg.com/400x300/3e/36/63/3e3663eba0b3e150204fdfc2df7809a0.jpg"},"600x":{"width":600,"height":800,"url":"https://i.pinimg.com/600x/3e/36/63/3e3663eba0b3e150204fdfc2df7809a0.jpg"},"1200x":{"width":1200,"height":1600,"url":"https://i.pinimg.com/1200x/3e/36/63/3e3663eba0b3e150204fdfc2df7809a0.jpg"},"originals":{"width":3456,"height":4608,"url":"https://i.pinimg.com/originals/3e/36/63/3e3663eba0b3e150204fdfc2df7809a0.jpg"}}},"title":"","board_id":"1022106146619699845","board_owner":{"username":"merleliukun"},"board_section_id":"5215148235938652082","id":"1022106077902203823","created_at":"2021-12-29T02:24:55","description":"City","alt_text":null}],"bookmark":null}`,
		),
	)

	boards, _ := bc.Pin.BoardResource.ListPinsOnBoardSection(boardID, sectionID, ListOptions{})
	bc.Equal(*boards.Items[0].ID, "1022106077902203823")
	bc.Nil(boards.Bookmark)
}
