package scheme

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/heromicro/omgind/internal/app/contextx"
	"github.com/heromicro/omgind/internal/gen/mainent"
	"github.com/heromicro/omgind/internal/gen/mainent/intercept"
)

// tenantColumn identifies tables whose rows belong to one organization.
func tenantColumn(typ string) string {
	switch typ {
	case "OrgOrgan":
		return "id"
	case "OrgStaff", "OrgDept", "OrgPosition", "SysRole", "SysUserRole", "SysAddress":
		return "org_id"
	}
	return ""
}

// InstallTenantGuards scopes queries and mutations, including transaction clients.
func InstallTenantGuards(client *mainent.Client) {
	client.Intercept(intercept.TraverseFunc(func(ctx context.Context, q intercept.Query) error {
		tenant, scoped := contextx.Tenant(ctx)
		if !scoped {
			return nil
		}
		if col := tenantColumn(q.Type()); col != "" {
			q.WhereP(sql.FieldEQ(col, tenant))
		}
		if q.Type() == "SysUser" && tenant != "" && !contextx.PlatformAdmin(ctx) {
			q.WhereP(func(s *sql.Selector) {
				members := sql.Table("org_staffs")
				s.Where(sql.In(s.C("id"), sql.Select(members.C("user_id")).From(members).Where(sql.EQ(members.C("org_id"), tenant))))
			})
		}
		if tenant != "" && (q.Type() == "SysLogging" || q.Type() == "SysJwtBlock" || q.Type() == "SysTeam" || q.Type() == "SysTeamUser" || q.Type() == "SysAnnex") {
			q.WhereP(sql.FieldEQ("id", ""))
		}
		if q.Type() == "SysRoleMenu" {
			q.WhereP(roleTenantPredicate(tenant))
		}
		return nil
	}))
	client.Use(func(next mainent.Mutator) mainent.Mutator {
		return mainent.MutateFunc(func(ctx context.Context, m mainent.Mutation) (mainent.Value, error) {
			tenant, scoped := contextx.Tenant(ctx)
			if !scoped {
				return next.Mutate(ctx, m)
			}
			col := tenantColumn(m.Type())
			if col == "" && m.Type() != "SysRoleMenu" {
				if contextx.PlatformAdmin(ctx) || tenant == "" {
					return next.Mutate(ctx, m)
				}
				return nil, fmt.Errorf("tenant mutation is not supported for %s", m.Type())
			}
			if tenant == "" {
				return nil, fmt.Errorf("tenant is required")
			}
			if m.Type() == "OrgOrgan" {
				if m.Op().Is(mainent.OpCreate) {
					return nil, fmt.Errorf("create organizations through the onboarding endpoint")
				}
				m.(interface{ WhereP(...func(*sql.Selector)) }).WhereP(sql.FieldEQ("id", tenant))
				if owner, ok := m.Field("owner_id"); ok {
					row, err := client.OrgOrgan.Get(ctx, tenant)
					if err != nil {
						return nil, err
					}
					if row.OwnerID == nil || owner != *row.OwnerID {
						return nil, fmt.Errorf("organization ownership cannot be changed here")
					}
				}
			}
			if col == "org_id" {
				if value, ok := m.Field("org_id"); ok && value != "" && value != tenant {
					return nil, fmt.Errorf("organization mismatch")
				}
				if m.Op().Is(mainent.OpCreate) || func() bool { v, ok := m.Field("org_id"); return ok && v == "" }() {
					if err := m.SetField("org_id", tenant); err != nil {
						return nil, err
					}
				}
				if !m.Op().Is(mainent.OpCreate) {
					m.(interface{ WhereP(...func(*sql.Selector)) }).WhereP(sql.FieldEQ(col, tenant))
				}
			}
			if m.Type() == "SysRoleMenu" {
				if !m.Op().Is(mainent.OpCreate) {
					m.(interface{ WhereP(...func(*sql.Selector)) }).WhereP(roleTenantPredicate(tenant))
				}
			}
			// Resolve references through the same scoped client to reject cross-tenant links.
			for field, typ := range map[string]string{"dept_id": "OrgDept", "posi_id": "OrgPosition", "parent_id": m.Type(), "role_id": "SysRole", "haddr_id": "SysAddress", "iden_addr_id": "SysAddress", "resi_addr_id": "SysAddress"} {
				value, ok := m.Field(field)
				if !ok || value == "" {
					continue
				}
				id, ok := value.(string)
				if !ok {
					continue
				}
				var exists bool
				var err error
				switch typ {
				case "OrgDept":
					exists, err = client.OrgDept.Query().Where(func(s *sql.Selector) { s.Where(sql.EQ(s.C("id"), id)) }).Exist(ctx)
				case "OrgPosition":
					exists, err = client.OrgPosition.Query().Where(func(s *sql.Selector) { s.Where(sql.EQ(s.C("id"), id)) }).Exist(ctx)
				case "SysAddress":
					exists, err = client.SysAddress.Query().Where(func(s *sql.Selector) { s.Where(sql.EQ(s.C("id"), id)) }).Exist(ctx)
				case "SysRole":
					exists, err = client.SysRole.Query().Where(func(s *sql.Selector) { s.Where(sql.EQ(s.C("id"), id)) }).Exist(ctx)
				default:
					continue
				}
				if err != nil {
					return nil, err
				}
				if !exists {
					return nil, fmt.Errorf("invalid tenant reference: %s", field)
				}
			}
			return next.Mutate(ctx, m)
		})
	})
}

func roleTenantPredicate(tenant string) func(*sql.Selector) {
	return func(s *sql.Selector) {
		roles := sql.Table("sys_roles")
		s.Where(sql.In(s.C("role_id"), sql.Select(roles.C("id")).From(roles).Where(sql.EQ(roles.C("org_id"), tenant))))
	}
}
