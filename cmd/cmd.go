package cmd

import (
	"fmt"
	"os"

	"github.com/piovani/digital_wallet_2/infra/config"
	"github.com/piovani/digital_wallet_2/infra/database/mysql"
	log_infra "github.com/piovani/digital_wallet_2/infra/log"
	"github.com/piovani/digital_wallet_2/infra/metric"
	"gorm.io/gorm"
)

var Log log_infra.Log

func Execute() {
	Log = *log_infra.NewLog("CMD")

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
	case "migrate":
		migrate()
	case "seed":
		seed()
	default:
		fmt.Println("NENHUM COMANDO")
	}
}

func CheckFatal(err error) {
	if err != nil {
		Log.Fatal(err.Error())
	}
}

func InitConfig() {
	CheckFatal(config.InitConfig())
}

func GetConnectionDB() *gorm.DB {
	db, err := mysql.GetClient()
	CheckFatal(err)

	return db
}

func GetMetric() *metric.Metric {
	metric, err := metric.NewMetric()
	CheckFatal(err)

	return metric
}
