package trackinfo

import (
	"audirvana-scrobbler/app/domain"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/samber/do"
	"github.com/samber/lo"
	"github.com/uptrace/bun"
)

type trackInfoRepositoryImpl struct {
	db *bun.DB
}

func NewTrackInfoRepository(i *do.Injector) (domain.TrackInfoRepository, error) {
	return &trackInfoRepositoryImpl{
		db: do.MustInvoke[*bun.DB](i),
	}, nil
}

func (r *trackInfoRepositoryImpl) GetAll(ctx context.Context) ([]domain.TrackInfo, error) {
	var ret []TrackInfoDBSchema

	if err := r.db.NewSelect().
		Model(&ret).
		Where("deleted_at IS NULL").
		Order("played_at").
		Scan(ctx); err != nil {
		return nil, err
	}

	tracks := lo.Map(ret, func(x TrackInfoDBSchema, _ int) domain.TrackInfo {
		return x.ToEntity()
	})
	return tracks, nil
}

func (r *trackInfoRepositoryImpl) Get(ctx context.Context, id int) (*domain.TrackInfo, error) {
	ret := TrackInfoDBSchema{ID: id}
	if err := r.db.NewSelect().
		Model(&ret).
		WherePK().
		Where("deleted_at IS NULL").
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return new(ret.ToEntity()), nil
}

func (r *trackInfoRepositoryImpl) Save(ctx context.Context, entity *domain.TrackInfo) error {
	col := NewTrackInfoDBSchema(*entity)
	_, err := r.db.NewInsert().
		Model(&col).
		On("CONFLICT (id) DO UPDATE").
		Set("artist = EXCLUDED.artist").
		Set("album = EXCLUDED.album").
		Set("track = EXCLUDED.track").
		Set("duration = EXCLUDED.duration").
		Set("played_at = EXCLUDED.played_at").
		Set("scrobbled_at = EXCLUDED.scrobbled_at").
		Exec(ctx)
	return err
}

func (r *trackInfoRepositoryImpl) SaveRange(ctx context.Context, entities []domain.TrackInfo) error {
	cols := lo.Map(entities, func(t domain.TrackInfo, _ int) TrackInfoDBSchema {
		return NewTrackInfoDBSchema(t)
	})

	if len(cols) == 0 {
		return nil
	}

	_, err := r.db.NewInsert().
		Model(&cols).
		On("CONFLICT (id) DO UPDATE").
		Set("artist = EXCLUDED.artist").
		Set("album = EXCLUDED.album").
		Set("track = EXCLUDED.track").
		Set("duration = EXCLUDED.duration").
		Set("played_at = EXCLUDED.played_at").
		Set("scrobbled_at = EXCLUDED.scrobbled_at").
		Exec(ctx)
	return err
}

func (r *trackInfoRepositoryImpl) Delete(ctx context.Context, id int) error {
	_, err := r.db.NewUpdate().
		Model(&TrackInfoDBSchema{}).
		Set("deleted_at = ?", time.Now().UTC()).
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
