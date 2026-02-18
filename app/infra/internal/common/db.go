package common

import (
	"audirvana-scrobbler/app/infra/internal/trackinfo"
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/samber/do"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
)

func NewDB(i *do.Injector) (*bun.DB, error) {
	configPath := do.MustInvoke[*ConfigPathProvider](i)
	dbPath := configPath.GetJoinedPath("local_tracks.db")

	sqldb, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	db := bun.NewDB(sqldb, sqlitedialect.New())

	if _, err := db.NewCreateTable().
		Model((*trackinfo.TrackInfoDBSchema)(nil)).
		IfNotExists().
		Exec(context.Background()); err != nil {
		_ = sqldb.Close()
		return nil, err
	}

	return db, nil
}
