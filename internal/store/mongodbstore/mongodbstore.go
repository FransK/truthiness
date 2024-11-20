package mongodbstore

import (
	"context"
	"log"

	"github.com/fransk/truthiness/internal/store"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewMongoDbStore(mongodb *mongo.Database) store.Storage {
	return &MongoDbStorage{
		db: mongodb,
	}
}

type MongoDbStorage struct {
	db *mongo.Database
}

func (store *MongoDbStorage) Experiments() store.ExperimentRepository {
	return MongoExperimentRepository{
		collection: store.db.Collection("experiments"),
	}
}

func (store *MongoDbStorage) Trials(experiment string) store.TrialRepository {
	return MongoTrialRepository{
		collection: store.db.Collection(experiment),
	}
}

func (store *MongoDbStorage) Users() store.UserRepository {
	return MongoUserRepository{
		collection: store.db.Collection("users"),
	}
}

type MongoExperimentRepository struct {
	collection *mongo.Collection
}

func (repo MongoExperimentRepository) Create(ctx context.Context) error {
	return nil
}

func (repo MongoExperimentRepository) GetAll() ([]store.Experiment, error) {
	cursor, err := repo.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var experiments []store.Experiment
	if err = cursor.All(context.Background(), &experiments); err != nil {
		log.Fatal(err)
	}

	return experiments, nil
}

type MongoUserRepository struct {
	collection *mongo.Collection
}

func (repo MongoUserRepository) GetById(ctx context.Context) (*store.User, error) {
	return nil, nil
}

func (repo MongoUserRepository) Create(ctx context.Context) error {
	return nil
}

type MongoTrialRepository struct {
	collection *mongo.Collection
}

func (repo MongoTrialRepository) GetAll(ctx context.Context) ([]store.Trial, error) {
	return nil, nil
}

func (repo MongoTrialRepository) Create(ctx context.Context) error {
	return nil
}
