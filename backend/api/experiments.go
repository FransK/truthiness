package api

import (
	"net/http"
)

func (app *Application) getExperimentsHandler(w http.ResponseWriter, r *http.Request) {
	experiments, err := app.Store.Experiments().GetAll(r.Context())
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, experiments); err != nil {
		app.internalServerError(w, r, err)
	}
}
