package download_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"mcmanager/internal/download"
	"mcmanager/internal/testutils"
	"net/http"
	"testing"
)

func TestDownload(t *testing.T) {
	client := testutils.NewTestClient(func(req *http.Request) *http.Response {
		testutils.Equals(t, req.URL.String(), "http://example.com/test")

		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString("data")),
			Header:     make(http.Header),
		}
	})

	data, err := download.Download("http://example.com/test", client)

	testutils.Ok(t, err)

	testutils.Equals(t, []byte("data"), data)
}

func TestDownloadJSON(t *testing.T) {
	client := testutils.NewTestClient(func(req *http.Request) *http.Response {
		testutils.Equals(t, req.URL.String(), "http://example.com/test")

		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`{"data": "test"}`)),
			Header:     make(http.Header),
		}
	})

	type TestStruct struct {
		Data string `json:"data"`
	}

	data, err := download.DownloadJSON("http://example.com/test", func(b []byte) (interface{}, error) {
		var data TestStruct

		json.Unmarshal(b, &data)

		return data, nil
	}, client)

	testutils.Ok(t, err)

	testutils.Equals(t, TestStruct{Data: "test"}, data)
}
