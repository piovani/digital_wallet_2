package operation

import (
	"fmt"

	"github.com/piovani/digital_wallet_2/domain"
)

type Input struct {
	UserName string  `json:"user_name"`
	Coin     string  `json:"coin"`
	Value    float64 `json:"value"`
}

type OperationOutput struct {
	ID        int64   `json:"id"`
	Username  string  `json:"user_name"`
	Type      string  `json:"type"`
	Coin      string  `json:"coin"`
	Value     float64 `json:"value"`
	CreatedAt string  `json:"created_at"`
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
