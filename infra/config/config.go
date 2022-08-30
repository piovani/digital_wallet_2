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

	MysqlUser     string `envconfig:"MYSQL_USER"`
	MysqlPassword string `envconfig:"MYSQL_PASSWORD"`
	MysqlHost     string `envconfig:"MYSQL_HOST"`
	MysqlPort     int    `envconfig:"MYSQL_PORT"`
	MysqlDatabase string `envconfig:"MYSQL_DATABASE"`

	Prometheuspushgateway string `envconfig:"PROMETHEUS_PUSHGATEWAY"`
}

var Env Config

func InitConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	return envconfig.Process("", &Env)
}
