package core

import (
	"context"
	"webapp-demo/app"
	"webapp-demo/config"
)

type AppContext struct {
	Context context.Context
	Config  *config.Config
	Repo    *app.Repositories
}
