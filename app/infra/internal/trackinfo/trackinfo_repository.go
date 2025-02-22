package trackinfo

import (
	"audirvana-scrobbler/app/domain"
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

func (r *trackInfoRepositoryImpl) Get(ctx context.Context, id int) (domain.TrackInfo, error) {
	var ret TrackInfoDBSchema

	query := r.db.WithContext(ctx).Model(&TrackInfoDBSchema{}).Where("scrobbled_at IS NULL")
	if err := query.First(&ret).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return domain.TrackInfo{}, nil
		default:
			return domain.TrackInfo{}, err
		}
	}

	return ret.ToEntity(), nil
}

func (r *trackInfoRepositoryImpl) Save(ctx context.Context, entity domain.TrackInfo) (domain.TrackInfo, error) {
	col := NewTrackInfoDBSchema(entity)
	err := r.db.WithContext(ctx).Model(&TrackInfoDBSchema{}).Save(&col).Error
	if err != nil {
		return entity, err
	}
	return col.ToEntity(), nil
}

func (r *trackInfoRepositoryImpl) MarkAsScrobbled(ctx context.Context, entities []domain.TrackInfo) (
	[]domain.TrackInfo, error) {
	ids := lo.Map(entities, func(t domain.TrackInfo, _ int) int { return t.ID() })
	scrobbledAt := time.Now().UTC()

	err := r.db.WithContext(ctx).Model(&TrackInfoDBSchema{}).
		Where("id IN ?", ids).
		Update("scrobbled_at", scrobbledAt).Error
	if err != nil {
		return entities, err
	}

	for i, entity := range entities {
		entities[i] = entity.MarkAsScrobbled(scrobbledAt)
	}

	return entities, nil
}

func (r *trackInfoRepositoryImpl) Delete(ctx context.Context, id int) error {
	err := r.db.WithContext(ctx).Delete(&TrackInfoDBSchema{ID: id}).Error
	if err != nil {
		return err
	}
	return nil
}
