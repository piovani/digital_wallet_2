package api

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/piovani/digital_wallet_2/api/controller/coin"
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

	// Controllers
	CoinController      *coin.CoinController
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
	ctx := context.TODO()

	a.CoinController = coin.NewCoinController(usecase.NewCoinPrice(ctx))
	a.OperationController = operation.NewOperationController(
		usecase.NewDeposit(a.Database),
		usecase.NewWithdraw(a.Database),
		usecase.NewHistoric(a.Database),
	)
}
