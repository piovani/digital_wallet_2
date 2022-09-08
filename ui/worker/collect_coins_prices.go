package worker

import (
	"context"
	"time"

	"github.com/piovani/digital_wallet_2/infra/price"
)

type CollectCoinsPrices struct{}

func NewCollectCoinsPrices() *CollectCoinsPrices {
	return &CollectCoinsPrices{}
}

func (c *CollectCoinsPrices) Execute() {
	price := price.NewPrice(context.TODO())

	for {
		go price.CollectBtcUsd()
		go price.CollectBtcEur()
		go price.CollectEthUsd()
		go price.CollecEtcEur()
		go price.CollectAdaUsd()
		go price.CollectAdaEur()
		go price.CollectXrpUsd()
		go price.CollectXrpEur()
		go price.CollectDogeUsd()
		go price.CollectDogeEur()

		time.Sleep(time.Second * 50)
	}
}
