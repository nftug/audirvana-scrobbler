package infra

import (
	"audirvana-scrobbler/app/internal/domain"
	"audirvana-scrobbler/app/shared/response"
	"context"
	"database/sql"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/kirsle/configdir"
	"github.com/samber/do"
	"github.com/samber/lo"

	_ "github.com/mattn/go-sqlite3"
)

const LAST_SCROBBLE_TIME_LOG = "lastScrobbleTime.log"

type audirvanaImporterImpl struct {
	configpath ConfigPath
	tempPath   TempPath
}

func NewAudirvanaImporter(i *do.Injector) (domain.AudirvanaImporter, error) {
	return &audirvanaImporterImpl{
		configpath: do.MustInvoke[ConfigPath](i),
		tempPath:   do.MustInvoke[TempPath](i),
	}, nil
}

func (a *audirvanaImporterImpl) GetAllTracks(ctx context.Context) ([]response.TrackInfo, error) {
	dbPath, err := a.copyDBFileToLocal()
	if err != nil {
		return nil, err
	}

	// Get last scrobble time
	lastScrobbleTime, err := a.readLastScrobbleTime()
	if err != nil {
		return nil, err
	}

	tracks, err := a.getScrobbleLog(ctx, *dbPath, *lastScrobbleTime)
	if err != nil {
		return nil, err
	}

	return tracks, nil
}

func (a *audirvanaImporterImpl) copyDBFileToLocal() (*string, error) {
	audirvanaConfigPath := configdir.LocalConfig("Audirvana")
	dbPathOriginal := filepath.Join(audirvanaConfigPath, "AudirvanaDatabase.sqlite")
	dbPathTemp := a.tempPath.GetJoinedPath("AudirvanaDatabase.sqlite")

	src, err := os.Open(dbPathOriginal)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer src.Close()

	dst, err := os.Create(dbPathTemp)
	if err != nil {
		return nil, fmt.Errorf("failed to create local DB file: %w", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return nil, fmt.Errorf("failed to copy DB file to local: %w", err)
	}

	return &dbPathTemp, nil
}

func (a *audirvanaImporterImpl) readLastScrobbleTime() (*string, error) {
	logPath := a.configpath.GetJoinedPath(LAST_SCROBBLE_TIME_LOG)
	file, err := os.OpenFile(logPath, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open last scrobble time log: %w", err)
	}
	defer file.Close()

	buf := make([]byte, 19)
	n, _ := file.Read(buf)

	lastScrobble := lo.Ternary(n == 0, "2459364.9074768517", strings.TrimSpace(string(buf[:n])))
	return lo.ToPtr(lastScrobble), nil
}

func (a *audirvanaImporterImpl) getScrobbleLog(ctx context.Context, dbFilePath string, lastScrobbleTime string) ([]response.TrackInfo, error) {
	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?mode=ro", dbFilePath))
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `
		SELECT ARTISTS.name AS artist, ALBUMS.title AS album, TRACKS.title,
			   date(TRACKS.last_played_date) AS date, time(TRACKS.last_played_date) AS time,
			   last_played_date
		FROM TRACKS
		JOIN ALBUMS ON ALBUMS.album_id = TRACKS.album_id
		JOIN ALBUMS_ARTISTS ON ALBUMS_ARTISTS.album_id = ALBUMS.album_id
		JOIN ARTISTS ON ALBUMS_ARTISTS.artist_id = ARTISTS.artist_id
		WHERE round(last_played_date, 4) > round(?, 4)
		ORDER BY last_played_date ASC
	`

	rows, err := db.QueryContext(ctx, query, lastScrobbleTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tracks []response.TrackInfo
	for rows.Next() {
		var track response.TrackInfo
		var playedDate, playedTime string
		if err := rows.Scan(&track.Artist, &track.Album, &track.Track, &playedDate, &playedTime, &track.PlayedAt); err != nil {
			return nil, err
		}

		playedAt, err := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s %s", playedDate, playedTime))
		if err != nil {
			return nil, err
		}
		track.PlayedAt = playedAt.Format(time.RFC3339)
		track.Id = uuid.New().String()

		tracks = append(tracks, track)
	}

	return tracks, nil
}
