package internal

import (
	"audirvana-scrobbler/app/bindings"
	"audirvana-scrobbler/app/domain"
	"time"

	"gorm.io/gorm"
)

type TrackInfoDBSchema struct {
	ID          string     `gorm:"primaryKey"`
	Artist      string     `gorm:"not null"`
	Album       string     `gorm:"not null"`
	Track       string     `gorm:"not null"`
	PlayedAt    time.Time  `gorm:"not null"`
	ScrobbledAt *time.Time `gorm:"null;default:null"`
	DeletedAt   gorm.DeletedAt
}

func (TrackInfoDBSchema) TableName() string {
	return "track_info"
}

func (t *TrackInfoDBSchema) ToResponse() bindings.TrackInfoResponse {
	return bindings.TrackInfoResponse{
		ID:       t.ID,
		Artist:   t.Artist,
		Album:    t.Album,
		Track:    t.Track,
		PlayedAt: t.PlayedAt.Format(time.RFC3339),
	}
}

func (t *TrackInfoDBSchema) ToEntity() *domain.TrackInfo {
	return domain.ReconstructTrackInfo(t.ID, t.Artist, t.Album, t.Track, t.PlayedAt, t.ScrobbledAt)
}
