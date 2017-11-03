package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type payload struct {
	Date string /*
		Location string
		Temp     string
		Humidity string
		Wind     string*/
	Status string
}

func index(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Local()
	send := payload{
		Date:   currentTime.Format("01-02-2006"),
		Status: "Active",
	}
	t, err := template.ParseFiles("static/templates/index.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}
	t.Execute(w, send)
}

func main() {

	http.HandleFunc("/", index)
	log.Println("Now serving on localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
