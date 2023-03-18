package router

import (
	"github.com/gin-gonic/gin"
	api_v2 "github.com/heromicro/omgind/internal/api/v2"
)

func (r *Router) initOrgPositionRouterV2(urg *gin.RouterGroup, api *api_v2.OrgPosition, pathcomp string) {
	gOrgPosition := urg.Group(pathcomp)
	{
		gOrgPosition.GET("", api.Query)
		gOrgPosition.GET(":id", api.Get)
		gOrgPosition.POST("", api.Create)
		gOrgPosition.PUT(":id", api.Update)
		gOrgPosition.DELETE(":id", api.Delete)
		gOrgPosition.PATCH(":id/enable", api.Enable)
		gOrgPosition.PATCH(":id/disable", api.Disable)
		gOrgPosition.GET(":id/view", api.View)
	}
}
