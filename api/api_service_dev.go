package vercel_functions

import (
	"fmt"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, os.Getenv("API_KEY"))
}
