package router

import (
	"github.com/gin-gonic/gin"
	api_v2 "github.com/heromicro/omgind/internal/api/v2"
)

func (r *Router) initOrgOrganRouterV2(urg *gin.RouterGroup, api *api_v2.OrgOrgan, pathcomp string) {
	gOrgOrgan := urg.Group(pathcomp)
	{
		gOrgOrgan.GET("", api.Query)
		gOrgOrgan.GET(":id", api.Get)
		gOrgOrgan.POST("", api.Create)
		gOrgOrgan.PUT(":id", api.Update)
		gOrgOrgan.DELETE(":id", api.Delete)
		gOrgOrgan.PATCH(":id/enable", api.Enable)
		gOrgOrgan.PATCH(":id/disable", api.Disable)
		gOrgOrgan.PATCH(":id/view", api.View)
	}
}
