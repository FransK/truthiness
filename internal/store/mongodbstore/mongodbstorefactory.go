package mongodbstore

import "github.com/fransk/truthiness/internal/store"

type MongoDbStoreFactory struct{}

func (factory *MongoDbStoreFactory) NewStore() store.Storage {
	return &MongoDbStorage{}
}
