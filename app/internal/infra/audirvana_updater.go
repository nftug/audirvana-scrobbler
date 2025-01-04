package infra

import (
	"audirvana-scrobbler/app/internal/domain"
	"audirvana-scrobbler/app/internal/infra/internal"
	"context"
	"database/sql"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/kirsle/configdir"
	"github.com/samber/do"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	_ "github.com/mattn/go-sqlite3"
)

const LAST_SCROBBLE_TIME_LOG = "lastScrobbleTime.log"

type audirvanaUpdaterImpl struct {
	configpath internal.ConfigPath
	tempPath   internal.TempPath
	db         *gorm.DB
}

func NewAudirvanaUpdater(i *do.Injector) (domain.AudirvanaUpdater, error) {
	return &audirvanaUpdaterImpl{
		configpath: do.MustInvoke[internal.ConfigPath](i),
		tempPath:   do.MustInvoke[internal.TempPath](i),
		db:         do.MustInvoke[*gorm.DB](i),
	}, nil
}

func (a *audirvanaUpdaterImpl) Update(ctx context.Context) error {
	dbPath, err := a.getTempDBPath()
	if err != nil {
		return err
	}

	// Get last scrobble time
	lastScrobbleTime, err := a.readLastScrobbleTime()
	if err != nil {
		return err
	}

	if err := a.updateCore(ctx, dbPath, lastScrobbleTime); err != nil {
		return err
	}

	return nil
}

func (a *audirvanaUpdaterImpl) getTempDBPath() (string, error) {
	audirvanaConfigPath := configdir.LocalConfig("Audirvana")
	originalDBPath := filepath.Join(audirvanaConfigPath, "AudirvanaDatabase.sqlite")
	tempDBPath := a.tempPath.GetJoinedPath("AudirvanaDatabase.sqlite")

	src, err := os.Open(originalDBPath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer src.Close()

	dst, err := os.Create(tempDBPath)
	if err != nil {
		return "", fmt.Errorf("failed to create local DB file: %w", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("failed to copy DB file to local: %w", err)
	}

	return tempDBPath, nil
}

func (a *audirvanaUpdaterImpl) readLastScrobbleTime() (string, error) {
	logPath := a.configpath.GetJoinedPath(LAST_SCROBBLE_TIME_LOG)
	file, err := os.OpenFile(logPath, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return "", fmt.Errorf("failed to open last scrobble time log: %w", err)
	}
	defer file.Close()

	buf := make([]byte, 19)
	n, _ := file.Read(buf)

	lastScrobble := lo.Ternary(n == 0, "2459364.9074768517", strings.TrimSpace(string(buf[:n])))
	return lastScrobble, nil
}

func (a *audirvanaUpdaterImpl) updateCore(ctx context.Context, dbFilePath string, lastScrobbleTime string) error {
	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?mode=ro", dbFilePath))
	if err != nil {
		return err
	}
	defer db.Close()

	query := `
		SELECT ARTISTS.name AS artist, ALBUMS.title AS album, TRACKS.title,
				date(TRACKS.last_played_date) AS date, time(TRACKS.last_played_date) AS time
		FROM TRACKS
		JOIN ALBUMS ON ALBUMS.album_id = TRACKS.album_id
		JOIN ALBUMS_ARTISTS ON ALBUMS_ARTISTS.album_id = ALBUMS.album_id
		JOIN ARTISTS ON ALBUMS_ARTISTS.artist_id = ARTISTS.artist_id
		WHERE round(last_played_date, 4) > round(?, 4)
		ORDER BY last_played_date ASC
	`

	rows, err := db.QueryContext(ctx, query, lastScrobbleTime)
	if err != nil {
		return err
	}
	defer rows.Close()

	var newTracks []internal.TrackInfoDBSchema
	for rows.Next() {
		var track internal.TrackInfoDBSchema
		var playedDate, playedTime string
		if err := rows.Scan(&track.Artist, &track.Album, &track.Track, &playedDate, &playedTime); err != nil {
			return err
		}

		track.PlayedAt, err = time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s %s", playedDate, playedTime))
		if err != nil {
			return err
		}

		track.ID = track.PlayedAt.Format(time.RFC3339)
		newTracks = append(newTracks, track)
	}

	// Update local DB
	if err := a.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoNothing: true,
	}).Create(&newTracks).Error; err != nil {
		return err
	}

	return nil
}
