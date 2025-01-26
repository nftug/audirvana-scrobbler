package domain

import (
	"audirvana-scrobbler/app/bindings"
	"audirvana-scrobbler/app/lib"
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type TrackInfo struct {
	id          int
	artist      string
	album       string
	track       string
	duration    float64
	playedAt    time.Time
	scrobbledAt lib.Nullable[time.Time]
}

func (t TrackInfo) ID() int                 { return t.id }
func (t TrackInfo) Artist() string          { return t.artist }
func (t TrackInfo) Album() string           { return t.album }
func (t TrackInfo) Track() string           { return t.track }
func (t TrackInfo) Duration() float64       { return t.duration }
func (t TrackInfo) PlayedAt() time.Time     { return t.playedAt }
func (t TrackInfo) ScrobbledAt() *time.Time { return t.scrobbledAt.ToCopiedPtr() }

func ReconstructTrackInfo(
	id int,
	artist string,
	album string,
	track string,
	duration float64,
	playedAt time.Time,
	scrobbledAt *time.Time,
) *TrackInfo {
	return &TrackInfo{
		id:          id,
		artist:      artist,
		album:       album,
		track:       track,
		playedAt:    playedAt,
		scrobbledAt: lib.NewNullable(scrobbledAt),
	}
}

func CreateTrackInfo(np NowPlaying, playedAt time.Time) *TrackInfo {
	return &TrackInfo{
		artist:   np.Artist,
		album:    np.Album,
		track:    np.Track,
		playedAt: playedAt,
	}
}

func (t TrackInfo) Update(form bindings.TrackInfoForm) (*TrackInfo, error) {
	var errData []bindings.ErrorData
	if err := validator.New().Struct(form); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errData = append(errData, bindings.ErrorData{
				Field:   strings.ToLower(err.Field()),
				Message: fmt.Sprintf("Validation error on %s: %s %s", err.Field(), err.Tag(), err.Param()),
			})
		}
	}
	if len(errData) > 0 {
		return nil, &bindings.ErrorResponse{
			Code: bindings.ValidationError,
			Data: errData,
		}
	}

	updated := &TrackInfo{
		id:          t.id,
		artist:      form.Artist,
		album:       form.Album,
		track:       form.Track,
		playedAt:    t.playedAt,
		scrobbledAt: t.scrobbledAt,
	}
	return updated, nil
}

func (t TrackInfo) MarkAsScrobbled(scrobbledAt time.Time) *TrackInfo {
	return &TrackInfo{
		id:          t.id,
		artist:      t.artist,
		album:       t.album,
		track:       t.track,
		playedAt:    t.playedAt,
		scrobbledAt: lib.NewNullableByVal(scrobbledAt),
	}
}
