package main

import "log"

func main() {
	app := application{}
	log.Fatal(app.Start())
}
