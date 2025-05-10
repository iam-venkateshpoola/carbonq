package main

import (
	"net/http"
	"log"

	"github.com/venkyGeek/carboq/function"
)

func main() {
	http.HandleFunc("/", function.Carbonquest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}



