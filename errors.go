package letscloud

import "errors"

var (
	ErrMakingRequest     = errors.New("error creating new request")
	ErrSendingRequest    = errors.New("error sending request")
	ErrDecodingResponse  = errors.New("error decoding response body")
	ErrBadStatus         = errors.New("error getting bad status")
	ErrNoLocation        = errors.New("error no location found for this name")
	ErrInvalidToken      = errors.New("error invalid token provided")
	ErrInvalidHttpClient = errors.New("error invalid http client provided")
	ErrInvalidTimeout    = errors.New("error invalid timeout provided")
	ErrCreatingInstance  = errors.New("error creating new instance")
)
