package trackinfo

import (
	"audirvana-scrobbler/app/bindings"
	"audirvana-scrobbler/app/domain"
	"context"

	"github.com/samber/do"
	"github.com/samber/lo"
)

type SaveTrackInfo interface {
	Execute(ctx context.Context, id int, form bindings.TrackInfoForm) *bindings.ErrorResponse
}

type saveTrackInfoImpl struct {
	repo domain.TrackInfoRepository
}

func NewSaveTrackInfo(i *do.Injector) (SaveTrackInfo, error) {
	return &saveTrackInfoImpl{
		repo: do.MustInvoke[domain.TrackInfoRepository](i),
	}, nil
}

func (s saveTrackInfoImpl) Execute(ctx context.Context, id int, form bindings.TrackInfoForm) *bindings.ErrorResponse {
	track, err := s.repo.Get(ctx, id)
	if err != nil {
		return bindings.NewInternalError("Error while getting track info from DB: %v", err)
	}
	if lo.IsEmpty(track) {
		return bindings.NewNotFoundError()
	}

	track, err = track.Update(form)
	if err != nil {
		return err.(*bindings.ErrorResponse)
	}

	if _, err := s.repo.Save(ctx, track); err != nil {
		return bindings.NewInternalError("Error while saving track info: %v", err)
	}

	return nil
}
