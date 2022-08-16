package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/piovani/digital_wallet_2/infra/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetClient() (*gorm.DB, error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Env.MysqlUser,
		config.Env.MysqlPassword,
		config.Env.MysqlHost,
		config.Env.MysqlPort,
		config.Env.MysqlDatabase,
	)

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		return db, err
	}

	return db, nil
}
