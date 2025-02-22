package trackinfo

import (
	"audirvana-scrobbler/app/bindings"
	"audirvana-scrobbler/app/domain"
	"context"

	"github.com/samber/do"
)

type GetTrackInfoList interface {
	Execute(ctx context.Context) ([]bindings.TrackInfoResponse, *bindings.ErrorResponse)
}

type getTrackInfoListImpl struct {
	qs domain.TrackInfoQueryService
}

func NewGetTrackInfoList(i *do.Injector) (GetTrackInfoList, error) {
	return &getTrackInfoListImpl{
		qs: do.MustInvoke[domain.TrackInfoQueryService](i),
	}, nil
}

func (g getTrackInfoListImpl) Execute(ctx context.Context) ([]bindings.TrackInfoResponse, *bindings.ErrorResponse) {
	result, err := g.qs.GetAll(ctx)
	if err != nil {
		return nil, bindings.NewInternalError("Error while getting scrobble log: %v", err.Error())
	}
	return result, nil
}
