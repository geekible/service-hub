package configuration

import (
	"github.com/go-redis/redis/v8"
)

type ServiceConfiguration struct {
	CacheDb *redis.Client
}

func NewServiceConfiguration() *ServiceConfiguration {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return &ServiceConfiguration{
		CacheDb: rdb,
	}
}
