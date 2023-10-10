package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App      `yaml:"app"`
	HTTP     `yaml:"http"`
	Database `yaml:"database"`
}

type App struct {
	Name        string `env-required:"true" env:"APP_NAME"`
	Host        string `env-required:"true" env:"APP_HOST"`
	Version     string `env-required:"true" env:"APP_VERSION"`
	Environment string `env-required:"true" env:"APP_Environment"`
}

type HTTP struct {
	Port string `env-required:"true" env:"HTTP_PORT"`
}

type Database struct {
	PrimaryHost  string `env-required:"true" env:"DATABASE_PRIMARY_HOST"`
	PrimaryPort  int    `env-required:"true" env:"DATABASE_PRIMARY_PORT"`
	ReplicaHost  string `env-required:"true" env:"DATABASE_REPLICA_HOST"`
	ReplicaPort  int    `env-required:"true" env:"DATABASE_REPLICA_PORT"`
	DatabaseName string `env-required:"true" env:"DATABASE_NAME"`
	User         string `env-required:"true" env:"DATABASE_USER"`
	Password     string `env-required:"true" env:"DATABASE_PASSWORD"`
}

func NewConfig() (*Config, error) {
	config := &Config{}

	err := cleanenv.ReadEnv(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
