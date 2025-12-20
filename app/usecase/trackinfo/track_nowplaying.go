package trackinfo

import (
	"audirvana-scrobbler/app/bindings"
	"audirvana-scrobbler/app/domain"
	"audirvana-scrobbler/app/lib/option"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"sync"

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

	npChan := make(chan option.Option[domain.NowPlaying], 5)
	errChan := make(chan error, 5)

	go t.tracker.StreamNowPlaying(ctx, npChan, errChan)

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

func (t *trackNowPlayingImpl) processNowPlaying(ctx context.Context, npOpt option.Option[domain.NowPlaying]) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	np, ok := npOpt.TryUnwrap()
	if !ok {
		t.app.EmitEvent(bindings.NotifyNowPlaying, nil, nil)
		return
	}

	cfg := t.cfgProvider.Get()
	shouldScrobble := t.lastfm.IsLoggedIn() && cfg.ScrobbleImmediately

	// Update now playing
	t.app.EmitEvent(bindings.NotifyNowPlaying, np.ToResponse(), nil)

	if shouldScrobble && !t.npPrev.IsNotified {
		if _, err := t.lastfm.UpdateNowPlaying(ctx, np); err != nil {
			t.notifyError("failed to update now playing: %v", err)
		}
		t.npPrev.IsNotified = true
	}

	// Update scrobble log
	percentage := int(np.Position / np.Duration * 100)
	shouldAdd := percentage >= cfg.PositionThreshold && !t.npPrev.IsAdded

	if shouldAdd {
		track := domain.CreateTrackInfo(np, time.Now().UTC())

		// Scrobble
		if shouldScrobble {
			ret, err := t.lastfm.Scrobble(ctx, []domain.TrackInfo{track})
			if err != nil {
				t.notifyError("scrobble failed: %v", err)
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
	if shouldAdd {
		t.npPrev.IsAdded = true
	}
}

func (t *trackNowPlayingImpl) notifyError(format string, a ...any) {
	t.app.EmitEvent(bindings.NotifyNowPlaying, nil, bindings.NewInternalError(format, a...))
}
