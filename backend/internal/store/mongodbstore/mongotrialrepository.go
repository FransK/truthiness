package mongodbstore

import (
	"context"
	"log"

	"github.com/fransk/truthiness/internal/store"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoTrialRepository struct {
	collection *mongo.Collection
}

func (repo *MongoTrialRepository) Create(ctx context.Context, trial *store.Trial) error {
	result, err := repo.collection.InsertOne(ctx, *trial)
	if err != nil {
		return err
	}

	log.Printf("Inserted new trial with ID %v\n", result.InsertedID)

	return nil
}

func (repo *MongoTrialRepository) CreateMany(ctx context.Context, trials []store.Trial) error {
	anytrials := make([]interface{}, len(trials))
	for i, val := range trials {
		anytrials[i] = val
	}

	result, err := repo.collection.InsertMany(ctx, anytrials)
	if err != nil {
		return err
	}

	log.Printf("successfully inserted %v new trials\n", len(result.InsertedIDs))

	return nil
}

func (repo *MongoTrialRepository) GetAll(ctx context.Context) ([]store.Trial, error) {
	cursor, err := repo.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var trials []store.Trial
	if err = cursor.All(ctx, &trials); err != nil {
		log.Fatal(err)
	}

	return trials, nil
}
