package api

import (
	"time"

	"github.com/fransk/truthiness/internal/store"
)

type Application struct {
	Config Config
	Store  store.Storage
}

type Config struct {
	Addr string
	DB   DbConfig
}

type DbConfig struct {
	Addr         string
	MaxOpenConns int
	MaxIdleTime  time.Duration
}
