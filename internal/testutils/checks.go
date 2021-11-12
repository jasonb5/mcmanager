package testutils

import (
	"net/http"
	"reflect"
	"testing"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func Ok(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}
}

func Equals(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %v got %v", expected, actual)
	}
}
