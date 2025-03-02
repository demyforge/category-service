package config

import (
	"fmt"

	"github.com/demyforge/utils/conf/toml"
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
	return toml.Load[Config](configPath)
}

func (c *Config) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		c.Postgres.Username,
		c.Postgres.Password,
		c.Postgres.Host,
		c.Postgres.Database)
}
