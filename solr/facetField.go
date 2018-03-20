package solr

type FacetValues struct {
	text  string
	count int
	// remove_url string
	// add_url string
}

type FacetField struct {
	name   string
	title  string
	values []FacetValues
}
