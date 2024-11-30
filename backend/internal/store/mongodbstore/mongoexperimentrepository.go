package mongodbstore

import (
	"context"
	"log"

	"github.com/fransk/truthiness/internal/store"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoExperimentRepository struct {
	collection *mongo.Collection
}

func (repo *MongoExperimentRepository) Create(ctx context.Context, experiment *store.Experiment) error {
	result, err := repo.collection.InsertOne(ctx, *experiment)
	if err != nil {
		return err
	}

	log.Printf("Inserted new experiment with ID %v\n", result.InsertedID)

	return nil
}

func (repo *MongoExperimentRepository) GetAll(ctx context.Context) ([]store.Experiment, error) {
	cursor, err := repo.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var experiments []store.Experiment
	if err = cursor.All(ctx, &experiments); err != nil {
		log.Fatal(err)
	}

	return experiments, nil
}
