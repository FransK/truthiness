package api

import (
	"encoding/csv"
	"log"
	"net/http"
	"strconv"

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

	// Create a records map to store the dominant type for each column
	records := make(map[string]int)

	trials := make([]store.Trial, 0, len(rows))
	typeCounts := make(map[string]map[int]int)       // Track counts of each type per column
	uniqueValues := make(map[string]map[string]bool) // Track unique values per columns

	for _, key := range keys {
		typeCounts[key] = map[int]int{
			store.DataTypeUnknown:     0,
			store.DataTypeNumeric:     0,
			store.DataTypeCategorical: 0,
		}
		uniqueValues[key] = make(map[string]bool)
	}

	for _, trial := range rows[1:] { // Skip header
		data := make(map[string]any)
		for i, key := range keys {
			if i >= len(trial) || trial[i] == "" {
				// Missing value
				data[key] = nil
				typeCounts[key][store.DataTypeUnknown]++
				continue
			}

			value := trial[i]
			uniqueValues[key][value] = true // Track all unique values

			numeric, err := strconv.ParseFloat(value, 64)
			if err != nil {
				// Value is not numeric
				data[key] = value
				typeCounts[key][store.DataTypeCategorical]++
			} else {
				// Value is numeric
				data[key] = numeric
				typeCounts[key][store.DataTypeNumeric]++
			}
		}

		trials = append(trials, store.Trial{
			Data: data,
		})
	}

	// Determine the final type for each columns
	for key, counts := range typeCounts {
		// If the columns has missing values only, classify as unknown
		if counts[store.DataTypeNumeric] == 0 && counts[store.DataTypeCategorical] == 0 {
			records[key] = store.DataTypeUnknown
			log.Printf("records[%s] is set to DataTypeUnknown", key)
			continue
		}

		// Check for small unique value sets
		if len(uniqueValues[key]) <= 5 {
			log.Printf("records[%s] is set to DataTypeCategorical with %d unique values", key, len(uniqueValues[key]))
			records[key] = store.DataTypeCategorical
		} else {
			// Otherwise, classify based on majority type
			if counts[store.DataTypeNumeric] > counts[store.DataTypeCategorical] {
				log.Printf("records[%s] is set to DataTypeNumeric with %d numerics and %d categorical", key, counts[store.DataTypeNumeric], counts[store.DataTypeCategorical])
				records[key] = store.DataTypeNumeric
			} else {
				log.Printf("records[%s] is set to DataTypeCategorical with %d numerics and %d categorical", key, counts[store.DataTypeNumeric], counts[store.DataTypeCategorical])
				records[key] = store.DataTypeCategorical
			}
		}
	}

	experiment := store.Experiment{
		Name:     experimentname,
		Date:     experimentdate,
		Location: experimentlocation,
		Records:  records,
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
	_, err = app.Store.WithTransaction(r.Context(), fn)
	if err != nil {
		app.internalServerError(w, r, err)
	} else {
		app.jsonResponse(w, http.StatusOK, "success")
	}
}
