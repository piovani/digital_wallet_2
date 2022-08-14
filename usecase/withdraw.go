package usecase

import (
	"github.com/piovani/digital_wallet_2/domain"
	"github.com/piovani/digital_wallet_2/infra/database"
	"github.com/piovani/digital_wallet_2/infra/database/mysql/repositories"
)

const (
	WITHDRAW_TYPE = "withdraw"
)

type Withdraw struct {
	Database *database.Database
	repo     *repositories.OperationRepository
}

func NewWithdraw(db *database.Database) *Withdraw {
	return &Withdraw{
		Database: db,
		repo:     repositories.NewOperationRepository(db.Connection),
	}
}

func (d *Withdraw) Execute(username string, value float64) error {
	opt := domain.NewOperation(username, WITHDRAW_TYPE, value)

	return d.repo.Insert(opt)
}
