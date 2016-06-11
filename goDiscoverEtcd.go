// Package goDiscoverEtcd provides a mechanism to get a discovery url.
package goDiscoverEtcd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetDiscoveryURL issues a GET request against the URL https://discovery.etcd.io/new
// returns a string containing your specific URL.
// returns a string of zero length and the relevant error
func GetDiscoveryURL(size int) (string, error) {

	if size < 1 {
		return "", errors.New("size must be greater than1")
	}

	discoveryURL := fmt.Sprintf("https://discovery.etcd.io/new?size=%v", size)

	resp, err := http.Get(discoveryURL)
	defer resp.Body.Close()

	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
