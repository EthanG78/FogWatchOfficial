package main

import (
	"github.com/EthanG78/fog_watch/payload"
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
	//Will eventually remove all of this
	toSend := payload.Payload{
		Date:     payload.GetPayloadField(firebase, "TestData", "Date"),
		Location: payload.GetPayloadField(firebase, "TestData", "Location"),
		Temp:     payload.GetPayloadField(firebase, "TestData", "Temp"),
		Humidity: payload.GetPayloadField(firebase, "TestData", "Humidity"),
		WindS:    payload.GetPayloadField(firebase, "TestData", "WindS"),
		Status:   payload.GetPayloadField(firebase, "TestData", "Status"),
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

	//API CALLS
	toSend := payload.Payload{}
	toSend.SetDate(currentTime.Format("01-02-2006"))
	toSend.SetStatus("Active")

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
