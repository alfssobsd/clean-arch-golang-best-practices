package redis_repository

import (
	"clean-arch-golang-best-practices/credit-library/loggerhelper"
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type SessionCacheRepository struct {
	logger      *loggerhelper.CustomLogger
	redisClient *redis.Client
}

type ISessionCacheRepository interface {
	CreateSession(ctx context.Context, key string, payload string) error
	GetSession(ctx context.Context, key string) (string, error)
}

func NewSessionCacheRepository(logger *loggerhelper.CustomLogger, redisClient *redis.Client) *SessionCacheRepository {
	return &SessionCacheRepository{logger: logger, redisClient: redisClient}
}

func (r *SessionCacheRepository) CreateSession(ctx context.Context, key string, payload string) error {
	err := r.redisClient.Set(ctx, key, payload, time.Second*300).Err()
	return err
}

func (r *SessionCacheRepository) GetSession(ctx context.Context, key string) (string, error) {
	value, err := r.redisClient.Get(ctx, key).Result()
	return value, err
}
