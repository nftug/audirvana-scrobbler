package usecase

import (
	"audirvana-scrobbler/app/internal/domain"
	"audirvana-scrobbler/app/shared/customerr"
	"audirvana-scrobbler/app/shared/response"
	"context"
	"time"

	"github.com/samber/do"
)

type GetTrackInfo interface {
	Execute(ctx context.Context) ([]response.TrackInfo, error)
}

type getTrackInfoImpl struct {
	importer domain.AudirvanaImporter
}

func NewGetTrackInfo(i *do.Injector) (GetTrackInfo, error) {
	return &getTrackInfoImpl{
		importer: do.MustInvoke[domain.AudirvanaImporter](i),
	}, nil
}

func (g *getTrackInfoImpl) Execute(ctx context.Context) ([]response.TrackInfo, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	result, err := g.importer.GetAllTracks(ctx)
	if err != nil {
		return nil, customerr.NewInternalError("Error while getting track info: %v", err.Error())
	}

	return result, nil
}
