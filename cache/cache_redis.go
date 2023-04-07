package cache

import (
	"context"
	"encoding/json"
	"reflect"
	"time"

	"github.com/go-redis/redis/v8"
)

type (
	redisImpl struct {
		client            *redis.Client
		defaultExpiration time.Duration
	}
)

func (r redisImpl) Ping(ctx context.Context) error {
	return r.client.Ping(ctx).Err()
}

func (r redisImpl) Get(ctx context.Context, key string, v interface{}) error {
	if reflect.TypeOf(v).Kind() != reflect.Ptr {
		return ErrValueIsNotPointer()
	}
	cmd := r.client.Get(ctx, key)
	if err := cmd.Err(); err != nil {
		return err
	}
	b := []byte(cmd.Val())
	return json.Unmarshal(b, v)
}

func (r redisImpl) Set(ctx context.Context, key string, v interface{}, expire time.Duration) error {
	bytes, err := json.Marshal(v)
	if err != nil {
		return err
	}
	if expire <= 0 {
		expire = r.defaultExpiration
	}
	return r.client.Set(ctx, key, bytes, expire).Err()
}

func (r redisImpl) Del(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}
