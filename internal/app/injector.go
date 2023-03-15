package app

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/hibiken/asynq"

	"github.com/heromicro/omgind/internal/app/service"
	"github.com/heromicro/omgind/pkg/auth"
	"github.com/heromicro/omgind/pkg/mw/rdb"
)

// InjectorSet 注入Injector
var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

// Injector 注入器(用于初始化完成之后的引用)
type Injector struct {
	Engine         *gin.Engine
	Auth           auth.Auther
	CasbinEnforcer *casbin.SyncedEnforcer
	MenuSrv        *service.Menu
	Rdb            *rdb.Redis
	//InfluxDb       influxdb2.Client
	//RabbitMq       *amqp.Connection
	AsynqCli *asynq.Client
	// AsyncqServer *asyncq.Server
}
