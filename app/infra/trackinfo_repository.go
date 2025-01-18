package infra

import (
	"audirvana-scrobbler/app/bindings"
	"audirvana-scrobbler/app/domain"
	"audirvana-scrobbler/app/infra/internal"
	"context"
	"time"

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

	query := r.db.WithContext(ctx).Model(&internal.TrackInfoDBSchema{}).
		Where("scrobbled_at IS NULL").Where("deleted_at IS NULL")
	if err := query.Order("played_at").Find(&ret).Error; err != nil {
		return nil, err
	}

	tracks := lo.Map(ret, func(x internal.TrackInfoDBSchema, _ int) bindings.TrackInfo {
		return x.ToResponse()
	})
	return tracks, nil
}

func (r *trackInfoRepositoryImpl) GetLatestPlayedAt(ctx context.Context) (time.Time, error) {
	var col internal.TrackInfoDBSchema

	query := r.db.WithContext(ctx).Model(&col).Unscoped().Order("played_at desc")
	if err := query.First(&col).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return time.Now().UTC(), nil
		} else {
			return time.Time{}, err
		}
	}

	return col.PlayedAt, nil
}

func (r *trackInfoRepositoryImpl) Save(ctx context.Context, id string, form bindings.TrackInfoForm) error {
	col := internal.TrackInfoDBSchema{
		Artist: form.Artist,
		Album:  form.Album,
		Track:  form.Track,
	}
	if err := r.db.Model(&internal.TrackInfoDBSchema{ID: id}).Updates(col).Error; err != nil {
		return err
	}
	return nil
}
