package main

import (
	"fmt"
	"net/http"

	"github.com/hugh-onf/serverless-functions/api"
)

var port = 43210

func main() {
	http.HandleFunc("/api/api_service_calls", api.HandleApiServiceBurst)
	fmt.Printf("Server is listening at :%d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
