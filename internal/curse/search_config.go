package curse

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

type SearchConfig struct {
	CategoryID, GameID, Index, PageSize, SectionID, Sort int
	GameVersion, SearchFilter                            string
}

var DefaultConfig = SearchConfig{
	CategoryID:   0,
	GameID:       432,
	GameVersion:  "",
	Index:        0,
	PageSize:     25,
	SearchFilter: "",
	SectionID:    4471,
	Sort:         0,
}

func (sc *SearchConfig) ToURLParams(url *url.URL) *url.URL {
	q := url.Query()

	v := reflect.ValueOf(*sc)

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)

		q.Add(strings.ToLower(v.Type().Field(i).Name), fmt.Sprintf("%v", f))
	}

	url.RawQuery = q.Encode()

	return url
}
