package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/heromicro/omgind/pkg/mw/rdb"
)

// Config redis配置参数
type Config struct {
	Addrs     []string // 地址(IP:Port)
	DB        int      // 数据库
	Password  string   // 密码
	KeyPrefix string   // 存储key的前缀
}

// NewStore 创建基于redis存储实例
func NewStore(cfg *Config, rdb *rdb.Redis) *Store {
	return &Store{
		cli:    rdb.Client(),
		prefix: cfg.KeyPrefix,
	}
}

// NewStoreWithClient 使用redis客户端创建存储实例
func NewStoreWithClient(cli redis.UniversalClient, keyPrefix string) *Store {
	return &Store{
		cli:    cli,
		prefix: keyPrefix,
	}
}

// NewStoreWithClusterClient 使用redis集群客户端创建存储实例
func NewStoreWithClusterClient(cli *redis.ClusterClient, keyPrefix string) *Store {
	return &Store{
		cli:    cli,
		prefix: keyPrefix,
	}
}

type redisClienter interface {
	Get(key string) *redis.StringCmd
	Set(key string, value any, expiration time.Duration) *redis.StatusCmd
	Expire(key string, expiration time.Duration) *redis.BoolCmd
	Exists(keys ...string) *redis.IntCmd
	TxPipeline() redis.Pipeliner
	Del(keys ...string) *redis.IntCmd
	Close() error
}

// Store redis存储
type Store struct {
	cli    redis.UniversalClient
	prefix string
}

func (s *Store) wrapperKey(key string) string {
	return fmt.Sprintf("%s%s", s.prefix, key)
}

// Set ...
func (s *Store) Set(ctx context.Context, tokenString string, expiration time.Duration) error {
	cmd := s.cli.Set(ctx, s.wrapperKey(tokenString), "1", expiration)
	return cmd.Err()
}

// Delete ...
func (s *Store) Delete(ctx context.Context, tokenString string) (bool, error) {
	cmd := s.cli.Del(ctx, s.wrapperKey(tokenString))
	if err := cmd.Err(); err != nil {
		return false, err
	}
	return cmd.Val() > 0, nil
}

// Check ...
func (s *Store) Check(ctx context.Context, tokenString string) (bool, error) {
	cmd := s.cli.Exists(ctx, s.wrapperKey(tokenString))
	if err := cmd.Err(); err != nil {
		return false, err
	}
	return cmd.Val() > 0, nil
}

// Close ...
func (s *Store) Close() error {
	return s.cli.Close()
}
