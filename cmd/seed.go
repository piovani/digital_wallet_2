package cmd

import (
	"math/rand"

	"github.com/piovani/digital_wallet_2/domain"
	"github.com/piovani/digital_wallet_2/usecase"
	"github.com/spf13/cobra"
)

var Seed = &cobra.Command{
	Use:                "seed",
	Short:              "Charge inital from database",
	Version:            "1.0.0",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		InitConfig()

		depositCase := usecase.NewDeposit(GetDatabase())
		userName := "john"

		for i := 0; i < 5; i++ {
			err := depositCase.Execute(userName, domain.COINS[rand.Intn(len(domain.COINS))], rand.Float64())
			CheckFatal(err)
		}
	},
}
