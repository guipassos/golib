//go:generate mockgen -source=${GOFILE} -package=${GOPACKAGE} -destination=${GOPACKAGE}_mock.go
package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	ProviderRedis = "redis"
)

type (
	Options struct {
		ProviderType string
		Addr         string
		Password     string
		DB           int
		Timeout      time.Duration
		Expiration   time.Duration
	}
	Cache interface {
		Ping(ctx context.Context) error
		Get(ctx context.Context, key string, v interface{}) error
		Set(ctx context.Context, key string, v interface{}, expire time.Duration) error
		Del(ctx context.Context, key string) error
	}
)

func New(opts Options) (Cache, error) {
	switch opts.ProviderType {
	case ProviderRedis:
		redisOpts := &redis.Options{
			Addr:        opts.Addr,
			Password:    opts.Password,
			DB:          opts.DB,
			IdleTimeout: opts.Timeout,
		}
		if opts.Timeout > 0 {
			redisOpts.DialTimeout = opts.Timeout
		}
		impl := &redisImpl{
			defaultExpiration: opts.Expiration,
			client:            redis.NewClient(redisOpts),
		}
		if err := impl.Ping(context.Background()); err != nil {
			return nil, ErrPing(err)
		}
		return impl, nil
	default:
		return nil, ErrInvlaidProvider(opts.ProviderType)
	}
}
