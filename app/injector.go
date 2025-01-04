package app

import (
	"audirvana-scrobbler/app/internal/infra"
	"audirvana-scrobbler/app/internal/usecase"

	"github.com/samber/do"
)

func BuildInjector() *do.Injector {
	i := do.New()

	infra.Inject(i)
	usecase.Inject(i)
	do.Provide(i, NewApp)

	return i
}
