package testtransport

import (
	"errors"
	"net/http"
)

type RoundTripFunc func(*http.Request) *http.Response

func (r RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	if resp := r(req); resp != nil {
		return resp, nil
	}
	return nil, errors.New("testtransport: nil response")
}

func NewHTTPClient(fn RoundTripFunc) *http.Client {
	return &http.Client{Transport: fn}
}
