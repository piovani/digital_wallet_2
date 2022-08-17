package usecase

import (
	"github.com/piovani/digital_wallet_2/domain"
	"github.com/piovani/digital_wallet_2/infra/database"
)

type Historic struct {
	Database *database.Database
}

func NewHistoric(db *database.Database) *Historic {
	return &Historic{
		Database: db,
	}
}

func (d *Historic) Execute(username string) ([]domain.Operation, error) {
	operations, err := d.Database.OperationRepository.FindByUserName(username)
	if err != nil {
		return operations, err
	}

	return operations, nil
}
