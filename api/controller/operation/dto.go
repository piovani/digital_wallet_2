package operation

import (
	"fmt"

	"github.com/piovani/digital_wallet_2/domain"
)

type Input struct {
	UserName string  `json:"user_name" example:"john"`
	Coin     string  `json:"coin" example:"btc, eth, ada, xrp, doge"`
	Value    float64 `json:"value" example:"1.23"`
}

// OperationOutput model info
// @Description User account information
// @Description with user id and username

type OperationOutput struct {
	ID        int64   `json:"id"`         // ID from Operation
	Username  string  `json:"user_name"`  // UserName from Operation
	Type      string  `json:"type"`       // Type from Operation
	Coin      string  `json:"coin"`       // Coin from Operation ex: "btc, eth, ada, xrp, doge"
	Value     float64 `json:"value"`      // Value from Operation
	CreatedAt string  `json:"created_at"` // When operation maked
}

type OperationsOutput struct {
	Operations []OperationOutput `json:"operations"`
}

func GetOpertationsOutput(operations []domain.Operation) OperationsOutput {
	var opertationsOutput []OperationOutput

	for _, elem := range operations {
		opt := OperationOutput{
			ID:        elem.ID,
			Username:  elem.UserName,
			Type:      elem.Type,
			Coin:      elem.Coin,
			Value:     elem.Value,
			CreatedAt: fmt.Sprintf("%s", elem.CreatedAt.Format("02-01-2006 15:04:05")),
		}

		opertationsOutput = append(opertationsOutput, opt)
	}

	return OperationsOutput{Operations: opertationsOutput}
}
