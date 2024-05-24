package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Consul struct {
		Address string `json:"address"`
		Port    int    `json:"port"`
	} `json:"consul"`
	Endpoint struct {
		Address string `json:"address"`
		Port    int    `json:"port"`
	} `json:"endpoint"`
}

func LoadConfig(filename string) (*Config, error) {
	configFile, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening config file: %v", err)
	}
	defer configFile.Close()

	var config Config
	err = json.NewDecoder(configFile).Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("error decoding config JSON: %v", err)
	}
	return &config, nil
}
