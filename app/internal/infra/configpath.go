package infra

import (
	"path/filepath"

	"github.com/kirsle/configdir"
	"github.com/samber/do"
)

type ConfigPath interface {
	GetJoinedPath(filename string) string
	GetLocalPath() string
}

type configPathImpl struct {
	localPath string
}

func NewConfigPath(i *do.Injector) (ConfigPath, error) {
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
