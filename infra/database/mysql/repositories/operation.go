package repositories

import (
	"database/sql"
	"fmt"

	"github.com/piovani/digital_wallet_2/domain"
)

const (
	AUX    = "?, ?, ?, ?, ?"
	FIELDS = "user_name, type, value, created_at, updated_at"
)

type OperationRepository struct {
	DB *sql.DB
}

func NewOperationRepository(db *sql.DB) *OperationRepository {
	return &OperationRepository{
		DB: db,
	}
}

func (o OperationRepository) Insert(opt *domain.Operation) error {
	statement, err := o.DB.Prepare(fmt.Sprintf("INSERT INTO operations (%s) VALUE (%s)", FIELDS, AUX))
	if err != nil {
		return err
	}
	defer statement.Close()

	r, err := statement.Exec(opt.UserName, opt.Type, opt.Value, opt.CreatedAt, opt.UpdatedAt)
	if err != nil {
		fmt.Println(err)
		return err
	}

	id, err := r.LastInsertId()
	if err != nil {
		return err
	}

	opt.ID = id

	return nil
}
