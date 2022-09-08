package cmd

import (
	"github.com/spf13/cobra"
)

var Migrate = &cobra.Command{
	Use:                "migrate",
	Short:              "Migrate initial charge from database",
	Version:            "1.0.0",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		InitConfig()
		CheckFatal(GetDatabase().Migrate())
	},
}
