package infra

import (
	"audirvana-scrobbler/app/internal/domain"
	"audirvana-scrobbler/app/internal/infra/internal"
	"audirvana-scrobbler/app/shared/response"
	"context"

	"github.com/samber/do"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type trackInfoRepositoryImpl struct {
	db *gorm.DB
}

func NewTrackInfoRepository(i *do.Injector) (domain.TrackInfoRepository, error) {
	return &trackInfoRepositoryImpl{
		db: do.MustInvoke[*gorm.DB](i),
	}, nil
}

func (r *trackInfoRepositoryImpl) GetAll(ctx context.Context) ([]response.TrackInfo, error) {
	var ret []internal.TrackInfoDBSchema

	query := r.db.WithContext(ctx).Where("scrobbled_at IS NULL")
	if err := query.Order("played_at").Find(&ret).Error; err != nil {
		return nil, err
	}

	tracks := lo.Map(ret, func(x internal.TrackInfoDBSchema, _ int) response.TrackInfo {
		return x.ToResponse()
	})
	return tracks, nil
}
