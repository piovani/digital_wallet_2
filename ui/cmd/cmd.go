package cmd

import (
	"github.com/piovani/digital_wallet_2/infra/config"
	"github.com/piovani/digital_wallet_2/infra/database"
	"github.com/piovani/digital_wallet_2/infra/database/mysql"
	"github.com/piovani/digital_wallet_2/infra/database/mysql/repositories"
	log_infra "github.com/piovani/digital_wallet_2/infra/log"
	"github.com/piovani/digital_wallet_2/infra/metric"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var Log log_infra.Log

func Execute() {
	Log = *log_infra.NewLog("CMD")

	cmd := &cobra.Command{
		Use:                   "digital-wallet-2",
		Short:                 "digital-wallet-2",
		Version:               "1.0.0",
		DisableFlagsInUseLine: true,
	}

	cmd.AddCommand(
		Api,
		CollectCoinsPrices,
		Grpc,
		Migrate,
		ReadCoinsPrices,
		Seed,
	)

	CheckFatal(cmd.Execute())
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

func GetDatabase() *database.Database {
	conn := GetConnectionDB()

	return database.NewDatabase(
		conn,
		repositories.NewWalletRepository(conn),
		repositories.NewOperationRepository(conn),
	)
}

func GetMetric() *metric.Metric {
	metric, err := metric.NewMetric()
	CheckFatal(err)

	return metric
}
