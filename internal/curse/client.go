package curse

import (
	"encoding/json"
	"fmt"
	mcmanager "mcmanager/internal"
	"net/http"
	"net/url"
)

const BASE_URL = "https://addons-ecs.forgesvc.net/api/v2/addon"
const SEARCH_URL = "%s/search"
const FILES_URL = "%s/%d/files"
const DOWNLOAD_URL = "%s/%d/file/%d/download-url"

func Search(c *SearchConfig) ([]ModPack, error) {
	return search(c, nil)
}

func search(c *SearchConfig, client *http.Client) ([]ModPack, error) {
	u, err := url.Parse(fmt.Sprintf(SEARCH_URL, BASE_URL))

	if err != nil {
		return nil, fmt.Errorf("error parsing url: %v", err)
	}

	c.ToURLParams(u)

	data, err := mcmanager.Download(u.String(), nil)

	if err != nil {
		return nil, fmt.Errorf("error searching for modpacks: %v", err)
	}

	var modpacks []ModPack

	if err := json.Unmarshal(data, &modpacks); err != nil {
		return nil, fmt.Errorf("error decoding json response: %v", err)
	}

	return modpacks, nil
}
