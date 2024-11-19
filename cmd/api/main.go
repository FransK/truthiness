package main

import (
	"log"

	"github.com/fransk/truthiness/internal/env"
	"github.com/fransk/truthiness/internal/store"
	"github.com/fransk/truthiness/internal/store/inmemorystore"
	"github.com/fransk/truthiness/internal/store/mongodbstore"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	var store store.Storage
	switch env.GetString("STORAGE_TYPE", "IN_MEMORY") {
	case "MONGODB":
		factory := mongodbstore.MongoDbStoreFactory{}
		store = factory.NewStore()
	case "IN_MEMORY":
		factory := inmemorystore.InMemoryStoreFactory{}
		store = factory.NewStore()
	default:
		factory := inmemorystore.InMemoryStoreFactory{}
		store = factory.NewStore()
	}

	log.Printf("Storage type: %T", store)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
