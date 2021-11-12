package curse

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"mcmanager/internal/download"
	"mcmanager/internal/mcmanager"
	"mcmanager/internal/runner"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

const BASE_URL = "https://addons-ecs.forgesvc.net/api/v2/addon"
const SEARCH_URL = "%s/search"
const MODPACK_URL = "%s/%d"
const VERSIONS_URL = "%s/%d/files"
const DOWNLOAD_URL = "%s/%d/file/%d/download-url"

type SearchConfig struct {
	CategoryID, GameID, Index, PageSize, SectionID, Sort int
	GameVersion, SearchFilter                            string
}

var DefaultSearchConfig = &SearchConfig{
	CategoryID:   0,
	GameID:       432,
	GameVersion:  "",
	Index:        0,
	PageSize:     25,
	SearchFilter: "",
	SectionID:    4471,
	Sort:         0,
}

func (sc *SearchConfig) UpdateParams(u *url.URL) {
	query := u.Query()
	val := reflect.ValueOf(*sc)

	for i := 0; i < val.NumField(); i++ {
		query.Add(strings.ToLower(val.Type().Field(i).Name), fmt.Sprintf("%v", val.Field(i)))

		log.Printf("%v = %v", val.Type().Field(i).Name, val.Field(i))
	}

	u.RawQuery = query.Encode()
}

func Search(c *SearchConfig, client *http.Client) ([]ModPack, error) {
	searchURL, err := url.Parse(fmt.Sprintf(SEARCH_URL, BASE_URL))

	if err != nil {
		return nil, fmt.Errorf("error parsing search url")
	}

	c.UpdateParams(searchURL)

	data, err := download.DownloadJSON(searchURL.String(), func(data []byte) (interface{}, error) {
		var modpacks []ModPack

		if err := json.Unmarshal(data, &modpacks); err != nil {
			return nil, fmt.Errorf("error converting json: %v", err)
		}

		return modpacks, nil
	}, client)

	if err != nil {
		return nil, fmt.Errorf("error downloading search results: %v", err)
	}

	return data.([]ModPack), nil
}

func InstallServer(modpack *ModPack, version *ModPackVersion, config *mcmanager.Config, client *http.Client) error {
	if err := download.DownloadExtract(version.DownloadURL, config.InstallPath, client); err != nil {
		return fmt.Errorf("error downloading and extracting file %v: %v", version.DownloadURL, err)
	}

	if version.ServerPackFileID == 0 {
		return fmt.Errorf("error %s does not have a server pack", version.DisplayName)
	}

	serverPackDownloadURL := fmt.Sprintf(DOWNLOAD_URL, BASE_URL, modpack.ID, version.ServerPackFileID)

	serverPackURL, err := download.Download(serverPackDownloadURL, client)

	if err != nil {
		return fmt.Errorf("error getting download url for %v: %v", serverPackDownloadURL, err)
	}

	if err := download.DownloadExtract(string(serverPackURL), config.InstallPath, client); err != nil {
		return fmt.Errorf("error downloading and extracting server files %v: %v", serverPackURL, err)
	}

	serverSetupConfig := filepath.Join(config.InstallPath, "server-setup-config.yaml")

	if _, err := os.Stat(serverSetupConfig); err == nil {
		startServer := filepath.Join(config.InstallPath, "startserver.sh")

		if err := runner.Run(startServer); err != nil {
			return fmt.Errorf("error running server starter: %v", err)
		}
	} else {
		return errors.New("installing non-serverstart has not implemented yet")
	}

	return nil
}

func InstallServerByID(modpackID, versionID int, config *mcmanager.Config, client *http.Client) error {
	modpack, err := GetModPack(modpackID, client)

	if err != nil {
		return fmt.Errorf("error getting modpack: %v", err)
	}

	version, err := modpack.GetVersion(versionID, client)

	if err != nil {
		return fmt.Errorf("error getting version %d: %v", versionID, err)
	}

	return InstallServer(modpack, version, config, client)
}
