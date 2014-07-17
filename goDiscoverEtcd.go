// Package goDiscoverEtcd provides a mechanism to get a discovery url.
package goDiscoverEtcd

import "net/http"
import "io/ioutil"

// GetDiscoveryURL issues a GET request against the URL https://discovery.etcd.io/new
// returns a string containing your specific URL.
// returns a string of zero length if it has a error
func GetDiscoveryURL() string {

	discoveryURL := "https://discovery.etcd.io/new"

	resp, err := http.Get(discoveryURL)
	defer resp.Body.Close()

	if err != nil {
		return ""
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	return string(body)
}
