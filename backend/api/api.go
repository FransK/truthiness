package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/fransk/truthiness/internal/auth"
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
	handler = app.checkAuthHeaders(handler)

	return handler
}

// Middleware to handle CORS
func (app *Application) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")        // Allow specific origin
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")          // Allow specific HTTP methods
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // Allow specific headers
		w.Header().Set("Access-Control-Allow-Credentials", "true")                    // Allow cookies/credentials

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

// JWT Authorization middleware
func (app *Application) checkAuthHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Skip preflight requests
		if r.Method == http.MethodOptions {
			next.ServeHTTP(w, r)
			return
		}

		// check to see if this path requires authentication
		log.Printf("checking auth headers. URL: %v", r.URL)

		pathRoles := map[string][]string{
			"/v1/upload": {"admin"},
		}

		allowedRoles, ok := pathRoles[r.URL.Path]
		if !ok {
			next.ServeHTTP(w, r)
			return
		}

		// get the user role from token
		userRole, err := auth.ValidateTokenAndGetRole(r)

		// if error, don't serve
		if err != nil {
			app.unauthorized(w, r, err)
			return
		}

		log.Printf("token validated. user role: %s", userRole)

		// check if the user role is in the allowed list
		for _, role := range allowedRoles {
			if userRole == role {
				next.ServeHTTP(w, r)
				return
			}
		}
		app.unauthorized(w, r, errors.New(fmt.Sprintf("role have: %s, role want: %v", userRole, allowedRoles)))
	})
}
