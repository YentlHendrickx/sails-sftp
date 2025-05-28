package core

import (
	"context"
	"sails-sftp/backend/protocols"
	"sails-sftp/backend/utils/logging"
)

type App struct {
	ctx     context.Context
	Logger  *logging.Logger
	Manager *protocols.ConnectionManager
}

func NewApp(logger *logging.Logger) *App {
	return &App{
		Logger:  logger,
		Manager: protocols.NewManager(context.Background(), logger),
	}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}
