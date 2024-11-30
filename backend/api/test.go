package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/fransk/truthiness/internal/store"
)

func (app *Application) testHandler(w http.ResponseWriter, r *http.Request) {
	// Testing insert
	user := store.User{
		ID: 1,
	}
	if err := app.Store.Users().Create(r.Context(), &user); err != nil {
		log.Println(err.Error())
	}

	_, err := app.Store.Users().GetById(r.Context(), 1)
	if err != nil {
		log.Println(err.Error())
	}

	// Testing insert
	experiment := store.Experiment{
		Name:     "scienceworld",
		Date:     "November 10 1993",
		Location: "sfu",
		Records:  []string{"pid", "answer", "answer2", "answer3"},
	}
	if err = app.Store.Experiments().Create(r.Context(), &experiment); err != nil {
		log.Println(err.Error())
	}

	_, err = app.Store.Experiments().GetAll(r.Context())
	if err != nil {
		log.Println(err.Error())
	}

	// Testing insert
	trials := make([]store.Trial, 10)
	for i := 0; i < 10; i++ {
		trials[i] = store.Trial{
			Data: map[string]string{
				"pid":     strconv.Itoa(i),
				"answer":  "test",
				"answer2": "false",
				"answer3": strconv.Itoa(i),
			},
		}
	}
	if err = app.Store.Trials("scienceworld").CreateMany(r.Context(), trials); err != nil {
		log.Println(err.Error())
	}
}
