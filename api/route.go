package api

import (
	"github.com/piovani/digital_wallet_2/api/controller/infra"
)

func (a *Api) getRoutes() {
	a.Service.GET("/", infra.Heart)

	groupOperation := a.Service.Group("/operations")
	groupOperation.POST("/deposit", a.OperationController.Deposit)
	groupOperation.POST("/withdraw", a.OperationController.Withdraw)
}
