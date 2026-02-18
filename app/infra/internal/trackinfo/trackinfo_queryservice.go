package trackinfo

import (
	"audirvana-scrobbler/app/bindings"
	"audirvana-scrobbler/app/domain"
	"context"

	"github.com/samber/do"
	"github.com/samber/lo"
	"github.com/uptrace/bun"
)

type trackInfoQueryServiceImpl struct {
	db *bun.DB
}

func NewTrackInfoQueryService(i *do.Injector) (domain.TrackInfoQueryService, error) {
	return &trackInfoQueryServiceImpl{
		db: do.MustInvoke[*bun.DB](i),
	}, nil
}

func (q *trackInfoQueryServiceImpl) GetAll(ctx context.Context) ([]bindings.TrackInfoResponse, error) {
	var ret []TrackInfoDBSchema

	if err := q.db.NewSelect().
		Model(&ret).
		Where("deleted_at IS NULL").
		Order("played_at DESC").
		Scan(ctx); err != nil {
		return nil, err
	}

	tracks := lo.Map(ret, func(t TrackInfoDBSchema, _ int) bindings.TrackInfoResponse {
		return t.ToResponse()
	})
	return tracks, nil
}
