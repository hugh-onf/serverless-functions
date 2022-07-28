package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

const MAX_BURST = 5000
const MIN_BURST = 100

func buildHttpRpcUrl(apiDomain string) ([]string, error) {
	apiKeys := getApiKeys()
	networks := getNetworks()
	var urls []string
	for _, k := range apiKeys {
		for _, n := range networks {
			urls = append(urls, fmt.Sprintf("https://%s.%s/rpc?apikey=%s", n, apiDomain, k))
		}
	}
	return urls, nil
}

func callHttpRpc(url string, rpc *JsonRpc) ([]byte, int, error) {
	postBody, _ := json.Marshal(rpc)
	payload := bytes.NewBuffer(postBody)
	resp, err := http.Post(url, "application/json", payload)
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if resp != nil {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}
		if resp.StatusCode >= 300 {
			return b, resp.StatusCode, errors.New(resp.Status)
		}
		return b, http.StatusOK, nil
	}
	return nil, http.StatusNoContent, nil
}

func BurstHttpRpc(apiDomain string, methods []*JsonRpc) error {
	burstVal := os.Getenv("BURST")
	burst, err := strconv.ParseUint(burstVal, 0, 16)
	// Clamp max burst
	if err != nil || burst > MAX_BURST {
		burst = MAX_BURST
	}
	if burst < MIN_BURST {
		burst = MIN_BURST
	}
	urls, err := buildHttpRpcUrl(apiDomain)
	if err != nil {
		return err
	}
	for _, url := range urls {
		for _, method := range methods {
			// Random call burst for more realistic data
			rand.Seed(time.Now().UnixNano())
			callsCount := rand.Intn(int(burst)-MIN_BURST+1) + MIN_BURST
			for i := 0; i < callsCount; i++ {
				// Fire and forget, hope for the best
				go func() {
					callHttpRpc(url, method)
				}()
			}
		}
	}
	return nil
}
