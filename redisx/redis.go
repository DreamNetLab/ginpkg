package redisx

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

type ClientConfig struct {
	Type int

	MaxActive   int
	MaxIdle     int
	MaxIdleTime time.Duration

	Addr     []string
	Password string
}

type RedisxClient struct {
	Rdb redis.UniversalClient
}

func Setup(config *ClientConfig) (*RedisxClient, error) {
	opts := redis.UniversalOptions{
		MaxActiveConns:  config.MaxActive,
		MaxIdleConns:    config.MaxIdle,
		ConnMaxIdleTime: config.MaxIdleTime,

		Addrs:    config.Addr,
		Password: config.Password,
	}

	redisClient := &RedisxClient{
		Rdb: redis.NewUniversalClient(&opts),
	}

	ctx := context.Background()

	_, err := redisClient.Rdb.Ping(ctx).Result()
	if err != nil {
		_ = redisClient.Rdb.Close()
		return nil, err
	}

	return redisClient, nil
}

func (rc *RedisxClient) SetString(key string, data string, expire int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := rc.Rdb.Set(ctx, key, data, time.Duration(expire)*time.Second).Err()
	if err != nil {
		return err
	}

	return nil
}

func (rc *RedisxClient) SetStringNx(key string, data any, expire int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := rc.Rdb.SetNX(ctx, key, data, time.Duration(expire)*time.Second).Err()
	if err != nil {
		return err
	}

	return nil
}

func (rc *RedisxClient) Exist(key string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := rc.Rdb.Get(ctx, key).Result()

	return err == nil
}

func (rc *RedisxClient) Get(key string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	value, err := rc.Rdb.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return []byte(value), nil
}

func (rc *RedisxClient) Delete(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := rc.Rdb.Del(ctx, key).Result()
	if err != nil {
		return err
	}

	return nil
}
