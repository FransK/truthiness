package inmemorystore

import (
	"context"
	"errors"

	"github.com/fransk/truthiness/internal/store"
)

type InMemoryUserRepository struct {
	users map[int64]store.User
}

func (repo *InMemoryUserRepository) Create(ctx context.Context, user *store.User) error {
	newUser := store.User{
		ID: user.ID,
	}
	repo.users[user.ID] = newUser
	return nil
}

func (repo *InMemoryUserRepository) GetById(ctx context.Context, id int64) (*store.User, error) {
	if v, ok := repo.users[id]; ok {
		newUser := store.User{
			ID: v.ID,
		}
		return &newUser, nil
	}
	return nil, errors.New("no user with that id")
}
