package scheme

import (
	"context"
	"errors"
	"github.com/google/wire"
	entity "github.com/heromicro/omgind/internal/scheme/entity/mongo"
	"github.com/heromicro/omgind/pkg/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"time"
)

// NewMongoClient 创建并验证 MongoDB 连接，返回释放连接的函数。
func NewMongoClient(cfg *config.AppConfig) (*mongo.Client, func(), error) {
	noop := func() {}
	if cfg == nil || cfg.Mongo.URI == "" || cfg.Mongo.Database == "" {
		return nil, noop, errors.New("Mongo.URI and Mongo.Database are required")
	}
	timeout := time.Duration(cfg.Mongo.ConnectTimeout) * time.Second
	if timeout <= 0 {
		timeout = 10 * time.Second
	}
	client, err := mongo.Connect(options.Client().ApplyURI(cfg.Mongo.URI).SetConnectTimeout(timeout))
	if err != nil {
		return nil, noop, err
	}
	cleanup := func() {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		_ = client.Disconnect(ctx)
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		cleanup()
		return nil, noop, err
	}
	entity.SetDatabaseName(cfg.Mongo.Database)
	return client, cleanup, nil
}

// MongoProviderSet 可与现有 Ent 或 GORM ProviderSet 组合。
var MongoProviderSet = wire.NewSet(NewMongoClient)
