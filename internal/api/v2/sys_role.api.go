package api_v2

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/heromicro/omgind/internal/app/ginx"
	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/app/service"
	"github.com/heromicro/omgind/internal/scheme/repo"
)

// RoleSet 注入Role
var RoleSet = wire.NewSet(wire.Struct(new(Role), "*"))

// Role 角色管理
type Role struct {
	RoleSrv *service.Role
}

// Query 查询数据
func (a *Role) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.RoleQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}

	params.Pagination = true
	result, err := a.RoleSrv.Query(ctx, params, schema.RoleQueryOptions{
		OrderFields: schema.NewOrderFields(schema.NewOrderField("sort", schema.OrderByDESC)),
	})
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResPage(c, result.Data, result.PageResult)
}

// QuerySelectPage 查询选择数据
func (a *Role) QuerySelectPage(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.RoleQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		// log.Printf(" ------- ------ %+v \n", err)

		ginx.ResError(c, err)
		return
	}

	// log.Printf(" ------- ------ %+v \n", params)
	params.Sort_Order = repo.OrderByASC.String()

	result, err := a.RoleSrv.QuerySelectPage(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResList(c, result.Data)
}

// Get 查询指定数据
func (a *Role) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.RoleSrv.Get(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

// Create 创建数据
func (a *Role) Create(c *gin.Context) {

	log.Println(" -------- 0000000 === ==== ")

	ctx := c.Request.Context()
	var item schema.Role
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.RoleSrv.Create(ctx, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	log.Println(" -------- vvvvvv == ===== ", result)

	ginx.ResSuccess(c, result)
}

// Update 更新数据
func (a *Role) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.Role
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.RoleSrv.Update(ctx, c.Param("id"), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

// Delete 删除数据
func (a *Role) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.RoleSrv.Delete(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c, "成功删除数据")
}

// Enable 启用数据
func (a *Role) Enable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.RoleSrv.UpdateActive(ctx, c.Param("id"), true)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c, "启用成功")
}

// Disable 禁用数据
func (a *Role) Disable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.RoleSrv.UpdateActive(ctx, c.Param("id"), false)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c, "禁用成功")
}
