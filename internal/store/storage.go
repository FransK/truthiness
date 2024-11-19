package store

import (
	"context"
)

// Trial has an unknown number of columns which represent
// the data from a single participant in an experiment
type Trial interface {
	GetData() map[string]string
}

// TrialRepository represents all the trials in a single
// experiment
type TrialRepository interface {
	GetAll(ctx context.Context) ([]Trial, error)
	Insert(ctx context.Context) error
}

// User needs to have an identifier and an auth token
// for performing actions
type User interface {
}

type UserRepository interface {
	GetById(ctx context.Context) (*User, error)
}

// Storage interface that combines all the repositories
type Storage interface {
	Trials() TrialRepository
	Users() UserRepository
}
