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
	"github.com/heromicro/omgind/internal/schema/repo"
)

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/heromicro/omgind/internal/app/swagger"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// Injectors from wire.go:

// BuildInjector 生成注入器
func BuildInjector() (*Injector, func(), error) {
	auther, cleanup, err := InitAuth()
	if err != nil {
		return nil, nil, err
	}
	client, cleanup2, err := InitEntClient()
	if err != nil {
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
	syncedEnforcer, cleanup3, err := InitCasbin(casbinAdapter)
	if err != nil {
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
		MenuRepo:               menu,
		MenuActionRepo:         menuAction,
		MenuActionResourceRepo: menuActionResource,
	}
	api_v2Menu := &api_v2.Menu{
		MenuSrv: serviceMenu,
	}
	serviceRole := &service.Role{
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
		Enforcer:     syncedEnforcer,
		UserRepo:     user,
		UserRoleRepo: userRole,
		RoleRepo:     role,
	}
	api_v2User := &api_v2.User{
		UserSrv: serviceUser,
	}
	cmdable, cleanup4, err := InitRedisCli()
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	vcode := InitVcode(cmdable)
	signIn := &service.SignIn{
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
	sysDistrict := &repo.SysDistrict{
		EntCli: client,
	}
	serviceSysDistrict := &service.SysDistrict{
		SysDistrictRepo: sysDistrict,
	}
	api_v2SysDistrict := &api_v2.SysDistrict{
		SysDistrictSrv: serviceSysDistrict,
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
	}
	engine := InitGinEngine(routerRouter)
	injector := &Injector{
		Engine:         engine,
		Auth:           auther,
		CasbinEnforcer: syncedEnforcer,
		MenuSrv:        serviceMenu,
		RedisCli:       cmdable,
	}
	return injector, func() {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
