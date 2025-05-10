package function

import (
	"fmt"
)

// Carbonquest is the Cloud Function entrypoint
func Carbonquest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Cloud Function!")
}
