package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ApiPort int64 `envconfig:"API_PORT"`

	RedisDSN      string `envconfig:"REDIS_DNS"`
	RedisUser     string `envconfig:"REDIS_USER"`
	RedisPassword string `envconfig:"REDIS_PASSWORD"`
}

var Env Config

func InitConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	return envconfig.Process("", &Env)
}
