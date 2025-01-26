package common

import (
	"audirvana-scrobbler/app/infra/internal/trackinfo"

	"github.com/glebarez/sqlite"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func NewDB(i *do.Injector) (*gorm.DB, error) {
	configPath := do.MustInvoke[*ConfigPathProvider](i)
	dbPath := configPath.GetJoinedPath("local_tracks.db")

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&trackinfo.TrackInfoDBSchema{}); err != nil {
		return nil, err
	}

	return db, nil
}
