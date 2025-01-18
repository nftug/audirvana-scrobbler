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
	Scrobble(ctx context.Context, tracks []bindings.TrackInfo) error
}

type TrackInfoRepository interface {
	GetAll(ctx context.Context) ([]bindings.TrackInfo, error)
	GetLatestPlayedAt(ctx context.Context) (time.Time, error)
	Save(ctx context.Context, id string, trackInfo bindings.TrackInfoForm) error
	MarkAsScrobbled(ctx context.Context, ids []string) error
	Delete(ctx context.Context, id string) error
}

type ConfigPathProvider interface {
	GetJoinedPath(filename string) string
	GetLocalPath() string
}

type ConfigProvider interface {
	Get() Config
	Write(cfg Config) error
}
