package trackinfo

import (
	"audirvana-scrobbler/app/domain"
	"context"
	"embed"
	"encoding/json"
	"os/exec"
	"time"

	"github.com/samber/do"
)

//go:embed script/*
var script embed.FS

type nowPlayingTrackerImpl struct{}

func NewNowPlayingTracker(i *do.Injector) (domain.NowPlayingTracker, error) {
	return &nowPlayingTrackerImpl{}, nil
}

func (n *nowPlayingTrackerImpl) StreamNowPlaying(
	ctx context.Context, npChan chan<- *domain.NowPlaying, errChan chan<- error) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			close(npChan)
			close(errChan)
			return
		case <-ticker.C:
			np, err := getNowPlaying(ctx)
			if err != nil {
				errChan <- err
				continue
			}
			npChan <- np
		}
	}
}

func getNowPlaying(ctx context.Context) (*domain.NowPlaying, error) {
	scriptByte, _ := script.ReadFile("script/getNowPlaying.applescript")
	ret, err := exec.CommandContext(ctx, "osascript", "-e", string(scriptByte)).Output()
	if err != nil {
		return nil, err
	}

	var np domain.NowPlaying
	if err := json.Unmarshal(ret, &np); err != nil {
		return nil, nil
	}

	return &np, nil
}
