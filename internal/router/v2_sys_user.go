package router

import (
	"github.com/gin-gonic/gin"
	api_v2 "github.com/heromicro/omgind/internal/api/v2"
)

func (r *Router) initUserRouterV2(urg *gin.RouterGroup, api *api_v2.User, pathcomp string) {

	gUser := urg.Group(pathcomp)
	{
		gUser.GET("", api.Query)
		gUser.GET(":id", api.Get)
		gUser.POST("", api.Create)
		gUser.PUT(":id", api.Update)
		gUser.DELETE(":id", api.Delete)
		gUser.PATCH(":id/enable", api.Enable)
		gUser.PATCH(":id/disable", api.Disable)
	}

}
