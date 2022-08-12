package api

import (
	"github.com/piovani/digital_wallet_2/api/controller/infra"
)

func (a *Api) GetRoutes() {
	a.Service.GET("/", infra.Heart)
}
