package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/piovani/digital_wallet_2/infra/config"
)

func GetClient() (*sql.DB, error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.Env.MysqlUser,
		config.Env.MysqlPassword,
		config.Env.MysqlHost,
		config.Env.MysqlPort,
		config.Env.MysqlDatabase,
	)

	db, err := sql.Open("mysql", url)
	if err != nil {
		return db, err
	}

	if err = db.Ping(); err != nil {
		return db, err
	}

	return db, nil
}
