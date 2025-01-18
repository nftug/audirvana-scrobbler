package app

import (
	"audirvana-scrobbler/app/bindings"
	"audirvana-scrobbler/app/usecase"
	"context"

	"github.com/samber/do"
)

type TrackListService struct {
	getTrackInfo usecase.GetTrackInfo
}

func NewApp(i *do.Injector) (*TrackListService, error) {
	// build injector
	return &TrackListService{
		getTrackInfo: do.MustInvoke[usecase.GetTrackInfo](i),
	}, nil
}

func (a *TrackListService) GetTrackInfo(ctx context.Context) ([]bindings.TrackInfo, *bindings.ErrorResponse) {
	return a.getTrackInfo.Execute(ctx)
}
