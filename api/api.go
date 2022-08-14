package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/piovani/digital_wallet_2/infra/config"
	"github.com/piovani/digital_wallet_2/infra/database"
	"github.com/piovani/digital_wallet_2/infra/redis"
)

type Api struct {
	Redis    *redis.Redis
	Database *database.Database
	Service  *echo.Echo
}

func NewApi(redis *redis.Redis, database *database.Database) *Api {
	return &Api{
		Redis:    redis,
		Database: database,
		Service:  echo.New(),
	}
}

func (a *Api) Start() error {
	a.GetRoutes()

	return a.Service.Start(fmt.Sprintf(":%d", config.Env.ApiPort))
}
