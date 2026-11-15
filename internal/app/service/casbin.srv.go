package service

import (
	"context"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/heromicro/omgind/pkg/global"
)

// LoadCasbinPolicy refreshes policy after a committed database mutation.
func LoadCasbinPolicy(ctx context.Context, e *casbin.SyncedEnforcer) error {
	if global.CFG == nil || !global.CFG.Casbin.Enable {
		return nil
	}
	if e == nil {
		return fmt.Errorf("Casbin is enabled but the policy enforcer was not injected")
	}
	if err := e.LoadPolicy(); err != nil {
		e.ClearPolicy()
		return err
	}
	return nil
}
