package mongodbstore

import (
	"context"

	"github.com/fransk/truthiness/internal/store"
)

type MongoDbStorage struct {
}

func (store *MongoDbStorage) Experiments() store.ExperimentRepository {
	return nil
}

func (store *MongoDbStorage) Trials(experiment string) store.TrialRepository {
	return nil
}

func (store *MongoDbStorage) Users() store.UserRepository {
	return nil
}

type MongoTrialRepository struct {
}

func (repo *MongoTrialRepository) GetAll(ctx context.Context) ([]store.Trial, error) {
	return nil, nil
}

func (repo *MongoTrialRepository) Insert(ctx context.Context) error {
	return nil
}
