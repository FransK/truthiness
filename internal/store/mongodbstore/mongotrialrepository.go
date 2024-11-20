package mongodbstore

import (
	"context"

	"github.com/fransk/truthiness/internal/store"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoTrialRepository struct {
	collection *mongo.Collection
}

func (repo MongoTrialRepository) GetAll(ctx context.Context) ([]store.Trial, error) {
	return nil, nil
}

func (repo MongoTrialRepository) Create(ctx context.Context) error {
	return nil
}
