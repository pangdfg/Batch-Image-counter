package services

import (
	"os"
	"runtime"

	"batch/image-counter/models"

	"gopkg.in/yaml.v3"
)

func LoadConfig(path string) (*models.Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg models.Config

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	if cfg.Workers <= 0 {
		cfg.Workers = runtime.NumCPU()
	}

	return &cfg, nil
}