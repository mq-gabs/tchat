package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"tchat.com/server/router/handlers"
)

func ProcessResponseData[T any](r *http.Response) (T, error) {
	var (
		zero     T
		response handlers.Response[T]
	)

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return zero, errors.Join(errCannotDecode, err)
	}

	switch r.StatusCode {
	case http.StatusOK:
		return response.Data, nil
	default:
		return zero, errInvalidResponse
	}
}

func AddQuery(req *http.Request, query map[string]string) {
	q := req.URL.Query()

	for key, val := range query {
		q.Set(key, val)
	}

	req.URL.RawQuery = q.Encode()
}

func NewGet(path string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, errors.Join(errCannotCreateRequest, err)
	}

	return req, nil
}

func NewPost(path string, body any) (*http.Request, error) {
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return nil, errors.Join(errCannotMarshalBody, err)
	}

	req, err := http.NewRequest(http.MethodPost, path, bytes.NewReader(bodyJSON))
	if err != nil {

		return nil, errors.Join(errCannotCreateRequest, err)
	}

	return req, nil
}
