package entity

import (
	"github.com/heromicro/omgind/internal/gen/ent"
	// "github.com/heromicro/omgind/internal/gen/ent/intercept"
)

func (s SysDict) Interceptors() []ent.Interceptor {

	return []ent.Interceptor{
		// intercept.TraverseFunc(func(ctx context.Context, q intercept.Query) error {
		// 	if skip, _: = ctx.Value(); skip {
		// 		return nil
		// 	}
		// 	return nil
		// }),
	}
}
