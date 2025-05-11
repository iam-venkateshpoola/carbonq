package function

import (
	"fmt"
	"net/http"
)

// Carbonquest is the Cloud Function entrypoint
func Carbonquest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from CarbonQ Cloud Function!")
}
