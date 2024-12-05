package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEnableCORS(t *testing.T) {
	app := &Application{}
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	handler := app.enableCORS(nextHandler)

	// Test CORS headers
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if got := rr.Header().Get("Access-Control-Allow-Origin"); got != "http://localhost:5173" {
		t.Errorf("Access-Control-Allow-Origin = %q; want %q", got, "http://localhost:5173")
	}

	// Test preflight request
	req = httptest.NewRequest(http.MethodOptions, "/", nil)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Preflight OPTIONS returned status %v; want %v", rr.Code, http.StatusOK)
	}
}
