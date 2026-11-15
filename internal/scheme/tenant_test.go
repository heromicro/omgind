package scheme

import (
	"context"
	"github.com/heromicro/omgind/internal/app/contextx"
	"github.com/heromicro/omgind/internal/gen/mainent/enttest"
	"testing"
)

func TestTenantGuards(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:tenant-guards?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	ctx := context.Background()
	a := client.OrgOrgan.Create().SetName("A").SaveX(ctx)
	b := client.OrgOrgan.Create().SetName("B").SaveX(ctx)
	deptB := client.OrgDept.Create().SetOrgID(b.ID).SetName("B department").SaveX(ctx)
	InstallTenantGuards(client)
	unselected := contextx.WithTenant(ctx, "")
	if rows, err := client.OrgDept.Query().All(unselected); err != nil || len(rows) != 0 {
		t.Fatal("organization rows visible without a selected tenant")
	}
	if _, err := client.OrgPosition.Create().SetName("Unscoped").Save(unselected); err == nil {
		t.Fatal("organization write accepted without a selected tenant")
	}
	tenant := contextx.WithTenant(ctx, a.ID)
	if _, err := client.OrgDept.Get(tenant, deptB.ID); err == nil {
		t.Fatal("cross-tenant row was visible")
	}
	if _, err := client.OrgDept.UpdateOneID(deptB.ID).SetName("intrusion").Save(tenant); err == nil {
		t.Fatal("cross-tenant mutation succeeded")
	}
	if _, err := client.OrgStaff.Create().SetOrgID(a.ID).SetDeptID(deptB.ID).Save(tenant); err == nil {
		t.Fatal("cross-tenant reference accepted")
	}
	own, err := client.OrgPosition.Create().SetName("Own position").Save(tenant)
	if err != nil {
		t.Fatal(err)
	}
	if own.OrgID == nil || *own.OrgID != a.ID {
		t.Fatal("tenant not assigned on create")
	}
	if _, err := client.OrgPosition.Create().SetOrgID(b.ID).Save(tenant); err == nil {
		t.Fatal("spoofed tenant accepted")
	}
	tx, err := client.Tx(tenant)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := tx.OrgDept.Get(tenant, deptB.ID); err == nil {
		t.Fatal("transaction bypassed tenant guard")
	}
	_ = tx.Rollback()
}
