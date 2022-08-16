package usecase

import (
	"github.com/piovani/digital_wallet_2/domain"
	"github.com/piovani/digital_wallet_2/infra/database"
)

const (
	DEPOSIT_TYPE = "deposit"
)

type Deposit struct {
	Database *database.Database
}

func NewDeposit(db *database.Database) *Deposit {
	return &Deposit{
		Database: db,
	}
}

func (d *Deposit) Execute(username, coin string, value float64) error {
	opt := domain.NewOperation(username, DEPOSIT_TYPE, coin, value)

	return d.Database.OperationRepository.Insert(opt)
}
