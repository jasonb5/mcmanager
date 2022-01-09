package curse

import (
	"net/http"
	"strconv"
)

type SearchConfig struct {
	CategoryID   int
	GameID       int
	GameVersion  string
	Index        int
	PageSize     int
	SearchFilter string
	SectionID    int
	Sort         int
}

var DefaultConfig = &SearchConfig{
	CategoryID: 0,
	GameID:     432,
	Index:      0,
	PageSize:   25,
	SectionID:  4471,
	Sort:       0,
}

func (c *SearchConfig) SetQueryParams(req *http.Request) {
	url := req.URL.Query()

	url.Add("categoryId", strconv.Itoa(c.CategoryID))
	url.Add("gameId", strconv.Itoa(c.GameID))
	url.Add("gameVersion", c.GameVersion)
	url.Add("index", strconv.Itoa(c.Index))
	url.Add("pageSize", strconv.Itoa(c.PageSize))
	url.Add("searchFilter", c.SearchFilter)
	url.Add("sectionId", strconv.Itoa(c.SectionID))
	url.Add("sort", strconv.Itoa(c.Sort))

	req.URL.RawQuery = url.Encode()
}
