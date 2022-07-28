package utils

import (
	"os"
	"strings"
)

var defaultNetwork = []string{"acala-dev"}

func getApiKeys() []string {
	return strings.Split(strings.TrimSpace(os.Getenv("API_KEYS")), ",")
}

func getNetworks() []string {
	envValue := strings.TrimSpace(os.Getenv("NETWORKS"))
	if len(envValue) == 0 {
		return defaultNetwork
	}
	return strings.Split(envValue, ",")
}
