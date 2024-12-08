package api

import (
	"log"
	"net/http"
)

func (app *Application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("bad request.\nmethod %v\npath %v\nerror %v", r.Method, r.URL.Path, err.Error())

	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *Application) forbidden(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("forbidden.\nmethod %v\npath %v\nerror %v", r.Method, r.URL.Path, err.Error())

	writeJSONError(w, http.StatusForbidden, "insufficient permissions")
}

func (app *Application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal error.\nmethod %v\npath %v\nerror %v", r.Method, r.URL.Path, err.Error())

	writeJSONError(w, http.StatusInternalServerError, "the server encountered a problem")
}

func (app *Application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("method not allowed.\nmethod %v\npath %v\nerror %v", r.Method, r.URL.Path, err.Error())

	writeJSONError(w, http.StatusMethodNotAllowed, err.Error())
}

func (app *Application) unauthorized(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("unauthorized.\nmethod %v\npath %v\nerror %v", r.Method, r.URL.Path, err.Error())

	writeJSONError(w, http.StatusUnauthorized, "unauthorized")
}
