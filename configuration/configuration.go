package configuration

import (
	"github.com/go-redis/redis/v8"
)

type ServiceConfiguration struct {
	CacheDb *redis.Client
}

func NewServiceConfiguration() *ServiceConfiguration {
	//sudo docker run -d -p 6379:6379 redis redis-server --requirepass "fiesta99"
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "fiesta99",
		DB:       0,
	})

	return &ServiceConfiguration{
		CacheDb: rdb,
	}
}
