package main

import (
	"log"
	"net/http"
)

type App struct {
}

// Start an HTTP server to respond to requests to upload data
// and get useful pieces of data for the UI
func (app App) Start() error {
	const serverAddr string = "0.0.0.0:3001"
	log.Printf("Starting HTTP server: %s\n", serverAddr)

	// http handlers
	http.HandleFunc("POST /api/upload", app.uploadData)

	// open the port
	return http.ListenAndServe(serverAddr, nil)
}

func (app App) uploadData(w http.ResponseWriter, r *http.Request) {
	log.Println("Upload Data invoked.")

	// file, _, err := r.FormFile("file")
	// if err != nil {
	// 	http.Error(w, "Unable to read file", http.StatusBadRequest)
	// 	log.Printf("Error reading file: %v", err)
	// 	return
	// }
	// defer file.Close()

	// reader := csv.NewReader(file)
	// rows, err := reader.ReadAll()
	// if err != nil {
	// 	http.Error(w, "Unable to parse CSV", http.StatusBadRequest)
	// 	log.Printf("Error parsing CSV: %v", err)
	// 	return
	// }
}
