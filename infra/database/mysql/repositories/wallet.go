package repositories

import (
	"github.com/piovani/digital_wallet_2/domain"
	"gorm.io/gorm"
)

type WalletRepository struct {
	DB *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{
		DB: db,
	}
}

func (o WalletRepository) FindByUserName(username string) (*domain.Wallet, error) {
	var w domain.Wallet

	tx := o.DB.Where("user_name = ?", username).First(&w)
	if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
		return &w, tx.Error
	}

	return &w, nil
}

func (o WalletRepository) Save(w *domain.Wallet) error {
	walletDB, err := o.FindByUserName(w.UserName)
	if err != nil {
		return err
	}

	if walletDB.ID == 0 {
		return o.Insert(w)
	}

	return o.Update(w)
}

func (o WalletRepository) Insert(w *domain.Wallet) error {
	tx := o.DB.Create(w)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (o WalletRepository) Update(w *domain.Wallet) error {
	tx := o.DB.Save(&w)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
