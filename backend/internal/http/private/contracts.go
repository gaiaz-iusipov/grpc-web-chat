package httpprivate

import "net/http"

type Controller interface {
	GetLive(rw http.ResponseWriter, req *http.Request)
	GetReady(rw http.ResponseWriter, req *http.Request)
	GetInfo(rw http.ResponseWriter, req *http.Request)
}
