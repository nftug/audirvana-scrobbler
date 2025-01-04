package infra

import "github.com/samber/do"

func Inject(i *do.Injector) {
	do.Provide(i, NewConfigPath)
	do.Provide(i, NewTempPath)
	do.Provide(i, NewAudirvanaImporter)
}
