package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/piovani/digital_wallet_2/infra/price"
)

func read() {
	InitConfig()

	ctx := context.TODO()
	p := price.NewPrice(ctx)

	for {
		fmt.Println(p.GetBtcUsd())
		fmt.Println(p.GetBtcEur())
		fmt.Println(p.GetEthUsd())
		fmt.Println(p.GetEthEur())
		fmt.Println(p.GetAdaUsd())
		fmt.Println(p.GetAdaEur())
		fmt.Println(p.GetXrpUsd())
		fmt.Println(p.GetXrpEur())
		fmt.Println(p.GetDogeUsd())
		fmt.Println(p.GetDogeEur())

		fmt.Println("=========================")

		time.Sleep(time.Second * 10)
	}
}
