package api_v2

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/heromicro/omgind/internal/app/ginx"
	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/app/service"
	"github.com/heromicro/omgind/pkg/errors"
)

// UserSet 注入User
var UserSet = wire.NewSet(wire.Struct(new(User), "*"))

// User 用户管理
type User struct {
	UserSrv *service.User
}

// Query 查询数据
func (a *User) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.UserQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}
	if v := c.Query("roleIDs"); v != "" {
		params.RoleIDs = strings.Split(v, ",")
	}

	params.Pagination = true
	result, err := a.UserSrv.QueryShow(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResPage(c, result.Data, result.PageResult)
}

// Get 查询指定数据
func (a *User) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.UserSrv.Get(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item.CleanSecure())
}

// Create 创建数据
func (a *User) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.User
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	} else if item.Password == "" {
		ginx.ResError(c, errors.New400Response("密码不能为空"))
		return
	}

	item.Creator = ginx.GetUserID(c)
	result, err := a.UserSrv.Create(ctx, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

// Update 更新数据
func (a *User) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.User
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	err := a.UserSrv.Update(ctx, c.Param("id"), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

// Delete 删除数据
func (a *User) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.UserSrv.Delete(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

// Enable 启用数据
func (a *User) Enable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.UserSrv.UpdateStatus(ctx, c.Param("id"), true)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

// Disable 禁用数据
func (a *User) Disable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.UserSrv.UpdateStatus(ctx, c.Param("id"), false)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}
