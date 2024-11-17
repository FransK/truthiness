package main

import "log"

func main() {
	app := App{}
	log.Fatal(app.Start())
}
