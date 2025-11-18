package testtransport

import (
	"errors"
	"net/http"
)

// RoundTripFunc implements http.RoundTripper as a function for tests.
type RoundTripFunc func(*http.Request) *http.Response

// RoundTrip satisfies the http.RoundTripper interface using the wrapped function.
func (r RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	if resp := r(req); resp != nil {
		return resp, nil
	}
	return nil, errors.New("testtransport: nil response")
}

// NewHTTPClient returns an http.Client that uses the provided RoundTripFunc.
func NewHTTPClient(fn RoundTripFunc) *http.Client {
	return &http.Client{Transport: fn}
}
