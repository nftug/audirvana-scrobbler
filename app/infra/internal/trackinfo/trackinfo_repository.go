package trackinfo

import (
	"audirvana-scrobbler/app/domain"
	"audirvana-scrobbler/app/lib/option"
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

func (r *trackInfoRepositoryImpl) Get(ctx context.Context, id int) (
	option.Option[domain.TrackInfo], error) {
	ret := TrackInfoDBSchema{ID: id}
	if err := r.db.NewSelect().
		Model(&ret).
		WherePK().
		Where("deleted_at IS NULL").
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return option.None[domain.TrackInfo](), nil
		}
		return option.None[domain.TrackInfo](), err
	}

	return option.Some(ret.ToEntity()), nil
}

func (r *trackInfoRepositoryImpl) Save(ctx context.Context, entity domain.TrackInfo) (domain.TrackInfo, error) {
	col := NewTrackInfoDBSchema(entity)
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

	if len(cols) == 0 {
		return entities, nil
	}

	if _, err := r.db.NewInsert().
		Model(&cols).
		On("CONFLICT (id) DO UPDATE").
		Set("artist = EXCLUDED.artist").
		Set("album = EXCLUDED.album").
		Set("track = EXCLUDED.track").
		Set("duration = EXCLUDED.duration").
		Set("played_at = EXCLUDED.played_at").
		Set("scrobbled_at = EXCLUDED.scrobbled_at").
		Exec(ctx); err != nil {
		return entities, err
	}

	return lo.Map(cols, func(t TrackInfoDBSchema, _ int) domain.TrackInfo {
		return t.ToEntity()
	}), nil
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
