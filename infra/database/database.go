package database

import (
	"database/sql"

	"github.com/piovani/digital_wallet_2/domain"
	"github.com/piovani/digital_wallet_2/infra/database/mysql/repositories"
)

type Database struct {
	Connection *sql.DB

	OperationRepository domain.OperationRepository
}

func NewDatabase(db *sql.DB) *Database {
	database := &Database{
		Connection: db,
	}

	database.getRepositores()

	return database
}

func (d *Database) getRepositores() {
	d.OperationRepository = repositories.NewOperationRepository(d.Connection)
}
