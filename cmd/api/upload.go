package main

import (
	"encoding/csv"
	"log"
	"net/http"

	"github.com/fransk/truthiness/internal/store"
)

func (app *application) uploadDataHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Upload Data invoked.")

	title := r.FormValue("title")
	if title == "" {
		http.Error(w, "No title provided", http.StatusBadRequest)
		return
	}
	log.Println("Title: ", title)

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

	// Get the column names. This CURRENTLY assumes data is formatted with column names
	// in first row and data in subsequent rows
	// TODO: Add ability for USER to determine col names and row where data starts
	// with some sort of previewer
	keys := make([]string, 0, len(rows[0]))
	keys = append(keys, rows[0]...)

	trials := make([]store.Trial, 0, len(rows))
	for _, row := range rows[1:] {
		data := make(map[string]string)
		i := 0
		for _, key := range keys {
			if i >= len(row) {
				data[key] = ""
			} else {
				data[key] = row[i]
				i++
			}
		}

		log.Println(data)

		trials = append(trials, store.Trial{
			Data: data,
		})
	}

	if err = app.store.Trials(title).CreateMany(r.Context(), trials); err != nil {
		log.Println(err.Error())
	}
}
