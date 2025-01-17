package main

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DBHost      string `envconfig:"DB_HOST" default:"localhost"`
	DBPort      uint16 `envconfig:"DB_PORT" default:"5432"`
	DBUser      string `envconfig:"DB_USER" default:"postgres"`
	DBPassword  string `envconfig:"DB_PASSWORD" default:"postgres"`
	DBName      string `envconfig:"DB_NAME" default:"postgres"`
	Environment string `envconfig:"ENVIRONMENT" default:"production"`
	Port        string `envconfig:"PORT" default:"8080"`
}

func GetConfig() (Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
