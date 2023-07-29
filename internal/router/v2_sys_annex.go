package router

import (
	"github.com/gin-gonic/gin"
	api_v2 "github.com/heromicro/omgind/internal/api/v2"
)

func (r *Router) initSysAnnexRouterV2(urg *gin.RouterGroup, api *api_v2.SysAnnex, pathcomp string) {
	gSysAnnex := urg.Group(pathcomp)
	{
		gSysAnnex.GET("", api.Query)
		gSysAnnex.GET(":id", api.Get)
		gSysAnnex.POST("", api.Create)
		gSysAnnex.PUT(":id", api.Update)
		gSysAnnex.DELETE(":id", api.Delete)
		gSysAnnex.PATCH(":id/enable", api.Enable)
		gSysAnnex.PATCH(":id/disable", api.Disable)
		gSysAnnex.GET(":id/view", api.View)
	}
}
