package config

import (
    "fmt"

    "github.com/BurntSushi/toml"
)

type Config struct {
    ListenAddr string
}

func Load(configPath string) (*Config, error) {
    var conf Config
    if _, err := toml.DecodeFile(configPath, &conf); err != nil {
        return nil, fmt.Errorf("failed to parse config file: %s, error: %v", configPath, err)
    }

    return &conf, nil
}
