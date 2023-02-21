package router

import (
	"github.com/gin-gonic/gin"
	api_v2 "github.com/heromicro/omgind/internal/api/v2"
)

func (r *Router) initRoleRouterV2(urg *gin.RouterGroup, api *api_v2.Role, pathcomp string) {
	gRole := urg.Group(pathcomp)
	{
		gRole.GET("", api.Query)
		gRole.GET(":id", api.Get)
		gRole.POST("", api.Create)
		gRole.PUT(":id", api.Update)
		gRole.DELETE(":id", api.Delete)
		gRole.PATCH(":id/enable", api.Enable)
		gRole.PATCH(":id/disable", api.Disable)
	}

}
