package router

import (
	"github.com/gin-gonic/gin"
	"github.com/heromicro/omgind/internal/app/middleware"
)

// RegisterAPI register api group router
func (a *Router) RegisterAPI(app *gin.Engine) {

	g := app.Group("/api")

	g.Use(middleware.UserAuthMiddleware(a.Auth,
		middleware.AllowPathPrefixSkipper("/api/v2/pub/signin"),
		middleware.AllowPathPrefixSkipper("/ws/"),
	))

	g.Use(middleware.CasbinMiddleware(a.CasbinEnforcer,
		middleware.AllowPathPrefixSkipper("/api/v2/pub"),
		middleware.AllowPathPrefixSkipper("/ws/"),
	))

	g.Use(middleware.RateLimiterMiddleware())

	v2 := g.Group("/v2")
	{
		pub := v2.Group("/pub")
		{
			gSignIn := pub.Group("/signin")
			{
				gSignIn.GET("captchaid", a.SignInAPIV2.GetCaptcha)
				gSignIn.GET("captcha", a.SignInAPIV2.ResCaptcha)
				gSignIn.POST("", a.SignInAPIV2.SignIn)
				gSignIn.POST("exit", a.SignInAPIV2.SignOut)
			}

			gCurrent := pub.Group("current")
			{
				gCurrent.PUT("password", a.SignInAPIV2.UpdatePassword)
				gCurrent.GET("user", a.SignInAPIV2.GetUserInfo)
				gCurrent.GET("menutree", a.SignInAPIV2.QueryUserMenuTree)
			}
			pub.POST("/refresh-token", a.SignInAPIV2.RefreshToken)
		}

		a.initMenuRouterV2(v2, a.MenuAPIV2, "sys-menus")
		v2.GET("/sys-menus.tree", a.MenuAPIV2.QueryTree)

		a.initRoleRouterV2(v2, a.RoleAPIV2, "sys-roles")
		v2.GET("/sys-roles.select.page", a.RoleAPIV2.QuerySelectPage)

		a.initUserRouterV2(v2, a.UserAPIV2, "sys-users")
		v2.GET("/sys-users.select.page", a.UserAPIV2.QuerySelectPage)

		a.initDictRouterV2(v2, a.DictApiV2, "sys-dicts")
		v2.GET("/sys-dicts.select", a.DictApiV2.QuerySelect)

		a.initSysDistrictRouterV2(v2, a.SysDistrictAPIV2, "sys-district")

		a.initSysAddressRouterV2(v2, a.SysAddressAPIV2, "sys-address")

		a.initDemoRouterV2(v2, a.DemoAPIV2, "demos")

		a.initOrgOrganRouterV2(v2, a.OrgOrganAPIV2, "org-organs")
		v2.GET("/sys-organs.select.page", a.OrgOrganAPIV2.QuerySelectPage)

		a.initOrgStaffRouterV2(v2, a.OrgStaffAPIV2, "org-staffs")

		a.initOrgPositionRouterV2(v2, a.OrgPositionAPIV2, "org-positions")

		a.initOrgDeptRouterV2(v2, a.OrgDeptAPIV2, "org-depts")

		a.initSysTeamRouterV2(v2, a.SysTeamAPIV2, "sys-teams")

		a.initSysAnnexRouterV2(v2, a.SysAnnexAPIV2, "sys-annexes")

	}

	// app.GET("/ws/*any", gin.WrapH(a.SockIO))
	// app.POST("/ws/*any", gin.WrapH(a.SockIO))

}
