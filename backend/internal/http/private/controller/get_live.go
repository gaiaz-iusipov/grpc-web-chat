package httpprivatecontroller

import (
	"net/http"
)

func (*Controller) GetLive(rw http.ResponseWriter, _ *http.Request) {
	rw.WriteHeader(http.StatusNoContent)
}
