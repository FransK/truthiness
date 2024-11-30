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
func (app *Application) mount() *http.ServeMux {
	mux := http.NewServeMux()

	// http handlers
	mux.HandleFunc("GET /v1/experiments/{experimentname}/trials", app.getTrialsHandler)
	mux.HandleFunc("GET /v1/experiments", app.getExperimentsHandler)
	mux.HandleFunc("GET /v1/health", app.healthCheckHandler)
	mux.HandleFunc("GET /v1/test", app.testHandler)
	mux.HandleFunc("POST /v1/upload", app.uploadDataHandler)

	return mux
}

// Start an HTTP server to respond to requests to upload data
// and get useful pieces of data for the UI
func (app *Application) run(mux *http.ServeMux) error {
	srv := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	// log info for user
	log.Printf("Starting HTTP server: %s\n", app.Config.Addr)

	// open the port
	return srv.ListenAndServe()
}
