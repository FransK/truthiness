package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func New(addr string, maxOpenConns uint64, maxIdleTime time.Duration) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	optns := options.Client()
	optns.ApplyURI(addr)
	optns.SetMaxConnecting(maxOpenConns)
	optns.SetMaxConnIdleTime(maxIdleTime)

	client, err := mongo.Connect(ctx, optns)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	log.Printf("Connected to mongdb server at %v", addr)

	return client, nil
}
