package core

import (
	"context"
	"webapp-demo/app/repository"
	"webapp-demo/config"
)

type AppContext struct {
	Context context.Context
	Config  *config.Config
	Repo    *repository.Repositories
}
