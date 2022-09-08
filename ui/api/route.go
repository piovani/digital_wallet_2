package api

import (
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/piovani/digital_wallet_2/docs"
	"github.com/piovani/digital_wallet_2/ui/api/controller/infra"
)

func (a *Api) getRoutes() {
	a.setDocs()

	a.Service.GET("/", infra.Heart)
	a.Service.GET("/swagger/*", echoSwagger.WrapHandler)

	a.Service.GET("coin/price", a.CoinController.Price)

	groupOperation := a.Service.Group("/operations")
	groupOperation.POST("/deposit", a.OperationController.Deposit)
	groupOperation.POST("/withdraw", a.OperationController.Withdraw)
	groupOperation.GET("/historic", a.OperationController.Historic)
	groupOperation.GET("/balance", a.OperationController.Balance)
}

func (a *Api) setDocs() {
	docs.SwaggerInfo.Title = "Digital Wallet 2"
	docs.SwaggerInfo.Description = "This is an application to simulate the operations of a digital sheep"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "https://github.com/piovani/digital_wallet_2"
	docs.SwaggerInfo.Schemes = []string{"http"}
}
