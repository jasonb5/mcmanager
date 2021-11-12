package curse

import (
	"encoding/json"
	"fmt"
	"mcmanager/internal/download"
	"net/http"
	"time"
)

type ModPack struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Authors []struct {
		Name              string      `json:"name"`
		URL               string      `json:"url"`
		ProjectID         int         `json:"projectId"`
		ID                int         `json:"id"`
		ProjectTitleID    interface{} `json:"projectTitleId"`
		ProjectTitleTitle interface{} `json:"projectTitleTitle"`
		UserID            int         `json:"userId"`
		TwitchID          int         `json:"twitchId"`
	} `json:"authors,omitempty"`
	Attachments []struct {
		ID           int    `json:"id"`
		ProjectID    int    `json:"projectId"`
		Description  string `json:"description"`
		IsDefault    bool   `json:"isDefault"`
		ThumbnailURL string `json:"thumbnailUrl"`
		Title        string `json:"title"`
		URL          string `json:"url"`
		Status       int    `json:"status"`
	} `json:"attachments,omitempty"`
	IssueTrackerURL string  `json:"issueTrackerUrl,omitempty"`
	WikiURL         string  `json:"wikiUrl,omitempty"`
	WebsiteURL      string  `json:"websiteUrl,omitempty"`
	GameID          int     `json:"gameId,omitempty"`
	Summary         string  `json:"summary,omitempty"`
	DefaultFileID   int     `json:"defaultFileId,omitempty"`
	DownloadCount   float64 `json:"downloadCount,omitempty"`
	LatestFiles     []struct {
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
			Type        int    `json:"type"`
		} `json:"modules"`
		PackageFingerprint  int64    `json:"packageFingerprint"`
		GameVersion         []string `json:"gameVersion"`
		SortableGameVersion []struct {
			GameVersionPadded      string    `json:"gameVersionPadded"`
			GameVersion            string    `json:"gameVersion"`
			GameVersionReleaseDate time.Time `json:"gameVersionReleaseDate"`
			GameVersionName        string    `json:"gameVersionName"`
			GameVersionTypeID      int       `json:"gameVersionTypeId"`
		} `json:"sortableGameVersion"`
		InstallMetadata            interface{} `json:"installMetadata"`
		Changelog                  interface{} `json:"changelog"`
		HasInstallScript           bool        `json:"hasInstallScript"`
		IsCompatibleWithClient     bool        `json:"isCompatibleWithClient"`
		CategorySectionPackageType int         `json:"categorySectionPackageType"`
		RestrictProjectFileAccess  int         `json:"restrictProjectFileAccess"`
		ProjectStatus              int         `json:"projectStatus"`
		RenderCacheID              int         `json:"renderCacheId"`
		FileLegacyMappingID        interface{} `json:"fileLegacyMappingId"`
		ProjectID                  int         `json:"projectId"`
		ParentProjectFileID        interface{} `json:"parentProjectFileId"`
		ParentFileLegacyMappingID  interface{} `json:"parentFileLegacyMappingId"`
		FileTypeID                 interface{} `json:"fileTypeId"`
		ExposeAsAlternative        interface{} `json:"exposeAsAlternative"`
		PackageFingerprintID       int         `json:"packageFingerprintId"`
		GameVersionDateReleased    time.Time   `json:"gameVersionDateReleased"`
		GameVersionMappingID       int         `json:"gameVersionMappingId"`
		GameVersionID              int         `json:"gameVersionId"`
		GameID                     int         `json:"gameId"`
		IsServerPack               bool        `json:"isServerPack"`
		ServerPackFileID           int         `json:"serverPackFileId"`
		GameVersionFlavor          interface{} `json:"gameVersionFlavor"`
		Hashes                     []struct {
			Algorithm int    `json:"algorithm"`
			Value     string `json:"value"`
		} `json:"hashes"`
		DownloadCount int `json:"downloadCount"`
	} `json:"latestFiles,omitempty"`
	Categories []struct {
		CategoryID   int       `json:"categoryId"`
		Name         string    `json:"name"`
		URL          string    `json:"url"`
		AvatarURL    string    `json:"avatarUrl"`
		ParentID     int       `json:"parentId"`
		RootID       int       `json:"rootId"`
		ProjectID    int       `json:"projectId"`
		AvatarID     int       `json:"avatarId"`
		GameID       int       `json:"gameId"`
		Slug         string    `json:"slug"`
		DateModified time.Time `json:"dateModified"`
	} `json:"categories,omitempty"`
	Status            int `json:"status,omitempty"`
	PrimaryCategoryID int `json:"primaryCategoryId,omitempty"`
	CategorySection   struct {
		ID                      int         `json:"id"`
		GameID                  int         `json:"gameId"`
		Name                    string      `json:"name"`
		PackageType             int         `json:"packageType"`
		Path                    string      `json:"path"`
		InitialInclusionPattern string      `json:"initialInclusionPattern"`
		ExtraIncludePattern     interface{} `json:"extraIncludePattern"`
		GameCategoryID          int         `json:"gameCategoryId"`
	} `json:"categorySection,omitempty"`
	Slug                   string `json:"slug"`
	GameVersionLatestFiles []struct {
		GameVersion       string      `json:"gameVersion"`
		ProjectFileID     int         `json:"projectFileId"`
		ProjectFileName   string      `json:"projectFileName"`
		FileType          int         `json:"fileType"`
		GameVersionFlavor interface{} `json:"gameVersionFlavor"`
		GameVersionTypeID int         `json:"gameVersionTypeId"`
	} `json:"gameVersionLatestFiles,omitempty"`
	IsFeatured         bool      `json:"isFeatured,omitempty"`
	PopularityScore    float64   `json:"popularityScore,omitempty"`
	GamePopularityRank int       `json:"gamePopularityRank,omitempty"`
	PrimaryLanguage    string    `json:"primaryLanguage,omitempty"`
	GameSlug           string    `json:"gameSlug,omitempty"`
	GameName           string    `json:"gameName,omitempty"`
	PortalName         string    `json:"portalName,omitempty"`
	DateModified       time.Time `json:"dateModified,omitempty"`
	DateCreated        time.Time `json:"dateCreated,omitempty"`
	DateReleased       time.Time `json:"dateReleased,omitempty"`
	IsAvailable        bool      `json:"isAvailable,omitempty"`
	IsExperiemental    bool      `json:"isExperiemental,omitempty"`
	SourceURL          string    `json:"sourceUrl,omitempty"`
	ModLoaders         []string  `json:"modLoaders,omitempty"`
}

