// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/heromicro/omgind/internal/api/v2"
	"github.com/heromicro/omgind/internal/app/module/adapter"
	"github.com/heromicro/omgind/internal/app/service"
	"github.com/heromicro/omgind/internal/router"
	"github.com/heromicro/omgind/internal/schema"
	"github.com/heromicro/omgind/internal/schema/repo"
	"github.com/heromicro/omgind/pkg/config"
	"github.com/heromicro/omgind/pkg/mw/rdb"
)

import (
	_ "github.com/heromicro/omgind/internal/app/swagger"
)

// Injectors from wire.go:

// BuildInjector 生成注入器
func BuildInjector(cfg *config.AppConfig) (*Injector, func(), error) {
	redis, cleanup, err := rdb.New(cfg)
	if err != nil {
		return nil, nil, err
	}
	auther, cleanup2, err := InitAuth(cfg, redis)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	client, cleanup3, err := schema.New(cfg)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	role := &repo.Role{
		EntCli: client,
	}
	roleMenu := &repo.RoleMenu{
		EntCli: client,
	}
	menuActionResource := &repo.MenuActionResource{
		EntCli: client,
	}
	user := &repo.User{
		EntCli: client,
	}
	userRole := &repo.UserRole{
		EntCli: client,
	}
	casbinAdapter := &adapter.CasbinAdapter{
		RoleRepo:         role,
		RoleMenuRepo:     roleMenu,
		MenuResourceRepo: menuActionResource,
		UserRepo:         user,
		UserRoleRepo:     userRole,
	}
	syncedEnforcer, cleanup4, err := InitCasbin(casbinAdapter)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	dict := &repo.Dict{
		EntCli: client,
	}
	dictItem := &repo.DictItem{
		EntCli: client,
	}
	serviceDict := &service.Dict{
		EntCli:       client,
		DictRepo:     dict,
		DictItemRepo: dictItem,
	}
	api_v2Dict := &api_v2.Dict{
		DictSrv: serviceDict,
	}
	demo := &repo.Demo{
		EntCli: client,
	}
	serviceDemo := &service.Demo{
		EntCli:   client,
		DemoRepo: demo,
	}
	api_v2Demo := &api_v2.Demo{
		DemoSrv: serviceDemo,
	}
	menu := &repo.Menu{
		EntCli: client,
	}
	menuAction := &repo.MenuAction{
		EntCli: client,
	}
	serviceMenu := &service.Menu{
		EntCli:                 client,
		MenuRepo:               menu,
		MenuActionRepo:         menuAction,
		MenuActionResourceRepo: menuActionResource,
	}
	api_v2Menu := &api_v2.Menu{
		MenuSrv: serviceMenu,
	}
	serviceRole := &service.Role{
		EntCli:                 client,
		Enforcer:               syncedEnforcer,
		RoleRepo:               role,
		RoleMenuRepo:           roleMenu,
		UserRepo:               user,
		MenuActionResourceRepo: menuActionResource,
	}
	api_v2Role := &api_v2.Role{
		RoleSrv: serviceRole,
	}
	serviceUser := &service.User{
		EntCli:       client,
		Enforcer:     syncedEnforcer,
		UserRepo:     user,
		UserRoleRepo: userRole,
		RoleRepo:     role,
	}
	api_v2User := &api_v2.User{
		UserSrv: serviceUser,
	}
	vcode := InitVcode(redis)
	signIn := &service.SignIn{
		EntCli:         client,
		Auth:           auther,
		UserRepo:       user,
		UserRoleRepo:   userRole,
		RoleRepo:       role,
		RoleMenuRepo:   roleMenu,
		MenuRepo:       menu,
		MenuActionRepo: menuAction,
		Vcode:          vcode,
	}
	api_v2SignIn := &api_v2.SignIn{
		SigninSrv: signIn,
		Vcode:     vcode,
	}
	queuer, cleanup5, err := InitQueue(cfg, redis)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	sysDistrict := &repo.SysDistrict{
		EntCli: client,
		Queue:  queuer,
	}
	consumer, cleanup6, err := InitAsynq(cfg, queuer)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	serviceSysDistrict := service.New(client, sysDistrict, queuer, consumer)
	api_v2SysDistrict := &api_v2.SysDistrict{
		SysDistrictSrv: serviceSysDistrict,
	}
	sysAddress := &repo.SysAddress{
		EntCli: client,
	}
	serviceSysAddress := &service.SysAddress{
		EntCli:         client,
		SysAddressRepo: sysAddress,
	}
	api_v2SysAddress := &api_v2.SysAddress{
		SysAddressSrv: serviceSysAddress,
	}
	orgOrgan := &repo.OrgOrgan{
		EntCli: client,
	}
	serviceOrgOrgan := &service.OrgOrgan{
		OrgOrganRepo: orgOrgan,
	}
	api_v2OrgOrgan := &api_v2.OrgOrgan{
		OrgOrganSrv: serviceOrgOrgan,
	}
	routerRouter := &router.Router{
		Auth:             auther,
		CasbinEnforcer:   syncedEnforcer,
		DictApiV2:        api_v2Dict,
		DemoAPIV2:        api_v2Demo,
		MenuAPIV2:        api_v2Menu,
		RoleAPIV2:        api_v2Role,
		UserAPIV2:        api_v2User,
		SignInAPIV2:      api_v2SignIn,
		SysDistrictAPIV2: api_v2SysDistrict,
		SysAddressAPIV2:  api_v2SysAddress,
		OrgOrganAPIV2:    api_v2OrgOrgan,
	}
	engine := InitGinEngine(routerRouter)
	injector := &Injector{
		Engine:         engine,
		Auth:           auther,
		CasbinEnforcer: syncedEnforcer,
		MenuSrv:        serviceMenu,
		Rdb:            redis,
		Queue:          queuer,
		Consumer:       consumer,
	}
	return injector, func() {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
