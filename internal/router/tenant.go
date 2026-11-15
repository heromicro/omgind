package router

import (
	"github.com/gin-gonic/gin"
	"github.com/heromicro/omgind/internal/app/ginx"
	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/mainent/orgorgan"
	"github.com/heromicro/omgind/internal/gen/mainent/orgstaff"
	"github.com/heromicro/omgind/pkg/errors"
)

// currentOrganizations exposes only organizations the authenticated account can enter.
func (a *Router) currentOrganizations(c *gin.Context) {
	uid := ginx.GetUserID(c)
	if uid == "" {
		ginx.ResError(c, errors.ErrInvalidToken, 401)
		return
	}
	ctx := c.Request.Context()
	query := a.EntCli.OrgOrgan.Query().Where(orgorgan.IsDelEQ(false), orgorgan.IsActiveEQ(true))
	if !schema.CheckIsRootUser(ctx, uid) {
		staff, err := a.EntCli.OrgStaff.Query().Where(orgstaff.UserIDEQ(uid), orgstaff.IsDelEQ(false), orgstaff.IsActiveEQ(true), orgstaff.EmpyStatEQ(1)).All(ctx)
		if err != nil {
			ginx.ResError(c, errors.WithStack(err))
			return
		}
		ids := make([]string, 0, len(staff))
		for _, member := range staff {
			if member.OrgID != nil {
				ids = append(ids, *member.OrgID)
			}
		}
		query.Where(orgorgan.IDIn(ids...))
	}
	rows, err := query.All(ctx)
	if err != nil {
		ginx.ResError(c, errors.WithStack(err))
		return
	}
	type organization struct {
		ID   string  `json:"id"`
		Name *string `json:"name"`
	}
	result := make([]organization, 0, len(rows))
	for _, row := range rows {
		result = append(result, organization{row.ID, row.Name})
	}
	ginx.ResSuccess(c, result)
}
