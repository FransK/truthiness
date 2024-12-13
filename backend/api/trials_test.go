package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fransk/truthiness/internal/store"
	"github.com/fransk/truthiness/internal/store/inmemorystore"
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
}
