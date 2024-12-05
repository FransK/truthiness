package api

import (
	"net/http"

	"github.com/fransk/truthiness/internal/store"
)

// Creates an http Handler which handles endpoints, CORS, auth, logging, etc
func NewServer(config *Config, store *store.Storage) http.Handler {
	app := &Application{
		Config: *config,
		Store:  *store,
	}
	mux := http.NewServeMux()

	// routes
	app.addRoutes(mux)

	// middleware
	var handler http.Handler = mux
	handler = app.enableCORS(handler)

	return handler
}

// Middleware to handle CORS
func (app *Application) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173") // Allow specific origin
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")   // Allow specific HTTP methods
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")         // Allow specific headers

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
