package router

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	api_v2 "github.com/heromicro/omgind/internal/api/v2"
	"github.com/heromicro/omgind/pkg/auth"
)

var _ IRouter = (*Router)(nil)

// RouterSet 注入router
var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"), wire.Bind(new(IRouter), new(*Router)))

// IRouter 注册路由
type IRouter interface {
	Register(app *gin.Engine) error
	Prefixes() []string
}

// Router 路由管理器
type Router struct {
	Auth           auth.Auther
	CasbinEnforcer *casbin.SyncedEnforcer

	// SockIO *socketio.Server

	DictApiV2 *api_v2.Dict
	DemoAPIV2 *api_v2.Demo
	MenuAPIV2 *api_v2.Menu
	RoleAPIV2 *api_v2.Role
	UserAPIV2 *api_v2.User

	SignInAPIV2      *api_v2.SignIn
	SysDistrictAPIV2 *api_v2.SysDistrict
	SysAddressAPIV2  *api_v2.SysAddress
	OrgOrganAPIV2    *api_v2.OrgOrgan
	OrgStaffAPIV2    *api_v2.OrgStaff
}

// Register 注册路由
func (a *Router) Register(app *gin.Engine) error {
	a.RegisterAPI(app)

	return nil
}

// Prefixes 路由前缀列表
func (a *Router) Prefixes() []string {
	return []string{
		"/api/",
	}
}
