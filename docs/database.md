# 数据库存储

默认 `scheme.ProviderSet` 为 `EntProviderSet`。客户端工厂均返回 `(client, cleanup, error)`，启动失败会释放已打开的连接，应用退出时调用 cleanup。

| 存储 | 工厂 | 配置 | Wire Provider |
| --- | --- | --- | --- |
| Ent | NewEnt / NewEntClient | Ent + MySQL/Postgres/Sqlite3 | EntProviderSet |
| GORM v1 | NewGormClient | Gorm + MySQL/Postgres/Sqlite3 | GormProviderSet |
| MongoDB v2 | NewMongoClient | Mongo.URI / Database / ConnectTimeout | MongoProviderSet |

Ent 与 GORM 的 DBType 支持 mysql、postgres、sqlite3。GORM 不自动迁移表；Ent 的迁移由 Ent.EnableAutoMigrate 控制。MongoDB 使用官方驱动，初始化时验证连接，不自动创建索引。

在 internal/wirex/wire.go 中将 scheme.ProviderSet 与 scheme.MongoProviderSet 组合，可供不同模型使用不同数据库。仅提供 Provider 不会创建未被业务依赖的客户端；添加消费者后需重新运行 Wire。将 Ent 替换为 GORM 时，已有服务的 Ent 仓储依赖也需要逐个替换，不能仅修改配置。

configs/config.toml 是开发配置的符号链接。开发和生产的 GORM 默认 PostgreSQL，Docker 配置默认 MySQL，与各自 Ent 配置一致。MongoDB 示例地址为 localhost:27017，部署时应覆盖为实际地址及认证信息。

数据库客户端和配置已提供；生成的 MongoDB/GORM repo 中的分页等公共辅助逻辑仍需补齐，客户端支持不表示这些业务仓储均已可运行。

## GORM system and organization entities

Run `go run ./cmd/generator/gorm` from the project root to regenerate the active Sys/Org models from the resolved Ent schema. Models preserve table names, storage keys, nullable fields, mixin fields and configured indexes. Foreign-key scalar fields are included; Ent navigation edges are not GORM associations. SysCasbinRule is commented out in the Ent source and is not generated. Client initialization does not automatically migrate these tables.
