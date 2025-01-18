package usecase

import "github.com/samber/do"

func Inject(i *do.Injector) {
	do.Provide(i, NewGetTrackInfo)
	do.Provide(i, NewSaveTrackInfo)
}
