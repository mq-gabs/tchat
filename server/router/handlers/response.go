package handlers

import (
	"encoding/json"
	"net/http"
)

type Response[T any] struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

func (r *Response[any]) WriteSelf(w http.ResponseWriter) {
	res, err := json.Marshal(r)
	if err != nil {
		WriteInternalServerError(w, err)
		return
	}

	w.WriteHeader(r.Code)
	w.Write(res)
}

func WriteResponse(w http.ResponseWriter, r *Response[any]) {
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
	r := Response[any]{
		Code:    500,
		Message: err.Error(),
	}

	WriteResponse(w, &r)
}

func WriteOKEmpty(w http.ResponseWriter, message string) {
	r := Response[any]{
		Code:    200,
		Message: message,
	}

	WriteResponse(w, &r)
}

func WriteNotFound(w http.ResponseWriter, msg string, err error) {
	r := Response[any]{
		Code:    404,
		Message: err.Error(),
	}

	WriteResponse(w, &r)
}

func WriteOKWithBody(w http.ResponseWriter, body any) {
	r := Response[any]{
		Code:    200,
		Message: "OK",
		Data:    body,
	}

	WriteResponse(w, &r)
}

func WriteBadRequest(w http.ResponseWriter, err error) {
	r := Response[any]{
		Code:    400,
		Message: err.Error(),
	}

	WriteResponse(w, &r)
}
