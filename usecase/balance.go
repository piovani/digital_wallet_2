package usecase

import (
	"github.com/piovani/digital_wallet_2/infra/database"
)

type Balance struct {
	DB *database.Database
}

type Bal struct {
	Input  float64
	Output float64
	Total  float64
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

		balanco[operation.Coin] = b
	}

	return balanco, nil
}
