package api

import (
	"github.com/fransk/truthiness/internal/store"
)

type Application struct {
	Config Config
	Store  store.Storage
}

type Config struct {
	Addr string
}
