package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/fransk/truthiness/internal/auth"
)

func (app *Application) authHandler(w http.ResponseWriter, r *http.Request) {
	// make sure post
	if r.Method != "POST" {
		app.methodNotAllowedResponse(w, r, errors.New(""))
		return
	}

	// read credentials
	var creds auth.Credentials
	err := readJSON(w, r, &creds)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// get user from data store
	user, err := app.Store.Users().GetByUsername(r.Context(), creds.Username)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	// check user credentials
	if user.Password != creds.Password {
		app.unauthorized(w, r, errors.New("unauthorized"))
		return
	}

	// create jwt token with role found in store
	tokenString, err := auth.CreateToken(user.Username, user.Role)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	// return the token
	w.Header().Set("Content-Type", "application/jwt")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, tokenString)
}
