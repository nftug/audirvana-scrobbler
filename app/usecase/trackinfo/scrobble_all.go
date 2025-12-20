package trackinfo

import (
	"audirvana-scrobbler/app/bindings"
	"audirvana-scrobbler/app/domain"
	"context"
	"time"

	"github.com/samber/do"
	"github.com/samber/lo"
)

type ScrobbleAll interface {
	Execute(ctx context.Context) *bindings.ErrorResponse
}

type scrobbleAllImpl struct {
	repo   domain.TrackInfoRepository
	lastfm domain.LastFMAPI
}

func NewScrobbleAll(i *do.Injector) (ScrobbleAll, error) {
	return &scrobbleAllImpl{
		repo:   do.MustInvoke[domain.TrackInfoRepository](i),
		lastfm: do.MustInvoke[domain.LastFMAPI](i),
	}, nil
}

func (s scrobbleAllImpl) Execute(ctx context.Context) *bindings.ErrorResponse {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	tracks, err := s.repo.GetAll(ctx)
	if err != nil {
		return bindings.NewInternalError("Error while getting tracks from DB: %v", err)
	}
	tracks = lo.Filter(tracks,
		func(t domain.TrackInfo, _ int) bool { return t.ScrobbledAt().IsNone() })

	if !s.lastfm.IsLoggedIn() {
		return bindings.NewNotLoggedInError()
	}

	chunks := lo.Chunk(tracks, 50)
	for i, tracks := range chunks {
		if _, err := s.lastfm.Scrobble(ctx, tracks); err != nil {
			return bindings.NewInternalError("Error while scrobbling tracks: %v", err)
		}

		updatedTracks := lo.Map(tracks, func(t domain.TrackInfo, _ int) domain.TrackInfo {
			return t.MarkAsScrobbled(time.Now().UTC())
		})

		if _, err := s.repo.SaveRange(ctx, updatedTracks); err != nil {
			return bindings.NewInternalError("Error while saving scrobbled tracks: %v", err)
		}

		if i < len(chunks)-1 {
			time.Sleep(2 * time.Second)
		}
	}

	return nil
}
