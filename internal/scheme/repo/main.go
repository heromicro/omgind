package repo

import (
	"github.com/google/wire"
)

var RepoSet = wire.NewSet(
	UserSet,
	RoleSet,
	UserRoleSet,
	MenuSet,
	MenuActionSet,
	MenuActionResourceSet,
	RoleMenuSet,

	DemoSet,
	DictSet,
	DictItemSet,
	SysDistrictSet,
	SysAddressSet,
	OrgOrganSet,
	OrgStaffSet,
	OrgPositionSet,
	OrgDeptSet,
	SysTeamSet,
	SysTeamUserSet,
)
