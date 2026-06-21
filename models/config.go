package models

type Config struct {
	InputFolder string `yaml:"input_folder"`
	Recursive   bool   `yaml:"recursive"`
	Workers     int    `yaml:"workers"`
}