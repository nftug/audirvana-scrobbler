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
	Scrobble(ctx context.Context, tracks []bindings.TrackInfoResponse) error
}

type TrackInfoRepository interface {
	Get(ctx context.Context, id string) (*TrackInfo, error)
	GetAll(ctx context.Context) ([]bindings.TrackInfoResponse, error)
	GetLatestPlayedAt(ctx context.Context) (time.Time, error)
	Save(ctx context.Context, entity *TrackInfo) error
	MarkAsScrobbled(ctx context.Context, ids []string) error
	Delete(ctx context.Context, id string) error
	UpdateRange(ctx context.Context, tracks []TrackInfo) error
}

type ConfigPathProvider interface {
	GetJoinedPath(filename string) string
	GetLocalPath() string
}

type ConfigProvider interface {
	Get() Config
	Write(cfg Config) error
}
