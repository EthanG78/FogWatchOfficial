package main

import (
	"log"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("./static/templates")))
	log.Println("Now serving on localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
