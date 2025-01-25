package infra

import (
	"audirvana-scrobbler/app/domain"
	"audirvana-scrobbler/app/infra/internal"
	"context"
	"database/sql"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/kirsle/configdir"
	"github.com/samber/do"
	"github.com/samber/lo"
	"github.com/soniakeys/meeus/v3/julian"
)

type audirvanaUpdaterImpl struct {
	configpath domain.ConfigPathProvider
	repo       domain.TrackInfoRepository
}

func NewAudirvanaUpdater(i *do.Injector) (domain.AudirvanaUpdater, error) {
	return &audirvanaUpdaterImpl{
		configpath: do.MustInvoke[domain.ConfigPathProvider](i),
		repo:       do.MustInvoke[domain.TrackInfoRepository](i),
	}, nil
}

func (a *audirvanaUpdaterImpl) Update(ctx context.Context) (err error) {
	tempDirPath, err := os.MkdirTemp("", "audirvana-scrobbler")
	if err != nil {
		return err
	}
	defer func() {
		if err = os.RemoveAll(tempDirPath); err != nil {
			err = fmt.Errorf("error deleting temp directory: %w", err)
		}
	}()

	dbPath, err := a.getCopiedDBPath(tempDirPath)
	if err != nil {
		return err
	}

	if err := a.updateCore(ctx, dbPath); err != nil {
		return err
	}

	return nil
}

func (a *audirvanaUpdaterImpl) getCopiedDBPath(tempDirPath string) (string, error) {
	audirvanaConfigPath := configdir.LocalConfig("Audirvana")
	originalDBPath := filepath.Join(audirvanaConfigPath, "AudirvanaDatabase.sqlite")
	destDBPath := filepath.Join(tempDirPath, "AudirvanaDatabase.sqlite")

	src, err := os.Open(originalDBPath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer src.Close()

	dst, err := os.Create(destDBPath)
	if err != nil {
		return "", fmt.Errorf("failed to create local DB file: %w", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("failed to copy DB file to local: %w", err)
	}

	return destDBPath, nil
}

func (a *audirvanaUpdaterImpl) updateCore(ctx context.Context, dbFilePath string) error {
	lastPlayedAt, err := a.repo.GetLatestPlayedAt(ctx)
	if err != nil {
		return err
	}

	db, err := sql.Open("sqlite", fmt.Sprintf("file:%s?mode=ro", dbFilePath))
	if err != nil {
		return err
	}
	defer db.Close()

	query := `
		SELECT ARTISTS.name AS artist, ALBUMS.title AS album, TRACKS.title, TRACKS.last_played_date
		FROM TRACKS
		JOIN ALBUMS ON ALBUMS.album_id = TRACKS.album_id
		JOIN ALBUMS_ARTISTS ON ALBUMS_ARTISTS.album_id = ALBUMS.album_id
		JOIN ARTISTS ON ALBUMS_ARTISTS.artist_id = ARTISTS.artist_id
		WHERE round(last_played_date, 4) > round(?, 4)
		ORDER BY last_played_date ASC
	`

	rows, err := db.QueryContext(ctx, query, julian.TimeToJD(lastPlayedAt))
	if err != nil {
		return err
	}
	defer rows.Close()

	var newTracks []internal.TrackInfoDBSchema
	for rows.Next() {
		var track internal.TrackInfoDBSchema
		var playedAt float64
		if err := rows.Scan(&track.Artist, &track.Album, &track.Track, &playedAt); err != nil {
			return err
		}
		track.PlayedAt = julian.JDToTime(playedAt)
		track.ID = track.PlayedAt.Format(time.RFC3339)
		newTracks = append(newTracks, track)
	}

	newTracksEntity :=
		lo.Map(newTracks, func(t internal.TrackInfoDBSchema, _ int) domain.TrackInfo { return *t.ToEntity() })
	if err := a.repo.UpdateRange(ctx, newTracksEntity); err != nil {
		return err
	}

	return nil
}
