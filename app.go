package main

import (
	"context"
	"fmt"
	"sails-sftp/backend/sftp"
)

// App struct
type App struct {
	ctx context.Context
	SFTP *sftp.SFTP // Embed SFTP struct
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.SFTP = sftp.NewSFTP(ctx) // Initialize SFTP with context
}


// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time we are very cool!", name)
}



