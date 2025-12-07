package handlers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (r *Response) WriteSelf(w http.ResponseWriter) {
	res, err := json.Marshal(r)
	if err != nil {
		WriteInternalServerError(w, err)
		return
	}

	w.WriteHeader(r.Code)
	w.Write(res)
}

func WriteResponse(w http.ResponseWriter, r *Response) {
	res, err := json.Marshal(r)
	if err != nil {
		WriteDefaultError(w)
		return
	}

	w.WriteHeader(r.Code)
	w.Write(res)
}

func WriteDefaultError(w http.ResponseWriter) {
	w.WriteHeader(500)
}

func WriteInternalServerError(w http.ResponseWriter, err error) {
	r := Response{
		Code:    500,
		Message: err.Error(),
	}

	WriteResponse(w, &r)
}

func WriteOKEmpty(w http.ResponseWriter, message string) {
	r := Response{
		Code:    200,
		Message: message,
	}

	WriteResponse(w, &r)
}

func WriteNotFound(w http.ResponseWriter, msg string, err error) {
	r := Response{
		Code:    404,
		Message: err.Error(),
	}

	WriteResponse(w, &r)
}

func WriteOKWithBody(w http.ResponseWriter, body any) {
	r := Response{
		Code:    200,
		Message: "OK",
		Data:    body,
	}

	WriteResponse(w, &r)
}

func WriteBadRequest(w http.ResponseWriter, err error) {
	r := Response{
		Code:    400,
		Message: err.Error(),
	}

	WriteResponse(w, &r)
}
