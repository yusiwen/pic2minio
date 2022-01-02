package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	EndPoint  string `yaml:"endpoint"`
	AccessKey string `yaml:"access-key"`
	SecretKey string `yaml:"secret-key"`
	Bucket    string `yaml:"bucket"`
	BaseDir   string `yaml:"base-dir"`
}

func InitConfig() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadFile(homeDir + "/.config/pic2minio.yaml")
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
