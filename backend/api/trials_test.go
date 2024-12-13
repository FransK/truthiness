package api

import (
	"context"
	"encoding/json"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fransk/truthiness/internal/stats"
	"github.com/fransk/truthiness/internal/store"
	"github.com/fransk/truthiness/internal/store/inmemorystore"
	"github.com/google/go-cmp/cmp"
)

func TestGetTrialsHandler(t *testing.T) {
	storage := inmemorystore.New()
	cfg := &Config{
		Addr: "testhost:1111",
	}
	experiments := []store.Experiment{
		{
			Name:     "testexperiment",
			Date:     "March 22, 2024",
			Location: "SFU",
			Records: map[string]int{
				"Age":        2,
				"Difference": 1,
			},
		},
	}
	trials := []store.Trial{
		{
			Data: map[string]any{
				"Age":        35,
				"Difference": 1.4,
			},
		},
		{
			Data: map[string]any{
				"Age":        21,
				"Difference": -0.4,
			},
		},
		{
			Data: map[string]any{
				"Age":        82,
				"Difference": -1.2,
			},
		},
		{
			Data: map[string]any{
				"Age":        72,
				"Difference": -1.4,
			},
		},
	}

	app := &Application{
		Config: *cfg,
		Store:  storage,
	}

	for _, experiment := range experiments {
		storage.Experiments().Create(context.TODO(), &experiment)
	}
	storage.Trials("testexperiment").CreateMany(context.TODO(), trials)

	tests := []struct {
		name           string
		experimentname string
		want           int
	}{
		{
			name:           "Experiment name must be provided",
			experimentname: "",
			want:           http.StatusBadRequest,
		},
		{
			name:           "Get trials for valid experiment",
			experimentname: "testexperiment",
			want:           http.StatusOK,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/v1/experiments/experimentname/trials", nil)
			req.SetPathValue("experimentname", test.experimentname)
			rr := httptest.NewRecorder()
			app.getTrialsHandler(rr, req)

			if rr.Code != test.want {
				t.Errorf("got %v; want %v", rr.Code, test.want)
			}
		})
	}

	t.Run("linear regression call", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/v1/experiments/experimentname/trials?model=linear&x_axis=Age&y_axis=Difference", nil)
		req.SetPathValue("experimentname", "testexperiment")
		rr := httptest.NewRecorder()

		app.getTrialsHandler(rr, req)

		want := http.StatusOK
		if rr.Code != want {
			t.Errorf("got %v; want %v", rr.Code, want)
		}

		type JSONResponse struct {
			Data stats.LinearLeastSquaresData `json:"data"`
		}

		var data JSONResponse
		decoder := json.NewDecoder(rr.Body)
		decoder.Decode(&data)
		lsd := data.Data

		const tolerance = .0001
		cmpr := cmp.Comparer(func(x, y float64) bool {
			diff := math.Abs(x - y)
			return diff < tolerance
		})

		mwant := -0.0293
		bwant := 1.1365
		r2want := 0.4474
		if !cmp.Equal(lsd.M, mwant, cmpr) {
			t.Errorf("slope m got %v; want %v", lsd.M, mwant)
		}
		if !cmp.Equal(lsd.B, bwant, cmpr) {
			t.Errorf("intercept b got %v; want %v", lsd.B, bwant)
		}
		if !cmp.Equal(lsd.R2, r2want, cmpr) {
			t.Errorf("coefficient R2 got %v; want %v", lsd.R2, r2want)
		}
	})
}
