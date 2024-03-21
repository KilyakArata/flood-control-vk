package config

import (
	"os"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Timer 		int64 `yaml:"timer"`
	CountMax    int64 `yaml:"countmax"`
	CheckUserID		int64 `yaml:"userID"`
}

func Read(filename string) (*Config, error) {
	var config Config

	information, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(information, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
