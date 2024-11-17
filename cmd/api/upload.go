package main

import (
	"encoding/csv"
	"log"
	"net/http"
)

func (app *application) uploadDataHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Upload Data invoked.")

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusBadRequest)
		log.Printf("Error reading file: %v", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		http.Error(w, "Unable to parse CSV", http.StatusBadRequest)
		log.Printf("Error parsing CSV: %v", err)
		return
	}

	for _, record := range rows {
		log.Println(record[0])
	}
}
