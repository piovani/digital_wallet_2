package cmd

import (
	"github.com/piovani/digital_wallet_2/api"
	"github.com/piovani/digital_wallet_2/infra/database"
	"github.com/piovani/digital_wallet_2/infra/database/mysql/repositories"
	"github.com/piovani/digital_wallet_2/infra/redis"
)

func startApi() {
	InitConfig()

	err := api.NewApi(redis.NewRedis(), getDatabase()).Start()
	CheckFatal(err)
}

func getDatabase() *database.Database {
	conn := GetConnectionDB()

	or := repositories.NewOperationRepository(conn)
	wr := repositories.NewWalletRepository(conn)

	return database.NewDatabase(conn, wr, or)
}
