package main

import (
	"log"
	"time"

	"github.com/fransk/truthiness/api"
	"github.com/fransk/truthiness/internal/db"
	"github.com/fransk/truthiness/internal/env"
	"github.com/fransk/truthiness/internal/store"
	"github.com/fransk/truthiness/internal/store/inmemorystore"
	"github.com/fransk/truthiness/internal/store/mongodbstore"
)

func main() {
	cfg := api.Config{
		Addr: env.GetString("ADDR", ":8080"),
		DB: api.DbConfig{
			Addr:         env.GetString("DB_ADDR", "mongodb://localhost:27017/"),
			MaxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			MaxIdleTime:  env.GetDuration("DB_MAX_IDLE_TIME", 15*time.Minute),
		},
	}

	var store store.Storage

	switch env.GetString("STORAGE_TYPE", "IN_MEMORY") {
	case "MONGODB":
		mydb, err := db.New(
			cfg.DB.Addr,
			cfg.DB.MaxOpenConns,
			cfg.DB.MaxIdleTime,
		)
		if err != nil {
			log.Panic(err)
		}
		defer db.Close(mydb)

		store = mongodbstore.New(mydb)
	default:
		store = inmemorystore.New()
	}

	log.Printf("Storage type: %T", store)

	app := &api.Application{
		Config: cfg,
		Store:  store,
	}

	log.Fatal(app.RunNew())
}
