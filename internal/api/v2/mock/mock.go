package mock

import (
	"github.com/google/wire"
)

// MockSet 注入mock
var MockSet = wire.NewSet(
	MenuSet,
	RoleSet,
	SignInSet,
	UserSet,
	DemoSet,
	DictSet,
	SysDistrictSet,
	SysAddressSet,
)
