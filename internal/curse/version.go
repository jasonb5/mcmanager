package curse

import (
	"encoding/json"
	"fmt"
	"mcmanager/internal/utils"
	"time"
)

type Versions []Version

type Version struct {
	ID              int           `json:"id"`
	DisplayName     string        `json:"displayName"`
	FileName        string        `json:"fileName"`
	FileDate        time.Time     `json:"fileDate"`
	FileLength      int           `json:"fileLength"`
	ReleaseType     int           `json:"releaseType"`
	FileStatus      int           `json:"fileStatus"`
	DownloadURL     string        `json:"downloadUrl"`
	IsAlternate     bool          `json:"isAlternate"`
	AlternateFileID int           `json:"alternateFileId"`
	Dependencies    []interface{} `json:"dependencies"`
	IsAvailable     bool          `json:"isAvailable"`
	Modules         []struct {
		Foldername  string `json:"foldername"`
		Fingerprint int    `json:"fingerprint"`
	} `json:"modules"`
	PackageFingerprint      int         `json:"packageFingerprint"`
	GameVersion             []string    `json:"gameVersion"`
	InstallMetadata         interface{} `json:"installMetadata"`
	ServerPackFileID        int         `json:"serverPackFileId"`
	HasInstallScript        bool        `json:"hasInstallScript"`
	GameVersionDateReleased time.Time   `json:"gameVersionDateReleased"`
	GameVersionFlavor       interface{} `json:"gameVersionFlavor"`
}

func GetVersionByID(modPackID, versionID int) (*Version, error) {
	var version *Version

	versions, err := GetVersions(modPackID)

	if err != nil {
		return nil, err
	}

	for _, item := range versions {
		if item.ServerPackFileID != 0 && item.ServerPackFileID == versionID {
			version = &item

			break
		}
	}

	if version == nil {
		return nil, fmt.Errorf("could not find version matching %d", versionID)
	}

	return version, nil
}

func GetVersions(modPackID int) (Versions, error) {
	versionURL := fmt.Sprintf("%s/%d/files", baseURL, modPackID)

	data, err := utils.Download(versionURL, nil)

	if err != nil {
		return nil, err
	}

	var versions Versions

	if err := json.Unmarshal(data, &versions); err != nil {
		return nil, err
	}

	return versions, nil
}
