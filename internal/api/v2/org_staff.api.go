package api_v2

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/heromicro/omgind/internal/app/ginx"
	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/app/service"
)

// OrgStaffSet 注入OrgStaff
var OrgStaffSet = wire.NewSet(wire.Struct(new(OrgStaff), "*"))

// OrgStaff 员工管理
type OrgStaff struct {
	OrgStaffSrv *service.OrgStaff
}

// Query 查询数据
func (a *OrgStaff) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.OrgStaffQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}

	params.Pagination = true
	result, err := a.OrgStaffSrv.Query(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	ginx.ResPage(c, result.Data, result.PageResult)
}

// Get 查询指定数据
func (a *OrgStaff) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.OrgStaffSrv.Get(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

// View 查询指定数据
func (a *OrgStaff) View(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.OrgStaffSrv.View(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

// Create 创建数据
func (a *OrgStaff) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.OrgStaff
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	item.Creator = ginx.GetUserID(c)
	result, err := a.OrgStaffSrv.Create(ctx, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

// Update 更新数据
func (a *OrgStaff) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.OrgStaff
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.OrgStaffSrv.Update(ctx, c.Param("id"), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

// Delete 删除数据
func (a *OrgStaff) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.OrgStaffSrv.Delete(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c, "成功删除数据")
}

// Enable 启用数据
func (a *OrgStaff) Enable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.OrgStaffSrv.UpdateActive(ctx, c.Param("id"), true)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c, "启用成功")
}

// Disable 禁用数据
func (a *OrgStaff) Disable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.OrgStaffSrv.UpdateActive(ctx, c.Param("id"), false)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c, "启用成功")
}
