package common

import (
	"audirvana-scrobbler/app/domain"
	"path/filepath"

	"github.com/kirsle/configdir"
	"github.com/samber/do"
)

type configPathProviderImpl struct {
	localPath string
}

func NewConfigPathProvider(i *do.Injector) (domain.ConfigPathProvider, error) {
	configPath := configdir.LocalConfig("AudirvanaScrobbler")
	if err := configdir.MakePath(configPath); err != nil {
		return nil, err
	}
	return &configPathProviderImpl{localPath: configPath}, nil
}

func (lp *configPathProviderImpl) GetJoinedPath(filename string) string {
	return filepath.Join(lp.GetLocalPath(), filename)
}

func (lp *configPathProviderImpl) GetLocalPath() string {
	return lp.localPath
}
