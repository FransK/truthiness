package inmemorystore

import (
	"context"
	"slices"

	"github.com/fransk/truthiness/internal/store"
)

type InMemoryExperimentRepository struct {
	experiments []store.Experiment
}

func (repo *InMemoryExperimentRepository) Create(ctx context.Context, experiment *store.Experiment) error {
	newExperiment := store.Experiment{
		Name:     experiment.Name,
		Date:     experiment.Date,
		Location: experiment.Location,
		Records:  slices.Clone(experiment.Records),
	}
	repo.experiments = append(repo.experiments, newExperiment)
	return nil
}

func (repo *InMemoryExperimentRepository) GetAll(ctx context.Context) ([]store.Experiment, error) {
	return slices.Clone(repo.experiments), nil
}
