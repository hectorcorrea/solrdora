package web

import (
	"gosiah/catalog"
	"gosiah/solr"
	"log"
	"net/http"
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

	params.Fl = settings.SearchFl
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
	cat := catalog.New(settings.SolrCoreUrl)
	record, err := cat.Get(values["id"], settings.ViewOneFl)

	s := NewSession(values, resp, req)
	if err != nil {
		renderError(s, "Error retrieving document from Solr", err)
	} else {
		renderTemplate(s, "views/one.html", record)
	}
}
