package inmemorystore

import "github.com/fransk/truthiness/internal/store"

// Store the Trials and Users in memory
type InMemoryStorage struct {
}

func (store *InMemoryStorage) Trials() store.TrialRepository {
	return nil
}

func (store *InMemoryStorage) Users() store.UserRepository {
	return nil
}
