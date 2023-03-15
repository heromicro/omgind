package store

import (
	"context"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/heromicro/omgind/pkg/mw/rdb"
	"github.com/mojocn/base64Captcha"
)

type RedisStore struct {
	cli redis.UniversalClient

	prefix     string
	expiration time.Duration
}

// NewRedisStoreWithCli create an instance of a redis store
func NewRedisStore(rdb *rdb.Redis, expiration time.Duration,
	prefix ...string) base64Captcha.Store {
	store := &RedisStore{
		cli:        rdb.Client(),
		expiration: expiration,
	}
	if len(prefix) > 0 {
		store.prefix = prefix[0]
	}
	return store
}

func (rs *RedisStore) Set(id string, value string) error {
	ctx := context.Background()
	_, err := rs.cli.Set(ctx, rs.prefix+":"+id, value, rs.expiration).Result()
	return err
}

func (rs *RedisStore) Get(id string, clear bool) string {
	ctx := context.Background()

	val := rs.cli.Get(ctx, rs.prefix+":"+id).Val()
	if clear && val != "" {
		rs.cli.Del(ctx, "captcha:"+id)
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
