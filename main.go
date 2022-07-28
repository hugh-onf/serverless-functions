package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/hugh-onf/serverless-functions/api"
	"github.com/hugh-onf/serverless-functions/utils"
)

var port = 43210

func main() {
	// Simple args for now
	args := os.Args[1:]
	if len(args) > 0 {
		if args[0] == "server" {
			http.HandleFunc("/api/api_service_calls", api.HandleApiServiceBurst)
			fmt.Printf("Server is listening at :%d", port)
			http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
		}

		if args[0] == "burst" {
			if total, err := utils.BurstHttpRpc(); err == nil {
				fmt.Printf("API Service burst calls done, requests sent: %d", total)
			} else {
				panic(err)
			}
		}
	} else {
		fmt.Println("Missing args: server | burst")
	}
}
