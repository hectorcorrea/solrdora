package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var router Router

func init() {
	router.Add("GET", "/catalog/:bib", catView)
	router.Add("GET", "/catalog", catSearch)
	// router.Add("GET", "/catalog/", catSearchPost, "catalog2")
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

func renderNotFound(s Session) {
	log.Printf(fmt.Sprintf("Not found (%s)", s.Req.URL.Path))
	t, err := template.New("layout").ParseFiles("views/layout.html", "views/notFound.html")
	if err != nil {
		log.Printf("Error rendering not found page :(")
		// perhaps render a hard coded string?
	} else {
		s.Resp.WriteHeader(http.StatusNotFound)
		t.Execute(s.Resp, nil)
	}
}

func renderError(s Session, title string, err error) {
	// TODO: log more about the Request
	log.Printf("ERROR: %s - %s (%s)", title, err, "TODO: s.req.URL.Path")
	// vm := viewModels.NewError(title, err, s.toViewModel())
	t, err := template.New("layout").ParseFiles("views/layout.html", "views/error.html")
	if err != nil {
		log.Printf("Error rendering error page :(")
		// perhaps render a hard coded string?
	} else {
		s.Resp.WriteHeader(http.StatusInternalServerError)
		t.Execute(s.Resp, nil)
	}
}

func loadTemplate(s Session, viewName string) (*template.Template, error) {
	t, err := template.New("layout").ParseFiles("views/layout.html", viewName)
	if err != nil {
		renderError(s, fmt.Sprintf("Loading view %s", viewName), err)
		return nil, err
	} else {
		log.Printf("Loaded template %s (%s)", viewName, "TODO s.req.URL.Path")
		return t, nil
	}
}

func renderTemplate(s Session, viewName string, viewModel interface{}) {
	t, err := loadTemplate(s, viewName)
	if err != nil {
		log.Printf("Error loading: %s, %s ", viewName, err)
	} else {
		err = t.Execute(s.Resp, viewModel)
		if err != nil {
			log.Printf("Error rendering: %s, %s ", viewName, err)
		}
	}
}
