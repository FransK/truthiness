package inmemorystore

import (
	"context"
	"maps"
	"slices"

	"github.com/fransk/truthiness/internal/store"
)

type InMemoryTrialRepository struct {
	trials []store.Trial
}

func (repo *InMemoryTrialRepository) Create(ctx context.Context, trial *store.Trial) error {
	newTrial := store.Trial{
		Data: maps.Clone(trial.Data),
	}
	repo.trials = append(repo.trials, newTrial)
	return nil
}

func (repo *InMemoryTrialRepository) CreateMany(ctx context.Context, trials []store.Trial) error {
	for _, v := range trials {
		newTrial := store.Trial{
			Data: maps.Clone(v.Data),
		}
		repo.trials = append(repo.trials, newTrial)
	}
	return nil
}

func (repo *InMemoryTrialRepository) Get(ctx context.Context, keys []string) ([]store.Trial, error) {
	response := make([]store.Trial, 0)
	for _, trial := range repo.trials {
		newTrial := store.Trial{
			Data: make(map[string]any),
		}
		for _, k := range keys {
			v, ok := trial.Data[k]
			if ok {
				newTrial.Data[k] = v
			}
		}
		response = append(response, newTrial)
	}
	return response, nil
}

func (repo *InMemoryTrialRepository) GetAll(ctx context.Context) ([]store.Trial, error) {
	return slices.Clone(repo.trials), nil
}
