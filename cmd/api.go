package cmd

import (
	"github.com/piovani/digital_wallet_2/api"
	"github.com/piovani/digital_wallet_2/infra/redis"
)

func startApi() {
	InitConfig()

	err := api.NewApi(redis.NewRedis(), GetDatabase(), GetMetric()).Start()
	CheckFatal(err)
}
