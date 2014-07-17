// Package goDiscoverEtcd provides a mechanism to get a discovery url.
package goDiscoverEtcd

import "net/http"
import "io/ioutil"

// GetURL issues a GET request against the URL https://discovery.etcd.io/new
// returns a string containing your specific URL.
// returns a string of zero length if it has a error
func GetURL() string {

	url := "https://discovery.etcd.io/new"
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	return string(body)
}
