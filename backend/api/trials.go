package api

import (
	"errors"
	"fmt"
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

	trials, err := app.Store.Trials(experimentname).GetAll(r.Context())
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	query := r.URL.Query()

	// Check to see if the user wants a regression model
	model := query.Get("model")
	xaxis := query.Get("x_axis")
	yaxis := query.Get("y_axis")
	if model == "linear" && xaxis != "" && yaxis != "" {
		var xs, ys []float64
		for _, trial := range trials {
			x, err := utils.GetFloat(trial.Data[xaxis])
			if err != nil {
				app.internalServerError(w, r, fmt.Errorf("unable to use x axis %s for regression: %w", xaxis, err))
				return
			}
			y, err := utils.GetFloat(trial.Data[yaxis])
			if err != nil {
				app.internalServerError(w, r, fmt.Errorf("unable to use y axis %s for regression: %w", yaxis, err))
				return
			}
			xs = append(xs, x)
			ys = append(ys, y)
		}
		regression, err := stats.LinearLeastSquares(xs, ys)
		if err != nil {
			app.internalServerError(w, r, fmt.Errorf("unable to compute least squares regression"))
			return
		}
		// regression successful, return result
		if err := app.jsonResponse(w, http.StatusOK, regression); err != nil {
			app.internalServerError(w, r, err)
		}
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, trials); err != nil {
		app.internalServerError(w, r, err)
	}
}
