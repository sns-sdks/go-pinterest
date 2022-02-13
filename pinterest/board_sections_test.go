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
	_, err := bc.Pin.BoardResource.ListBoardSections(boardID, ListBoardSectionOpts{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/boards/"+boardID+"/sections",
		httpmock.NewStringResponder(
			200,
			`{"items":[{"id":"5215150022519213435","name":"Night"},{"id":"5215148235938652082","name":"Big"}],"bookmark":null}`,
		),
	)

	bs, _ := bc.Pin.BoardResource.ListBoardSections(boardID, ListBoardSectionOpts{})
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
