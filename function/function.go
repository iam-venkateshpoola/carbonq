package function

import (
	"fmt"
	"net/http"
)

// Function is the Cloud Function entrypoint
func Function(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1><strong>Hello from Cloud Function!</strong></h1>")
}
