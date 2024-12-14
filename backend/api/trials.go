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

	var xs, ys []float64
	for _, trial := range trials {
		v, ok := trial.Data[xaxis]
		if !ok {
			log.Printf("trial does not contain value for x axis: %s", xaxis)
			continue
		}
		x, err := utils.GetFloat(v)
		if err != nil {
			log.Printf("unable to use x axis from trial data for regression: %v", err)
			continue
		}

		v, ok = trial.Data[yaxis]
		if !ok {
			log.Printf("trial does not contain value for y axis: %s", yaxis)
			continue
		}
		y, err := utils.GetFloat(v)
		if err != nil {
			log.Printf("unable to use y axis from trial data for regression: %v", err)
			continue
		}

		xs = append(xs, x)
		ys = append(ys, y)
	}
	regression, err := stats.LinearLeastSquares(xs, ys)
	if err != nil {
		app.internalServerError(w, r, fmt.Errorf("unable to compute least squares regression"))
		return
	}
	// regression successful, append results
	for _, trial := range trials {
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
