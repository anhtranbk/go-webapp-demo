package app

import (
	"context"
	"webapp-demo/config"
)

type AppContext struct {
	Context context.Context
	Config  *config.Config
	Repo    *Repositories
}
