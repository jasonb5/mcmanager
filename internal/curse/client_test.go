package curse_test

import (
	"bytes"
	"encoding/json"
	"io"
	"mcmanager/internal/curse"
	"mcmanager/internal/testutils"
	"net/http"
	"net/url"
	"testing"
)

func TestUpdateParams(t *testing.T) {
	config := curse.DefaultSearchConfig

	testURL := &url.URL{}

	config.UpdateParams(testURL)

	exp := "?categoryid=0&gameid=432&gameversion=&index=0&pagesize=25&searchfilter=&sectionid=4471&sort=0"

	testutils.Equals(t, exp, testURL.String())
}

func TestSearch(t *testing.T) {
	exp := "https://addons-ecs.forgesvc.net/api/v2/addon/search?categoryid=0&gameid=432&gameversion=&index=0&pagesize=25&searchfilter=&sectionid=4471&sort=0"

	client := testutils.NewTestClient(func(req *http.Request) *http.Response {
		testutils.Equals(t, exp, req.URL.String())

		pack := []curse.ModPack{
			{
				ID:   0,
				Name: "test",
				Slug: "test",
			},
		}

		data, _ := json.Marshal(pack)

		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBuffer(data)),
			Header:     make(http.Header),
		}
	})

	config := curse.DefaultSearchConfig

	data, err := curse.Search(config, client)

	testutils.Ok(t, err)

	expData := []curse.ModPack{
		{
			ID:   0,
			Name: "test",
			Slug: "test",
		},
	}

	testutils.Equals(t, expData, data)
}
