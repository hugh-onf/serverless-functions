package utils

import "fmt"

func BuildHttpRpcUrl(apiDomain string) []string {
	apiKey := getApiKey()
	if len(apiKey) == 0 {
		fmt.Println("no api key set, skip this run")
		return nil
	}
	var urls []string
	networks := getNetworks()
	for _, n := range networks {
		urls = append(urls, fmt.Sprintf("https://%s.%s/rpc?apiKey=%s", n, apiDomain, apiKey))
	}
	return urls
}
