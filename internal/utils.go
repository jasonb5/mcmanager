package mcmanager

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
		return nil, fmt.Errorf("error executing request to %v: %v", req.Host, err)
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("error reading data from response body: %v", err)
	}

	return data, err
}
