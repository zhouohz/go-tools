package store

import (
	"context"
	"strconv"
	"sync"
	"time"
)

// MemoryCache struct represents an in-memory cache using sync.Map.
type MemoryCache struct {
	cache sync.Map
}

// NewMemoryCache creates a new instance of MemoryCache.
func NewMemoryCache() *MemoryCache {
	return &MemoryCache{}
}

// Get retrieves a value from the in-memory cache using the provided key.
func (m *MemoryCache) Get(ctx context.Context, key string) string {
	val, found := m.cache.Load(key)
	if !found {
		return ""
	}

	item, ok := val.(cachedItem)
	if !ok || item.expired() {
		return ""
	}

	return item.value
}

// Set sets a key-value pair in the in-memory cache with an optional expiration time.
func (m *MemoryCache) Set(ctx context.Context, key string, val string, expiresInSeconds int) {
	expirationTime := time.Now().Add(time.Duration(expiresInSeconds) * time.Second)
	m.cache.Store(key, cachedItem{value: val, expirationTime: expirationTime})
}

// Delete deletes a key from the in-memory cache.
func (m *MemoryCache) Delete(ctx context.Context, key string) {
	m.cache.Delete(key)
}

// Exists checks if a key exists in the in-memory cache.
func (m *MemoryCache) Exists(ctx context.Context, key string) bool {
	val, found := m.cache.Load(key)
	if !found {
		return false
	}

	item, ok := val.(cachedItem)
	return ok && !item.expired()
}

// GetType returns the type of the cache, which is "memory".
func (m *MemoryCache) GetType(ctx context.Context) string {
	return "memory"
}

// Increment increments the value of a key in the in-memory cache by a specified value.
func (m *MemoryCache) Increment(ctx context.Context, key string, val int64) int64 {
	item, found := m.cache.Load(key)
	if !found {
		return 0
	}

	cachedItem, ok := item.(cachedItem)
	if !ok || cachedItem.expired() {
		return 0
	}

	// Check if the current value can be converted to int64
	currentValue, err := strconv.ParseInt(cachedItem.value, 10, 64)
	if err != nil {
		return 0
	}

	// Increment the value
	newValue := currentValue + val
	cachedItem.value = strconv.FormatInt(newValue, 10)
	m.cache.Store(key, cachedItem)

	return newValue
}

// cachedItem struct represents an item in the cache.
type cachedItem struct {
	value          string
	expirationTime time.Time
}

// expired returns true if the cached item has expired.
func (item *cachedItem) expired() bool {
	return !item.expirationTime.IsZero() && time.Now().After(item.expirationTime)
}
