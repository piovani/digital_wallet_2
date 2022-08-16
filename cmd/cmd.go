package cmd

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/piovani/digital_wallet_2/infra/config"
	"github.com/piovani/digital_wallet_2/infra/database/mysql"
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

func GetConnectionDB() *sql.DB {
	db, err := mysql.GetClient()
	CheckFatal(err)

	return db
}