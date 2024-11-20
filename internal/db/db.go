package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func New(addr string, maxOpenConns uint64, maxIdleTime time.Duration) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	clientOpts := options.Client()
	clientOpts.ApplyURI(addr)
	clientOpts.SetMaxConnecting(maxOpenConns)
	clientOpts.SetMaxConnIdleTime(maxIdleTime)

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	log.Printf("Connected to mongdb server at %v", addr)

	database := client.Database("truthiness")

	return database, nil
}

func Close(db *mongo.Database) error {
	if err := db.Client().Disconnect(context.TODO()); err != nil {
		return err
	}

	log.Println("MongoDB client disconnected.")
	return nil
}
