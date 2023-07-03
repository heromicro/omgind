package router

import (
	"github.com/gin-gonic/gin"
	api_v2 "github.com/heromicro/omgind/internal/api/v2"
)

func (r *Router) initSysTeamRouterV2(urg *gin.RouterGroup, api *api_v2.SysTeam, pathcomp string) {
	gSysTeam := urg.Group(pathcomp)
	{
		gSysTeam.GET("", api.Query)
		gSysTeam.GET(":id", api.Get)
		gSysTeam.POST("", api.Create)
		gSysTeam.PUT(":id", api.Update)
		gSysTeam.DELETE(":id", api.Delete)
		gSysTeam.PATCH(":id/enable", api.Enable)
		gSysTeam.PATCH(":id/disable", api.Disable)
		gSysTeam.GET(":id/view", api.View)
	}
}
