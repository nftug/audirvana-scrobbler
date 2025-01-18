package internal

import (
	"audirvana-scrobbler/app/domain"
	"fmt"
	"reflect"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"github.com/samber/do"
)

const CONFIG_FILENAME = "env.cfg"

type configProviderImpl struct {
	filepath string
	config   domain.Config
}

func NewConfigProvider(i *do.Injector) (domain.ConfigProvider, error) {
	configPath := do.MustInvoke[domain.ConfigPathProvider](i)
	filepath := configPath.GetJoinedPath(CONFIG_FILENAME)
	if err := godotenv.Load(filepath); err != nil {
		return nil, err
	}

	var cfg domain.Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &configProviderImpl{filepath, cfg}, nil
}

func (c *configProviderImpl) Get() domain.Config {
	return c.config
}

func (c *configProviderImpl) Write(cfg domain.Config) error {
	envmap, err := structToEnvMap(cfg)
	if err != nil {
		return err
	}
	if err := godotenv.Write(envmap, c.filepath); err != nil {
		return err
	}
	c.config = cfg
	return nil
}

func structToEnvMap(cfg interface{}) (map[string]string, error) {
	envMap := make(map[string]string)

	v := reflect.ValueOf(cfg)
	t := reflect.TypeOf(cfg)

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a struct, got %T", cfg)
	}

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("env")
		if tag == "" {
			continue
		}

		value := v.Field(i)

		var strValue string
		switch value.Kind() {
		case reflect.String:
			strValue = value.String()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			strValue = fmt.Sprintf("%d", value.Int())
		case reflect.Bool:
			if value.Bool() {
				strValue = "true"
			} else {
				strValue = "false"
			}
		default:
			return nil, fmt.Errorf("unsupported field type: %s", value.Kind())
		}

		envMap[tag] = strValue
	}

	return envMap, nil
}
