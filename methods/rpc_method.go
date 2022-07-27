package methods

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/hugh-onf/serverless-functions/utils"
)

func RpcMethod() error {
	apis, err := utils.BuildHttpRpcUrl("api.onfinality.me")
	if err != nil {
		return err
	}
	postBody, _ := json.Marshal(map[string]interface{}{
		"id":      1,
		"jsonrpc": "2.0",
		"method":  "rpc_methods",
	})
	responseBody := bytes.NewBuffer(postBody)
	allErrors := ""
	for _, api := range apis {
		resp, err := http.Post(api, "application/json", responseBody)
		if err != nil {
			allErrors += fmt.Sprintf("\n%s", err.Error())
		}
		resp.Body.Close()
	}
	if len(allErrors) > 0 {
		return errors.New(allErrors)
	}
	return nil
}
