package main

import (
	"net/http"
	"log"

	// "github.com/venkyGeek/carboq/deploy-carbonquest"
)

func main() {
	http.HandleFunc("/", function.Carbonquest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}



