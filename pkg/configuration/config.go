package configuration

import (
	"errors"
	"os"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

const (
	configsDir = "configs"
	configName = "main"
)

func New(cfg interface{}) interface{} {
	if err := parseEnv(); err != nil {
		panic(err)
	}
	if err := parseConfigFile(configsDir, configName); err != nil {
		panic(err)
	}

	env := viper.GetString("app.env")
	configOverride := viper.GetString("app.config.override")
	if err := parseConfigOverride(configsDir+"/"+env, configOverride); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&cfg, viper.DecodeHook(decodeHook)); err != nil {
		panic(err)
	}

	return cfg
}

type Config struct {
	Folder       string
	Name         string
	OverrideName string
}

func NewChain(cfg interface{}, configs ...Config) interface{} {
	var err error

	if err = parseEnv(); err != nil {
		panic(err)
	}

	for _, config := range configs {
		cfg, err = LoadData(cfg, config)
		if err != nil {
			panic(err)
		}
	}

	return cfg
}

func LoadData(cfg interface{}, config Config) (interface{}, error) {
	configFolder := configsDir

	if len(config.Folder) > 0 {
		configFolder = configFolder + "/" + config.Folder
	}

	env := viper.GetString("app.env")

	if err := parseConfigFile(configFolder, config.Name); err != nil {
		return nil, err
	}

	if err := parseConfigOverride(configFolder+"/"+env, config.OverrideName); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&cfg, viper.DecodeHook(decodeHook)); err != nil {
		return nil, err
	}

	return cfg, nil
}

func parseEnv() error {
	viper.AutomaticEnv()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	return nil
}

func parseConfigFile(folder, configName string) error {
	viper.AddConfigPath(folder)
	viper.SetConfigName(configName)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.MergeInConfig()
}

func parseConfigOverride(folder, configOverride string) error {
	if len(configOverride) == 0 {
		return nil
	}

	if _, err := os.Stat(folder + "/" + configOverride + ".yml"); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
	} else {
		viper.AddConfigPath(folder)
		viper.SetConfigName(configOverride)

		if err = viper.ReadInConfig(); err != nil {
			return err
		}
	}

	return viper.MergeInConfig()
}

func decodeHook(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
	if f.Kind() == reflect.String {
		value := data.(string)

		value = os.ExpandEnv(value)

		return value, nil
	}

	return data, nil
}
