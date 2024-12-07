package inmemorystore

import (
	"context"

	"github.com/fransk/truthiness/internal/store"
)

/* InMemoryStorage implements the storage interface to be used by our truthiness api*/
type InMemoryStorage struct {
	experiments *InMemoryExperimentRepository
	trials      map[string]*InMemoryTrialRepository
	users       *InMemoryUserRepository
}

func New() *InMemoryStorage {
	return &InMemoryStorage{
		experiments: &InMemoryExperimentRepository{
			experiments: make([]store.Experiment, 0),
		},
		trials: make(map[string]*InMemoryTrialRepository),
		users: &InMemoryUserRepository{
			users: make(map[int64]store.User),
		},
	}
}

func (storage *InMemoryStorage) Experiments() store.ExperimentRepository {
	return storage.experiments
}

func (storage *InMemoryStorage) Trials(trialname string) store.TrialRepository {
	if v, ok := storage.trials[trialname]; ok {
		return v
	}
	storage.trials[trialname] = &InMemoryTrialRepository{
		trials: make([]store.Trial, 0),
	}
	return storage.trials[trialname]
}

func (storage *InMemoryStorage) Users() store.UserRepository {
	return storage.users
}

func (storage *InMemoryStorage) WithTransaction(ctx context.Context, fn func() (interface{}, error)) (interface{}, error) {
	// TODO: Can add transaction logic later for InMemoryStorage, testing for failures along way of executing fn
	return fn()
}
