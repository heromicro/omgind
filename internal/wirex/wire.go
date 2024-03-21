//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package wirex

import (
	api_v2 "github.com/heromicro/omgind/internal/api/v2"
	"github.com/heromicro/omgind/internal/app/service"

	"github.com/heromicro/omgind/internal/scheme"
	"github.com/heromicro/omgind/internal/scheme/repo"
	"github.com/heromicro/omgind/pkg/config"

	// "github.com/heromicro/omgind/pkg/mw/asyncq"
	"github.com/heromicro/omgind/pkg/mw/rdb"

	// "github.com/heromicro/omgind/internal/app/api_v2/mock"
	"github.com/google/wire"
	"github.com/heromicro/omgind/internal/router"
	// "github.com/heromicro/omgind/pkg/ws/sockio"
)

// BuildInjector 生成注入器
func BuildInjector(cfg *config.AppConfig) (*Injector, func(), error) {

	wire.Build(
		// mock.MockSet,
		// config.ProviderSet,
		scheme.ProviderSet,
		rdb.ProviderSet,

		InitQueue,
		InitAsynq,

		// asyncq.ProviderSet,

		InitVcode,
		//InitInfluxDB,
		//InitRabbitMQ,
		repo.RepoSet,
		InitAuth,
		// InitCasbin,
		InitGinEngine,
		service.ServiceSet,

		api_v2.APIV2Set,
		router.RouterSet,
		// adapter.CasbinAdapterSet,
		InjectorSet,
		// sockio.ProviderSet,

	)
	return new(Injector), nil, nil
}
