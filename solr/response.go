package solr

type ResponseHeader struct {
	Status int               `json:"status"`
	QTime  int               `json:"QTime"`
	Params map[string]string `json:"params"`
}

type Response struct {
	NumFound  int        `json:"numFound"`
	Start     int        `json:"start"`
	Documents []Document `json:"docs"`
}

type SolrResponse struct {
	ResponseHeader ResponseHeader `json:"responseHeader"`
	Response       Response       `json:"response"`
}
