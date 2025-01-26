package usecase

import (
	"audirvana-scrobbler/app/bindings"
	"audirvana-scrobbler/app/domain"
	"context"
	"time"

	"github.com/samber/do"
)

type GetTrackInfo interface {
	Execute(ctx context.Context) ([]bindings.TrackInfoResponse, *bindings.ErrorResponse)
}

type getTrackInfoImpl struct {
	updater domain.AudirvanaUpdater
	qs      domain.TrackInfoQueryService
}

func NewGetTrackInfo(i *do.Injector) (GetTrackInfo, error) {
	return &getTrackInfoImpl{
		updater: do.MustInvoke[domain.AudirvanaUpdater](i),
		qs:      do.MustInvoke[domain.TrackInfoQueryService](i),
	}, nil
}

func (g *getTrackInfoImpl) Execute(ctx context.Context) ([]bindings.TrackInfoResponse, *bindings.ErrorResponse) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := g.updater.Update(ctx); err != nil {
		return nil, bindings.NewInternalError("Error while updating scrobble log: %v", err.Error())
	}

	result, err := g.qs.GetAll(ctx)
	if err != nil {
		return nil, bindings.NewInternalError("Error while getting scrobble log: %v", err.Error())
	}

	return result, nil
}
