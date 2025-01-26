package main

import (
	"audirvana-scrobbler/app"
	"audirvana-scrobbler/app/usecase/trackinfo"
	"embed"
	"fmt"
	"os"
	"runtime"

	"github.com/samber/do"
	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/icons"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	if runtime.GOOS != "darwin" {
		fmt.Fprintln(os.Stderr, "This app is only for MacOS.")
		os.Exit(-1)
	}

	// Create an instance of the app structure
	injector := app.BuildInjector()
	service := do.MustInvoke[*app.TrackInfoService](injector)

	app := application.New(application.Options{
		Name:        "Audirvana Scrobbler",
		Description: "Last.fm scrobbler for MacOS.",

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

	// do.Invoke経由で使う
	do.Provide(injector, func(i *do.Injector) (*application.App, error) {
		return app, nil
	})

	window := app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title:  "Audirvana Scrobbler",
		Width:  1024,
		Height: 768,
		URL:    "/",
		ShouldClose: func(window *application.WebviewWindow) bool {
			window.Hide()
			return false
		},
	})

	do.Provide(injector, func(i *do.Injector) (*application.WebviewWindow, error) {
		return window, nil
	})

	// Start nowplaying tracker
	tracker := do.MustInvoke[trackinfo.TrackNowPlaying](injector)
	go tracker.Execute(app)

	systray := app.NewSystemTray()
	systray.SetTemplateIcon(icons.SystrayMacTemplate)
	systray.OnClick(func() {
		window.Show()
	})

	if err := app.Run(); err != nil {
		println("Error:", err.Error())
	}
}
