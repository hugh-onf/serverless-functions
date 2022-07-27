package methods

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/hugh-onf/serverless-functions/utils"
)

func RpcMethod(apiDomain string) error {
	apis, err := utils.BuildHttpRpcUrl(apiDomain)
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
		eachError := ""
		if err != nil {
			eachError = err.Error()
		}
		if resp != nil && resp.StatusCode >= 300 {
			eachError = resp.Status
		}
		if len(eachError) > 0 {
			allErrors += fmt.Sprintf("%s - %s\n", api, eachError)
		}
		resp.Body.Close()
	}

	if len(allErrors) > 0 {
		return errors.New(allErrors)
	}

	return nil
}
