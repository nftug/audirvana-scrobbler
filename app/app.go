package app

import (
	"audirvana-scrobbler/app/internal/usecase"
	"audirvana-scrobbler/app/shared/customerr"
	"audirvana-scrobbler/app/shared/response"
	"context"

	"github.com/samber/do"
)

type App struct {
	getTrackInfo usecase.GetTrackInfo
}

func NewApp(i *do.Injector) (*App, error) {
	return &App{
		getTrackInfo: do.MustInvoke[usecase.GetTrackInfo](i),
	}, nil
}

func (a *App) GetTrackInfo(ctx context.Context) ([]response.TrackInfo, *customerr.ErrorResponse) {
	return a.getTrackInfo.Execute(ctx)
}
