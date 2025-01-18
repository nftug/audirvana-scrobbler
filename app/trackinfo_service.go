package app

import (
	"audirvana-scrobbler/app/bindings"
	"audirvana-scrobbler/app/usecase"
	"context"

	"github.com/samber/do"
)

type TrackInfoService struct {
	getTrackInfo    usecase.GetTrackInfo
	saveTrackInfo   usecase.SaveTrackInfo
	deleteTrackInfo usecase.DeleteTrackInfo
	scrobbleAll     usecase.ScrobbleAll
}

func NewApp(i *do.Injector) (*TrackInfoService, error) {
	return &TrackInfoService{
		getTrackInfo:    do.MustInvoke[usecase.GetTrackInfo](i),
		saveTrackInfo:   do.MustInvoke[usecase.SaveTrackInfo](i),
		deleteTrackInfo: do.MustInvoke[usecase.DeleteTrackInfo](i),
		scrobbleAll:     do.MustInvoke[usecase.ScrobbleAll](i),
	}, nil
}

func (t *TrackInfoService) GetTrackInfo(ctx context.Context) ([]bindings.TrackInfo, *bindings.ErrorResponse) {
	return t.getTrackInfo.Execute(ctx)
}

func (t *TrackInfoService) SaveTrackInfo(ctx context.Context, id string, form bindings.TrackInfoForm) *bindings.ErrorResponse {
	return t.saveTrackInfo.Execute(ctx, id, form)
}

func (t *TrackInfoService) DeleteTrackInfo(ctx context.Context, id string) *bindings.ErrorResponse {
	return t.deleteTrackInfo.Execute(ctx, id)
}

func (t *TrackInfoService) ScrobbleAll(ctx context.Context) *bindings.ErrorResponse {
	return t.scrobbleAll.Execute(ctx)
}
