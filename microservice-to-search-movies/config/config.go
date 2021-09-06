package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"sync"
)

type Config struct {
	AppPort    string `envconfig:"APP_PORT" default:"8090"`
	GrpcPort   string `envconfig:"GRPC_PORT" default:"9000"`
	OMDBHost   string `envconfig:"OMDB_HOST" default:"http://www.omdbapi.com/"`
	OMDBAPIKey string `envconfig:"OMDB_API_KEY" default:"faf7e5bb"`
}

var once sync.Once
var instance Config

func GetConfig() Config {
	once.Do(func() {
		if err := godotenv.Load(".env"); err != nil {
			zap.Error(errors.New("failed to read from .env file"))
		}

		err := envconfig.Process("", &instance)
		if err != nil {
			zap.Error(err)
		}
	})

	return instance
}
