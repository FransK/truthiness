package inmemorystore

import "github.com/fransk/truthiness/internal/store"

/* InMemoryStorage implements the storage interface to be used by our truthiness api*/
type InMemoryStorage struct {
}

func New() store.Storage {
	return InMemoryStorage{}
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
