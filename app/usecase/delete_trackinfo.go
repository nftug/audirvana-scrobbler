package usecase

import (
	"audirvana-scrobbler/app/bindings"
	"audirvana-scrobbler/app/domain"
	"context"
	"time"

	"github.com/samber/do"
)

type DeleteTrackInfo interface {
	Execute(ctx context.Context, id string) *bindings.ErrorResponse
}

type deleteTrackInfoImpl struct {
	repo domain.TrackInfoRepository
}

func NewDeleteTrackInfo(i *do.Injector) (DeleteTrackInfo, error) {
	return &deleteTrackInfoImpl{
		repo: do.MustInvoke[domain.TrackInfoRepository](i),
	}, nil
}

func (d *deleteTrackInfoImpl) Execute(ctx context.Context, id string) *bindings.ErrorResponse {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := d.repo.Delete(ctx, id); err != nil {
		return bindings.NewInternalError("Error while deleting track info: %v", err)
	}

	return nil
}
