package router

import (
	"github.com/gin-gonic/gin"
	api_v2 "github.com/heromicro/omgind/internal/api/v2"
)

func (r *Router) initOrgDepartmentRouterV2(urg *gin.RouterGroup, api *api_v2.OrgDepartment, pathcomp string) {
	gOrgDepartment := urg.Group(pathcomp)
	{
		gOrgDepartment.GET("", api.Query)
		gOrgDepartment.GET(":id", api.Get)
		gOrgDepartment.POST("", api.Create)
		gOrgDepartment.PUT(":id", api.Update)
		gOrgDepartment.DELETE(":id", api.Delete)
		gOrgDepartment.PATCH(":id/enable", api.Enable)
		gOrgDepartment.PATCH(":id/disable", api.Disable)
		gOrgDepartment.GET(":id/view", api.View)
	}
}
