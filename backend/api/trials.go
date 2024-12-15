package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/fransk/truthiness/internal/stats"
	"github.com/fransk/truthiness/internal/utils"
)

func (app *Application) getTrialsHandler(w http.ResponseWriter, r *http.Request) {
	experimentname := r.PathValue("experimentname")
	if experimentname == "" {
		app.badRequestResponse(w, r, errors.New("bad request"))
		return
	}

	query := r.URL.Query()

	// Check to see if the user wants a regression model
	model := query.Get("model")
	xaxis := query.Get("x_axis")
	yaxis := query.Get("y_axis")
	if xaxis == "" || yaxis == "" {
		trials, err := app.Store.Trials(experimentname).GetAll(r.Context())
		if err != nil {
			app.internalServerError(w, r, err)
			return
		}

		if err := app.jsonResponse(w, http.StatusOK, trials); err != nil {
			app.internalServerError(w, r, err)
		}
		return
	}

	trials, err := app.Store.Trials(experimentname).Get(r.Context(), []string{xaxis, yaxis})
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if model == "" {
		if err := app.jsonResponse(w, http.StatusOK, trials); err != nil {
			app.internalServerError(w, r, err)
		}
		return
	}

	// User has specified a model and 2 axes, perform regression
	var xs, ys []float64
	missingData := 0
	for _, trial := range trials {
		v, ok := trial.Data[xaxis]
		if !ok || v == nil {
			missingData++
			continue
		}
		x, err := utils.GetFloat(v)
		if err != nil {
			missingData++
			continue
		}

		v, ok = trial.Data[yaxis]
		if !ok || v == nil {
			missingData++
			continue
		}
		y, err := utils.GetFloat(v)
		if err != nil {
			missingData++
			continue
		}

		xs = append(xs, x)
		ys = append(ys, y)
	}

	log.Printf("skipped %d out of %d trials due to missing data", missingData, len(trials))

	regression, err := stats.LinearLeastSquares(xs, ys)
	if err != nil {
		app.internalServerError(w, r, fmt.Errorf("unable to compute least squares regression"))
		return
	}
	// regression successful, append results
	for _, trial := range trials {
		if trial.Data[xaxis] == nil {
			continue
		}

		x, err := utils.GetFloat(trial.Data[xaxis])
		if err != nil {
			log.Printf("unable to get float from trial data: %v", err)
			continue
		}
		trial.Data["LineY"] = regression.M*x + regression.B
	}
	if err := app.jsonResponse(w, http.StatusOK, trials); err != nil {
		app.internalServerError(w, r, err)
	}
}
