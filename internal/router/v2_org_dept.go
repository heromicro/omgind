package router

import (
	"github.com/gin-gonic/gin"
	api_v2 "github.com/heromicro/omgind/internal/api/v2"
)

func (r *Router) initOrgDeptRouterV2(urg *gin.RouterGroup, api *api_v2.OrgDept, pathcomp string) {
	gOrgDept := urg.Group(pathcomp)
	{
		gOrgDept.GET("", api.Query)
		gOrgDept.GET(":id", api.Get)
		gOrgDept.POST("", api.Create)
		gOrgDept.PUT(":id", api.Update)
		gOrgDept.DELETE(":id", api.Delete)
		gOrgDept.PATCH(":id/enable", api.Enable)
		gOrgDept.PATCH(":id/disable", api.Disable)
		gOrgDept.GET(":id/view", api.View)
		gOrgDept.GET(":id/subs", api.GetAllSubs)
		gOrgDept.GET(":id/tree", api.QueryTree)
	}
}
