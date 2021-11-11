package curse

import (
	"net/url"
	"testing"
)

func TestToURLParams(t *testing.T) {
	sc := DefaultConfig
	url := &url.URL{}
	expected := "categoryid=0&gameid=432&gameversion=&index=0&pagesize=25&searchfilter=&sectionid=4471&sort=0"

	sc.ToURLParams(url)

	if url.RawQuery != expected {
		t.Fatalf("Got %v Expected %v", url.RawQuery, expected)
	}
}
