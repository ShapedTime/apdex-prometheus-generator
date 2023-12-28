package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Apdexes []Apdex `yaml:"apdex"`
}

type Apdex struct {
	Name             string            `yaml:"name"`
	Subsystem        string            `yaml:"subsystem,omitempty"`
	System           string            `yaml:"system"`
	GoodBucket       string            `yaml:"good_bucket"`
	ToleratingBucket string            `yaml:"tolerating_bucket"`
	Total            string            `yaml:"total"`
	ExtraLabels      map[string]string `yaml:"extra_labels"`
	Windows          []string          `yaml:"windows"`
}

// ReadConfig reads a YAML configuration file and unmarshals it into a Config struct.
func ReadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
