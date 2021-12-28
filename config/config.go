package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Environment string `json:"environment"`
	Server      struct {
		Port int `json:"port"`
	}
	Store struct {
		Name string `json:"name"`
		Dsn  string `json:"dsn"`
	}
}

// load config from a json file
func LoadConf(path string) (*Config, error) {
	configFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(configFile)
	config := &Config{}
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}
	return config, nil
}
