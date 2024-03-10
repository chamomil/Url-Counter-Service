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

type UserCredentials struct {
	Username string
	Password string
}

type Config struct {
	Postgres   PostgresConfig
	ServerPort uint `yaml:"server_port"`
	Users      []UserCredentials
}

var Data Config

func ReadConfig(configPath string) (*Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, &Data); err != nil {
		return nil, err
	}

	return &Data, nil
}
