package store

import (
	"context"
)

// Storage interface that combines all the repositories
type Storage interface {
	Experiments() ExperimentRepository
	Trials(experiment string) TrialRepository
	Users() UserRepository

	WithTransaction(ctx context.Context, fn func() (interface{}, error)) (interface{}, error)
}

// Experiment is a single run of an experiment that
// will have recruited participants to a number of trials
type Experiment struct {
	Name     string
	Date     string
	Location string
	Records  []string
}

// ExperimentRepository contains a list of all the
// experiments contained within the database
type ExperimentRepository interface {
	Create(ctx context.Context, experiment *Experiment) error
	GetAll(ctx context.Context) ([]Experiment, error)
}

// Trial has an unknown number of columns which represent
// the data from a single participant in an experiment
type Trial struct {
	Data map[string]string
}

// TrialRepository represents all the trials in a single
// experiment
type TrialRepository interface {
	Create(ctx context.Context, trial *Trial) error
	CreateMany(ctx context.Context, trials []Trial) error
	GetAll(ctx context.Context) ([]Trial, error)
}

// User needs to have an identifier
type User struct {
	ID          int64  `bson:"_id"`
	Username    string `bson:"username"`
	Password    string
	Role        string
	Permissions []string
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetAll(ctx context.Context) ([]User, error)
	GetById(ctx context.Context, id int64) (*User, error)
	GetByUsername(ctx context.Context, username string) (*User, error)
}
