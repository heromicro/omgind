package router

import (
	"github.com/gin-gonic/gin"
	api_v2 "github.com/heromicro/omgind/internal/api/v2"
)

func (r *Router) initSysDistrictRouterV2(urg *gin.RouterGroup, api *api_v2.SysDistrict, pathcomp string) {
	gSysDistrict := urg.Group(pathcomp)
	{
		gSysDistrict.GET("", api.Query)
		gSysDistrict.GET(":id", api.Get)
		gSysDistrict.POST("", api.Create)
		gSysDistrict.PUT(":id", api.Update)
		gSysDistrict.DELETE(":id", api.Delete)
		gSysDistrict.PATCH(":id/enable", api.Enable)
		gSysDistrict.PATCH(":id/disable", api.Disable)
		gSysDistrict.GET(":id/substricts", api.GetAllSubDistricts)
		gSysDistrict.GET(":id/view", api.View)
		gSysDistrict.GET(":id/tree", api.QueryTree)

	}
}
