package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Db    string `json:"db",omitempty`
	Redis string `json:"redis",omitempty`
	Kafka string `json:"kafka,omitempty"`
}

// type MixdList []*Config

func LoadFile(filename string) (*Config, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = json.Unmarshal(content, &cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
