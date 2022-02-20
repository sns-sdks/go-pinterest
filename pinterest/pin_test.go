package pinterest

import "github.com/jarcoal/httpmock"

func (bc *BCSuite) TestCreatePin() {
	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/pins",
		httpmock.NewStringResponder(
			400,
			`{"code":400,"message":"Pin is not a call-to-create response"}`,
		),
	)
	_, err := bc.Pin.Pin.CreatePin(CreatePinOpts{BoardID: "123456", MediaSource: CreatePinMediaSourceOpts{SourceType: "image_url", Url: "https://xxx.com/image.png"}})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/pins",
		httpmock.NewStringResponder(
			200,
			`{"title":"","board_id":"1022106146619699845","media":{"media_type":"image","images":{"150x150":{"width":150,"height":150,"url":"https://i.pinimg.com/150x150/39/90/d9/3990d935052091b45865fb001609b97e.jpg"},"400x300":{"width":400,"height":300,"url":"https://i.pinimg.com/400x300/39/90/d9/3990d935052091b45865fb001609b97e.jpg"},"600x":{"width":600,"height":893,"url":"https://i.pinimg.com/600x/39/90/d9/3990d935052091b45865fb001609b97e.jpg"},"1200x":{"width":1200,"height":1786,"url":"https://i.pinimg.com/1200x/39/90/d9/3990d935052091b45865fb001609b97e.jpg"},"originals":{"width":1920,"height":2858,"url":"https://i.pinimg.com/originals/39/90/d9/3990d935052091b45865fb001609b97e.jpg"}}},"board_section_id":null,"id":"1022106077902810180","board_owner":{"username":"merleliukun"},"description":" ","alt_text":null,"link":null,"created_at":"2022-02-14T02:54:38"}`,
		),
	)

	pin, _ := bc.Pin.Pin.CreatePin(CreatePinOpts{BoardID: "123456", MediaSource: CreatePinMediaSourceOpts{SourceType: "image_url", Url: "https://xxx.com/image.png"}})
	bc.Equal(*pin.BoardID, "1022106146619699845")
	bc.Equal(*pin.Media.MediaType, "image")
	bc.Equal(*pin.Media.Images["150x150"].Width, 150)
}

func (bc *BCSuite) TestGetPin() {
	pinID := "1022106077902810180"
	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/pins/"+pinID,
		httpmock.NewStringResponder(
			404,
			`{"code":404,"message":"Pin not found."}`,
		),
	)
	_, err := bc.Pin.Pin.GetPin(pinID, "")
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/pins/"+pinID,
		httpmock.NewStringResponder(
			200,
			`{"title":"","board_id":"1022106146619699845","media":{"media_type":"image","images":{"150x150":{"width":150,"height":150,"url":"https://i.pinimg.com/150x150/39/90/d9/3990d935052091b45865fb001609b97e.jpg"},"400x300":{"width":400,"height":300,"url":"https://i.pinimg.com/400x300/39/90/d9/3990d935052091b45865fb001609b97e.jpg"},"600x":{"width":600,"height":893,"url":"https://i.pinimg.com/600x/39/90/d9/3990d935052091b45865fb001609b97e.jpg"},"1200x":{"width":1200,"height":1786,"url":"https://i.pinimg.com/1200x/39/90/d9/3990d935052091b45865fb001609b97e.jpg"},"originals":{"width":1920,"height":2858,"url":"https://i.pinimg.com/originals/39/90/d9/3990d935052091b45865fb001609b97e.jpg"}}},"board_section_id":null,"id":"1022106077902810180","board_owner":{"username":"merleliukun"},"description":" ","alt_text":null,"link":null,"created_at":"2022-02-14T02:54:38"}`,
		),
	)

	pin, _ := bc.Pin.Pin.GetPin(pinID, "")
	bc.Equal(*pin.BoardID, "1022106146619699845")
	bc.Equal(*pin.Media.MediaType, "image")

	pin, _ = bc.Pin.Pin.GetPin(pinID, "123456")
	bc.Equal(*pin.BoardID, "1022106146619699845")

}

func (bc *BCSuite) TestDeletePin() {
	pinID := "1022106146619729163"
	httpmock.RegisterResponder(
		HttpDelete, Baseurl+"/pins/"+pinID,
		httpmock.NewStringResponder(
			404,
			`{"code":404,"message":"Board not found."}`,
		),
	)
	err := bc.Pin.Pin.DeletePin(pinID)
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpDelete, Baseurl+"/pins/"+pinID,
		httpmock.NewStringResponder(
			204,
			``,
		),
	)

	err = bc.Pin.Pin.DeletePin(pinID)
	bc.Nil(err)
}
