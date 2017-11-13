package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

var templates = template.Must(template.ParseGlob("static/*.html"))

//For if statements in templates
var FuncMap = template.FuncMap{
	"eq": func(a, b interface{}) bool {
		return a == b
	},
}

//This struct will be removed in the future
type payload struct {
	Date string
	/*
		Location string
		Temp     string
		Humidity string
		Wind     string
	*/
	Status string
}

func index(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Local()
	toSend := payload{
		Date:   currentTime.Format("01-02-2006"),
		Status: "Active",
	}
	t, err := template.ParseFiles("static/index.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}
	t.Execute(w, toSend)
}

//Just for serving static files
func about(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/about.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}
	t.Execute(w, nil)
}

func partners(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/partners.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}
	t.Execute(w, nil)
}

func monitorStatus(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Local()
	toSend := payload{
		Date:   currentTime.Format("01-02-2006"),
		Status: "Active",
	}
	t, err := template.ParseFiles("static/status.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}
	t.Execute(w, toSend)
}

func main() {
	templates = templates.Funcs(FuncMap)

	http.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir("static/media"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("static/js"))))

	http.HandleFunc("/", index)
	http.HandleFunc("/about.html", about)
	http.HandleFunc("/partners.html", partners)
	http.HandleFunc("/status.html", monitorStatus)

	log.Println("Now serving on localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

//TODO: Design how to present the firebase data on the site using payload struct
