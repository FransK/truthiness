package inmemorystore

import "github.com/fransk/truthiness/internal/store"

type InMemoryStoreFactory struct{}

func (factory *InMemoryStoreFactory) NewStore() store.Storage {
	return &InMemoryStorage{}
}
