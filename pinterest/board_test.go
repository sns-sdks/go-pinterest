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
