package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/heromicro/omgind/internal/app/contextx"
	"github.com/heromicro/omgind/internal/app/ginx"
	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/mainent"
	"github.com/heromicro/omgind/internal/gen/mainent/orgorgan"
	"github.com/heromicro/omgind/internal/gen/mainent/orgstaff"
	"github.com/heromicro/omgind/pkg/errors"
	"strings"
)

// TenantMiddleware verifies membership before accepting X-Org-ID.
func TenantMiddleware(client *mainent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetHeader("X-Org-ID")
		uid := ginx.GetUserID(c)
		path := c.Request.URL.Path
		if (c.Request.Method == "POST" && path == "/api/v2/org-organs") || path == "/api/v2/pub/current/password" || path == "/api/v2/pub/current/user" || path == "/api/v2/pub/current/organs" || path == "/api/v2/pub/refresh-token" || path == "/api/v2/pub/signin/exit" {
			c.Next()
			return
		}
		if uid == "" {
			c.Next()
			return
		}

		if id == "" && (strings.HasPrefix(path, "/api/v2/org-") || strings.HasPrefix(path, "/api/v2/sys-addresses")) {
			ginx.ResError(c, errors.ErrNoPerm, 403)
			return
		}
		if id != "" {
			ctx := c.Request.Context()
			active, err := client.OrgOrgan.Query().Where(orgorgan.IDEQ(id), orgorgan.IsDelEQ(false), orgorgan.IsActiveEQ(true)).Exist(ctx)
			if err != nil {
				ginx.ResError(c, errors.WithStack(err))
				return
			}
			if !active {
				ginx.ResError(c, errors.ErrNoPerm, 403)
				return
			}
			member := schema.CheckIsRootUser(ctx, uid)
			if !member {
				member, err = client.OrgStaff.Query().Where(orgstaff.OrgIDEQ(id), orgstaff.UserIDEQ(uid), orgstaff.IsDelEQ(false), orgstaff.IsActiveEQ(true), orgstaff.EmpyStatEQ(1)).Exist(ctx)
				if err != nil {
					ginx.ResError(c, errors.WithStack(err))
					return
				}
			}
			if !member {
				ginx.ResError(c, errors.ErrNoPerm, 403)
				return
			}
		}
		ctx := contextx.WithTenant(c.Request.Context(), id)
		if schema.CheckIsRootUser(ctx, uid) {
			ctx = contextx.WithPlatformAdmin(ctx)
		}
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
