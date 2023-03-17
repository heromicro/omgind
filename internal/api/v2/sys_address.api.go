package api_v2

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/heromicro/omgind/internal/app/ginx"
	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/app/service"
)

// SysAddressSet 注入SysAddress
var SysAddressSet = wire.NewSet(wire.Struct(new(SysAddress), "*"))

// SysAddress 地址管理
type SysAddress struct {
	SysAddressSrv *service.SysAddress
}

// Query 查询数据
func (a *SysAddress) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.SysAddressQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}

	params.Pagination = true
	result, err := a.SysAddressSrv.Query(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	ginx.ResPage(c, result.Data, result.PageResult)
}

// Get 查询指定数据
func (a *SysAddress) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.SysAddressSrv.Get(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

// View 查询指定数据
func (a *SysAddress) View(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.SysAddressSrv.View(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

// Create 创建数据
func (a *SysAddress) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.SysAddress
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	item.Creator = ginx.GetUserID(c)
	result, err := a.SysAddressSrv.Create(ctx, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result.ID)
}

// Update 更新数据
func (a *SysAddress) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.SysAddress
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.SysAddressSrv.Update(ctx, c.Param("id"), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

// Delete 删除数据
func (a *SysAddress) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.SysAddressSrv.Delete(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c, "成功删除数据")
}

// Enable 启用数据
func (a *SysAddress) Enable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.SysAddressSrv.UpdateActive(ctx, c.Param("id"), true)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c, "启用成功")
}

// Disable 禁用数据
func (a *SysAddress) Disable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.SysAddressSrv.UpdateActive(ctx, c.Param("id"), false)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c, "禁用成功")
}
