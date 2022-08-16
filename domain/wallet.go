package domain

type Wallet struct {
	ID       int64   `gorm:"primary_key;autoIncrement"`
	UserName string  `gorm:"type:varchar(255);column:user_name"`
	Btc      float64 `gorm:"type:double;column:btc"`
	Eth      float64 `gorm:"type:double;column:eth"`
	Ada      float64 `gorm:"type:double;column:ada"`
	Xrp      float64 `gorm:"type:double;column:xrp"`
	Doge     float64 `gorm:"type:double;column:doge"`
}

func NewWallet(username string, btc, eth, ada, xrp, doge float64) *Wallet {
	return &Wallet{
		UserName: username,
		Btc:      btc,
		Eth:      eth,
		Ada:      ada,
		Xrp:      xrp,
		Doge:     doge,
	}
}

type WalletRepository interface {
	FindByUserName(username string) (*Wallet, error)
	Save(w *Wallet) error
	Insert(w *Wallet) error
	Update(w *Wallet) error
}
