package api

import (
	"github.com/piovani/digital_wallet_2/api/controller/infra"
)

func (a *Api) getRoutes() {
	a.Service.GET("/", infra.Heart)

	a.Service.GET("coin/price", a.CoinController.Price)

	groupOperation := a.Service.Group("/operations")
	groupOperation.POST("/deposit", a.OperationController.Deposit)
	groupOperation.POST("/withdraw", a.OperationController.Withdraw)
}
