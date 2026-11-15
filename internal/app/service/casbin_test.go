package service

import (
	"context"
	"github.com/casbin/casbin/v2"
	casbinmodel "github.com/casbin/casbin/v2/model"
	"github.com/heromicro/omgind/internal/gen/mainent/enttest"
	repo "github.com/heromicro/omgind/internal/scheme/repo/entgo"
	"github.com/heromicro/omgind/pkg/config"
	"github.com/heromicro/omgind/pkg/global"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

type countingPolicyAdapter struct{ loads int }

func (a *countingPolicyAdapter) LoadPolicy(casbinmodel.Model) error                      { a.loads++; return nil }
func (*countingPolicyAdapter) SavePolicy(casbinmodel.Model) error                        { return nil }
func (*countingPolicyAdapter) AddPolicy(string, string, []string) error                  { return nil }
func (*countingPolicyAdapter) RemovePolicy(string, string, []string) error               { return nil }
func (*countingPolicyAdapter) RemoveFilteredPolicy(string, string, int, ...string) error { return nil }

func TestMenuStartupPolicyRefresh(t *testing.T) {
	previous := global.CFG
	global.CFG = &config.AppConfig{}
	global.CFG.Casbin.Enable = true
	defer func() { global.CFG = previous }()
	client := enttest.Open(t, "sqlite3", "file:menu-policy-startup?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	ctx := context.Background()
	client.SysMenu.Create().SetName("Menu").SetIcon("").SetRouter("").SetLevel(1).SaveX(ctx)
	adapter := &countingPolicyAdapter{}
	enforcer, err := casbin.NewSyncedEnforcer("../../../configs/rbac_model.conf", adapter)
	if err != nil {
		t.Fatal(err)
	}
	menu := &Menu{MenuRepo: &repo.Menu{EntCli: client}, Enforcer: enforcer}
	before := adapter.loads
	if err := menu.InitData(ctx, "unused.yaml"); err != nil {
		t.Fatal(err)
	}
	if adapter.loads != before+1 {
		t.Fatal("startup did not refresh policy")
	}
	menu.Enforcer = nil
	if err := menu.InitData(ctx, "unused.yaml"); err == nil {
		t.Fatal("missing enforcer should return an error")
	}
	global.CFG.Casbin.Enable = false
	if err := menu.InitData(ctx, "unused.yaml"); err != nil {
		t.Fatal(err)
	}
}
