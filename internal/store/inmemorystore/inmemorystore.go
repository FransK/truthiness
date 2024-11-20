package inmemorystore

import "github.com/fransk/truthiness/internal/store"

// Store the Trials and Users in memory

func NewInMemoryStore() store.Storage {
	return InMemoryStorage{}
}

type InMemoryStorage struct {
}

func (store InMemoryStorage) Experiments() store.ExperimentRepository {
	return nil
}

func (store InMemoryStorage) Trials(trialname string) store.TrialRepository {
	return nil
}

func (store InMemoryStorage) Users() store.UserRepository {
	return nil
}
