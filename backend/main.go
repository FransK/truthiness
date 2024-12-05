package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/fransk/truthiness/api"
	"github.com/fransk/truthiness/internal/db"
	"github.com/fransk/truthiness/internal/env"
	"github.com/fransk/truthiness/internal/store"
	"github.com/fransk/truthiness/internal/store/inmemorystore"
	"github.com/fransk/truthiness/internal/store/mongodbstore"
)

type DbConfig struct {
	Addr         string
	MaxOpenConns int
	MaxIdleTime  time.Duration
}

func run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	var store store.Storage

	switch env.GetString("STORAGE_TYPE", "IN_MEMORY") {
	case "MONGODB":
		dbConfig := &DbConfig{
			Addr:         env.GetString("DB_ADDR", "mongodb://localhost:27017/"),
			MaxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			MaxIdleTime:  env.GetDuration("DB_MAX_IDLE_TIME", 15*time.Minute),
		}

		mydb, err := db.New(
			dbConfig.Addr,
			dbConfig.MaxOpenConns,
			dbConfig.MaxIdleTime,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error connecting to mongo database: %s\n", err)
		}
		defer db.Close(mydb)

		store = mongodbstore.New(mydb)
	default:
		store = inmemorystore.New()
	}

	log.Printf("storage type: %T", store)

	cfg := &api.Config{
		Addr: env.GetString("ADDR", ":8080"),
	}
	srv := api.NewServer(cfg, &store)
	httpServer := &http.Server{
		Addr:         cfg.Addr,
		Handler:      srv,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	go func() {
		log.Printf("starting HTTP server on: %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()
	return nil
}

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
