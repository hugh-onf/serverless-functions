package methods

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

	allErrors := ""

	for _, api := range apis {
		payload := bytes.NewBuffer(postBody)
		resp, err := http.Post(api, "application/json", payload)
		eachError := ""
		if err != nil {
			eachError = err.Error()
		}
		if resp != nil && resp.StatusCode >= 300 {
			eachError = resp.Status
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				eachError += " - unknown error"
			}
			eachError += fmt.Sprintf(" - %s", string(b))
		}
		if len(eachError) > 0 {
			allErrors += fmt.Sprintf("%s - %s\n", api, eachError)
		}
		if resp != nil {
			resp.Body.Close()
		}
	}

	if len(allErrors) > 0 {
		return errors.New(allErrors)
	}

	return nil
}
