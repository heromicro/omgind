package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/heromicro/omgind/internal/app/contextx"
	"github.com/heromicro/omgind/internal/app/ginx"
	"github.com/heromicro/omgind/pkg/auth"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/global"
	"github.com/heromicro/omgind/pkg/logger"
)

func wrapUserAuthContext(c *gin.Context, userID string) {
	ginx.SetUserID(c, userID)
	ctx := contextx.NewUserID(c.Request.Context(), userID)
	ctx = logger.NewUserIDContext(ctx, userID)
	c.Request = c.Request.WithContext(ctx)
}

// UserAuthMiddleware 用户授权中间件
func UserAuthMiddleware(a auth.Auther, skippers ...SkipperFunc) gin.HandlerFunc {
	if !global.CFG.JWTAuth.Enable {
		return func(c *gin.Context) {
			wrapUserAuthContext(c, global.CFG.Root.UserName)
			c.Next()
		}
	}

	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		userID, err := a.ParseUserID(c.Request.Context(), ginx.GetToken(c))
		if err != nil {
			if err == auth.ErrInvalidToken {
				if global.CFG.IsDebugMode() {
					wrapUserAuthContext(c, global.CFG.Root.UserName)
					c.Next()
					return
				}
				ginx.ResError(c, errors.ErrInvalidToken)
				return
			}
			ginx.ResError(c, errors.WithStack(err))
			return
		}

		wrapUserAuthContext(c, userID)
		c.Next()
	}
}
