package api

import (
	"fmt"
	"net/http"

	"github.com/hugh-onf/serverless-functions/utils"
)

func HandleApiServiceBurst(w http.ResponseWriter, r *http.Request) {
	if total, err := utils.BurstHttpRpc(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("OK, requests sent: %d", total)))
	}
}
