package pinterest

import (
	"github.com/jarcoal/httpmock"
)

func (bc *BCSuite) TestListMediaUploads() {
	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/media",
		httpmock.NewStringResponder(
			401,
			`{"code":2,"message":"Authentication failed.","status":"failure"}`,
		),
	)
	_, err := bc.Pin.Media.ListMediaUploads(ListOptions{})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/media",
		httpmock.NewStringResponder(
			200,
			`{"items":[{"media_id":"5216397987686698987","media_type":"video","status":"registered"}],"bookmark":null}`,
		),
	)

	mediaUpload, _ := bc.Pin.Media.ListMediaUploads(ListOptions{})
	bc.Equal(*mediaUpload.Items[0].MediaID, "5216397987686698987")
	bc.Nil(mediaUpload.Bookmark)
}

func (bc *BCSuite) TestGetMediaUploadDetail() {
	muID := "5216391720465385860"

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/media/"+muID,
		httpmock.NewStringResponder(
			404,
			`{"code":404,"message":"Media upload not found"}`,
		),
	)
	_, err := bc.Pin.Media.GetMediaUploadDetail(muID)
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpGet, Baseurl+"/media/"+muID,
		httpmock.NewStringResponder(
			200,
			`{"media_id":"5216391720465385860","media_type":"video","status":"registered"}`,
		),
	)

	mediaUpload, _ := bc.Pin.Media.GetMediaUploadDetail(muID)
	bc.Equal(*mediaUpload.MediaID, "5216391720465385860")
	bc.Equal(*mediaUpload.MediaType, "video")
}

func (bc *BCSuite) TestRegisterMediaUpload() {
	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/media",
		httpmock.NewStringResponder(
			401,
			`{"code":2,"message":"Authentication failed.","status":"failure"}`,
		),
	)
	_, err := bc.Pin.Media.RegisterMediaUpload(RegisterMediaUploadOpts{MediaType: "video"})
	bc.IsType(&APIError{}, err)

	httpmock.RegisterResponder(
		HttpPost, Baseurl+"/media",
		httpmock.NewStringResponder(
			200,
			`{"media_id":"5216393791692388749","media_type":"video","upload_url":"https://pinterest-media-upload.s3-accelerate.amazonaws.com/","upload_parameters":{"x-amz-date":"20220220T082536Z","x-amz-signature":"signature","x-amz-security-token":"token","x-amz-algorithm":"AWS4-HMAC-SHA256","key":"key","policy":"policy","x-amz-credential":"credential","Content-Type":"multipart/form-data"}}`,
		),
	)
	mediaUploadResponse, _ := bc.Pin.Media.RegisterMediaUpload(RegisterMediaUploadOpts{MediaType: "video"})
	bc.Equal(*mediaUploadResponse.MediaID, "5216393791692388749")
	bc.Equal(mediaUploadResponse.UploadParameters["x-amz-date"], "20220220T082536Z")
}
