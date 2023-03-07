package config

import (
	"github.com/xsicx/highload/pkg/configuration"
)

type Application struct {
	DB DBConfig
}

type DBConfig struct {
	DSN                string `mapstructure:"dsn"`
	ConnMaxLifeTime    int    `mapstructure:"conn_max_life_time"`
	ConnMaxIdleTime    int    `mapstructure:"conn_max_idle_time"`
	MaxOpenConnections int    `mapstructure:"max_open_connections"`
	MaxIdleConnections int    `mapstructure:"max_idle_connections"`
}

func Initialize() *Application {
	if cfg, ok := configuration.New(&Application{}).(*Application); ok {
		return cfg
	}

	panic("couldn't initialize application config")
}
