package gateway

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type CacheGW interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string) error
}

type RedisGateway struct {
	redisClient *redis.Client
}

func NewCacheGateway(client *redis.Client) *RedisGateway {
	return &RedisGateway{
		redisClient: client,
	}
}

func (gw *RedisGateway) Get(ctx context.Context, key string) (string, error) {
	val, err := gw.redisClient.Get(ctx, key).Result()
	switch {
	case err == redis.Nil:
		return "", nil
	case err != nil:
		return "", err
	}
	return val, nil
}

func (gw *RedisGateway) Set(ctx context.Context, key string, value string) error {
	err := gw.redisClient.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}
