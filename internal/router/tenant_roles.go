package router

import (
	"github.com/gin-gonic/gin"
	"github.com/heromicro/omgind/internal/app/contextx"
	"github.com/heromicro/omgind/internal/app/ginx"
	"github.com/heromicro/omgind/internal/app/service"
	"github.com/heromicro/omgind/internal/gen/mainent"
	"github.com/heromicro/omgind/internal/gen/mainent/sysrole"
	"github.com/heromicro/omgind/internal/gen/mainent/sysuserrole"
	repo "github.com/heromicro/omgind/internal/scheme/repo/entgo"
	"github.com/heromicro/omgind/pkg/errors"
)

// assignStaffRoles changes only role assignments in the verified tenant.
func (a *Router) assignStaffRoles(c *gin.Context) {
	ctx := c.Request.Context()
	tenant, ok := contextx.Tenant(ctx)
	if !ok || tenant == "" {
		ginx.ResError(c, errors.ErrNoPerm, 403)
		return
	}
	var input struct {
		RoleIDs []string `json:"role_ids"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		ginx.ResError(c, errors.New400Response("Invalid role assignment"))
		return
	}
	staff, err := a.EntCli.OrgStaff.Get(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, errors.ErrNotFound)
		return
	}
	if staff.UserID == nil || *staff.UserID == "" {
		ginx.ResError(c, errors.New400Response("Staff member has no linked account"))
		return
	}
	ids := map[string]bool{}
	for _, id := range input.RoleIDs {
		ids[id] = true
	}
	roles, err := a.EntCli.SysRole.Query().Where(sysrole.IDIn(input.RoleIDs...), sysrole.IsDelEQ(false), sysrole.IsActiveEQ(true)).All(ctx)
	if err != nil {
		ginx.ResError(c, errors.WithStack(err))
		return
	}
	if len(roles) != len(ids) {
		ginx.ResError(c, errors.ErrNoPerm, 403)
		return
	}
	err = repo.WithTx(ctx, a.EntCli, func(tx *mainent.Tx) error {
		if _, err := tx.SysUserRole.Delete().Where(sysuserrole.UserIDEQ(*staff.UserID)).Exec(ctx); err != nil {
			return err
		}
		for id := range ids {
			if _, err := tx.SysUserRole.Create().SetUserID(*staff.UserID).SetRoleID(id).SetOrgID(tenant).Save(ctx); err != nil {
				return err
			}
		}
		return nil
	})
	if err == nil {
		err = service.LoadCasbinPolicy(ctx, a.CasbinEnforcer)
	}
	if err != nil {
		ginx.ResError(c, errors.WithStack(err))
		return
	}
	ginx.ResSuccess(c, nil)
}
