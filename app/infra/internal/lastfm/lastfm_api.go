package lastfm

import (
	"audirvana-scrobbler/app/domain"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/samber/do"
	"github.com/samber/lo"
)

type lastFMAPIImpl struct {
	cfg        lastFMConfig
	sessionKey string
}

func NewLastFMAPI(i *do.Injector) (domain.LastFMAPI, error) {
	cfgProvider := do.MustInvoke[domain.ConfigProvider](i)
	cfg := cfgProvider.Get()
	if cfg.APIKey == "" {
		return nil, errors.New("api key is empty")
	}
	if cfg.APISecret == "" {
		return nil, errors.New("api secret is empty")
	}

	return &lastFMAPIImpl{
		cfg: lastFMConfig{apiKey: cfg.APIKey, apiSecret: cfg.APISecret},
	}, nil
}

func (l *lastFMAPIImpl) IsLoggedIn() bool { return l.sessionKey != "" }

func (l *lastFMAPIImpl) Login(ctx context.Context, username, password string) error {
	if username == "" {
		return errors.New("username is empty")
	}
	if password == "" {
		return errors.New("password is empty")
	}

	params := map[string]string{
		"username": username,
		"password": password,
	}

	result, err := l.callPostWithoutSk(ctx, "auth.getmobilesession", params)
	if err != nil {
		return err
	}

	session, ok := result["session"].(map[string]any)
	if !ok {
		errBody, _ := json.Marshal(result)
		return fmt.Errorf("invalid response: %v", string(errBody))
	}
	key, ok := session["key"].(string)
	if !ok {
		return errors.New("session key not found")
	}

	l.sessionKey = key

	return nil
}

func (l *lastFMAPIImpl) Scrobble(ctx context.Context, tracks []domain.TrackInfo) (map[string]any, error) {
	// ScrobbleされていないTracksだけを選択
	tracks = lo.Filter(tracks, func(t domain.TrackInfo, _ int) bool { return t.ScrobbledAt() == nil })

	if len(tracks) > 50 {
		return nil, errors.New("number of tracks is more than 50")
	}

	params := map[string]string{}
	for i, track := range tracks {
		params[fmt.Sprintf("artist[%d]", i)] = track.Artist()
		params[fmt.Sprintf("track[%d]", i)] = track.Track()
		params[fmt.Sprintf("album[%d]", i)] = track.Album()
		params[fmt.Sprintf("timestamp[%d]", i)] = strconv.FormatInt(track.PlayedAt().Unix(), 10)
	}

	result, err := l.callPostWithSk(ctx, "track.scrobble", params)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (l *lastFMAPIImpl) UpdateNowPlaying(ctx context.Context, np domain.NowPlaying) (map[string]any, error) {
	params := map[string]string{
		"artist":   np.Artist,
		"track":    np.Track,
		"album":    np.Album,
		"duration": strconv.Itoa(int(np.Duration)),
	}

	result, err := l.callPostWithSk(ctx, "track.updateNowPlaying", params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
