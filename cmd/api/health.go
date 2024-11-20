package main

import (
	"log"
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))

	// Testing connection
	experiments, err := app.store.Experiments().GetAll()
	if err != nil {
		log.Println(err.Error())
	}
	log.Printf("Experiments: %v", experiments)
}
