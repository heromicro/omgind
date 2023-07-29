package api_v2

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/heromicro/omgind/internal/app/ginx"
	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/app/service"
)

// SysAnnexSet 注入SysAnnex
var SysAnnexSet = wire.NewSet(wire.Struct(new(SysAnnex), "*"))

// SysAnnex
type SysAnnex struct {
	SysAnnexSrv *service.SysAnnex
}

// Query 查询数据
func (a *SysAnnex) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.SysAnnexQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}

	params.Pagination = true
	result, err := a.SysAnnexSrv.Query(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	ginx.ResPage(c, result.Data, result.PageResult)
}

// Get 查询指定数据
func (a *SysAnnex) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.SysAnnexSrv.Get(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

// View 查询指定数据
func (a *SysAnnex) View(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.SysAnnexSrv.View(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

// Create 创建数据
func (a *SysAnnex) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.SysAnnex
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	// item.Creator = ginx.GetUserID(c)
	result, err := a.SysAnnexSrv.Create(ctx, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

// Update 更新数据
func (a *SysAnnex) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.SysAnnex
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.SysAnnexSrv.Update(ctx, c.Param("id"), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

// Delete 删除数据
func (a *SysAnnex) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.SysAnnexSrv.Delete(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c, "成功删除数据")
}

// Enable 启用数据
func (a *SysAnnex) Enable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.SysAnnexSrv.UpdateActive(ctx, c.Param("id"), true)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c, "启用成功")
}

// Disable 禁用数据
func (a *SysAnnex) Disable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.SysAnnexSrv.UpdateActive(ctx, c.Param("id"), false)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c, "启用成功")
}
