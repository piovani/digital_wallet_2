package database

import "database/sql"

type Database struct {
	Connection *sql.DB
}

func NewDatabase(db *sql.DB) *Database {
	return &Database{
		Connection: db,
	}
}
