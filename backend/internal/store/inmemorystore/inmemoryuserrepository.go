package inmemorystore

import (
	"context"
	"errors"
	"maps"
	"slices"

	"github.com/fransk/truthiness/internal/store"
)

type InMemoryUserRepository struct {
	users map[int64]store.User
}

func (repo *InMemoryUserRepository) Create(ctx context.Context, user *store.User) error {
	newUser := store.User{
		ID:          user.ID,
		Username:    user.Username,
		Password:    user.Password,
		Role:        user.Role,
		Permissions: user.Permissions,
	}
	repo.users[user.ID] = newUser
	return nil
}

func (repo *InMemoryUserRepository) GetAll(ctx context.Context) ([]store.User, error) {
	return slices.Collect(maps.Values(repo.users)), nil
}

func (repo *InMemoryUserRepository) GetById(ctx context.Context, id int64) (*store.User, error) {
	if user, ok := repo.users[id]; ok {
		newUser := store.User{
			ID:          user.ID,
			Username:    user.Username,
			Password:    user.Password,
			Role:        user.Role,
			Permissions: user.Permissions,
		}
		return &newUser, nil
	}
	return nil, errors.New("no user with that id")
}

func (repo *InMemoryUserRepository) GetByUsername(ctx context.Context, username string) (*store.User, error) {
	for _, user := range repo.users {
		if user.Username == username {
			newUser := store.User{
				ID:          user.ID,
				Username:    user.Username,
				Password:    user.Password,
				Role:        user.Role,
				Permissions: user.Permissions,
			}
			return &newUser, nil
		}
	}
	return nil, errors.New("no user with that username")
}
