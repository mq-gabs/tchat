package api

import "errors"

var (
	errCannotDecode      = errors.New("cannot decode")
	errInvalidResponse   = errors.New("invalid response")
	errCannotMarshalBody = errors.New("cannot marshal body")
)
