package curse

import (
	"encoding/json"
	"errors"
	"fmt"
	"mcmanager/internal/utils"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

const baseURL = "https://addons-ecs.forgesvc.net/api/v2/addon"

func queryUpdater(c *SearchConfig) func(*http.Request) error {
	return func(req *http.Request) error {
		c.SetQueryParams(req)

		return nil
	}
}

func Search(c *SearchConfig) (ModPacks, error) {
	searchURL := fmt.Sprintf("%s/search", baseURL)

	data, err := utils.Download(searchURL, queryUpdater(c))

	if err != nil {
		return nil, nil
	}

	var modpacks ModPacks

	if err := json.Unmarshal(data, &modpacks); err != nil {
		return nil, nil
	}

	return modpacks, nil
}

func getModDownloadURL(modPackID, modID int) (string, error) {
	downloadURL := fmt.Sprintf("%s/%d/file/%d/download-url", baseURL, modPackID, modID)

	data, err := utils.Download(downloadURL, nil)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

func DownloadMod(modPackID, modID int, outputPath string) (string, error) {
	downloadURL, err := getModDownloadURL(modPackID, modID)

	if err != nil {
		return "", err
	}

	filename := filepath.Base(downloadURL)

	filePath := path.Join(outputPath, filename)

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		_, err := utils.DownloadFile(downloadURL, filePath)

		if err != nil {
			return "", err
		}

		return filePath, nil
	}

	return filePath, nil
}
