package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

const (
	configEnvKey = "CONFIG_FILE"
)

type Config struct {
	Service struct {
		Host         string        `yaml:"host"`
		Port         string        `yaml:"port"`
		ReadTimeout  time.Duration `yaml:"read_timeout"`
		WriteTimeout time.Duration `yaml:"write_timeout"`
		IdleTimeout  time.Duration `yaml:"idle_timeout"`
	} `yaml:"service"`
	DB struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"db_name"`
		ShardID  uint8  `yaml:"shard_id"`
		DCID     uint8  `yaml:"dc_id"`
	} `yaml:"db"`
}

func LoadConfig() (*Config, error) {
	f, err := os.Open(os.Getenv(configEnvKey))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	config := &Config{}
	if err := yaml.NewDecoder(f).Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}
