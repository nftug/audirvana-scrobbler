package domain

import (
	"audirvana-scrobbler/app/bindings"
	"context"
)

type NowPlayingTracker interface {
	StreamNowPlaying(ctx context.Context, npChan chan<- NowPlaying, errChan chan<- error)
}

type TrackInfoRepository interface {
	Get(ctx context.Context, id int) (TrackInfo, error)
	GetAll(ctx context.Context) ([]TrackInfo, error)
	Save(ctx context.Context, entity TrackInfo) (TrackInfo, error)
	MarkAsScrobbled(ctx context.Context, entities []TrackInfo) ([]TrackInfo, error)
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
	IsLoggedIn() bool
	Login(ctx context.Context, username, password string) error
	Scrobble(ctx context.Context, tracks []TrackInfo) (map[string]any, error)
	UpdateNowPlaying(ctx context.Context, np NowPlaying) (map[string]any, error)
}
