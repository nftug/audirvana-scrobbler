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
	"github.com/wailsapp/wails/v3/pkg/events"
	"github.com/wailsapp/wails/v3/pkg/icons"
)

//go:embed frontend/dist
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
			ActivationPolicy: application.ActivationPolicyAccessory,
		},
	})

	// Start now playing tracker
	tracker := do.MustInvoke[trackinfo.TrackNowPlaying](injector)
	go tracker.Execute(app)

	window := app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:  "Audirvana Scrobbler",
		Width:  1024,
		Height: 768,
		URL:    "/",
		Hidden: true,
	})

	window.RegisterHook(events.Mac.WindowShouldClose, func(event *application.WindowEvent) {
		event.Cancel()
		window.Hide()
	})

	showWindow := func() {
		if window.NativeWindow() == nil {
			return
		}
		window.Show()
		window.Focus()
	}

	systray := app.SystemTray.New()
	systray.SetTemplateIcon(icons.SystrayMacTemplate)
	systray.OnClick(showWindow)

	menu := app.NewMenu()
	menu.Add("Open app").OnClick(func(ctx *application.Context) {
		showWindow()
	})
	menu.Add("Quit").OnClick(func(ctx *application.Context) {
		app.Quit()
	})
	systray.SetMenu(menu)

	if err := app.Run(); err != nil {
		println("Error:", err.Error())
	}
}
