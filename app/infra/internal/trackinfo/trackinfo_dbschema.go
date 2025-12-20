package trackinfo

import (
	"audirvana-scrobbler/app/bindings"
	"audirvana-scrobbler/app/domain"
	"audirvana-scrobbler/app/lib/option"
	"time"

	"github.com/samber/lo"
	"gorm.io/gorm"
)

type TrackInfoDBSchema struct {
	ID          int        `gorm:"primaryKey"`
	Artist      string     `gorm:"not null"`
	Album       string     `gorm:"not null"`
	Track       string     `gorm:"not null"`
	Duration    float64    `gorm:"not null"`
	PlayedAt    time.Time  `gorm:"not null"`
	ScrobbledAt *time.Time `gorm:"null;default:null"`
	DeletedAt   gorm.DeletedAt
}

func (TrackInfoDBSchema) TableName() string {
	return "track_info"
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
