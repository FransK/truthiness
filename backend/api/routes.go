package api

import "net/http"

func (app *Application) addRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /v1/experiments/{experimentname}/trials", app.getTrialsHandler)
	mux.HandleFunc("GET /v1/experiments", app.getExperimentsHandler)
	mux.HandleFunc("GET /v1/health", app.healthCheckHandler)
	mux.HandleFunc("GET /v1/test", app.testHandler)
	mux.HandleFunc("POST /v1/upload", app.uploadDataHandler)
}
