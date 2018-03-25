package web

import (
	"log"
	"net/http"
)

var router Router

func init() {
	router.Add("GET", "/catalog/:bib", viewOne)
	router.Add("GET", "/catalog", search)
	router.Add("GET", "/", home)
}

func StartWebServer(address string) {
	log.Printf("Listening for requests at http://%s", address)

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/favicon.ico", fs)
	http.Handle("/robots.txt", fs)
	http.Handle("/public/", http.StripPrefix("/public/", fs))
	http.HandleFunc("/catalog", dispacher)
	http.HandleFunc("/", dispacher)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal("Failed to start the web server: ", err)
	}
}

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
