package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/heromicro/omgind/internal/app/contextx"
	"github.com/heromicro/omgind/internal/app/ginx"
	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/global"
)

// CasbinMiddleware casbin中间件
func CasbinMiddleware(enforcer *casbin.SyncedEnforcer, skippers ...SkipperFunc) gin.HandlerFunc {
	cfg := global.CFG.Casbin
	if !cfg.Enable {
		return EmptyMiddleware()
	}

	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		userID := ginx.GetUserID(c)
		if userID != "" && schema.CheckIsRootUser(c.Request.Context(), userID) {
			c.Next()
			return
		}
		path := c.Request.URL.Path
		method := c.Request.Method
		public := (method == "POST" && path == "/api/v2/pub/signin") || (method == "GET" && (path == "/api/v2/pub/signin/captcha" || path == "/api/v2/pub/signin/captchaid"))
		if public {
			c.Next()
			return
		}
		if userID == "" {
			ginx.ResError(c, errors.ErrInvalidToken, 401)
			return
		}
		own := (method == "GET" && (path == "/api/v2/pub/current/organs" || path == "/api/v2/pub/current/user" || path == "/api/v2/pub/current/menutree")) || (method == "PUT" && path == "/api/v2/pub/current/password") || (method == "POST" && (path == "/api/v2/pub/signin/exit" || path == "/api/v2/pub/refresh-token"))
		if own || (method == "POST" && path == "/api/v2/org-organs") {
			c.Next()
			return
		}
		if enforcer == nil {
			ginx.ResError(c, errors.ErrNoPerm, 403)
			return
		}
		tenant, _ := contextx.Tenant(c.Request.Context())
		p := c.Request.URL.Path
		m := c.Request.Method
		if b, err := enforcer.Enforce(ginx.GetUserID(c), tenant, p, m); err != nil {
			ginx.ResError(c, errors.WithStack(err))
			return
		} else if !b {
			ginx.ResError(c, errors.ErrNoPerm, 403)
			return
		}
		c.Next()
	}
}
