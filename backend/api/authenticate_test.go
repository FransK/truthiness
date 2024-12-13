package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fransk/truthiness/internal/auth"
	"github.com/fransk/truthiness/internal/store"
	"github.com/fransk/truthiness/internal/store/inmemorystore"
)

func TestAuthHandler(t *testing.T) {
	storage := inmemorystore.New()
	cfg := &Config{
		Addr: "testhost:1111",
	}
	users := []store.User{
		{
			ID:          1,
			Username:    "testuser",
			Password:    "testpassword",
			Role:        "user",
			Permissions: []string{http.MethodGet},
		},
		{
			ID:          2,
			Username:    "testadmin",
			Password:    "testadminpass",
			Role:        "admin",
			Permissions: []string{http.MethodGet, http.MethodPost},
		},
	}
	for _, user := range users {
		storage.Users().Create(context.TODO(), &user)
	}

	app := &Application{
		Config: *cfg,
		Store:  storage,
	}

	t.Run("method must be POST", func(t *testing.T) {
		methods := []string{
			http.MethodConnect,
			http.MethodDelete,
			http.MethodGet,
			http.MethodHead,
			http.MethodOptions,
			http.MethodPatch,
			http.MethodPut,
			http.MethodTrace}
		for _, m := range methods {
			t.Run(fmt.Sprintf("method: %s", m), func(t *testing.T) {
				req := httptest.NewRequest(m, "/authenticate", nil)
				rr := httptest.NewRecorder()
				app.authHandler(rr, req)

				want := http.StatusMethodNotAllowed
				if rr.Code != want {
					t.Errorf("got %v; want %v", rr.Code, want)
				}
			})
		}

		t.Run(fmt.Sprintf("method: %s", http.MethodPost), func(t *testing.T) {
			user, err := storage.Users().GetByUsername(context.TODO(), "testadmin")
			if err != nil {
				t.Errorf("failed to get testuser")
			}
			creds := auth.Credentials{
				Username: user.Username,
				Password: user.Password,
			}
			body, _ := json.Marshal(creds)
			req := httptest.NewRequest(http.MethodPost, "/authenticate", bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()
			app.authHandler(rr, req)

			want := http.StatusOK
			if rr.Code != want {
				t.Errorf("got %v; want %v", rr.Code, want)
			}
		})
	})

	t.Run("credentials validation", func(t *testing.T) {
		user, err := storage.Users().GetByUsername(context.TODO(), "testuser")
		if err != nil {
			t.Errorf("failed to get testuser")
		}

		creds := auth.Credentials{
			Username: user.Username,
			Password: user.Password,
		}
		body, _ := json.Marshal(creds)

		req := httptest.NewRequest(http.MethodPost, "/authenticate", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		app.authHandler(rr, req)

		want := http.StatusOK
		if rr.Code != want {
			t.Errorf("got %v; want %v", rr.Code, want)
		}

		auth.ValidateTokenAndGetRole(req)
	})
}
