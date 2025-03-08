package trackinfo

import (
	"audirvana-scrobbler/app/bindings"
	"audirvana-scrobbler/app/domain"
	"context"

	"github.com/samber/do"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type trackInfoQueryServiceImpl struct {
	db *gorm.DB
}

func NewTrackInfoQueryService(i *do.Injector) (domain.TrackInfoQueryService, error) {
	return &trackInfoQueryServiceImpl{
		db: do.MustInvoke[*gorm.DB](i),
	}, nil
}

func (q *trackInfoQueryServiceImpl) GetAll(ctx context.Context) ([]bindings.TrackInfoResponse, error) {
	var ret []TrackInfoDBSchema

	query := q.db.WithContext(ctx).Model(&TrackInfoDBSchema{})
	if err := query.Order("played_at DESC").Find(&ret).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return []bindings.TrackInfoResponse{}, nil
		}
		return nil, err
	}

	tracks := lo.Map(ret, func(t TrackInfoDBSchema, _ int) bindings.TrackInfoResponse {
		return t.ToResponse()
	})
	return tracks, nil
}
