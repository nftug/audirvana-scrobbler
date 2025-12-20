package trackinfo

import (
	"audirvana-scrobbler/app/domain"
	"audirvana-scrobbler/app/lib/option"
	"context"

	"github.com/samber/do"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type trackInfoRepositoryImpl struct {
	db *gorm.DB
}

func NewTrackInfoRepository(i *do.Injector) (domain.TrackInfoRepository, error) {
	return &trackInfoRepositoryImpl{
		db: do.MustInvoke[*gorm.DB](i),
	}, nil
}

func (r *trackInfoRepositoryImpl) GetAll(ctx context.Context) ([]domain.TrackInfo, error) {
	var ret []TrackInfoDBSchema

	query := r.db.WithContext(ctx).Model(&TrackInfoDBSchema{})
	if err := query.Order("played_at").Find(&ret).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return []domain.TrackInfo{}, nil
		}
		return nil, err
	}

	tracks := lo.Map(ret, func(x TrackInfoDBSchema, _ int) domain.TrackInfo {
		return x.ToEntity()
	})
	return tracks, nil
}

func (r *trackInfoRepositoryImpl) Get(ctx context.Context, id int) (
	option.Option[domain.TrackInfo], error) {
	var ret TrackInfoDBSchema

	query := r.db.WithContext(ctx).Model(&TrackInfoDBSchema{})
	if err := query.First(&ret).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return option.None[domain.TrackInfo](), nil
		default:
			return option.None[domain.TrackInfo](), err
		}
	}

	return option.Some(ret.ToEntity()), nil
}

func (r *trackInfoRepositoryImpl) Save(ctx context.Context, entity domain.TrackInfo) (domain.TrackInfo, error) {
	col := NewTrackInfoDBSchema(entity)
	err := r.db.WithContext(ctx).Model(&TrackInfoDBSchema{}).Save(&col).Error
	if err != nil {
		return entity, err
	}
	return col.ToEntity(), nil
}

func (r *trackInfoRepositoryImpl) SaveRange(ctx context.Context, entities []domain.TrackInfo) (
	[]domain.TrackInfo, error) {
	cols := lo.Map(entities, func(t domain.TrackInfo, _ int) TrackInfoDBSchema {
		return NewTrackInfoDBSchema(t)
	})

	if err := r.db.WithContext(ctx).Model(&TrackInfoDBSchema{}).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		UpdateAll: true,
	}).Create(&cols).Error; err != nil {
		return entities, err
	}

	return lo.Map(cols, func(t TrackInfoDBSchema, _ int) domain.TrackInfo {
		return t.ToEntity()
	}), nil
}

func (r *trackInfoRepositoryImpl) Delete(ctx context.Context, id int) error {
	err := r.db.WithContext(ctx).Delete(&TrackInfoDBSchema{ID: id}).Error
	if err != nil {
		return err
	}
	return nil
}
