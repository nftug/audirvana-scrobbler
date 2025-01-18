package internal

import (
	"audirvana-scrobbler/app/domain"
	"path/filepath"

	"github.com/kirsle/configdir"
	"github.com/samber/do"
)

type configPathImpl struct {
	localPath string
}

func NewConfigPath(i *do.Injector) (domain.ConfigPathProvider, error) {
	configPath := configdir.LocalConfig("AudirvanaScrobbler")
	if err := configdir.MakePath(configPath); err != nil {
		return nil, err
	}
	return &configPathImpl{localPath: configPath}, nil
}

func (lp *configPathImpl) GetJoinedPath(filename string) string {
	return filepath.Join(lp.GetLocalPath(), filename)
}

func (lp *configPathImpl) GetLocalPath() string {
	return lp.localPath
}
