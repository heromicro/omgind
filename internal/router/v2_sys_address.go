package router

import (
	"github.com/gin-gonic/gin"
	api_v2 "github.com/heromicro/omgind/internal/api/v2"
)

func (r *Router) initSysAddressRouterV2(urg *gin.RouterGroup, api *api_v2.SysAddress, pathcomp string) {
	gSysAddress := urg.Group(pathcomp)
	{
		gSysAddress.GET("", api.Query)
		gSysAddress.GET(":id", api.Get)
		gSysAddress.POST("", api.Create)
		gSysAddress.PUT(":id", api.Update)
		gSysAddress.DELETE(":id", api.Delete)
		gSysAddress.PATCH(":id/enable", api.Enable)
		gSysAddress.PATCH(":id/disable", api.Disable)
		gSysAddress.GET(":id/view", api.View)
	}
}
