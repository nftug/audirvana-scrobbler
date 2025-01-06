package usecase

import (
	"audirvana-scrobbler/app/internal/domain"
	"audirvana-scrobbler/app/shared/customerr"
	"audirvana-scrobbler/app/shared/response"
	"context"
	"time"

	"github.com/samber/do"
)

type GetTrackInfo interface {
	Execute(ctx context.Context) ([]response.TrackInfo, *customerr.ErrorResponse)
}

type getTrackInfoImpl struct {
	updater domain.AudirvanaUpdater
	repo    domain.TrackInfoRepository
}

func NewGetTrackInfo(i *do.Injector) (GetTrackInfo, error) {
	return &getTrackInfoImpl{
		updater: do.MustInvoke[domain.AudirvanaUpdater](i),
		repo:    do.MustInvoke[domain.TrackInfoRepository](i),
	}, nil
}

func (g *getTrackInfoImpl) Execute(ctx context.Context) ([]response.TrackInfo, *customerr.ErrorResponse) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := g.updater.Update(ctx); err != nil {
		return nil, customerr.NewInternalError("Error while updating scrobble log: %v", err.Error())
	}

	result, err := g.repo.GetAll(ctx)
	if err != nil {
		return nil, customerr.NewInternalError("Error while getting scrobble log: %v", err.Error())
	}

	return result, nil
}
