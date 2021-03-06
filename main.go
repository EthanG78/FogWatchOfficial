package main

import (
	"encoding/json"
	"fmt"
	"github.com/EthanG78/FogWatchOfficial/payload"
	"html/template"
	"log"
	"net/http"
	"time"
)

const (
	firebase = "https://fogwatch-45fe5.firebaseio.com/"
)

var templates = template.Must(template.ParseGlob("static/*.html"))

//For if statements in templates
var FuncMap = template.FuncMap{
	"eq": func(a, b interface{}) bool {
		return a == b
	},
}

func index(w http.ResponseWriter, r *http.Request) {
	dataKey := time.Now().Format("01-02-2006:15") //Fetching data in this format (m-d-y:hour)
	//Will eventually remove all of this
	toSend := payload.GetPayload(firebase, dataKey)

	t, err := template.ParseFiles("static/index.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}
	t.Execute(w, toSend)
}

//For AJAX update
func updatePayload(w http.ResponseWriter, r *http.Request) {
	dataKey := time.Now().Format("01-02-2006:15")
	updatedPayload := payload.GetPayload(firebase, dataKey)

	data, err := json.Marshal(updatedPayload)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(data)
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

func privact_policy(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/privacypolicy.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}
	t.Execute(w, nil)
}

func main() {
	time.LoadLocation("AST")
	templates = templates.Funcs(FuncMap)

	http.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir("static/media"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("static/js"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))
	http.Handle("/favicons/", http.StripPrefix("/favicons/", http.FileServer(http.Dir("static/favicons"))))

	http.HandleFunc("/", index)
	http.HandleFunc("/update", updatePayload)
	http.HandleFunc("/about.html", about)
	http.HandleFunc("/partners.html", partners)
	http.HandleFunc("/privacy.html", privact_policy)
	log.Println("Now serving on port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
