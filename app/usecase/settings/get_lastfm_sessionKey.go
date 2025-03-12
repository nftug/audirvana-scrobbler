package settings

import (
	"audirvana-scrobbler/app/bindings"
	"audirvana-scrobbler/app/domain"
	"context"

	"github.com/samber/do"
)

type GetLastFMSessionKey interface {
	Execute(ctx context.Context, username, password string) *bindings.ErrorResponse
}

type getLastFMSessionKey struct {
	lastfm      domain.LastFMAPI
	cfgProvider domain.ConfigProvider
}

func NewGetLastFMSessionKey(i *do.Injector) (GetLastFMSessionKey, error) {
	return &getLastFMSessionKey{
		lastfm:      do.MustInvoke[domain.LastFMAPI](i),
		cfgProvider: do.MustInvoke[domain.ConfigProvider](i),
	}, nil
}

func (l *getLastFMSessionKey) Execute(ctx context.Context, username, password string) *bindings.ErrorResponse {
	sessionKey, err := l.lastfm.GetSessionKey(ctx, username, password)
	if err != nil {
		return bindings.NewInternalError("Failed to login.")
	}

	cfg := l.cfgProvider.Get()
	cfg.SessionKey = sessionKey
	l.cfgProvider.Write(cfg)

	l.lastfm.LoginWithSessionKey(sessionKey)
	return nil
}
