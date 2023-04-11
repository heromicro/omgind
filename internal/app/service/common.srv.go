package service

import (
	"github.com/heromicro/omgind/internal/app/schema"
)

type BaseService struct {
}

func (u *BaseService) getQueryOption(opts ...schema.UserQueryOptions) schema.UserQueryOptions {
	var opt schema.UserQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}
