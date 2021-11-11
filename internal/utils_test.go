package mcmanager

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestDownload(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		Equals(t, req.URL.String(), "http://example.com/test")

		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString("data")),
			Header:     make(http.Header),
		}
	})

	data, err := Download("http://example.com/test", client)

	Ok(t, err)

	Equals(t, []byte("data"), data)
}
