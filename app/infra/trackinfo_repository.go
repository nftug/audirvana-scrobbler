package infra

import (
	"audirvana-scrobbler/app/bindings"
	"audirvana-scrobbler/app/domain"
	"audirvana-scrobbler/app/infra/internal"
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

func (r *trackInfoRepositoryImpl) GetAll(ctx context.Context) ([]bindings.TrackInfo, error) {
	var ret []internal.TrackInfoDBSchema

	query := r.db.WithContext(ctx).
		Where("scrobbled_at IS NULL").Where("deleted_at IS NULL")
	if err := query.Order("played_at").Find(&ret).Error; err != nil {
		return nil, err
	}

	tracks := lo.Map(ret, func(x internal.TrackInfoDBSchema, _ int) bindings.TrackInfo {
		return x.ToResponse()
	})
	return tracks, nil
}
