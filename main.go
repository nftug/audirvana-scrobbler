package main

import (
	"audirvana-scrobbler/app"
	"embed"
	"fmt"

	"github.com/samber/do"
	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	injector := app.BuildInjector()
	service := do.MustInvoke[*app.TrackInfoService](injector)

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
	})

	window := app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title:  "Audirvana Scrobbler",
		Width:  1024,
		Height: 768,
		URL:    "/",
		ShouldClose: func(window *application.WebviewWindow) bool {
			window.Minimise()
			return false
			// return true
		},
	})

	app.OnApplicationEvent(events.Mac.ApplicationDidBecomeActive, func(event *application.ApplicationEvent) {
		fmt.Println("Restored")
		window.Show()
	})

	systray := app.NewSystemTray()
	systray.OnClick(func() {
		window.Show()
	})

	if err := app.Run(); err != nil {
		println("Error:", err.Error())
	}
}
