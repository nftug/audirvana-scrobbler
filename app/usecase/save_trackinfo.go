package usecase

import (
	"audirvana-scrobbler/app/bindings"
	"audirvana-scrobbler/app/domain"
	"context"

	"github.com/samber/do"
)

type SaveTrackInfo interface {
	Execute(ctx context.Context, id string, form bindings.TrackInfoForm) *bindings.ErrorResponse
}

type saveTrackInfoImpl struct {
	repo domain.TrackInfoRepository
}

func NewSaveTrackInfo(i *do.Injector) (SaveTrackInfo, error) {
	return &saveTrackInfoImpl{
		repo: do.MustInvoke[domain.TrackInfoRepository](i),
	}, nil
}

func (s *saveTrackInfoImpl) Execute(ctx context.Context, id string, form bindings.TrackInfoForm) *bindings.ErrorResponse {
	track, err := s.repo.Get(ctx, id)
	if err != nil {
		return bindings.NewInternalError("Error while getting track info from DB: %v", err)
	}
	if track == nil {
		return bindings.NewNotFoundError()
	}

	track, err = track.Update(form)
	if err != nil {
		return err.(*bindings.ErrorResponse)
	}

	if err := s.repo.Save(ctx, track); err != nil {
		return bindings.NewInternalError("Error while saving track info: %v", err)
	}

	return nil
}
