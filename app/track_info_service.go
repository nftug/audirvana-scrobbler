package app

import (
	"audirvana-scrobbler/app/bindings"
	"audirvana-scrobbler/app/usecase"
	"context"

	"github.com/samber/do"
)

type TrackInfoService struct {
	getTrackInfo  usecase.GetTrackInfo
	saveTrackInfo usecase.SaveTrackInfo
}

func NewApp(i *do.Injector) (*TrackInfoService, error) {
	// build injector
	return &TrackInfoService{
		getTrackInfo:  do.MustInvoke[usecase.GetTrackInfo](i),
		saveTrackInfo: do.MustInvoke[usecase.SaveTrackInfo](i),
	}, nil
}

func (t *TrackInfoService) GetTrackInfo(ctx context.Context) ([]bindings.TrackInfo, *bindings.ErrorResponse) {
	return t.getTrackInfo.Execute(ctx)
}

func (t *TrackInfoService) SaveTrackInfo(ctx context.Context, id string, form bindings.TrackInfoForm) *bindings.ErrorResponse {
	return t.saveTrackInfo.Execute(ctx, id, form)
}
