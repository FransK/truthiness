package api

import (
	"encoding/csv"
	"log"
	"net/http"
	"slices"

	"github.com/fransk/truthiness/internal/store"
)

func (app *Application) uploadDataHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("upload Data invoked.")

	experimentname := r.FormValue("experiment")
	if experimentname == "" {
		http.Error(w, "no experiment name provided", http.StatusBadRequest)
		return
	}
	experimentdate := r.FormValue("date")
	if experimentdate == "" {
		http.Error(w, "no experiment date provided", http.StatusBadRequest)
		return
	}
	experimentlocation := r.FormValue("location")
	if experimentlocation == "" {
		http.Error(w, "no experiment location provided", http.StatusBadRequest)
		return
	}
	log.Printf("experiment: %v - %v at %v", experimentname, experimentdate, experimentlocation)

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "unable to read file", http.StatusBadRequest)
		log.Printf("error reading file: %v", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		http.Error(w, "unable to parse CSV", http.StatusBadRequest)
		log.Printf("error parsing CSV: %v", err)
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

		trials = append(trials, store.Trial{
			Data: data,
		})
	}

	experiment := store.Experiment{
		Name:     experimentname,
		Date:     experimentdate,
		Location: experimentlocation,
		Records:  slices.Clone(keys),
	}

	fn := func() (interface{}, error) {
		if err = app.Store.Experiments().Create(r.Context(), &experiment); err != nil {
			log.Println(err.Error())
			return nil, err
		}

		if err = app.Store.Trials(experimentname).CreateMany(r.Context(), trials); err != nil {
			log.Println(err.Error())
			return nil, err
		}

		return "success", nil
	}
	app.Store.WithTransaction(r.Context(), fn)

	app.jsonResponse(w, http.StatusOK, "success")
}
