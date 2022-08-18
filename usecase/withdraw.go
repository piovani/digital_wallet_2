package usecase

import (
	"fmt"

	"github.com/piovani/digital_wallet_2/domain"
	"github.com/piovani/digital_wallet_2/infra/database"
)

const (
	WITHDRAW_TYPE = "withdraw"
)

type Withdraw struct {
	Database *database.Database
}

func NewWithdraw(db *database.Database) *Withdraw {
	return &Withdraw{
		Database: db,
	}
}

func (d *Withdraw) Execute(username, coin string, value float64) error {
	wallet, err := d.Database.WalletRepository.FindByUserName(username)
	if err != nil {
		return err
	}

	opt, err := domain.NewOperation(username, WITHDRAW_TYPE, coin, value)
	if err != nil {
		return err
	}

	if !d.checkThereValueInWallet(wallet, coin, value) {
		return fmt.Errorf("Insufficient portfolio value")
	}

	d.subValueInWallet(wallet, coin, value)

	err = d.Database.WalletRepository.Save(wallet)
	if err != nil {
		return err
	}

	return d.Database.OperationRepository.Insert(opt)
}

func (d *Withdraw) checkThereValueInWallet(w *domain.Wallet, coin string, value float64) bool {
	inWallet := float64(0)

	switch coin {
	case "btc":
		inWallet = w.Btc
	case "eth":
		inWallet = w.Eth
	case "ada":
		inWallet = w.Ada
	case "xrp":
		inWallet = w.Xrp
	case "doge":
		inWallet = w.Doge
	}

	return (inWallet - value) > 0
}

func (d *Withdraw) subValueInWallet(w *domain.Wallet, coin string, value float64) {
	switch coin {
	case "btc":
		w.Btc -= value
	case "eth":
		w.Eth -= value
	case "ada":
		w.Ada -= value
	case "xrp":
		w.Xrp -= value
	case "doge":
		w.Doge -= value
	}
}
