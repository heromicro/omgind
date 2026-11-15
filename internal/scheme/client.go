package scheme

import (
	"github.com/google/wire"
	"github.com/heromicro/omgind/internal/gen/mainent"
	"github.com/heromicro/omgind/pkg/config"
)

// New 保留原有调用入口；新代码请使用 NewEnt。
func New(cfg *config.AppConfig) (*mainent.Client, func(), error) {
	return NewEnt(cfg)
}

var EntProviderSet = wire.NewSet(NewEnt)

// ProviderSet 默认保留现有 Ent 存储。需要 MongoDB 时在应用 Wire 中组合
// EntProviderSet, MongoProviderSet；GORM 项目则组合 GormProviderSet, MongoProviderSet。
var ProviderSet = EntProviderSet
