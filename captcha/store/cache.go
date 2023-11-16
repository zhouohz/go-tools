package store

import "context"

type Cache interface {
	Get(ctx context.Context, key string) string
	Set(ctx context.Context, key string, val string, expiresInSeconds int)
	Delete(ctx context.Context, key string)
	Exists(ctx context.Context, key string) bool
	GetType(ctx context.Context) string
	Increment(ctx context.Context, key string, val int64) int64
}
