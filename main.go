package main

import (
	"context"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"sails-sftp/backend/core"
	"sails-sftp/backend/utils/logging"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	logger := logging.NewLogger(logging.LevelInfo, "logs/app.log")

	logger.Log(logging.LevelInfo, "Starting Sails SFTP application...")

	app := core.NewApp(logger)
	// app.SFTP = sftp.NewSFTP(co() // Initialize SFTP with a background context

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "sails-sftp",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			app.Startup(ctx)
		},
		Bind: []any{
			app,
			app.SFTP,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
