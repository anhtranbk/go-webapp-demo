package core

import (
	"webapp-demo/config"
	"webapp-demo/repository"
)

type AppContext struct {
	Config *config.Config
	Repo   *repository.Repositories
}
