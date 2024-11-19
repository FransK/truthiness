package mongodbstore

import "github.com/fransk/truthiness/internal/store"

type MongoDbStorage struct{}

func (store *MongoDbStorage) Trials() store.TrialRepository {
	return nil
}

func (store *MongoDbStorage) Users() store.UserRepository {
	return nil
}
