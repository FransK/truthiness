package mongodbstore

import (
	"context"

	"github.com/fransk/truthiness/internal/store"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewMongoDbStore(client *mongo.Client) store.Storage {
	return MongoDbStorage{}
}

type MongoDbStorage struct {
}

func (store MongoDbStorage) Experiments() store.ExperimentRepository {
	return MongoExperimentRepository{}
}

func (store MongoDbStorage) Trials(experiment string) store.TrialRepository {
	return MongoTrialRepository{}
}

func (store MongoDbStorage) Users() store.UserRepository {
	return MongoUserRepository{}
}

type MongoExperimentRepository struct {
}

func (repo MongoExperimentRepository) Create(ctx context.Context) error {
	return nil
}

func (repo MongoExperimentRepository) GetExperiments() ([]store.Experiment, error) {
	return nil, nil
}

type MongoUserRepository struct {
}

func (repo MongoUserRepository) GetById(ctx context.Context) (*store.User, error) {
	return nil, nil
}

func (repo MongoUserRepository) Create(ctx context.Context) error {
	return nil
}

type MongoTrialRepository struct {
}

func (repo MongoTrialRepository) GetAll(ctx context.Context) ([]store.Trial, error) {
	return nil, nil
}

func (repo MongoTrialRepository) Create(ctx context.Context) error {
	return nil
}
