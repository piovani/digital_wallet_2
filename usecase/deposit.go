package usecase

import (
	"github.com/piovani/digital_wallet_2/domain"
	"github.com/piovani/digital_wallet_2/infra/database"
)

const (
	DEPOSIT_TYPE = "deposit"
)

type Deposit struct {
	Database *database.Database
}

func NewDeposit(db *database.Database) *Deposit {
	return &Deposit{
		Database: db,
	}
}

func (d *Deposit) Execute(username, coin string, value float64) error {
	opt, err := domain.NewOperation(username, DEPOSIT_TYPE, coin, value)
	if err != nil {
		return err
	}

	wallet, err := d.Database.WalletRepository.FindByUserName(username)
	if err != nil {
		return err
	}

	if wallet.ID == 0 {
		wallet = domain.NewWallet(username, 0, 0, 0, 0, 0)
	}

	d.walletReceivesValue(wallet, coin, value)

	err = d.Database.WalletRepository.Save(wallet)
	if err != nil {
		return err
	}

	return d.Database.OperationRepository.Insert(opt)
}

func (d *Deposit) walletReceivesValue(w *domain.Wallet, coin string, value float64) {
	switch coin {
	case "btc":
		w.Btc += value
	case "eth":
		w.Eth += value
	case "ada":
		w.Ada += value
	case "xrp":
		w.Xrp += value
	case "doge":
		w.Doge += value
	}
}
