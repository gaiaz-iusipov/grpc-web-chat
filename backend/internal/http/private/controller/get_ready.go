package httpprivatecontroller

import (
	"net/http"
)

func (c *Controller) GetReady(rw http.ResponseWriter, _ *http.Request) {
	if c.ready.Load() {
		rw.WriteHeader(http.StatusNoContent)
		return
	}

	rw.WriteHeader(http.StatusServiceUnavailable)
}
