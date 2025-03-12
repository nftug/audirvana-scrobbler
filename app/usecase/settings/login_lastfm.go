package settings

import (
	"audirvana-scrobbler/app/domain"
	"context"

	"github.com/samber/do"
)

type LoginLastFM interface {
	Execute(ctx context.Context) bool
}

type loginLastFMImpl struct {
	lastfm      domain.LastFMAPI
	cfgProvider domain.ConfigProvider
}

func NewLoginLastFM(i *do.Injector) (LoginLastFM, error) {
	return &loginLastFMImpl{
		lastfm:      do.MustInvoke[domain.LastFMAPI](i),
		cfgProvider: do.MustInvoke[domain.ConfigProvider](i),
	}, nil
}

func (l *loginLastFMImpl) Execute(ctx context.Context) bool {
	cfg := l.cfgProvider.Get()
	if err := l.lastfm.LoginWithSessionKey(cfg.SessionKey); err != nil {
		return false
	}
	return true
}
