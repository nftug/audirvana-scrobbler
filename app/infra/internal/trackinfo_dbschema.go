package internal

import (
	"audirvana-scrobbler/app/bindings"
	"time"
)

type TrackInfoDBSchema struct {
	ID          string     `gorm:"primaryKey"`
	Artist      string     `gorm:"not null"`
	Album       string     `gorm:"not null"`
	Track       string     `gorm:"not null"`
	PlayedAt    time.Time  `gorm:"not null"`
	ScrobbledAt *time.Time `gorm:"null;default:null"`
	DeletedAt   *time.Time `gorm:"null;default:null"`
}

func (TrackInfoDBSchema) TableName() string {
	return "track_info"
}

func (t *TrackInfoDBSchema) ToResponse() bindings.TrackInfo {
	return bindings.TrackInfo{
		ID:       t.ID,
		Artist:   t.Artist,
		Album:    t.Album,
		Track:    t.Track,
		PlayedAt: t.PlayedAt.Format(time.RFC3339),
	}
}
