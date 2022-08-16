package repositories

import (
	"fmt"

	"github.com/piovani/digital_wallet_2/domain"
	"gorm.io/gorm"
)

const (
	AUX    = "?, ?, ?, ?, ?"
	FIELDS = "user_name, type, value, created_at, updated_at"
)

type OperationRepository struct {
	DB *gorm.DB
}

func NewOperationRepository(db *gorm.DB) *OperationRepository {
	return &OperationRepository{
		DB: db,
	}
}

func (o OperationRepository) Insert(opt *domain.Operation) error {
	tx := o.DB.Create(opt)
	if tx.Error != nil {
		return tx.Error
	}

	fmt.Println(tx)

	return nil
}
