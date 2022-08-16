package usecase

import (
	"context"

	"github.com/piovani/digital_wallet_2/infra/price"
)

type CoinPrice struct {
	ctx   context.Context
	price *price.Price
}

func NewCoinPrice(ctx context.Context) *CoinPrice {
	return &CoinPrice{
		ctx:   ctx,
		price: price.NewPrice(ctx),
	}
}

func (c *CoinPrice) Execute() (map[string]float64, error) {
	return map[string]float64{
		"BTC_USD": c.price.GetBtcUsd(),
		"BTC_EUR": c.price.GetBtcEur(),

		"ETH_USD": c.price.GetEthUsd(),
		"ETH_EUR": c.price.GetEthEur(),

		"ADA_USD": c.price.GetAdaUsd(),
		"ADA_EUR": c.price.GetAdaEur(),

		"XRP_USD": c.price.GetXrpUsd(),
		"XRP_EUR": c.price.GetXrpEur(),

		"DOG_USD": c.price.GetDogeUsd(),
		"DOG_EUR": c.price.GetDogeEur(),
	}, nil
}
