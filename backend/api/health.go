package api

import (
	"net/http"
)

func (app *Application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if err := app.jsonResponse(w, http.StatusOK, "OK"); err != nil {
		app.internalServerError(w, r, err)
	}
}
