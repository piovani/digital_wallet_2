package cmd

import (
	"github.com/piovani/digital_wallet_2/api"
	"github.com/piovani/digital_wallet_2/infra/database"
	"github.com/piovani/digital_wallet_2/infra/database/mysql/repositories"
	"github.com/piovani/digital_wallet_2/infra/redis"
)

func startApi() {
	InitConfig()

	err := api.NewApi(redis.NewRedis(), getDatabase(), GetMetric()).Start()
	CheckFatal(err)
}

func getDatabase() *database.Database {
	conn := GetConnectionDB()

	return database.NewDatabase(
		conn,
		repositories.NewWalletRepository(conn),
		repositories.NewOperationRepository(conn),
	)
}
