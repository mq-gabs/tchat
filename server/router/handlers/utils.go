package handlers

import (
	"encoding/json"
	"net/http"
)

func ReadBody(r *http.Request, model any) error {
	return json.NewDecoder(r.Body).Decode(model)
}
