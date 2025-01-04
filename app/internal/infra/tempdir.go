package infra

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/samber/do"
)

type TempPath interface {
	GetJoinedPath(filename string) string
	Shutdown() error
}

type tempPathImpl struct {
	tempPath string
}

func NewTempPath(i *do.Injector) (TempPath, error) {
	tempPath, err := os.MkdirTemp("", "audirvana-scrobbler")
	if err != nil {
		return nil, err
	}
	return &tempPathImpl{tempPath: tempPath}, nil
}

func (t *tempPathImpl) GetJoinedPath(filename string) string {
	return filepath.Join(t.tempPath, filename)
}

func (t *tempPathImpl) Shutdown() error {
	err := os.RemoveAll(t.tempPath) // ディレクトリ内のすべてを削除
	if err != nil {
		return fmt.Errorf("error deleting temp directory: %w", err)
	}
	return nil
}
