package app

import (
	"audirvana-scrobbler/app/infra"
	"audirvana-scrobbler/app/usecase"

	"github.com/samber/do"
)

func BuildInjector() *do.Injector {
	i := do.New()

	infra.Inject(i)
	usecase.Inject(i)
	do.Provide(i, NewApp)

	return i
}
