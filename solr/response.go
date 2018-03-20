package solr

type Header struct {
	Status int               `json:"status"`
	QTime  int               `json:"QTime"`
	Params map[string]string `json:"params"`
}

type Response struct {
	NumFound  int        `json:"numFound"`
	Start     int        `json:"start"`
	Documents []Document `json:"docs"`
}

type Error struct {
	Trace string `json:"trace"`
	Code  int    `json:"code"`
}

type SolrResponse struct {
	Header   Header   `json:"responseHeader"`
	Response Response `json:"response"`
	Error    Error    `json:"error"`
}
