package router

import (
	"github.com/gin-gonic/gin"
	api_v2 "github.com/heromicro/omgind/internal/api/v2"
)

func (r *Router) initDictRouterV2(urg *gin.RouterGroup, api *api_v2.Dict, pathcomp string) {
	gDemo := urg.Group(pathcomp)
	{
		gDemo.GET("", api.Query)
		gDemo.GET(":id", api.Get)
		gDemo.POST("", api.Create)
		gDemo.PUT(":id", api.Update)
		gDemo.DELETE(":id", api.Delete)
		gDemo.PATCH(":id/enable", api.Enable)
		gDemo.PATCH(":id/disable", api.Disable)
		// gDemo.GET(":id/view", api.View)
		gDemo.GET(":id/items", api.QueryItems)

	}
}
