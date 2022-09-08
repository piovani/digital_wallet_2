package usecase

import (
	"context"

	"github.com/piovani/digital_wallet_2/infra/database"
	"github.com/piovani/digital_wallet_2/infra/price"
)

type Balance struct {
	DB *database.Database
}

type Bal struct {
	Input    float64
	Output   float64
	Total    float64
	USD      float64
	USDTotal float64
	EUR      float64
	EURTotal float64
}

func NewBalance(db *database.Database) *Balance {
	return &Balance{
		DB: db,
	}
}

func (c *Balance) Execute(username string) (map[string]Bal, error) {
	balanco := map[string]Bal{}

	operations, err := c.DB.OperationRepository.FindByUserName(username)
	if err != nil {
		return balanco, err
	}

	for _, operation := range operations {
		b := balanco[operation.Coin]

		if operation.Type == "deposit" {
			b.Input += operation.Value
			b.Total += operation.Value
		}

		if operation.Type == "withdraw" {
			b.Output += operation.Value
			b.Total -= operation.Value
		}

		b.USD = c.GetValueCurrent("usd", operation.Coin)
		b.USDTotal = b.USD * b.Total

		b.EUR = c.GetValueCurrent("eur", operation.Coin)
		b.EURTotal = b.EUR * b.Total

		balanco[operation.Coin] = b
	}

	return balanco, nil
}

func (c *Balance) GetValueCurrent(base, coin string) float64 {
	price := price.NewPrice(context.TODO())

	switch coin {
	case "btc":
		switch base {
		case "usd":
			return price.GetBtcUsd()
		case "eur":
			return price.GetBtcEur()
		}
	case "eth":
		switch base {
		case "usd":
			return price.GetEthUsd()
		case "eur":
			return price.GetEthEur()
		}
	case "ada":
		switch base {
		case "usd":
			return price.GetAdaUsd()
		case "eur":
			return price.GetAdaEur()
		}
	case "xrp":
		switch base {
		case "usd":
			return price.GetXrpUsd()
		case "eur":
			return price.GetXrpEur()
		}
	case "doge":
		switch base {
		case "usd":
			return price.GetDogeUsd()
		case "eur":
			return price.GetDogeEur()
		}
	}

	return 1
}
