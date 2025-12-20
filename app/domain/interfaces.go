package domain

import (
	"audirvana-scrobbler/app/bindings"
	"audirvana-scrobbler/app/lib/option"
	"context"
)

type NowPlayingTracker interface {
	StreamNowPlaying(ctx context.Context, npChan chan<- option.Option[NowPlaying], errChan chan<- error)
}

type TrackInfoRepository interface {
	Get(ctx context.Context, id int) (option.Option[TrackInfo], error)
	GetAll(ctx context.Context) ([]TrackInfo, error)
	Save(ctx context.Context, entity TrackInfo) (TrackInfo, error)
	SaveRange(ctx context.Context, entities []TrackInfo) ([]TrackInfo, error)
	Delete(ctx context.Context, id int) error
}

type TrackInfoQueryService interface {
	GetAll(ctx context.Context) ([]bindings.TrackInfoResponse, error)
}

type ConfigProvider interface {
	Get() Config
	Write(cfg Config) error
}

type LastFMAPI interface {
	GetSessionKey(ctx context.Context, username, password string) (string, error)
	LoginWithSessionKey(sessionKey string) error
	RemoveSessionKey()
	IsLoggedIn() bool
	Scrobble(ctx context.Context, tracks []TrackInfo) (map[string]any, error)
	UpdateNowPlaying(ctx context.Context, np NowPlaying) (map[string]any, error)
}
