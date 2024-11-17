package main

import (
	"log"
	"net/http"
)

type App struct {
}

func (app App) Start() error {
	const serverAddr string = "0.0.0.0:3001"
	log.Printf("Starting HTTP server: %s\n", serverAddr)

	return http.ListenAndServe(serverAddr, nil)
}
