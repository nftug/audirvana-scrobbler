package internal

import (
	"github.com/samber/do"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB(i *do.Injector) (*gorm.DB, error) {
	configPath := do.MustInvoke[ConfigPath](i)
	dbPath := configPath.GetJoinedPath("local_tracks.db")

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&TrackInfoDBSchema{}); err != nil {
		return nil, err
	}

	return db, nil
}
