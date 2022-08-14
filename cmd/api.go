package cmd

import (
	"github.com/piovani/digital_wallet_2/api"
	"github.com/piovani/digital_wallet_2/infra/database"
	"github.com/piovani/digital_wallet_2/infra/redis"
)

func startApi() {
	InitConfig()

	connection := GetConnectionDB()

	err := api.NewApi(redis.NewRedis(), database.NewDatabase(connection)).Start()
	CheckFatal(err)
}