type ModPackVersion struct {
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

func GetModPack(id int, client *http.Client) (*ModPack, error) {
	url := fmt.Sprintf(MODPACK_URL, BASE_URL, id)

	data, err := download.DownloadJSON(url, func(b []byte) (interface{}, error) {
		var pack ModPack

		if err := json.Unmarshal(b, &pack); err != nil {
			return nil, fmt.Errorf("error decoding json: %v", err)
		}

		return pack, nil
	}, client)

	if err != nil {
		return nil, fmt.Errorf("error getting modpack: %v", err)
	}

	modpack := data.(ModPack)

	return &modpack, nil
}

func (m *ModPack) GetVersion(versionID int, client *http.Client) (*ModPackVersion, error) {
	versions, err := m.GetVersions()

	if err != nil {
		return nil, fmt.Errorf("error getting versions: %v", err)
	}

	var version *ModPackVersion

	for _, x := range versions {
		if x.ID == versionID {
			version = &x

			break
		}
	}

	if version == nil {
		return nil, fmt.Errorf("error finding match for version %v", versionID)
	}

	return version, nil
}

func (m *ModPack) GetVersions() ([]ModPackVersion, error) {
	url := fmt.Sprintf(VERSIONS_URL, BASE_URL, m.ID)

	data, err := download.DownloadJSON(url, func(b []byte) (interface{}, error) {
		var versions []ModPackVersion

		if err := json.Unmarshal(b, &versions); err != nil {
			return nil, fmt.Errorf("error decoding modpack versions: %v", err)
		}

		return versions, nil
	}, nil)

	if err != nil {
		return nil, fmt.Errorf("error getting modpack versions: %v", err)
	}

	return data.([]ModPackVersion), nil
}
