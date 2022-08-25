package cmd

import (
	"math/rand"

	"github.com/piovani/digital_wallet_2/domain"
	"github.com/piovani/digital_wallet_2/infra/database"
	"github.com/piovani/digital_wallet_2/infra/database/mysql"
	"github.com/piovani/digital_wallet_2/infra/database/mysql/repositories"
	"github.com/piovani/digital_wallet_2/usecase"
)

func seed() {
	InitConfig()
	db, err := mysql.GetClient()
	CheckFatal(err)

	database := database.NewDatabase(db, repositories.NewWalletRepository(db), repositories.NewOperationRepository(db))

	depositCase := usecase.NewDeposit(database)
	userName := "john"

	for i := 0; i < 5; i++ {
		err := depositCase.Execute(userName, domain.COINS[rand.Intn(len(domain.COINS))], rand.Float64())
		CheckFatal(err)
	}
}
