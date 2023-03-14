package app

import (
	"errors"

	"github.com/go-redis/redis/v7"
	"github.com/heromicro/omgind/pkg/global"
	"github.com/heromicro/omgind/pkg/logger"
)

func InitRedisCli() (redis.UniversalClient, func(), error) {

	cli, cleanFunc, err := NewRedisCli()
	if err != nil {
		return nil, cleanFunc, err
	}

	return cli, cleanFunc, nil
}

func NewRedisCli() (redis.UniversalClient, func(), error) {
	cfg := global.CFG
	if !cfg.Redis.Enable {
		return nil, func() {}, errors.New("redis数据库未启用")
	}
	cli := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		DB:       cfg.Redis.Database,
		Password: cfg.Redis.Password,
	})
	cleanFunc := func() {
		err := cli.Close()
		if err != nil {
			logger.Errorf("redis close error: %s", err.Error())
		}
	}
	_, err := cli.Ping().Result()
	if err != nil {
		return nil, cleanFunc, err
	}

	return cli, cleanFunc, nil
}
