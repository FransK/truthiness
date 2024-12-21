package mongodbstore

import (
	"context"
	"log"

	"github.com/fransk/truthiness/internal/store"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	collection *mongo.Collection
}

func (repo *MongoUserRepository) GetAll(ctx context.Context) ([]store.User, error) {
	cursor, err := repo.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := cursor.Close(ctx); closeErr != nil {
			log.Printf("warning: failed to close cursor: %v", closeErr)
		}
	}()

	users := make([]store.User, 0)
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (repo *MongoUserRepository) GetById(ctx context.Context, id int64) (*store.User, error) {
	var result store.User
	if err := repo.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (repo *MongoUserRepository) GetByUsername(ctx context.Context, username string) (*store.User, error) {
	var result store.User
	if err := repo.collection.FindOne(ctx, bson.M{"username": username}).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (repo *MongoUserRepository) Create(ctx context.Context, user *store.User) error {
	result, err := repo.collection.InsertOne(ctx, *user)
	if err != nil {
		return err
	}

	log.Printf("Inserted new user with ID %v\n", result.InsertedID)

	return nil
}
