package redis

import (
	"context"
	"errors"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/heromicro/omgind/pkg/config"
)

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

var ClientSet = wire.NewSet(NewClient)
