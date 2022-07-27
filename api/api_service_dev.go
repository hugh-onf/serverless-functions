package vercel_functions

import (
	"net/http"

	"github.com/hugh-onf/serverless-functions/methods"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	methods.RpcMethod()
}
