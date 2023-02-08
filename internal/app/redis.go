package app

import (
	"errors"

	"github.com/go-redis/redis"
	"github.com/heromicro/omgind/pkg/global"
	"github.com/heromicro/omgind/pkg/logger"
)

func InitRedisCli() (redis.Cmdable, func(), error) {

	cli, cleanFunc, err := NewRedisCli()
	if err != nil {
		return nil, cleanFunc, err
	}

	return cli, cleanFunc, nil
}

func NewRedisCli() (redis.Cmdable, func(), error) {
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
