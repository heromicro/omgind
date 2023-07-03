package api_v2

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/heromicro/omgind/internal/app/ginx"
	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/app/service"
)

// OrgDeptSet 注入OrgDept
var OrgDeptSet = wire.NewSet(wire.Struct(new(OrgDept), "*"))

// OrgDept 部门管理
type OrgDept struct {
	OrgDeptSrv *service.OrgDept
}

// Query 查询数据
func (a *OrgDept) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.OrgDeptQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}

	params.Pagination = true
	result, err := a.OrgDeptSrv.Query(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	ginx.ResPage(c, result.Data, result.PageResult)
}

// Get 查询指定数据
func (a *OrgDept) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.OrgDeptSrv.Get(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

// View 查询指定数据
func (a *OrgDept) View(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.OrgDeptSrv.View(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

// Get 所有的行政区
func (a *OrgDept) GetAllSubs(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.OrgDeptQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}
	var pid string = c.Param("id")
	params.Current = 1
	params.PageSize = 1000

	log.Println(" --- ---- === ==== pid ", pid)
	log.Println(" ----- -- ==== === URL ", c.Request.URL)

	result, err := a.OrgDeptSrv.GetAllSubs(ctx, pid, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResPage(c, result.Data, result.PageResult)
}

// QueryTree 查询菜单树
func (a *OrgDept) QueryTree(c *gin.Context) {

	ctx := c.Request.Context()
	var params schema.OrgDeptQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}
	var pid string = c.Param("id")

	params.PageSize = 200

	result, err := a.OrgDeptSrv.GetTree(ctx, pid, params)

	if err != nil {
		ginx.ResError(c, err)
		return
	}

	ginx.ResSuccess(c, result)
}

// Create 创建数据
func (a *OrgDept) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.OrgDept
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.OrgDeptSrv.Create(ctx, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

// Update 更新数据
func (a *OrgDept) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.OrgDept
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.OrgDeptSrv.Update(ctx, c.Param("id"), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

// Delete 删除数据
func (a *OrgDept) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.OrgDeptSrv.Delete(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c, "成功删除数据")
}

// Enable 启用数据
func (a *OrgDept) Enable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.OrgDeptSrv.UpdateActive(ctx, c.Param("id"), true)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c, "启用成功")
}

// Disable 禁用数据
func (a *OrgDept) Disable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.OrgDeptSrv.UpdateActive(ctx, c.Param("id"), false)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c, "启用成功")
}
