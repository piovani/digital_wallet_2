package cmd

import (
	"github.com/piovani/digital_wallet_2/domain"
	"github.com/piovani/digital_wallet_2/infra/database/mysql"
)

func migrate() {
	InitConfig()

	db, err := mysql.GetClient()
	CheckFatal(err)

	err = db.AutoMigrate(domain.Operation{})
	CheckFatal(err)
}
