package lib

import (
	"io/ioutil"
	"net/http"
)

// Get returns a response from a url
func Get(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}
