package trackinfo

import (
	"audirvana-scrobbler/app/bindings"
	"audirvana-scrobbler/app/domain"
	"context"
	"fmt"
	"time"

	"github.com/samber/do"
	"github.com/wailsapp/wails/v3/pkg/application"
)

type TrackNowPlaying interface {
	Execute(app *application.App)
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

func (t *trackNowPlayingImpl) Execute(app *application.App) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	npChan := make(chan *domain.NowPlaying)
	errChan := make(chan error)

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
			app.EmitEvent(string(bindings.NotifyNowPlaying), np, nil)
			if np == nil || !t.lastfm.IsLoggedIn() {
				continue
			}

			cfg := t.cfgProvider.Get()

			// Update nowplaying
			if cfg.ScrobbleImmediately && !npPrev.IsNotified {
				if _, err := t.lastfm.UpdateNowPlaying(ctx, *np); err != nil {
					t.notifyError(app, "failed to update nowplaying: %v", err)
				}
				npPrev.IsNotified = true
				// fmt.Println("Notified", np)
			}

			// Update scrobble log
			// 5秒間隔で実行する
			if time.Since(lastProcessedTime) >= 5*time.Second {
				lastProcessedTime = time.Now()
				if err := t.saveTrackAndScrobble(ctx, np, &npPrev); err != nil {
					t.notifyError(app, "%v", err)
					continue
				}
			}

			if !np.Equals(npPrev) {
				npPrev = *np
			}

		case err := <-errChan:
			fmt.Printf("Error: %v\n", err)
			t.notifyError(app, "error while getting nowplaying: %v", err)
		}
	}
}

func (t *trackNowPlayingImpl) saveTrackAndScrobble(
	ctx context.Context, np *domain.NowPlaying, npPrev *domain.NowPlaying) (err error) {
	cfg := t.cfgProvider.Get()

	percentage := int(np.Position / np.Duration * 100)
	if percentage < cfg.PositionThreshold || npPrev.IsSaved {
		return
	}

	track := domain.CreateTrackInfo(*np, time.Now().UTC())

	// Scrobble
	if cfg.ScrobbleImmediately {
		tracks := []domain.TrackInfo{*track}
		if _, err = t.lastfm.Scrobble(ctx, tracks); err != nil {
			return fmt.Errorf("scrobbling failed: %v", err)
		}
		track.MarkAsScrobbled(time.Now().UTC())
	}

	// Save
	if err = t.repo.Save(ctx, track); err != nil {
		return fmt.Errorf("failed to save play log: %v", err)
	}

	npPrev.IsSaved = true

	// fmt.Println("Scrobbled", track)
	return
}

func (t *trackNowPlayingImpl) notifyError(app *application.App, format string, a ...any) {
	app.EmitEvent(string(bindings.NotifyNowPlaying),
		nil, bindings.NewInternalError(format, a...))
}
