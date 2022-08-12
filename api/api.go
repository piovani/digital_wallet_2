package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/piovani/digital_wallet_2/infra/config"
	"github.com/piovani/digital_wallet_2/infra/redis"
)

type Api struct {
	Redis   *redis.Redis
	Service *echo.Echo
}

func NewApi(redis *redis.Redis) *Api {
	return &Api{
		Redis:   redis,
		Service: echo.New(),
	}
}

func (a *Api) Start() {
	a.GetRoutes()

	a.Service.Start(fmt.Sprintf(":%d", config.Env.ApiPort))
}
