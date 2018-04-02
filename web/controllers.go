package web

import (
	"gosiah/catalog"
	"gosiah/solr"
	"log"
	"net/http"
	// "net/url"
)

func home(values RouteValues, resp http.ResponseWriter, req *http.Request) {
	s := NewSession(values, resp, req)
	renderTemplate(s, "views/index.html", nil)
}

func search(values RouteValues, resp http.ResponseWriter, req *http.Request) {

	params := solr.NewSearchParams(
		req.URL.Query(),
		settings.SolrOptions,
		settings.SolrFacets)

	cat := catalog.New(settings.SolrCoreUrl)
	results, err := cat.Search(params)

	s := NewSession(values, resp, req)
	if err != nil {
		renderError(s, "Error during search", err)
	} else {
		renderTemplate(s, "views/results.html", results)
	}
}

func viewOne(values RouteValues, resp http.ResponseWriter, req *http.Request) {
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
