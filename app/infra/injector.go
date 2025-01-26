package infra

import (
	"audirvana-scrobbler/app/infra/internal/common"
	"audirvana-scrobbler/app/infra/internal/lastfm"
	"audirvana-scrobbler/app/infra/internal/trackinfo"

	"github.com/samber/do"
)

func Inject(i *do.Injector) {
	do.Provide(i, common.NewConfigPathProvider)
	do.Provide(i, common.NewDB)
	do.Provide(i, common.NewConfigProvider)

	do.Provide(i, trackinfo.NewAudirvanaUpdater)
	do.Provide(i, trackinfo.NewTrackInfoRepository)
	do.Provide(i, trackinfo.NewTrackInfoQueryService)
	do.Provide(i, lastfm.NewLastFMAPI)
}
