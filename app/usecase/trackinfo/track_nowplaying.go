package trackinfo

import (
	"audirvana-scrobbler/app/bindings"
	"audirvana-scrobbler/app/domain"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"sync"

	"github.com/samber/do"
	"github.com/samber/lo"
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
	app         *application.App
	mutex       sync.Mutex
	npPrev      domain.NowPlaying
}

func NewTrackNowPlaying(i *do.Injector) (TrackNowPlaying, error) {
	return &trackNowPlayingImpl{
		repo:        do.MustInvoke[domain.TrackInfoRepository](i),
		tracker:     do.MustInvoke[domain.NowPlayingTracker](i),
		lastfm:      do.MustInvoke[domain.LastFMAPI](i),
		cfgProvider: do.MustInvoke[domain.ConfigProvider](i),
		mutex:       sync.Mutex{},
		npPrev:      domain.NowPlaying{},
	}, nil
}

func (t *trackNowPlayingImpl) Execute(app *application.App) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	t.app = app

	npChan := make(chan domain.NowPlaying, 5)
	errChan := make(chan error, 5)

	go t.tracker.StreamNowPlaying(ctx, npChan, errChan)

	cfg := t.cfgProvider.Get()
	if !t.lastfm.IsLoggedIn() {
		if err := t.lastfm.Login(ctx, cfg.UserName, cfg.Password); err != nil {
			t.notifyError("failed to login: %v", err)
		}
	}

	for {
		select {
		case np := <-npChan:
			t.processNowPlaying(ctx, np)
		case err := <-errChan:
			fmt.Printf("Error: %v\n", err)
			t.notifyError("error while getting nowplaying: %v", err)
		}
	}
}

func (t *trackNowPlayingImpl) processNowPlaying(ctx context.Context, np domain.NowPlaying) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if lo.IsEmpty(np) {
		t.app.EmitEvent(bindings.NotifyNowPlaying, nil, nil)
		return
	}

	cfg := t.cfgProvider.Get()
	shouldScrobble := t.lastfm.IsLoggedIn() && cfg.ScrobbleImmediately

	// Update nowplaying
	t.app.EmitEvent(bindings.NotifyNowPlaying, np.ToResponse(), nil)

	if shouldScrobble && !t.npPrev.IsNotified {
		if _, err := t.lastfm.UpdateNowPlaying(ctx, np); err != nil {
			t.notifyError("failed to update nowplaying: %v", err)
		}
		t.npPrev.IsNotified = true
	}

	// Update scrobble log
	percentage := int(np.Position / np.Duration * 100)
	if percentage >= cfg.PositionThreshold && !t.npPrev.IsAdded {
		track := domain.CreateTrackInfo(np, time.Now().UTC())

		// Scrobble
		if shouldScrobble {
			ret, err := t.lastfm.Scrobble(ctx, []domain.TrackInfo{track})
			if err != nil {
				t.notifyError("scrobbling failed: %v", err)
			} else {
				track = track.MarkAsScrobbled(time.Now().UTC())
			}

			// For debug
			b, _ := json.Marshal(ret)
			log.Println("track.scrobble response: ", string(b))
		}

		// Save
		if _, err := t.repo.Save(ctx, track); err != nil {
			t.notifyError("failed to save play log: %v", err)
			return
		}

		t.app.EmitEvent(bindings.NotifyAdded)
	}

	// Update npPrev
	if !np.Equals(t.npPrev) {
		t.npPrev = np
	}
	if percentage >= cfg.PositionThreshold {
		t.npPrev.IsAdded = true
	}
}

func (t *trackNowPlayingImpl) notifyError(format string, a ...any) {
	t.app.EmitEvent(bindings.NotifyNowPlaying, nil, bindings.NewInternalError(format, a...))
}
