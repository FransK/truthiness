package mongodbstore

import (
	"context"

	"github.com/fransk/truthiness/internal/store"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	collection *mongo.Collection
}

func (repo MongoUserRepository) GetById(ctx context.Context) (*store.User, error) {
	return nil, nil
}

func (repo MongoUserRepository) Create(ctx context.Context) error {
	return nil
}
