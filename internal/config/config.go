package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	ListenAddr string

	Postgres struct {
		Host     string
		Database string
		Username string
		Password string
	}
}

func Load(configPath string) (*Config, error) {
	var conf Config
	if _, err := toml.DecodeFile(configPath, &conf); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %s, error: %v", configPath, err)
	}

	return &conf, nil
}

func (c *Config) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		c.Postgres.Username,
		c.Postgres.Password,
		c.Postgres.Host,
		c.Postgres.Database)
}
