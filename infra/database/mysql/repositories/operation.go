package repositories

import (
	"github.com/piovani/digital_wallet_2/domain"
	"gorm.io/gorm"
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

	return nil
}
