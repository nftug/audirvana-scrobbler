package infra

import (
	"audirvana-scrobbler/app/bindings"
	"audirvana-scrobbler/app/domain"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/samber/do"
	"github.com/samber/lo"
	"github.com/shkh/lastfm-go/lastfm"
)

type scrobblerImpl struct {
	api            *lastfm.Api
	repo           domain.TrackInfoRepository
	configProvider domain.ConfigProvider
}

func NewScrobbler(i *do.Injector) (domain.Scrobbler, error) {
	configProvider := do.MustInvoke[domain.ConfigProvider](i)
	cfg := configProvider.Get()

	return &scrobblerImpl{
		api:            lastfm.New(cfg.APIKey, cfg.APISecret),
		repo:           do.MustInvoke[domain.TrackInfoRepository](i),
		configProvider: configProvider,
	}, nil
}

func (s *scrobblerImpl) Scrobble(ctx context.Context, tracks []bindings.TrackInfoResponse) error {
	cfg := s.configProvider.Get()
	if cfg.UserName == "" {
		return errors.New("user name is empty")
	} else if cfg.Password == "" {
		return errors.New("password is empty")
	}

	if err := s.api.Login(cfg.UserName, cfg.Password); err != nil {
		return err
	}

	chunks := lo.Chunk(tracks, 50)
	for i, tracks := range chunks {
		data := lastfm.P{}
		for j, track := range tracks {
			data[fmt.Sprintf("artist[%d]", j)] = track.Artist
			data[fmt.Sprintf("albumArtist[%d]", j)] = track.Artist
			data[fmt.Sprintf("track[%d]", j)] = track.Track
			data[fmt.Sprintf("album[%d]", j)] = track.Album
			data[fmt.Sprintf("timestamp[%d]", j)] = track.PlayedAt
		}

		_, err := s.api.Track.Scrobble(data)
		if err != nil {
			return err
		}

		ids := lo.Map(tracks, func(t bindings.TrackInfoResponse, _ int) string { return t.ID })
		if err := s.repo.MarkAsScrobbled(ctx, ids); err != nil {
			return err
		}

		if i < len(chunks)-1 {
			time.Sleep(5 * time.Second)
		}
	}

	return nil
}
