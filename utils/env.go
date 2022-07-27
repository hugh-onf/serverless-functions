package utils

import (
	"encoding/json"
	"os"
	"strings"
)

var defaultNetwork = []string{"acala-dev"}

func getApiKey() string {
	return os.Getenv("API_KEY")
}

func getNetworks() []string {
	envValue := strings.TrimSpace(os.Getenv("NETWORKS"))
	if len(envValue) == 0 {
		return defaultNetwork
	}
	var networks *[]string
	err := json.Unmarshal([]byte(envValue), networks)
	if err != nil {
		return defaultNetwork
	}
	return *networks
}
