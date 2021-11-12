package download

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func Download(url string, client *http.Client) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("error creating GET request: %v", err)
	}

	var res *http.Response

	if client == nil {
		res, err = http.DefaultClient.Do(req)
	} else {
		res, err = client.Do(req)
	}

	if err != nil {
		return nil, fmt.Errorf("error request from %v: %v", req.Host, err)
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	log.Printf("Downloaded %v", url)

	return data, nil
}

func DownloadFile(url, path string, client *http.Client) (string, error) {
	filename := filepath.Base(url)

	data, err := Download(url, client)

	if err != nil {
		return "", fmt.Errorf("error downloading %v: %v", url, err)
	}

	downloadPath := filepath.Join(path, filename)

	outputFile, err := os.Create(downloadPath)

	if err != nil {
		return "", fmt.Errorf("error creating output file %v: %v", downloadPath, err)
	}

	defer outputFile.Close()

	_, err = outputFile.Write(data)

	if err != nil {
		return "", fmt.Errorf("error writing data to %v: %v", downloadPath, err)
	}

	log.Printf("Wrote data from %v -> %v", url, downloadPath)

	return downloadPath, nil
}

func DownloadExtract(url, path string, client *http.Client) error {
	downloadPath, err := DownloadFile(url, path, client)

	if err != nil {
		return fmt.Errorf("error downloading %v: %v", url, err)
	}

	if err := Unzip(downloadPath, path); err != nil {
		return fmt.Errorf("error unziping %v: %v", downloadPath, err)
	}

	return nil
}

func DownloadJSON(url string, convert func([]byte) (interface{}, error), client *http.Client) (interface{}, error) {
	data, err := Download(url, client)

	if err != nil {
		return nil, fmt.Errorf("error downloading json data: %v", err)
	}

	return convert(data)
}
