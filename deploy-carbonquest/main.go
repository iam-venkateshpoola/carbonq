package main

import (
	"log"
	"net/http"

	"github.com/venkyGeek/carboq/function"
)

func main() {
	http.HandleFunc("/", function.Carbonquest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
