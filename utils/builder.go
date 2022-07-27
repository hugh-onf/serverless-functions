package utils

import (
	"errors"
	"fmt"
)

func BuildHttpRpcUrl(apiDomain string) ([]string, error) {
	apiKey := getApiKey()
	if len(apiKey) == 0 {
		return nil, errors.New("no api key set, skip this run")
	}
	var urls []string
	networks := getNetworks()
	for _, n := range networks {
		urls = append(urls, fmt.Sprintf("https://%s.%s/rpc?apiKey=%s", n, apiDomain, apiKey))
	}
	return urls, nil
}
