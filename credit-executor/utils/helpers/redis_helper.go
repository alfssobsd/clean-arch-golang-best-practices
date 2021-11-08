package helpers

import (
	"clean-arch-golang-best-practices/credit-executor/utils/appconfig"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func NewRedisClient(redisConfig appconfig.RedisConfiguration) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.Database,
		PoolSize: redisConfig.PoolSize,
	})
	return rdb
}
