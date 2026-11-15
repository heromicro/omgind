package entity

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"time"
)

// Model 是 MongoDB 实体的公共字段。
type Model struct {
	ID        string    `bson:"_id"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

func (Model) CollectionName(name string) string { return name }

type collectionNamer interface{ CollectionName() string }

func GetCollection(ctx context.Context, client *mongo.Client, model collectionNamer) *mongo.Collection {
	if databaseName == "" {
		panic("mongo: call SetDatabaseName before accessing collections")
	}
	return client.Database(databaseName).Collection(model.CollectionName())
}

// databaseName 由连接配置显式设置，避免误写默认数据库。
var databaseName string

func SetDatabaseName(name string) { databaseName = name }

func (Model) CreateIndexes(ctx context.Context, client *mongo.Client, model collectionNamer, indexes []mongo.IndexModel) error {
	_, err := GetCollection(ctx, client, model).Indexes().CreateMany(ctx, indexes)
	return err
}
