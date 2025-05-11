package function

import (
	"fmt"
	"net/http"
)

// Carbonquest is the Cloud Function entrypoint
func Carbonquest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1><strong>Hello from CarbonQuest Cloud Function!</strong></h1>")
}
