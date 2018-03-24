package web

import (
	"net/http"
)

// Session aggregates information about the current request/response
// objects plus a few other data points that we care about
type Session struct {
	Resp      http.ResponseWriter
	Req       *http.Request
	URLValues map[string]string
}

// NewSession creates a new session
func NewSession(values map[string]string, resp http.ResponseWriter, req *http.Request) Session {
	return Session{URLValues: values, Resp: resp, Req: req}
}
