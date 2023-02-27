package api_v2

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/heromicro/omgind/internal/app/ginx"
	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/app/service"
)

// SysDistrictSet 注入SysDistrict
var SysDistrictSet = wire.NewSet(wire.Struct(new(SysDistrict), "*"))

// SysDistrict 行政区域
type SysDistrict struct {
	SysDistrictSrv *service.SysDistrict
}

// Query 查询数据
func (a *SysDistrict) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.SysDistrictQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}

	params.Pagination = true
	result, err := a.SysDistrictSrv.Query(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	ginx.ResPage(c, result.Data, result.PageResult)
}

// Get 查询指定数据
func (a *SysDistrict) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.SysDistrictSrv.Get(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

// Get 查询指定数据
func (a *SysDistrict) GetSubDistrict(c *gin.Context) {
	ctx := c.Request.Context()

	var params schema.SysDistrictQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}

	pid := c.Param("pid")
	params.ParentID = pid
	result, err := a.SysDistrictSrv.Query(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResPage(c, result.Data, result.PageResult)
}

// Create 创建数据
func (a *SysDistrict) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.SysDistrict
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	item.Creator = ginx.GetUserID(c)
	result, err := a.SysDistrictSrv.Create(ctx, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result.ID)
}

// Update 更新数据
func (a *SysDistrict) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.SysDistrict
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	_, err := a.SysDistrictSrv.Update(ctx, c.Param("id"), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

// Delete 删除数据
func (a *SysDistrict) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.SysDistrictSrv.Delete(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

// Enable 启用数据
func (a *SysDistrict) Enable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.SysDistrictSrv.UpdateActive(ctx, c.Param("id"), true)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

// Disable 禁用数据
func (a *SysDistrict) Disable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.SysDistrictSrv.UpdateActive(ctx, c.Param("id"), false)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}
