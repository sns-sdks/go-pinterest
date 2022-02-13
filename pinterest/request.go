package pinterest

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	goquery "github.com/google/go-querystring/query"
	"strings"
)

// APIError represents the error response
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

/*
	functions for http requests
*/
func ParseDataResponse(response *resty.Response, d interface{}) *APIError {
	var err error
	if response.IsSuccess() {
		switch d := d.(type) {
		case nil:
		default:
			err = json.Unmarshal(response.Body(), d)
		}
		if err != nil {
			return &APIError{Code: -1, Message: err.Error()}
		}
		return nil
	}
	apiErr := new(APIError)
	err = json.Unmarshal(response.Body(), &apiErr)
	if err != nil {
		return &APIError{Code: -1, Message: err.Error()}
	}
	return apiErr
}

func (r *Client) Do(method, path string, queryParams interface{}, jsonParams interface{}, d interface{}) *APIError {
	req := r.Cli.R()

	// parse struct params
	if queryParams != nil {
		v, err := goquery.Values(queryParams)
		if err != nil {
			apiError := APIError{Code: -1, Message: err.Error()}
			return &apiError
		}
		req.SetQueryParamsFromValues(v)
	}
	if jsonParams != nil {
		req.SetBody(jsonParams)
		req.SetHeader("Content-Type", "application/json")
	}

	// If the only path, add the domain.
	var url string
	if strings.HasPrefix(path, "http") {
		url = path
	} else {
		url = Baseurl + path
	}

	resp, err := req.Execute(method, url)
	if err != nil {
		apiError := APIError{Code: -1, Message: err.Error()}
		return &apiError
	}
	apiError := ParseDataResponse(resp, d)
	return apiError
}

func (r *Client) DoGet(path string, queryParams interface{}, d interface{}) *APIError {
	return r.Do(HttpGet, path, queryParams, nil, d)
}

func (r *Client) DoPost(path string, jsonParams interface{}, d interface{}) *APIError {
	return r.Do(HttpPost, path, nil, jsonParams, d)
}

func (r *Client) DoPatch(path string, jsonParams interface{}, d interface{}) *APIError {
	return r.Do(HttpPatch, path, nil, jsonParams, d)
}

func (r *Client) DoDelete(path string, d interface{}) *APIError {
	return r.Do(HttpDelete, path, nil, nil, d)
}
