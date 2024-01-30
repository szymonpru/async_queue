package utils

import (
	"log"

	"github.com/caarlos0/env/v10"
	_ "github.com/joho/godotenv/autoload" // Autoloads .env file
)

var (
	Config Configuration
)

type Configuration struct {
	ApiHost string `env:"API_HOST" envDefault:"localhost"`
	ApiPort string `env:"API_PORT" envDefault:"8080"`
	IsDebug bool   `env:"DEBUG" envDefault:"False"`

	RedisHost string `env:"REDIS_HOST" envDefault:"localhost"`
	RedisPort string `env:"REDIS_PORT" envDefault:"6379"`
}

func init() {
	if err := env.Parse(&Config); err != nil {
		log.Fatalf("%+v\n", err)
	}
}
