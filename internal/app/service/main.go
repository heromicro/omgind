package service

import "github.com/google/wire"

// ServiceSet 注入
var ServiceSet = wire.NewSet(
	UserSet,
	MenuSet,
	RoleSet,
	SignInSet,

	DictSet,
	DemoSet,
)
