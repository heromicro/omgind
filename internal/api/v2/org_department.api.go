package api_v2

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/heromicro/omgind/internal/app/ginx"
	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/app/service"
)

// OrgDepartmentSet 注入OrgDepartment
var OrgDepartmentSet = wire.NewSet(wire.Struct(new(OrgDepartment), "*"))

// OrgDepartment 部门管理
type OrgDepartment struct {
	OrgDepartmentSrv *service.OrgDepartment
}

// Query 查询数据
func (a *OrgDepartment) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.OrgDepartmentQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}

	params.Pagination = true
	result, err := a.OrgDepartmentSrv.Query(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	ginx.ResPage(c, result.Data, result.PageResult)
}

// Get 查询指定数据
func (a *OrgDepartment) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.OrgDepartmentSrv.Get(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

// View 查询指定数据
func (a *OrgDepartment) View(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.OrgDepartmentSrv.View(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

// Get 所有的行政区
func (a *OrgDepartment) GetAllSubs(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.OrgDepartmentQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}
	var pid string = c.Param("id")
	params.Current = 1
	params.PageSize = 1000

	log.Println(" --- ---- === ==== pid ", pid)
	log.Println(" ----- -- ==== === URL ", c.Request.URL)

	result, err := a.OrgDepartmentSrv.GetAllSubs(ctx, pid, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResPage(c, result.Data, result.PageResult)
}

// QueryTree 查询菜单树
func (a *OrgDepartment) QueryTree(c *gin.Context) {

	ctx := c.Request.Context()
	var params schema.OrgDepartmentQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}
	var pid string = c.Param("id")

	params.PageSize = 200

	result, err := a.OrgDepartmentSrv.GetTree(ctx, pid, params)

	if err != nil {
		ginx.ResError(c, err)
		return
	}

	ginx.ResSuccess(c, result)
}

// Create 创建数据
func (a *OrgDepartment) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.OrgDepartment
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	item.Creator = ginx.GetUserID(c)
	result, err := a.OrgDepartmentSrv.Create(ctx, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

// Update 更新数据
func (a *OrgDepartment) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.OrgDepartment
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.OrgDepartmentSrv.Update(ctx, c.Param("id"), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

// Delete 删除数据
func (a *OrgDepartment) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.OrgDepartmentSrv.Delete(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c, "成功删除数据")
}

// Enable 启用数据
func (a *OrgDepartment) Enable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.OrgDepartmentSrv.UpdateActive(ctx, c.Param("id"), true)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c, "启用成功")
}

// Disable 禁用数据
func (a *OrgDepartment) Disable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.OrgDepartmentSrv.UpdateActive(ctx, c.Param("id"), false)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c, "启用成功")
}
