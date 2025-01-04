package domain

import (
	"audirvana-scrobbler/app/shared/response"
	"context"
)

type AudirvanaUpdater interface {
	Update(ctx context.Context) error
}

type TrackInfoRepository interface {
	GetAll(ctx context.Context) ([]response.TrackInfo, error)
	// Save(ctx context.Context, track response.TrackInfo) error
}
