package trackinfo

import (
	"audirvana-scrobbler/app/bindings"
	"audirvana-scrobbler/app/domain"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/samber/do"
	"github.com/samber/lo"
	"github.com/wailsapp/wails/v3/pkg/application"
)

type TrackNowPlaying interface {
	Run(app *application.App)
}

type trackNowPlayingImpl struct {
	repo        domain.TrackInfoRepository
	tracker     domain.NowPlayingTracker
	lastfm      domain.LastFMAPI
	cfgProvider domain.ConfigProvider
}

func NewTrackNowPlaying(i *do.Injector) (TrackNowPlaying, error) {
	return &trackNowPlayingImpl{
		repo:        do.MustInvoke[domain.TrackInfoRepository](i),
		tracker:     do.MustInvoke[domain.NowPlayingTracker](i),
		lastfm:      do.MustInvoke[domain.LastFMAPI](i),
		cfgProvider: do.MustInvoke[domain.ConfigProvider](i),
	}, nil
}

func (t *trackNowPlayingImpl) Run(app *application.App) {
	go t.execute(app)
}

func (t *trackNowPlayingImpl) execute(app *application.App) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	npChan := make(chan domain.NowPlaying, 5)
	errChan := make(chan error, 5)

	go t.tracker.StreamNowPlaying(ctx, npChan, errChan)

	cfg := t.cfgProvider.Get()
	if !t.lastfm.IsLoggedIn() {
		if err := t.lastfm.Login(ctx, cfg.UserName, cfg.Password); err != nil {
			t.notifyError(app, "failed to login: %v", err)
		}
	}

	npPrev := domain.NowPlaying{}
	lastProcessedTime := time.Now()

	for {
		select {
		case np := <-npChan:
			if lo.IsEmpty(np) || !t.lastfm.IsLoggedIn() {
				app.EmitEvent(bindings.NotifyNowPlaying, nil, nil)
				continue
			}

			cfg := t.cfgProvider.Get()

			// Update nowplaying
			app.EmitEvent(bindings.NotifyNowPlaying, np, nil)

			if cfg.ScrobbleImmediately && !npPrev.IsNotified {
				if _, err := t.lastfm.UpdateNowPlaying(ctx, np); err != nil {
					t.notifyError(app, "failed to update nowplaying: %v", err)
				}
				npPrev.IsNotified = true
			}

			// Update scrobble log
			// 5秒間隔で実行する
			if time.Since(lastProcessedTime) >= 5*time.Second {
				lastProcessedTime = time.Now()

				percentage := int(np.Position / np.Duration * 100)
				if percentage < cfg.PositionThreshold || npPrev.IsSaved {
					continue
				}

				track := domain.CreateTrackInfo(np, time.Now().UTC())

				// Scrobble
				if cfg.ScrobbleImmediately {
					tracks := []domain.TrackInfo{track}
					ret, err := t.lastfm.Scrobble(ctx, tracks)
					if err != nil {
						t.notifyError(app, "scrobbling failed: %v", err)
						continue
					}

					// For debug
					b, _ := json.Marshal(ret)
					log.Println("track.scrobble response: ", string(b))

					track = track.MarkAsScrobbled(time.Now().UTC())
				}

				// Save
				if _, err := t.repo.Save(ctx, track); err != nil {
					t.notifyError(app, "failed to save play log: %v", err)
					continue
				}

				npPrev.IsSaved = true
			}

			if !np.Equals(npPrev) {
				npPrev = np
			}

		case err := <-errChan:
			fmt.Printf("Error: %v\n", err)
			t.notifyError(app, "error while getting nowplaying: %v", err)
		}
	}
}

func (t *trackNowPlayingImpl) notifyError(app *application.App, format string, a ...any) {
	app.EmitEvent(bindings.NotifyNowPlaying, nil, bindings.NewInternalError(format, a...))
}
