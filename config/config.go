package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type PostgresConfig struct {
	Host          string
	Port          uint16
	User          string
	Password      string
	Database      string
	RunMigrations bool `yaml:"run_migrations"`
}

type Config struct {
	Postgres PostgresConfig
}

func ReadConfig(configPath string) (*Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
