package usecase

import (
	"audirvana-scrobbler/app/bindings"
	"audirvana-scrobbler/app/domain"
	"context"

	"github.com/samber/do"
)

type ScrobbleAll interface {
	Execute(ctx context.Context) *bindings.ErrorResponse
}

type scrobbleAllImpl struct {
	repo      domain.TrackInfoRepository
	scrobbler domain.Scrobbler
}

func NewScrobbleAll(i *do.Injector) (ScrobbleAll, error) {
	return &scrobbleAllImpl{
		repo:      do.MustInvoke[domain.TrackInfoRepository](i),
		scrobbler: do.MustInvoke[domain.Scrobbler](i),
	}, nil
}

func (s *scrobbleAllImpl) Execute(ctx context.Context) *bindings.ErrorResponse {
	tracks, err := s.repo.GetAll(ctx)
	if err != nil {
		return bindings.NewInternalError("Error while getting tracks from DB: %v", err)
	}
	if err := s.scrobbler.Scrobble(ctx, tracks); err != nil {
		return bindings.NewInternalError("Error while scrobbling: %v", err)
	}
	return nil
}
