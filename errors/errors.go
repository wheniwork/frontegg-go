package errors

import "errors"

var (
	ErrInvalidRegion        = errors.New("invalid region")
	ErrAuthenticationFailed = errors.New("authentication failed")
	ErrInvalidResponse      = errors.New("invalid response")

	ErrInvalidToken = errors.New("invalid token")
)
