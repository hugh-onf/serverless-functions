package api

import (
	"net/http"
	"os"
	"strings"

	"github.com/hugh-onf/serverless-functions/utils"
)

func HandleApiServiceBurst(w http.ResponseWriter, r *http.Request) {
	methods := strings.Split(strings.TrimSpace(os.Getenv("METHODS")), ",")
	var rpcs []*utils.JsonRpc
	for _, method := range methods {
		rpcs = append(rpcs, utils.NewJsonRpcV2(method, nil))
	}
	if err := utils.BurstHttpRpc(os.Getenv("API_DOMAIN"), rpcs); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
}
