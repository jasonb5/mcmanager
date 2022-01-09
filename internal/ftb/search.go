package ftb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Result struct {
	Packs      []int `json:"packs"`
	Curseforge []int `json:"curseforge"`
	Total      int   `json:"total"`
	Limit      int   `json:"limit"`
	Refreshed  int   `json:"refreshed"`
}

type Version struct {
	Specs struct {
		ID          int `json:"id"`
		Minimum     int `json:"minimum"`
		Recommended int `json:"recommended"`
	} `json:"specs"`
	Targets []struct {
		Version string `json:"version"`
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Type    string `json:"type"`
		Updated int    `json:"updated"`
	} `json:"targets"`
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Updated int    `json:"updated"`
}

type Manifest struct {
	Synopsis    string `json:"synopsis"`
	Description string `json:"description"`
	Art         []struct {
		Width      int           `json:"width"`
		Height     int           `json:"height"`
		Compressed bool          `json:"compressed"`
		URL        string        `json:"url"`
		Mirrors    []interface{} `json:"mirrors"`
		Sha1       string        `json:"sha1"`
		Size       int           `json:"size"`
		ID         int           `json:"id"`
		Type       string        `json:"type"`
		Updated    int           `json:"updated"`
	} `json:"art"`
	Links   []interface{} `json:"links"`
	Authors []struct {
		Website string `json:"website"`
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Type    string `json:"type"`
		Updated int    `json:"updated"`
	} `json:"authors"`
	Versions []Version `json:"versions"`
	Installs int       `json:"installs"`
	Plays    int       `json:"plays"`
	Tags     []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"tags"`
	Featured     bool   `json:"featured"`
	Refreshed    int    `json:"refreshed"`
	Notification string `json:"notification"`
	Rating       struct {
		ID             int  `json:"id"`
		Configured     bool `json:"configured"`
		Verified       bool `json:"verified"`
		Age            int  `json:"age"`
		Gambling       bool `json:"gambling"`
		Frightening    bool `json:"frightening"`
		Alcoholdrugs   bool `json:"alcoholdrugs"`
		Nuditysexual   bool `json:"nuditysexual"`
		Sterotypeshate bool `json:"sterotypeshate"`
		Language       bool `json:"language"`
		Violence       bool `json:"violence"`
	} `json:"rating"`
	Status  string `json:"status"`
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Updated int    `json:"updated"`
}

func getSearchResult(term string) (*Result, error) {
	searchURL := fmt.Sprintf("https://api.modpacks.ch/public/modpack/search/5?term=%s", term)

	res, err := http.Get(searchURL)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var result Result

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func getModPackManifest(id int) (*Manifest, error) {
	manifestURL := fmt.Sprintf("https://api.modpacks.ch/public/modpack/%d", id)

	res, err := http.Get(manifestURL)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var manifest Manifest

	if err := json.Unmarshal(data, &manifest); err != nil {
		return nil, err
	}

	return &manifest, nil
}

func Search(term string) ([]*Manifest, error) {
	result, err := getSearchResult(term)

	if err != nil {
		return nil, err
	}

	var manifests []*Manifest

	for _, item := range result.Packs {
		manifest, err := getModPackManifest(item)

		if err != nil {
			return nil, err
		}

		manifests = append(manifests, manifest)
	}

	return manifests, nil
}
