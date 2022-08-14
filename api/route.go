package api

import (
	"github.com/piovani/digital_wallet_2/api/controller/infra"
	"github.com/piovani/digital_wallet_2/api/controller/operation"
)

func (a *Api) GetRoutes() {
	a.Service.GET("/", infra.Heart)

	a.Service.POST("/operations", operation.Create)
}
