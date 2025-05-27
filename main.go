package main

import (
	"context"
	"embed"
	"sails-sftp/backend/sftp"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := NewApp()
	app.SFTP = sftp.NewSFTP(context.Background()) // Initialize SFTP with a background context

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "sails-sftp",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []any{
			app,
			app.SFTP,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
