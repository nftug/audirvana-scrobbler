package settings

import (
	"audirvana-scrobbler/app/bindings"
	"audirvana-scrobbler/app/domain"
	"context"

	"github.com/samber/do"
)

type LogoutLastFM interface {
	Execute(ctx context.Context) *bindings.ErrorResponse
}

type logoutLastFMImpl struct {
	lastfm      domain.LastFMAPI
	cfgProvider domain.ConfigProvider
}

func NewLogoutLastFM(i *do.Injector) (LogoutLastFM, error) {
	return &logoutLastFMImpl{
		lastfm:      do.MustInvoke[domain.LastFMAPI](i),
		cfgProvider: do.MustInvoke[domain.ConfigProvider](i),
	}, nil
}

func (l *logoutLastFMImpl) Execute(ctx context.Context) *bindings.ErrorResponse {
	cfg := l.cfgProvider.Get()
	cfg.SessionKey = ""
	if err := l.cfgProvider.Write(cfg); err != nil {
		return bindings.NewInternalError("failed to logout: %v", err)
	}

	l.lastfm.RemoveSessionKey()
	return nil
}
