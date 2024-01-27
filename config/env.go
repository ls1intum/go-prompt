package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Database struct {
		Host     string `env:"HOST,notEmpty" envDefault:"localhost"`
		Port     int    `env:"PORT,notEmpty" envDefault:"5432"`
		Db       string `env:"DB,notEmpty" envDefault:"postgres"`
		User     string `env:"USER,notEmpty" envDefault:"postgres"`
		Password string `env:"PASSWORD,notEmpty" envDefault:"postgres"`
	} `envPrefix:"POSTGRES_"`
}

var GlobalConfig Config

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.WithError(err).Warn("Error loading .env file")
	}

	err = env.Parse(&GlobalConfig)
	if err != nil {
		log.WithError(err).Fatal("Error parsing environment variables")
	}
}
