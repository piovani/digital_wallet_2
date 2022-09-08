package cmd

import (
	"github.com/piovani/digital_wallet_2/ui/worker"
	"github.com/spf13/cobra"
)

var ReadCoinsPrices = &cobra.Command{
	Use:                "read",
	Short:              "Read all values conins collected and print",
	Version:            "1.0.0",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		InitConfig()
		worker.NewReadCoinsPrices().Execute()
	},
}
