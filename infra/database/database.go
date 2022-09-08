package database

import (
	"github.com/piovani/digital_wallet_2/domain"
	"gorm.io/gorm"
)

type Database struct {
	Connection *gorm.DB

	WalletRepository    domain.WalletRepository
	OperationRepository domain.OperationRepository
}

func NewDatabase(db *gorm.DB, wr domain.WalletRepository, or domain.OperationRepository) *Database {
	return &Database{
		Connection: db,

		WalletRepository:    wr,
		OperationRepository: or,
	}
}

func (d *Database) Migrate() error {
	return d.Connection.AutoMigrate(
		domain.Operation{},
		domain.Wallet{},
	)
}
