package grpc

import (
	"fmt"

	"github.com/piovani/digital_wallet_2/infra/database"
	"github.com/piovani/digital_wallet_2/infra/redis"
)

type Grpc struct {
	Redis    *redis.Redis
	Database *database.Database
}

func NewGrpc(r *redis.Redis, db *database.Database) *Grpc {
	return &Grpc{
		Redis:    r,
		Database: db,
	}
}

func (g *Grpc) Start() error {
	fmt.Println("AQQUi2")

	return nil
}
