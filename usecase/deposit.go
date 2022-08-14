package usecase

import (
	"github.com/piovani/digital_wallet_2/domain"
	"github.com/piovani/digital_wallet_2/infra/database"
	"github.com/piovani/digital_wallet_2/infra/database/mysql/repositories"
)

const (
	DEPOSIT_TYPE = "deposit"
)

type Deposit struct {
	Database *database.Database
	repo     *repositories.OperationRepository
}

func NewDeposit(db *database.Database) *Deposit {
	return &Deposit{
		Database: db,
		repo:     repositories.NewOperationRepository(db.Connection),
	}
}

func (d *Deposit) Execute(username string, value float64) error {
	opt := domain.NewOperation(username, DEPOSIT_TYPE, value)

	return d.repo.Insert(opt)
}
