package main

import (
	"fmt"
	"net/http"
)

func Carbonquest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Cloud Function!")
}
