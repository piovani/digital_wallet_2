package worker

import (
	"context"
	"fmt"
	"time"

	"github.com/piovani/digital_wallet_2/infra/price"
)

type ReadCoinsPrices struct{}

func NewReadCoinsPrices() *ReadCoinsPrices {
	return &ReadCoinsPrices{}
}

func (r *ReadCoinsPrices) Execute() {
	price := price.NewPrice(context.TODO())

	for {
		fmt.Println(price.GetBtcUsd())
		fmt.Println(price.GetBtcEur())
		fmt.Println(price.GetEthUsd())
		fmt.Println(price.GetEthEur())
		fmt.Println(price.GetAdaUsd())
		fmt.Println(price.GetAdaEur())
		fmt.Println(price.GetXrpUsd())
		fmt.Println(price.GetXrpEur())
		fmt.Println(price.GetDogeUsd())
		fmt.Println(price.GetDogeEur())

		fmt.Println("=========================")

		time.Sleep(time.Second * 10)
	}
}
