package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		DBName   string `json:"dbname"`
		SSLMode  string `json:"sslmode"`
	} `json:"database"`
	Consul struct {
		Address string `json:"address"`
		Port    int    `json:"port"`
	} `json:"consul"`
	Endpoint struct {
		Name    string `json:"name"`
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
