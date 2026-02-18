package trackinfo

import (
	"audirvana-scrobbler/app/bindings"
	"audirvana-scrobbler/app/domain"
	"audirvana-scrobbler/app/lib/option"
	"time"

	"github.com/samber/lo"
	"github.com/uptrace/bun"
)

type TrackInfoDBSchema struct {
	bun.BaseModel `bun:"table:track_info"`
	ID            int        `bun:"id,pk,autoincrement,nullzero"`
	Artist        string     `bun:"artist,notnull"`
	Album         string     `bun:"album,notnull"`
	Track         string     `bun:"track,notnull"`
	Duration      float64    `bun:"duration,notnull"`
	PlayedAt      time.Time  `bun:"played_at,notnull"`
	ScrobbledAt   *time.Time `bun:"scrobbled_at,nullzero"`
	DeletedAt     *time.Time `bun:"deleted_at,nullzero"`
}

func NewTrackInfoDBSchema(entity domain.TrackInfo) TrackInfoDBSchema {
	return TrackInfoDBSchema{
		ID:          entity.ID(),
		Artist:      entity.Artist(),
		Album:       entity.Album(),
		Track:       entity.Track(),
		Duration:    entity.Duration(),
		PlayedAt:    entity.PlayedAt(),
		ScrobbledAt: entity.ScrobbledAt().Unwrap(),
	}
}

func (t TrackInfoDBSchema) ToResponse() bindings.TrackInfoResponse {
	return bindings.TrackInfoResponse{
		ID:       t.ID,
		Artist:   t.Artist,
		Album:    t.Album,
		Track:    t.Track,
		PlayedAt: t.PlayedAt.Format(time.RFC3339),
		ScrobbledAt: lo.TernaryF(
			t.ScrobbledAt != nil,
			func() *string { return lo.ToPtr(t.ScrobbledAt.Format(time.RFC3339)) },
			func() *string { return nil },
		),
	}
}

func (t TrackInfoDBSchema) ToEntity() domain.TrackInfo {
	return domain.HydrateTrackInfo(
		t.ID, t.Artist, t.Album, t.Track, t.Duration, t.PlayedAt, option.NewOption(t.ScrobbledAt))
}
