package helpers

import (
	"clean-arch-golang-best-practices/credit-executor/utils/appconfig"
	"clean-arch-golang-best-practices/credit-library/loggerhelper"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func NewRedisClient(logger *loggerhelper.CustomLogger, redisConfig appconfig.RedisConfiguration) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.Database,
		PoolSize: redisConfig.PoolSize,
	})
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		logger.SugarNoTracing().Errorf("Redis ping error = %s", err)
	}
	return redisClient
}
