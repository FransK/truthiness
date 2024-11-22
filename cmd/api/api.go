package main

import (
	"log"
	"net/http"
	"time"

	"github.com/fransk/truthiness/internal/store"
)

type application struct {
	config config
	store  store.Storage
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleTime  time.Duration
}

// Create an HTTP request multiplexer
func (app *application) mount() *http.ServeMux {
	mux := http.NewServeMux()

	// http handlers
	mux.HandleFunc("GET /v1/health", app.healthCheckHandler)
	mux.HandleFunc("GET /v1/experiment/{experimentname}/trials", app.getExperimentHandler)
	mux.HandleFunc("GET /v1/test", app.testHandler)
	mux.HandleFunc("POST /v1/upload", app.uploadDataHandler)

	return mux
}

// Start an HTTP server to respond to requests to upload data
// and get useful pieces of data for the UI
func (app *application) run(mux *http.ServeMux) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	// log info for user
	log.Printf("Starting HTTP server: %s\n", app.config.addr)

	// open the port
	return srv.ListenAndServe()
}
