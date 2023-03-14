package store

import (
	"strings"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/mojocn/base64Captcha"
)

type RedisStore struct {
	cli redis.UniversalClient

	prefix     string
	expiration time.Duration
}

// NewStore create an instance of a redis store
func NewStore(opts *redis.Options, expiration time.Duration,
	prefix ...string) base64Captcha.Store {
	if opts == nil {
		panic("options cannot be nil")
	}
	return NewRedisStore(
		redis.NewClient(opts),
		expiration,
		prefix...,
	)
}

// NewRedisStoreWithCli create an instance of a redis store
func NewRedisStore(cli redis.UniversalClient, expiration time.Duration,
	prefix ...string) base64Captcha.Store {
	store := &RedisStore{
		cli:        cli,
		expiration: expiration,
	}
	if len(prefix) > 0 {
		store.prefix = prefix[0]
	}
	return store
}

func (rs *RedisStore) Set(id string, value string) error {
	_, err := rs.cli.Set(rs.prefix+":"+id, value, rs.expiration).Result()
	return err
}

func (rs *RedisStore) Get(id string, clear bool) string {
	val := rs.cli.Get(rs.prefix + ":" + id).Val()
	if clear && val != "" {
		rs.cli.Del("captcha:" + id)
	}
	return val
}

func (rs *RedisStore) Verify(id, answer string, clear bool) bool {
	vv := rs.Get(id, clear)
	vv = strings.TrimSpace(strings.ToLower(vv))
	return vv == strings.TrimSpace(strings.ToLower(answer))
}

// Logger Define the log output interface
type Logger interface {
	Printf(format string, args ...any)
}
