package api

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fransk/truthiness/internal/auth"
	"github.com/fransk/truthiness/internal/store"
	"github.com/fransk/truthiness/internal/store/inmemorystore"
)

func TestRoutes(t *testing.T) {
	experiment := store.Experiment{
		Name:     "TestExperiment",
		Date:     "March 22 2024",
		Location: "SFU",
		Records:  []string{"Age", "Difference"},
	}
	trial := store.Trial{
		Data: map[string]string{"Age": "20", "Difference": "0.2"},
	}
	var storage store.Storage
	storage = inmemorystore.New()
	storage.Experiments().Create(context.TODO(), &experiment)
	storage.Trials(experiment.Name).Create(context.TODO(), &trial)
	cfg := &Config{
		Addr: "testhost:1111",
	}

	srv := NewServer(cfg, &storage)

	// Verify routes are set correctly
	testCases := []struct {
		method string
		addr   string
		body   io.Reader
		want   int
	}{
		{http.MethodGet, "/v1/authenticate", nil, http.StatusMethodNotAllowed},
		{http.MethodPost, "/v1/authenticate", nil, http.StatusBadRequest},
		{http.MethodGet, "/v1/experiments", nil, http.StatusOK},
		{http.MethodGet, "/v1/experiments/TestExperiment/trials", nil, http.StatusOK},
		{http.MethodGet, "/v1/experiments/NonexistentExperiment/trials", nil, http.StatusOK},
		{http.MethodGet, "/v1/health", nil, http.StatusOK},
		{http.MethodGet, "/v1/upload", nil, http.StatusMethodNotAllowed},
		{http.MethodPost, "/v1/upload", nil, http.StatusBadRequest},
		{http.MethodGet, "/", nil, http.StatusNotFound},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%s at %s", tc.method, tc.addr), func(t *testing.T) {
			t.Parallel()
			req := httptest.NewRequest(tc.method, tc.addr, tc.body)
			rr := httptest.NewRecorder()
			srv.ServeHTTP(rr, req)
			if rr.Code != tc.want {
				t.Errorf("got %v; want %v", rr.Code, tc.want)
			}
		})
	}
}

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

func TestCheckAuthHeaders(t *testing.T) {
	app := &Application{}
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	handler := app.checkAuthHeaders(nextHandler)

	// Test no authentication headers on protected area
	t.Run("protected-notoken", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/upload", nil)
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		want := http.StatusUnauthorized
		if got := rr.Code; got != want {
			t.Errorf("got %v; want %v", got, want)
		}
	})

	// Test proper authenication headers on protected area
	t.Run("protected-token-good", func(t *testing.T) {
		tokenStr, err := auth.CreateToken("me", "admin")
		if err != nil {
			t.Errorf("failed to create token")
		}

		req := httptest.NewRequest(http.MethodPost, "/upload", nil)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tokenStr))
		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)

		dontwant := http.StatusUnauthorized
		if got := rr.Code; got == dontwant {
			t.Errorf("got %v; don't want %v", got, dontwant)
		}
	})

	// Test invalid authentication headers on protected area
	t.Run("protected-token-invalid", func(t *testing.T) {
		tokenStr, err := auth.CreateToken("me", "user")
		if err != nil {
			t.Errorf("failed to create token")
		}

		req := httptest.NewRequest(http.MethodPost, "/upload", nil)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tokenStr))
		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)

		want := http.StatusUnauthorized
		if got := rr.Code; got != want {
			t.Errorf("got %v; don't want %v", got, want)
		}
	})

	// Test authentication headers on unprotected area
	t.Run("unprotected", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/health", nil)
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		want := http.StatusOK
		if got := rr.Code; got != want {
			t.Errorf("got %v; want %v", got, want)
		}
	})
}
