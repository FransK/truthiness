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
