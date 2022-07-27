package utils

import (
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
	return strings.Split(envValue, ",")
}
