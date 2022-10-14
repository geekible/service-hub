package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"org.geekible.servicehub/configuration"
)

type CacheService struct {
	client *redis.Client
}

func NewCacheService(config *configuration.ServiceConfiguration) *CacheService {
	return &CacheService{client: config.CacheDb}
}

func (cache *CacheService) SetKey(key, value string) error {
	if err := cache.client.Set(context.Background(), key, value, 0).Err(); err != nil {
		return err
	}
	return nil
}

func (cache *CacheService) GetKey(key string) (string, error) {
	val, err := cache.client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

func (cache *CacheService) RemoveKey(key string) error {
	if err := cache.client.Del(context.Background(), key).Err(); err != nil {
		return err
	}
	return nil
}
