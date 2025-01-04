package domain

import (
	"audirvana-scrobbler/app/shared/response"
	"context"
)

type AudirvanaImporter interface {
	GetAllTracks(ctx context.Context) ([]response.TrackInfo, error)
}
