package main

import (
	"log"
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))

	// Testing connection
	experiments, err := app.store.Experiments().GetAll(r.Context())
	if err != nil {
		log.Println(err.Error())
	}
	log.Printf("Experiments: %v", experiments)

	/*
		// Testing insert
		trial := store.Trial{
			Data: map[string]string{
				"pid":    "2",
				"answer": "test",
			},
		}
		if err = app.store.Trials("scienceworld").CreateMany(r.Context(), []store.Trial{trial}); err != nil {
			log.Println(err.Error())
		}

		// Testing insert
		experiment := store.Experiment{
			Name:     "labtest1",
			Date:     "November 10 1993",
			Location: "sfu",
		}
		if err = app.store.Experiments().Create(r.Context(), experiment); err != nil {
			log.Println(err.Error())
		}

		experiments, err = app.store.Experiments().GetAll(r.Context())
		if err != nil {
			log.Println(err.Error())
		}
		log.Printf("Experiments: %v", experiments)
	*/
}
