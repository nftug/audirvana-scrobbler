package main

import (
	"audirvana-scrobbler/app"
	"embed"

	"github.com/samber/do"
	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	i := app.BuildInjector()
	service := do.MustInvoke[*app.App](i)

	app := application.New(application.Options{
		Name:        "Audirvana Scrobbler",
		Description: "Last.fm scrobbler for Audirvana.",

		Services: []application.Service{
			application.NewService(service),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: false,
		},
		OnShutdown: func() { i.Shutdown() },
	})

	app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title:  "Audirvana Scrobbler",
		Width:  1024,
		Height: 768,
		URL:    "/",
	})

	err := app.Run()

	if err != nil {
		println("Error:", err.Error())
	}
}
