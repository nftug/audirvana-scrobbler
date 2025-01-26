package trackinfo

import (
	"audirvana-scrobbler/app/domain"
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
	return &trackInfoRepositoryImpl{
		db: do.MustInvoke[*gorm.DB](i),
	}, nil
}

func (r *trackInfoRepositoryImpl) GetAll(ctx context.Context) ([]domain.TrackInfo, error) {
	var ret []TrackInfoDBSchema

	query := r.db.WithContext(ctx).Model(&TrackInfoDBSchema{}).Where("scrobbled_at IS NULL")
	if err := query.Order("played_at").Find(&ret).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return []domain.TrackInfo{}, nil
		}
		return nil, err
	}

	tracks := lo.Map(ret, func(x TrackInfoDBSchema, _ int) domain.TrackInfo {
		return *x.ToEntity()
	})
	return tracks, nil
}

func (r *trackInfoRepositoryImpl) GetLatestPlayedAt(ctx context.Context) (time.Time, error) {
	var col TrackInfoDBSchema

	query := r.db.WithContext(ctx).Model(&col).Unscoped().Order("played_at desc")
	if err := query.First(&col).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return time.Now().UTC(), nil
		default:
			return time.Time{}, err
		}
	}

	return col.PlayedAt, nil
}

func (r *trackInfoRepositoryImpl) Get(ctx context.Context, id string) (*domain.TrackInfo, error) {
	var ret TrackInfoDBSchema

	query := r.db.WithContext(ctx).Model(&TrackInfoDBSchema{}).Where("scrobbled_at IS NULL")
	if err := query.First(&ret).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, nil
		default:
			return nil, err
		}
	}

	return ret.ToEntity(), nil
}

func (r *trackInfoRepositoryImpl) Save(ctx context.Context, entity *domain.TrackInfo) error {
	col := TrackInfoDBSchema{
		Artist: entity.Artist(),
		Album:  entity.Album(),
		Track:  entity.Track(),
	}
	err := r.db.WithContext(ctx).Model(&TrackInfoDBSchema{ID: entity.ID()}).Updates(col).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *trackInfoRepositoryImpl) MarkAsScrobbled(ctx context.Context, entities []domain.TrackInfo) error {
	ids := lo.Map(entities, func(t domain.TrackInfo, _ int) string { return t.ID() })
	scrobbledAt := time.Now().UTC()

	err := r.db.WithContext(ctx).Model(&TrackInfoDBSchema{}).
		Where("id IN ?", ids).
		Update("scrobbled_at", scrobbledAt).Error
	if err != nil {
		return err
	}

	for i, entity := range entities {
		entities[i] = *entity.MarkAsScrobbled(scrobbledAt)
	}

	return nil
}

func (r *trackInfoRepositoryImpl) Delete(ctx context.Context, id string) error {
	err := r.db.WithContext(ctx).Delete(&TrackInfoDBSchema{ID: id}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *trackInfoRepositoryImpl) CreateRange(ctx context.Context, tracks []domain.TrackInfo) error {
	if len(tracks) == 0 {
		return nil
	}

	items := lo.Map(
		tracks, func(t domain.TrackInfo, _ int) TrackInfoDBSchema { return NewTrackInfoDBSchema(t) })

	if err := r.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoNothing: true,
	}).Create(&items).Error; err != nil {
		return err
	}
	return nil
}
