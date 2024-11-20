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

func (repo MongoExperimentRepository) Create(ctx context.Context) error {
	return nil
}

func (repo MongoExperimentRepository) GetAll() ([]store.Experiment, error) {
	cursor, err := repo.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var experiments []store.Experiment
	if err = cursor.All(context.Background(), &experiments); err != nil {
		log.Fatal(err)
	}

	return experiments, nil
}
