package errors

import "net/http"

type FronteggError struct {
	BaseError error

	message string

	Code int

	HttpResponse *http.Response
}

func NewFronteggError(resp *http.Response, message string, baseError error) *FronteggError {
	return &FronteggError{
		HttpResponse: resp,
		Code:         resp.StatusCode,
		message:      message,
		BaseError:    baseError,
	}
}

func (e *FronteggError) Error() string {
	return e.message
}
