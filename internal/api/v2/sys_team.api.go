package api_v2

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/heromicro/omgind/internal/app/ginx"
	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/app/service"
)

// SysTeamSet 注入SysTeam
var SysTeamSet = wire.NewSet(wire.Struct(new(SysTeam), "*"))

// SysTeam 地址管理
type SysTeam struct {
	SysTeamSrv *service.SysTeam
}

// Query 查询数据
func (a *SysTeam) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.SysTeamQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}

	params.Pagination = true
	result, err := a.SysTeamSrv.Query(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	ginx.ResPage(c, result.Data, result.PageResult)
}

// Get 查询指定数据
func (a *SysTeam) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.SysTeamSrv.Get(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

// View 查询指定数据
func (a *SysTeam) View(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.SysTeamSrv.View(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

// Create 创建数据
func (a *SysTeam) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.SysTeam
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.SysTeamSrv.Create(ctx, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

// Update 更新数据
func (a *SysTeam) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.SysTeam
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.SysTeamSrv.Update(ctx, c.Param("id"), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

// Delete 删除数据
func (a *SysTeam) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.SysTeamSrv.Delete(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c, "成功删除数据")
}

// Enable 启用数据
func (a *SysTeam) Enable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.SysTeamSrv.UpdateActive(ctx, c.Param("id"), true)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c, "启用成功")
}

// Disable 禁用数据
func (a *SysTeam) Disable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.SysTeamSrv.UpdateActive(ctx, c.Param("id"), false)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c, "启用成功")
}
