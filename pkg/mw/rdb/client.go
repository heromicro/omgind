package rdb

import (
	"context"
	"errors"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/heromicro/omgind/pkg/config"
)

// c
type Redis struct {
	client redis.UniversalClient
	addrs  []string
	DB     int
}

// n
func New(cfg *config.AppConfig) (*Redis, func(), error) {

	cli := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    []string{cfg.Redis.Addr},
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.Database,
	})

	cleanFunc := func() {
		err := cli.Close()
		if err != nil {
			log.Fatalf("redis close error: %s", err.Error())
		}
	}

	ctx := context.Background()
	result, err := cli.Ping(ctx).Result()
	if err != nil {
		return nil, cleanFunc, err
	}
	if result != "PONG" {
		return nil, cleanFunc, errors.New("redis conn error")
	}

	rdb := &Redis{
		addrs:  []string{cfg.Redis.Addr},
		DB:     cfg.Redis.Database,
		client: cli,
	}
	return rdb, cleanFunc, nil
}

// Client is to return underlying redis interface
func (r *Redis) Client() redis.UniversalClient {
	return r.client
}

// MakeRedisClient is used to fulfill asynq's interface
func (r *Redis) MakeRedisClient() any {
	return r.client
}

// n
func NewClient(cfg *config.AppConfig) (redis.UniversalClient, func(), error) {

	cli := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    []string{cfg.Redis.Addr},
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.Database,
	})

	cleanFunc := func() {
		err := cli.Close()
		if err != nil {
			log.Fatalf("redis close error: %s", err.Error())
		}
	}

	ctx := context.Background()
	result, err := cli.Ping(ctx).Result()
	if err != nil {
		return nil, cleanFunc, err
	}
	if result != "PONG" {
		return nil, cleanFunc, errors.New("redis conn error")
	}
	return cli, cleanFunc, nil
}

var ProviderSet = wire.NewSet(NewClient)

// var ProviderSet = wire.NewSet(New)
