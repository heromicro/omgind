package router

import (
	"github.com/gin-gonic/gin"
	api_v2 "github.com/heromicro/omgind/internal/api/v2"
)

func (r *Router) initOrgStaffRouterV2(urg *gin.RouterGroup, api *api_v2.OrgStaff, pathcomp string) {
	gOrgStaff := urg.Group(pathcomp)
	{
		gOrgStaff.GET("", api.Query)
		gOrgStaff.GET(":id", api.Get)
		gOrgStaff.POST("", api.Create)
		gOrgStaff.PUT(":id", api.Update)
		gOrgStaff.DELETE(":id", api.Delete)
		gOrgStaff.PATCH(":id/enable", api.Enable)
		gOrgStaff.PATCH(":id/disable", api.Disable)
		gOrgStaff.PATCH(":id/view", api.View)
	}
}
