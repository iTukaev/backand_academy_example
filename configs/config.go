package configs

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

//go:embed default_config.yaml
var defaultConfig []byte

type Config struct {
	Words []Word `yaml:"words"`
}

type Word struct {
	Word string `yaml:"word"`
	Hint string `yaml:"hint"`
}

func Init(configPath string) (*Config, error) {
	v := viper.New()

	v.SetConfigType("yaml")
	if configPath != "" {
		v.SetConfigFile(configPath)
	}

	if err := v.ReadConfig(bytes.NewBuffer(defaultConfig)); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	if err := v.MergeInConfig(); err != nil {
		if errors.Is(err, &viper.ConfigParseError{}) {
			return nil, fmt.Errorf("parse config file: %w", err)
		}
	}

	cfg := &Config{}

	if err := v.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config file: %w", err)
	}

	return cfg, nil
}
