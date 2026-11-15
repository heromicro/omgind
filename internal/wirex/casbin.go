package wirex

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/heromicro/omgind/pkg/config"
	"time"
)

// InitCasbin loads database policies before serving protected requests.
func InitCasbin(cfg *config.AppConfig, adapter persist.Adapter) (*casbin.SyncedEnforcer, func(), error) {
	noop := func() {}
	if cfg == nil {
		return nil, noop, fmt.Errorf("app config is required")
	}
	if cfg.Casbin.Enable && !cfg.JWTAuth.Enable {
		return nil, noop, fmt.Errorf("JWTAuth must be enabled when Casbin is enabled")
	}
	if !cfg.Casbin.Enable {
		m, err := model.NewModelFromString("[request_definition]\nr = sub, obj, act\n[policy_definition]\np = sub, obj, act\n[policy_effect]\ne = some(where (p.eft == allow))\n[matchers]\nm = r.sub == p.sub\n")
		if err != nil {
			return nil, noop, err
		}
		e, err := casbin.NewSyncedEnforcer(m)
		if err != nil {
			return nil, noop, err
		}
		e.EnableEnforce(false)
		return e, noop, nil
	}
	if cfg.Casbin.ModelFile == "" {
		return nil, noop, fmt.Errorf("Casbin.ModelFile is required when Casbin is enabled")
	}
	e, err := casbin.NewSyncedEnforcer(cfg.Casbin.ModelFile, adapter)
	if err != nil {
		return nil, noop, err
	}
	e.EnableLog(cfg.Casbin.Debug)
	if cfg.Casbin.AutoLoad {
		if cfg.Casbin.AutoLoadInternal <= 0 {
			return nil, noop, fmt.Errorf("Casbin.AutoLoadInternal must be positive")
		}
		e.StartAutoLoadPolicy(time.Duration(cfg.Casbin.AutoLoadInternal) * time.Second)
		return e, func() { e.StopAutoLoadPolicy() }, nil
	}
	return e, noop, nil
}
