package cmd

import (
	"github.com/piovani/digital_wallet_2/api"
	"github.com/piovani/digital_wallet_2/infra/redis"
)

func startApi() {
	InitConfig()

	serivce := api.NewApi(redis.NewRedis())
	serivce.Start()
}
