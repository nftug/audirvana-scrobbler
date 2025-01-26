package usecase

import (
	"audirvana-scrobbler/app/usecase/trackinfo"

	"github.com/samber/do"
)

func Inject(i *do.Injector) {
	do.Provide(i, trackinfo.NewGetTrackInfoList)
	do.Provide(i, trackinfo.NewSaveTrackInfo)
	do.Provide(i, trackinfo.NewDeleteTrackInfo)
	do.Provide(i, trackinfo.NewScrobbleAll)
	do.Provide(i, trackinfo.NewTrackNowPlaying)
}
