package domain

import (
	"audirvana-scrobbler/app/bindings"
)

type NowPlaying struct {
	AppName    string  `json:"appName"`
	Track      string  `json:"track"`
	Artist     string  `json:"artist"`
	Album      string  `json:"album"`
	Duration   float64 `json:"duration"`
	Position   float64 `json:"position"`
	IsNotified bool
	IsSaved    bool
}

func (np NowPlaying) ToResponse() bindings.NowPlayingResponse {
	return bindings.NowPlayingResponse{
		AppName:  np.AppName,
		Track:    np.Track,
		Artist:   np.Artist,
		Album:    np.Album,
		Duration: np.Duration,
		Position: np.Position,
	}
}

func (np NowPlaying) Equals(other NowPlaying) bool {
	return np.Track == other.Track && np.Artist == other.Artist && np.Album == other.Album
}
