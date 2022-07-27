package api

import (
	"net/http"

	"github.com/hugh-onf/serverless-functions/methods"
)

func ApiServiceDev(w http.ResponseWriter, r *http.Request) {
	if err := methods.RpcMethod("api.onfinality.me"); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
}