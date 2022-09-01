package api

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	echo_middleware "github.com/labstack/echo/v4/middleware"

	"github.com/piovani/digital_wallet_2/api/controller/coin"
	"github.com/piovani/digital_wallet_2/api/controller/operation"
	"github.com/piovani/digital_wallet_2/infra/config"
	"github.com/piovani/digital_wallet_2/infra/database"
	"github.com/piovani/digital_wallet_2/infra/metric"
	"github.com/piovani/digital_wallet_2/infra/redis"
	"github.com/piovani/digital_wallet_2/usecase"
)

const METRIC_KEY = "metric_http"

type Api struct {
	Redis    *redis.Redis
	Database *database.Database
	Metric   *metric.Metric
	Service  *echo.Echo

	// Controllers
	CoinController      *coin.CoinController
	OperationController *operation.OperationController
}

func NewApi(redis *redis.Redis, database *database.Database, metric *metric.Metric) *Api {
	return &Api{
		Redis:    redis,
		Database: database,
		Metric:   metric,
		Service:  echo.New(),
	}
}

func (a *Api) Start() error {
	a.useMiddleware()
	a.getControllers()
	a.getRoutes()

	return a.Service.Start(fmt.Sprintf(":%d", config.Env.ApiPort))
}

func (a *Api) useMiddleware() {
	a.Service.Pre(a.Before)
	a.Service.Use(echo_middleware.BodyDump(a.CollectMetric))
}

func (a *Api) getControllers() {
	ctx := context.TODO()

	a.CoinController = coin.NewCoinController(usecase.NewCoinPrice(ctx))
	a.OperationController = operation.NewOperationController(
		usecase.NewDeposit(a.Database),
		usecase.NewWithdraw(a.Database),
		usecase.NewHistoric(a.Database),
		usecase.NewBalance(a.Database),
	)
}

func (a *Api) Before(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		http := metric.NewHTTP(c.Request().Host, c.Request().Method)
		http.Started()

		c.Set(METRIC_KEY, http)

		return next(c)
	}
}

func (a *Api) CollectMetric(c echo.Context, reqBody, resBody []byte) {
	http, ok := (c.Get(METRIC_KEY)).(*metric.HTTP)
	if ok {
		http.Finished()
		a.Metric.SaveHTTP(http)
	}
}
