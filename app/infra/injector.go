package infra

import (
	"audirvana-scrobbler/app/infra/internal"

	"github.com/samber/do"
)

func Inject(i *do.Injector) {
	do.Provide(i, internal.NewConfigPath)
	do.Provide(i, internal.NewDB)
	do.Provide(i, internal.NewConfigProvider)

	do.Provide(i, NewAudirvanaUpdater)
	do.Provide(i, NewTrackInfoRepository)
	do.Provide(i, NewScrobbler)
}
