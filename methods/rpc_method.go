package methods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hugh-onf/serverless-functions/utils"
)

func RpcMethod() {
	defer func() {
		fmt.Println("start running `rpc_methods` call...DONE")
	}()
	fmt.Println("start running `rpc_methods` call...")
	apis := utils.BuildHttpRpcUrl("api.onfinality.me")
	if apis == nil {
		fmt.Println("no endpoints to call")
		return
	}
	postBody, _ := json.Marshal(map[string]interface{}{
		"id":      1,
		"jsonrpc": "2.0",
		"method":  "rpc_methods",
	})
	responseBody := bytes.NewBuffer(postBody)
	for _, api := range apis {
		resp, err := http.Post(api, "application/json", responseBody)
		if err != nil {
			fmt.Printf("ERR: %s", err.Error())
		}
		resp.Body.Close()
	}
}
