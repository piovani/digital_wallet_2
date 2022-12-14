package cmd

import (
	"github.com/piovani/digital_wallet_2/infra/redis"
	"github.com/piovani/digital_wallet_2/ui/api"
	"github.com/piovani/digital_wallet_2/ui/grpc"
	"github.com/spf13/cobra"
)

var (
	Api = &cobra.Command{
		Use:                   "api",
		Short:                 "Start listen http API",
		Version:               "1.0.0",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			InitConfig()
			CheckFatal(api.NewApi(redis.NewRedis(), GetDatabase(), GetMetric()).Start())
		},
	}

	Grpc = &cobra.Command{
		Use:                   "grpc",
		Short:                 "Start listen http gRPC",
		Version:               "1.0.0",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			InitConfig()
			CheckFatal(grpc.NewGrpc(redis.NewRedis(), GetDatabase()).Start())
		},
	}
)
