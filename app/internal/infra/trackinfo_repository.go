package infra

import (
	"audirvana-scrobbler/app/internal/domain"
	"audirvana-scrobbler/app/internal/infra/internal"
	"audirvana-scrobbler/app/shared/response"
	"context"
	"time"

	"github.com/samber/do"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type trackInfoRepositoryImpl struct {
	db *gorm.DB
}

func NewTrackInfoRepository(i *do.Injector) (domain.TrackInfoRepository, error) {
	return &trackInfoRepositoryImpl{db: do.MustInvoke[*gorm.DB](i)}, nil
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

func (r *trackInfoRepositoryImpl) Update(ctx context.Context, tracks []internal.TrackInfoDBSchema) error {
	for i, track := range tracks {
		tracks[i].ID = track.PlayedAt.Format(time.RFC3339)
	}

	if err := r.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}}, // IDが重複した場合
		DoNothing: true,
	}).Create(&tracks).Error; err != nil {
		return err
	}

	return nil
}
