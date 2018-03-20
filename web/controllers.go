package web

import (
	"gosiah/catalog"
	"gosiah/solr"
	"log"
	"net/http"
	"strings"
)

func home(values RouteValues, resp http.ResponseWriter, req *http.Request) {
	s := NewSession(values, resp, req)
	renderTemplate(s, "views/index.html", nil)
}

func catSearch(values RouteValues, resp http.ResponseWriter, req *http.Request) {
	url := "http://localhost:8983/solr/bibdata"
	cat := catalog.New(url)
	q := strings.Join(req.URL.Query()["q"], " ")
	options := map[string]string{
		"defType": "edismax",
		"qf":      "authorsAll title^100",
	}
	params := solr.SearchParams{Q: q, Options: options}
	records, err := cat.Search(params)

	s := NewSession(values, resp, req)
	if err != nil {
		renderError(s, "Error during search", err)
	} else {
		renderTemplate(s, "views/results.html", records)
	}
}

func catView(values RouteValues, resp http.ResponseWriter, req *http.Request) {
	bib := values["bib"]
	log.Printf("fetching bib: %s", bib)

	url := "http://localhost:8983/solr/bibdata"
	cat := catalog.New(url)
	record, err := cat.Get(bib)

	s := NewSession(values, resp, req)
	if err != nil {
		renderError(s, "Error retrieving document from Solr", err)
	} else {
		renderTemplate(s, "views/one.html", record)
	}
}
