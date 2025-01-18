package domain

import (
	"audirvana-scrobbler/app/bindings"
	"context"
)

type AudirvanaUpdater interface {
	Update(ctx context.Context) error
}

type TrackInfoRepository interface {
	GetAll(ctx context.Context) ([]bindings.TrackInfo, error)
	// Save(ctx context.Context, track response.TrackInfo) error
}

type ConfigPathProvider interface {
	GetJoinedPath(filename string) string
	GetLocalPath() string
}
