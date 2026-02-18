package common

import (
	"audirvana-scrobbler/app/domain"
	"encoding/json"
	"os"

	"github.com/samber/do"
)

const CONFIG_FILENAME = "config.json"

type configProviderImpl struct {
	filepath string
	config   domain.Config
}

func NewConfigProvider(i *do.Injector) (domain.ConfigProvider, error) {
	configPath := do.MustInvoke[*ConfigPathProvider](i)
	filepath := configPath.GetJoinedPath(CONFIG_FILENAME)
	cfg, err := loadConfig(filepath)
	if err != nil {
		return nil, err
	}
	return &configProviderImpl{
		filepath: filepath,
		config:   cfg,
	}, nil
}

func (c *configProviderImpl) Get() domain.Config {
	return c.config
}

func (c *configProviderImpl) Write(cfg domain.Config) error {
	if err := saveConfig(c.filepath, cfg); err != nil {
		return err
	}
	c.config = cfg
	return nil
}

func loadConfig(filepath string) (domain.Config, error) {
	data, err := os.ReadFile(filepath)

	// If file does not exist, return default config
	if os.IsNotExist(err) {
		return domain.DefaultConfig, nil
	} else if err != nil {
		return domain.Config{}, err
	}
	cfg := domain.DefaultConfig
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return domain.Config{}, err
	}
	return cfg, nil
}

func saveConfig(filepath string, config domain.Config) error {
	data, err := json.MarshalIndent(&config, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
