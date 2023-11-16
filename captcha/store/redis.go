package store

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisCache struct {
	Cache *redis.Client
}

func NewRedisCache(r *redis.Client) Cache {
	return &RedisCache{Cache: r}
}

func (l *RedisCache) Get(ctx context.Context, key string) string {
	return l.Cache.Get(ctx, key).Val()
}

func (l *RedisCache) Set(ctx context.Context, key string, val string, expiresInSeconds int) {
	l.Cache.Set(ctx, key, val, time.Duration(expiresInSeconds)*time.Second)
}

func (l *RedisCache) Delete(ctx context.Context, key string) {
	l.Cache.Del(ctx, key)
}

func (l *RedisCache) Exists(ctx context.Context, key string) bool {
	return l.Cache.Exists(ctx, key).Val() > 0
}

func (l *RedisCache) GetType(ctx context.Context) string {
	return "redis"
}

func (l *RedisCache) Increment(ctx context.Context, key string, val int64) int64 {
	return l.Cache.IncrBy(ctx, key, val).Val()
}
