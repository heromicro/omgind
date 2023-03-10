package app

import (
	"github.com/go-redis/redis/v7"
	"github.com/heromicro/omgind/internal/app/service"
	"github.com/heromicro/omgind/pkg/auth"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// InjectorSet 注入Injector
var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

// Injector 注入器(用于初始化完成之后的引用)
type Injector struct {
	Engine         *gin.Engine
	Auth           auth.Auther
	CasbinEnforcer *casbin.SyncedEnforcer
	MenuSrv        *service.Menu
	RedisCli       redis.Cmdable
	//InfluxDb       influxdb2.Client
	//RabbitMq       *amqp.Connection
}
