package domain

import (
	"audirvana-scrobbler/app/bindings"
	"audirvana-scrobbler/app/lib/option"
	"time"
)

type TrackInfo struct {
	id          int
	artist      string
	album       string
	track       string
	duration    float64
	playedAt    time.Time
	scrobbledAt option.Option[time.Time]
}

func (t TrackInfo) ID() int                               { return t.id }
func (t TrackInfo) Artist() string                        { return t.artist }
func (t TrackInfo) Album() string                         { return t.album }
func (t TrackInfo) Track() string                         { return t.track }
func (t TrackInfo) Duration() float64                     { return t.duration }
func (t TrackInfo) PlayedAt() time.Time                   { return t.playedAt }
func (t TrackInfo) ScrobbledAt() option.Option[time.Time] { return t.scrobbledAt }

func HydrateTrackInfo(
	id int,
	artist string,
	album string,
	track string,
	duration float64,
	playedAt time.Time,
	scrobbledAt option.Option[time.Time],
) TrackInfo {
	return TrackInfo{
		id:          id,
		artist:      artist,
		album:       album,
		track:       track,
		duration:    duration,
		playedAt:    playedAt,
		scrobbledAt: scrobbledAt,
	}
}

func NewTrackInfo(np NowPlaying, playedAt time.Time) TrackInfo {
	return TrackInfo{
		artist:   np.Artist,
		album:    np.Album,
		track:    np.Track,
		duration: np.Duration,
		playedAt: playedAt,
	}
}

func (t *TrackInfo) Update(form bindings.TrackInfoForm) error {
	if err := bindings.Validate(form); err != nil {
		return err
	}

	t.artist = form.Artist
	t.album = form.Album
	t.track = form.Track
	return nil
}

func (t *TrackInfo) MarkAsScrobbled(scrobbledAt time.Time) {
	t.scrobbledAt = option.Some(scrobbledAt)
}

func (t *TrackInfo) Equals(other TrackInfo) bool {
	return t.artist == other.artist && t.album == other.album && t.track == other.track
}
