package main

import (
	"log"
	"time"

	"github.com/fransk/truthiness/internal/db"
	"github.com/fransk/truthiness/internal/env"
	"github.com/fransk/truthiness/internal/store"
	"github.com/fransk/truthiness/internal/store/inmemorystore"
	"github.com/fransk/truthiness/internal/store/mongodbstore"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "mongodb://localhost:27017"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleTime:  env.GetDuration("DB_MAX_IDLE_TIME", 15*time.Minute),
		},
	}

	var store store.Storage

	switch env.GetString("STORAGE_TYPE", "MONGODB") {
	case "MONGODB":
		db, err := db.New(
			cfg.db.addr,
			uint64(cfg.db.maxOpenConns),
			cfg.db.maxIdleTime,
		)
		if err != nil {
			log.Panic(err)
		}

		store = mongodbstore.New(db)
	default:
		store = inmemorystore.New()
	}

	log.Printf("Storage type: %T", store)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
