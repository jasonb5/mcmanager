package curse

import (
	"bytes"
	"io/ioutil"
	mcmanager "mcmanager/internal"
	"net/http"
	"testing"
)

func TestSearch(t *testing.T) {
	client := mcmanager.NewTestClient(func(req *http.Request) *http.Response {
		mcmanager.Equals(t, req.URL.String(), "http://example.com/")

		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString("data")),
			Header:     make(http.Header),
		}
	})

	c := &DefaultConfig

	_, err := search(c, client)

	mcmanager.Ok(t, err)
}
