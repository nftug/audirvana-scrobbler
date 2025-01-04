package app

import (
	"audirvana-scrobbler/app/internal/usecase"
	"audirvana-scrobbler/app/shared/response"
	"context"

	"github.com/samber/do"
)

type App struct {
	ctx          context.Context
	injector     *do.Injector
	getTrackInfo usecase.GetTrackInfo
}

func NewApp(i *do.Injector) (*App, error) {
	return &App{
		injector:     i,
		getTrackInfo: do.MustInvoke[usecase.GetTrackInfo](i),
	}, nil
}

func (a *App) OnDomReady(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Shutdown(ctx context.Context) {
	a.injector.Shutdown()
}

func (a *App) GetTrackInfo() ([]response.TrackInfo, error) {
	return a.getTrackInfo.Execute(a.ctx)
}
