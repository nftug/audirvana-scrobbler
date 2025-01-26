package common

import (
	"path/filepath"

	"github.com/kirsle/configdir"
	"github.com/samber/do"
)

type ConfigPathProvider struct {
	localPath string
}

func NewConfigPathProvider(i *do.Injector) (*ConfigPathProvider, error) {
	configPath := configdir.LocalConfig("AudirvanaScrobbler")
	if err := configdir.MakePath(configPath); err != nil {
		return nil, err
	}
	return &ConfigPathProvider{localPath: configPath}, nil
}

func (p *ConfigPathProvider) GetJoinedPath(filename string) string {
	return filepath.Join(p.GetLocalPath(), filename)
}

func (p *ConfigPathProvider) GetLocalPath() string {
	return p.localPath
}
