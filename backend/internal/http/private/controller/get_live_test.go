package httpprivatecontroller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	httpprivatecontroller "github.com/gaiaz-iusipov/grpc-web-chat/internal/http/private/controller"
)

func TestController_GetLive(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/live", nil)
	controller := &httpprivatecontroller.Controller{}

	controller.GetLive(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Errorf("expected status code %v but got %v", http.StatusOK, rec.Code)
	}
}
