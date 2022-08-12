package cmd

import (
	"fmt"
	"os"

	"github.com/piovani/digital_wallet_2/infra/config"
)

func Execute() {
	cmd := ""

	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	switch cmd {
	case "api":
		startApi()
	case "collect":
		collect()
	case "read":
		read()
	default:
		fmt.Println("NENHUM COMANDO")
	}
}

func CheckFatal(err error) {
	if err != nil {
		fmt.Errorf(err.Error())
	}
}

func InitConfig() {
	CheckFatal(config.InitConfig())
}
