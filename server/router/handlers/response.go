package handlers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Details any    `json:"details,omitempty"`
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
		Success: false,
		Code:    500,
		Message: "Internal Server Error",
		Details: err,
	}

	WriteResponse(w, &r)
}

func WriteOKEmpty(w http.ResponseWriter, details any) {
	r := Response{
		Success: true,
		Code:    200,
		Message: "OK",
		Details: details,
	}

	WriteResponse(w, &r)
}

func WriteNotFound(w http.ResponseWriter, msg string, err error) {
	r := Response{
		Success: false,
		Code:    404,
		Message: msg,
		Details: err,
	}

	WriteResponse(w, &r)
}

func WriteOKWithBody(w http.ResponseWriter, body any) {
	r := Response{
		Success: true,
		Code:    200,
		Message: "OK",
		Data:    body,
	}

	WriteResponse(w, &r)
}

func WriteBadRequest(w http.ResponseWriter, err error) {
	r := Response{
		Success: false,
		Code:    400,
		Message: "Bad Request",
		Details: err,
	}

	WriteResponse(w, &r)
}
