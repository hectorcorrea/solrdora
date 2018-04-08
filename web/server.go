package web

import (
	"log"
	"net/http"
	"solrdora/models"
)

var router Router
var settings models.Settings

func init() {
	router.Add("GET", "/view/:id", viewOne)
	router.Add("GET", "/about", about)
	router.Add("GET", "/search", search)
	router.Add("GET", "/", home)
}

func StartWebServer(settingsFile string) {
	var err error
	settings, err = models.LoadSettings(settingsFile)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Loaded settings from: %s", settingsFile)
	log.Printf("Listening for requests at http://%s", settings.ServerAddress)

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/favicon.ico", fs)
	http.Handle("/robots.txt", fs)
	http.Handle("/public/", http.StripPrefix("/public/", fs))
	http.HandleFunc("/", dispacher)

	err = http.ListenAndServe(settings.ServerAddress, nil)
	if err != nil {
		log.Fatal("Failed to start the web server: ", err)
	}
}

// Dispatches the request to one our custom routes
func dispacher(resp http.ResponseWriter, req *http.Request) {
	found, route := router.FindRoute(req.Method, req.URL.Path)
	if found {
		values := route.UrlValues(req.URL.Path)
		route.handler(values, resp, req)
	} else {
		s := NewSession(map[string]string{}, resp, req)
		renderNotFound(s)
	}
}
