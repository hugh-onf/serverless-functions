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
	"strings"
	"sync"
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
	postBody, err := json.Marshal(rpc)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	payload := bytes.NewBuffer(postBody)
	resp, err := http.Post(url, "application/json", payload)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
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

func BurstHttpRpc() (int, error) {
	methods := strings.Split(strings.TrimSpace(os.Getenv("METHODS")), ",")
	var rpcs []*JsonRpc
	for _, method := range methods {
		rpcs = append(rpcs, NewJsonRpcV2(method, nil))
	}
	maxBurstVal := os.Getenv("BURST_MAX")
	maxBurst, err := strconv.ParseUint(maxBurstVal, 0, 16)
	if err != nil || maxBurst > MAX_BURST {
		maxBurst = MAX_BURST
	}

	minBurstVal := os.Getenv("BURST_MIN")
	minBurst, err := strconv.ParseUint(minBurstVal, 0, 16)
	if err != nil || minBurst < MIN_BURST {
		minBurst = MIN_BURST
	}

	urls, err := buildHttpRpcUrl(os.Getenv("API_DOMAIN"))
	if err != nil {
		return 0, err
	}
	// Random call burst for more realistic data
	rand.Seed(time.Now().UnixNano())
	callsCount := rand.Intn(int(maxBurst)-int(minBurst)+1) + int(minBurst)
	totalRequests := len(urls) * len(rpcs) * callsCount

	wg := &sync.WaitGroup{}
	wg.Add(totalRequests)
	for _, url := range urls {
		for _, rpc := range rpcs {
			for i := 0; i < callsCount; i++ {
				// Fire and forget, hope for the best
				go func() {
					defer wg.Done()
					msg := fmt.Sprintf("%s > %s | %v => ", url, rpc.Method, rpc.Params)
					body, _, err := callHttpRpc(url, rpc)
					if err != nil {
						msg += fmt.Sprintf("%s - %s", err.Error(), string(body))
					} else {
						msg += "OK"
					}
					fmt.Println(msg)
				}()
			}
		}
	}
	wg.Wait()

	return totalRequests, nil
}
