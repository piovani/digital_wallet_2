package database

import (
	"github.com/piovani/digital_wallet_2/domain"
	"github.com/piovani/digital_wallet_2/infra/database/mysql/repositories"
	"gorm.io/gorm"
)

type Database struct {
	Connection *gorm.DB

	WalletRepository    domain.WalletRepository
	OperationRepository domain.OperationRepository
}

func NewDatabase(db *gorm.DB) *Database {
	database := &Database{
		Connection: db,
	}

	database.getRepositores()

	return database
}

func (d *Database) getRepositores() {
	d.OperationRepository = repositories.NewOperationRepository(d.Connection)
	d.WalletRepository = repositories.NewWalletRepository(d.Connection)
}
