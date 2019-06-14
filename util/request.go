package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Get(queryUrl string) (string, error) {
	res, err := http.Get(queryUrl)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	return string(body), nil
}
