package mongodbstore

import (
	"context"
	"fmt"
	"log"

	"github.com/fransk/truthiness/internal/store"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (repo *MongoTrialRepository) Get(ctx context.Context, keys []string) ([]store.Trial, error) {
	if len(keys) == 0 {
		return nil, fmt.Errorf("keys cannot be empty for projection")
	}

	queryFields := bson.M{}
	queryFields["_id"] = 0
	for _, key := range keys {
		prefixedKey := "data." + key
		queryFields[prefixedKey] = 1
		log.Printf("query field added: %s", prefixedKey)
	}
	log.Printf("final projection map: %+v", queryFields)

	opts := options.Find().SetProjection(queryFields)

	cursor, err := repo.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to execute find query with projection %v: %w", queryFields, err)
	}
	defer func() {
		if closeErr := cursor.Close(ctx); closeErr != nil {
			log.Printf("warning: failed to close cursor: %v", closeErr)
		}
	}()

	trials := make([]store.Trial, 0)
	if err = cursor.All(ctx, &trials); err != nil {
		return nil, fmt.Errorf("failed to decode cursor results into store.Trial slice: %w", err)
	}

	log.Printf("successfully retrieved %d trials", len(trials))
	return trials, nil
}

func (repo *MongoTrialRepository) GetAll(ctx context.Context) ([]store.Trial, error) {
	cursor, err := repo.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	trials := make([]store.Trial, 0)
	if err = cursor.All(ctx, &trials); err != nil {
		return nil, err
	}

	return trials, nil
}
