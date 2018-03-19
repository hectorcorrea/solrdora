package web

import (
	"net/http"
)

// Aggregates information about the current request/response
// objects plus a few other data points that we care about
type Session struct {
	Resp      http.ResponseWriter
	Req       *http.Request
	UrlValues map[string]string
}

func NewSession(values map[string]string, resp http.ResponseWriter, req *http.Request) Session {
	return Session{UrlValues: values, Resp: resp, Req: req}
}
