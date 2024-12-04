package api

import (
	"log"
	"net/http"
	"time"
)

func (app *Application) RunNew() error {
	return app.run(app.mount())
}

// Create an HTTP request multiplexer
func (app *Application) mount() http.Handler {
	mux := http.NewServeMux()

	// routes
	mux.HandleFunc("GET /v1/experiments/{experimentname}/trials", app.getTrialsHandler)
	mux.HandleFunc("GET /v1/experiments", app.getExperimentsHandler)
	mux.HandleFunc("GET /v1/health", app.healthCheckHandler)
	mux.HandleFunc("GET /v1/test", app.testHandler)
	mux.HandleFunc("POST /v1/upload", app.uploadDataHandler)

	// middleware
	var handler http.Handler = mux
	handler = app.enableCORS(handler)

	return handler
}

// Start an HTTP server to respond to requests to upload data
// and get useful pieces of data for the UI
func (app *Application) run(handler http.Handler) error {
	srv := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      handler,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	// log info for user
	log.Printf("Starting HTTP server: %s\n", app.Config.Addr)

	// open the port
	return srv.ListenAndServe()
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
