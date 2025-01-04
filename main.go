package main

import (
	"audirvana-scrobbler/app"
	"audirvana-scrobbler/app/shared/enums"
	"embed"
	"math"

	"github.com/samber/do"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	i := app.BuildInjector()
	app := do.MustInvoke[*app.App](i)

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "Audirvana Scrobbler",
		Width:     1024,
		Height:    768,
		MaxWidth:  math.MaxInt16,
		MaxHeight: math.MaxInt16,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        nil,
		OnDomReady:       app.OnDomReady,
		OnShutdown:       app.Shutdown,
		Bind: []interface{}{
			app,
		},
		EnumBind: []interface{}{
			enums.ErrorCodes,
		},
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				FullSizeContent: true,
			},
			About: &mac.AboutInfo{
				Title:   "Audirvana Scrobbler",
				Message: "Last.fm scrobbler for Audirvana.",
			},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
