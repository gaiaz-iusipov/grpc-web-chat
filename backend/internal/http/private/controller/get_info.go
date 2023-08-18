package httpprivatecontroller

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-deeper/app"

	httpheader "github.com/gaiaz-iusipov/grpc-web-chat/internal/http/header"
)

func (*Controller) GetInfo(rw http.ResponseWriter, req *http.Request) {
	response, marshalErr := json.Marshal(app.GetInfo())
	if marshalErr != nil {
		marshalErr = fmt.Errorf("json marshal: %w", marshalErr)
		http.Error(rw, marshalErr.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set(httpheader.ContentType, httpheader.ContentTypeJSON)
	if _, respErr := rw.Write(response); respErr != nil {
		slog.ErrorContext(req.Context(), "failed to write http response", "error", respErr)
	}
}
