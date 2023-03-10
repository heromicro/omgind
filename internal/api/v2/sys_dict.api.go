package api_v2

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/heromicro/omgind/internal/app/ginx"
	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/app/service"
)

// DictSet 注入Dict
var DictSet = wire.NewSet(wire.Struct(new(Dict), "*"))

// Dict 字典
type Dict struct {
	DictSrv *service.Dict
}

// Query 查询数据
func (a *Dict) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.DictQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}

	params.Pagination = true
	result, err := a.DictSrv.Query(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	ginx.ResPage2(c, result.Data, result.PageResult)
}

// Get 查询指定数据
func (a *Dict) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.DictSrv.Get(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	ginx.ResSuccess2(c, item)
}

// Create 创建数据
func (a *Dict) Create(c *gin.Context) {

	ctx := c.Request.Context()
	var item schema.Dict

	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	item.Creator = ginx.GetUserID(c)
	result, err := a.DictSrv.Create(ctx, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess2(c, result)
}

// Update 更新数据
func (a *Dict) Update(c *gin.Context) {

	ctx := c.Request.Context()
	var item schema.Dict
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.DictSrv.Update(ctx, c.Param("id"), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	ginx.ResSuccess2(c, result)
}

// Delete 删除数据
func (a *Dict) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.DictSrv.Delete(ctx, c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK2(c, "成功删除数据")
}

// Enable 启用数据
func (a *Dict) Enable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.DictSrv.UpdateStatus(ctx, c.Param("id"), true)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK2(c, "启用成功")
}

// Disable 禁用数据
func (a *Dict) Disable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.DictSrv.UpdateStatus(ctx, c.Param("id"), false)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	ginx.ResOK2(c, "启用成功")
}
