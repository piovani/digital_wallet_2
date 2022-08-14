package api

import (
	"github.com/piovani/digital_wallet_2/api/controller/infra"
)

func (a *Api) getRoutes() {
	a.Service.GET("/", infra.Heart)

	a.Service.POST("/operations", a.OperationController.Create)
}
