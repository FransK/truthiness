package api

import (
	"errors"
	"net/http"
)

func (app *Application) getTrialsHandler(w http.ResponseWriter, r *http.Request) {
	experimentname := r.PathValue("experimentname")
	if experimentname == "" {
		app.badRequestResponse(w, r, errors.New("bad request"))
		return
	}

	trials, err := app.Store.Trials(experimentname).GetAll(r.Context())
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, trials); err != nil {
		app.internalServerError(w, r, err)
	}
}
