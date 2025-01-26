package domain

import (
	"audirvana-scrobbler/app/bindings"
	"context"
	"time"
)

type AudirvanaUpdater interface {
	Update(ctx context.Context) error
}

type Scrobbler interface {
	Scrobble(ctx context.Context, tracks []TrackInfo) error
}

type TrackInfoRepository interface {
	Get(ctx context.Context, id string) (*TrackInfo, error)
	GetAll(ctx context.Context) ([]TrackInfo, error)
	GetLatestPlayedAt(ctx context.Context) (time.Time, error)
	Save(ctx context.Context, entity *TrackInfo) error
	MarkAsScrobbled(ctx context.Context, entities []TrackInfo) error
	Delete(ctx context.Context, id string) error
	CreateRange(ctx context.Context, tracks []TrackInfo) error
}

type TrackInfoQueryService interface {
	GetAll(ctx context.Context) ([]bindings.TrackInfoResponse, error)
}

type ConfigPathProvider interface {
	GetJoinedPath(filename string) string
	GetLocalPath() string
}

type ConfigProvider interface {
	Get() Config
	Write(cfg Config) error
}

type LastFMAPI interface {
	Login(ctx context.Context, username, password string) error
	Scrobble(ctx context.Context, tracks []TrackInfo) (map[string]any, error)
}
