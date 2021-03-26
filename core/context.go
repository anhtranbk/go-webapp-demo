package core

import (
	"context"
	"webapp-demo/config"
	"webapp-demo/repository"
)

type AppContext struct {
	Context context.Context
	Config  *config.Config
	Repo    *repository.Repositories
}
