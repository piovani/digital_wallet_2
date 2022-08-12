package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/piovani/digital_wallet_2/infra/config"
)

type Redis struct {
	Client *redis.Client
}

func NewRedis() *Redis {
	return &Redis{
		Client: redis.NewClient(&redis.Options{
			Addr:     config.Env.RedisDSN,
			Username: config.Env.RedisUser,
			Password: config.Env.RedisPassword,
		}),
	}
}

func (r *Redis) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	return r.Client.Set(ctx, key, value, ttl).Err()
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}
