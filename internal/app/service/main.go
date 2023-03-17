package service

import "github.com/google/wire"

// ServiceSet 注入
var ServiceSet = wire.NewSet(
	SysUserSet,
	SysMenuSet,
	SysRoleSet,
	SignInSet,

	DictSet,
	DemoSet,

	SysDistrictSet,
	SysAddressSet,
	OrgOrganSet,
	OrgStaffSet,
)
