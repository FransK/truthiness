package store

import (
	"context"
)

// Experiment is a single run of an experiment that
// will have recruited participants to a number of trials
type Experiment struct {
	Name     string
	Date     string
	Location string
}

// ExperimentRepository contains a list of all the
// experiments contained within the database
type ExperimentRepository interface {
	Create(ctx context.Context) error
	GetAll() ([]Experiment, error)
}

// Trial has an unknown number of columns which represent
// the data from a single participant in an experiment
type Trial struct {
	Data map[string]string
}

// TrialRepository represents all the trials in a single
// experiment
type TrialRepository interface {
	Create(ctx context.Context) error
	GetAll(ctx context.Context) ([]Trial, error)
}

// User needs to have an identifier
type User struct {
	ID int64
}

type UserRepository interface {
	Create(ctx context.Context) error
	GetById(ctx context.Context) (*User, error)
}

// Storage interface that combines all the repositories
type Storage interface {
	Experiments() ExperimentRepository
	Trials(experiment string) TrialRepository
	Users() UserRepository
}
