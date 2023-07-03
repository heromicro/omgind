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
	"github.com/heromicro/omgind/internal/scheme"
	"github.com/heromicro/omgind/internal/scheme/repo"
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
	client, cleanup3, err := scheme.New(cfg)
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
	serviceDict := &service.Dict{
		EntCli:   client,
		DictRepo: dict,
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
	queuer, cleanup5, err := InitQueue(cfg, redis)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
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
	serviceMenu := service.NewMenuSrv(client, menu, menuAction, menuActionResource, queuer, consumer)
	api_v2Menu := &api_v2.Menu{
		MenuSrv: serviceMenu,
	}
	serviceRole := &service.Role{
		Enforcer:               syncedEnforcer,
		EntCli:                 client,
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
		EntCli:       client,
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
	sysDistrict := &repo.SysDistrict{
		EntCli: client,
		Queue:  queuer,
	}
	serviceSysDistrict := service.NewSysDistrictSrv(client, sysDistrict, queuer, consumer)
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
	orgDept := &repo.OrgDept{
		EntCli: client,
	}
	serviceOrgOrgan := &service.OrgOrgan{
		EntCli:       client,
		OrgOrganRepo: orgOrgan,
		OrgDeptRepo:  orgDept,
	}
	api_v2OrgOrgan := &api_v2.OrgOrgan{
		OrgOrganSrv: serviceOrgOrgan,
	}
	orgStaff := &repo.OrgStaff{
		EntCli: client,
	}
	serviceOrgStaff := &service.OrgStaff{
		EntCli:       client,
		OrgStaffRepo: orgStaff,
	}
	api_v2OrgStaff := &api_v2.OrgStaff{
		OrgStaffSrv: serviceOrgStaff,
	}
	orgPosition := &repo.OrgPosition{
		EntCli: client,
	}
	serviceOrgPosition := &service.OrgPosition{
		EntCli:          client,
		OrgPositionRepo: orgPosition,
	}
	api_v2OrgPosition := &api_v2.OrgPosition{
		OrgPositionSrv: serviceOrgPosition,
	}
	serviceOrgDept := service.NewOrgDeptSrv(client, orgDept, queuer, consumer)
	api_v2OrgDept := &api_v2.OrgDept{
		OrgDeptSrv: serviceOrgDept,
	}
	sysTeam := &repo.SysTeam{
		EntCli: client,
	}
	serviceSysTeam := &service.SysTeam{
		SysTeamRepo: sysTeam,
	}
	api_v2SysTeam := &api_v2.SysTeam{
		SysTeamSrv: serviceSysTeam,
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
		OrgStaffAPIV2:    api_v2OrgStaff,
		OrgPositionAPIV2: api_v2OrgPosition,
		OrgDeptAPIV2:     api_v2OrgDept,
		SysTeamAPIV2:     api_v2SysTeam,
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
