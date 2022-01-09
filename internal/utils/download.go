package utils

import (
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
)

type QueryUpdater = func(*http.Request) error

func Download(url string, queryUpdater QueryUpdater) ([]byte, error) {
	log.Printf("Downloading %s", url)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, nil
	}

	if queryUpdater != nil {
		if err := queryUpdater(req); err != nil {
			return nil, nil
		}
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, nil
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, nil
	}

	log.Printf("Downloaded %d bytes from %s", len(data), url)

	return data, nil
}

func DownloadFile(url, dst string) (string, error) {
	log.Printf("Downloading %s -> %s", url, dst)

	res, err := http.DefaultClient.Get(url)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	if err := ioutil.WriteFile(dst, data, fs.ModePerm); err != nil {
		return "", err
	}

	log.Printf("Wrote %d bytes to %s", len(data), dst)

	return dst, nil
}
