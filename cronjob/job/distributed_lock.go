package job

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type DistributedLock struct {
	rdb    *redis.Client
	key    string
	expiry time.Duration
}

func NewDistributedLock(rdb *redis.Client, key string, expiry time.Duration) *DistributedLock {
	return &DistributedLock{
		rdb:    rdb,
		key:    key,
		expiry: expiry,
	}
}

func (l *DistributedLock) Acquire(ctx context.Context) (bool, error) {
	return l.rdb.SetNX(ctx, l.key, "1", l.expiry).Result()
}

func (l *DistributedLock) Release(ctx context.Context) error {
	return l.rdb.Del(ctx, l.key).Err()
}