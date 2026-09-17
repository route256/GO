package cache

import (
	"fmt"
	"workshopcache/config"

	"github.com/redis/go-redis/v9"
)

func NewCacheClient(cfg *config.CacheConfig) *redis.Client {
	var addr string
	if cfg != nil {
		addr = fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	}

	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return client
}
