package router

import (
	"github.com/gin-gonic/gin"
	api_v2 "github.com/heromicro/omgind/internal/api/v2"
)

func (r *Router) initMenuRouterV2(urg *gin.RouterGroup, api *api_v2.Menu, pathcomp string) {
	gMenu := urg.Group(pathcomp)
	{
		gMenu.GET("", api.Query)
		gMenu.GET(":id", api.Get)
		gMenu.POST("", api.Create)
		gMenu.PUT(":id", api.Update)
		gMenu.DELETE(":id", api.Delete)
		gMenu.PATCH(":id/enable", api.Enable)
		gMenu.PATCH(":id/disable", api.Disable)
	}

}
