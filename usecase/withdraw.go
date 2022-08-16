package usecase

import (
	"github.com/piovani/digital_wallet_2/domain"
	"github.com/piovani/digital_wallet_2/infra/database"
)

const (
	WITHDRAW_TYPE = "withdraw"
)

type Withdraw struct {
	Database *database.Database
}

func NewWithdraw(db *database.Database) *Withdraw {
	return &Withdraw{
		Database: db,
	}
}

func (d *Withdraw) Execute(username, coin string, value float64) error {
	// NECESSARIO VALIDAR VALOR NA WALLET
	opt := domain.NewOperation(username, WITHDRAW_TYPE, coin, value)

	return d.Database.OperationRepository.Insert(opt)
}
