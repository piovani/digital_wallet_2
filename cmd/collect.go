package cmd

import (
	"context"

	"github.com/piovani/digital_wallet_2/infra/price"
)

func collect() {
	InitConfig()

	ctx := context.TODO()
	price := price.NewPrice(ctx)

	price.Collect()
}
