package pinterest

/*
	Media API
*/

type MediaResource Resource

func newMediaResource(cli *Client) *MediaResource {
	return &MediaResource{Cli: cli}
}

// MediaUpload represents the media upload info.
type MediaUpload struct {
	MediaID   *string `json:"media_id"`
	MediaType *string `json:"media_type"`
	Status    *string `json:"status"`
}

func (m MediaUpload) String() string {
	return Stringify(m)
}

// MediaUploadsResponse represents the response for list media uploads.
type MediaUploadsResponse struct {
	Items    []*MediaUpload `json:"items"`
	Bookmark *string        `json:"bookmark"`
}

func (m MediaUploadsResponse) String() string {
	return Stringify(m)
}

// ListMediaUploads List media uploads filtered by given parameters.
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/media/list
func (r *MediaResource) ListMediaUploads(args ListOptions) (*MediaUploadsResponse, *APIError) {
	path := "/media"

	resp := new(MediaUploadsResponse)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetMediaUploadDetail Get details for a registered media upload, including its current status.
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/media/get
func (r *MediaResource) GetMediaUploadDetail(mediaID string) (*MediaUpload, *APIError) {
	path := "/media/" + mediaID

	resp := new(MediaUpload)
	err := r.Cli.DoGet(path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// RegisterMediaUploadOpts represents the parameters for register media upload.
type RegisterMediaUploadOpts struct {
	MediaType string `json:"media_type"`
}

type UploadParameters struct {
	XAmzDate          string `json:"x-amz-date"`
	XAmzSignature     string `json:"x-amz-signature"`
	XAmzSecurityToken string `json:"x-amz-security-token"`
	XAmzAlgorithm     string `json:"x-amz-algorithm"`
	Key               string `json:"key"`
	Policy            string `json:"policy"`
	XAmzCredential    string `json:"x-amz-credential"`
	ContentType       string `json:"Content-Type"`
}

type RegisterMediaUploadResponse struct {
	MediaID          *string            `json:"media_id"`
	MediaType        *string            `json:"media_type"`
	UploadURL        *string            `json:"upload_url"`
	UploadParameters map[string]*string `json:"upload_parameters"`
}

// RegisterMediaUpload Register your intent to upload media.
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/media/create
func (r *MediaResource) RegisterMediaUpload(args RegisterMediaUploadOpts) {

}
