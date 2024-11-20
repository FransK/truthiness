package mongodbstore

import (
	"github.com/fransk/truthiness/internal/store"
	"go.mongodb.org/mongo-driver/mongo"
)

func New(mongodb *mongo.Database) store.Storage {
	return &MongoDbStore{
		db: mongodb,
	}
}

type MongoDbStore struct {
	db *mongo.Database
}

func (store *MongoDbStore) Experiments() store.ExperimentRepository {
	return MongoExperimentRepository{
		collection: store.db.Collection("experiments"),
	}
}

func (store *MongoDbStore) Trials(experiment string) store.TrialRepository {
	return MongoTrialRepository{
		collection: store.db.Collection(experiment),
	}
}

func (store *MongoDbStore) Users() store.UserRepository {
	return MongoUserRepository{
		collection: store.db.Collection("users"),
	}
}
