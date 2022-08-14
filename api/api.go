package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/piovani/digital_wallet_2/api/controller/operation"
	"github.com/piovani/digital_wallet_2/infra/config"
	"github.com/piovani/digital_wallet_2/infra/database"
	"github.com/piovani/digital_wallet_2/infra/redis"
	"github.com/piovani/digital_wallet_2/usecase"
)

type Api struct {
	Redis    *redis.Redis
	Database *database.Database
	Service  *echo.Echo

	// Usecases
	Deposit *usecase.Deposit

	// Controllers
	OperationController *operation.OperationController
}

func NewApi(redis *redis.Redis, database *database.Database) *Api {
	return &Api{
		Redis:    redis,
		Database: database,
		Service:  echo.New(),
	}
}

func (a *Api) Start() error {
	a.getControllers()
	a.getRoutes()

	return a.Service.Start(fmt.Sprintf(":%d", config.Env.ApiPort))
}

func (a *Api) getControllers() {
	a.OperationController = operation.NewOperationController(usecase.NewDeposit(a.Database))
}
