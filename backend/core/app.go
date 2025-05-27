package core

import (
	"context"
	"sails-sftp/backend/sftp"
	"sails-sftp/backend/utils/logging"
)

type App struct {
	ctx    context.Context
	SFTP   *sftp.SFTP // App will have a reference to SFTP -> this should be more modular (Connections?)
	Logger *logging.Logger
}

func NewApp(logger *logging.Logger) *App {
	return &App{
		SFTP:   sftp.NewSFTP(context.Background(), logger),
		Logger: logger,
	}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}
