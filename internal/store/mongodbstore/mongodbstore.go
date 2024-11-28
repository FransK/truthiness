package mongodbstore

import (
	"context"

	"github.com/fransk/truthiness/internal/store"
	"go.mongodb.org/mongo-driver/mongo"
)

/* MongoDbStore implements the storage interface to be used by our truthiness api*/
type MongoDbStore struct {
	db *mongo.Database
}

/* New creates a container for a MongoDB database which can be used for queries */
func New(mongodb *mongo.Database) store.Storage {
	return &MongoDbStore{
		db: mongodb,
	}
}

/* Experiments returns the truthiness collection for the list of Experiments */
func (store *MongoDbStore) Experiments() store.ExperimentRepository {
	return &MongoExperimentRepository{
		collection: store.db.Collection("experiments"),
	}
}

/* Trials returns the truthiness collection of trials for a specific experiment */
func (store *MongoDbStore) Trials(experiment string) store.TrialRepository {
	return &MongoTrialRepository{
		collection: store.db.Collection(experiment),
	}
}

/* Users returns the truthiness collection for the list of users */
func (store *MongoDbStore) Users() store.UserRepository {
	return &MongoUserRepository{
		collection: store.db.Collection("users"),
	}
}

func (store *MongoDbStore) WithTransaction(ctx context.Context, fn func() (interface{}, error)) (interface{}, error) {
	session, err := store.db.Client().StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(ctx)

	if err = session.StartTransaction(); err != nil {
		return nil, err
	}

	result, err := fn()
	if err != nil {
		session.AbortTransaction(ctx)
		return nil, err
	}

	if err = session.CommitTransaction(ctx); err != nil {
		session.AbortTransaction(ctx)
		return nil, err
	}

	return result, nil
}
