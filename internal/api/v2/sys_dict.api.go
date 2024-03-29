package api_v2

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/heromicro/omgind/internal/app/ginx"
	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/app/service"
	"github.com/heromicro/omgind/internal/scheme/repo"
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
		ginx.ResErrorCode(c, -1000, err)
		return
	}

	params.Pagination = true
	result, err := a.DictSrv.Query(ctx, params)
	if err != nil {
		ginx.ResErrorCode(c, -1010, err)
		return
	}

	ginx.ResPage(c, result.Data, result.PageResult)
}

// QuerySelect 查询选择数据
func (a *Dict) QuerySelect(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.DictQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResErrorCode(c, -1000, err)
		return
	}

	// log.Printf(" ------- ------ %+v \n", params)
	params.Sort_Order = repo.OrderByASC.String()
	result, err := a.DictSrv.Query(ctx, params)
	if err != nil {
		ginx.ResErrorCode(c, -1016, err)
		return
	}
	ginx.ResList(c, result.Data)
}

// Query 查询数据
func (a *Dict) QueryItems(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.DictQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResErrorCode(c, -1000, err)
		return
	}

	result, err := a.DictSrv.QueryItems(ctx, c.Param("id"), params)
	if err != nil {
		ginx.ResErrorCode(c, -1018, err)
		return
	}

	ginx.ResSuccess(c, result.Data)
}

// Get 查询指定数据
func (a *Dict) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.DictSrv.Get(ctx, c.Param("id"))
	if err != nil {
		ginx.ResErrorCode(c, -1020, err)
		return
	}

	ginx.ResSuccess(c, item)
}

// View 查询指定数据
func (a *Dict) View(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.DictSrv.View(ctx, c.Param("id"))
	if err != nil {
		ginx.ResErrorCode(c, -1030, err)
		return
	}

	ginx.ResSuccess(c, item)
}

// Create 创建数据
func (a *Dict) Create(c *gin.Context) {

	ctx := c.Request.Context()
	var item schema.Dict

	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResErrorCode(c, -1000, err)
		return
	}

	result, err := a.DictSrv.Create(ctx, item)
	if err != nil {
		ginx.ResErrorCode(c, -1040, err)
		return
	}
	ginx.ResSuccess(c, result)
}

// Update 更新数据
func (a *Dict) Update(c *gin.Context) {

	ctx := c.Request.Context()
	var item schema.Dict
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResErrorCode(c, -1000, err)
		return
	}

	result, err := a.DictSrv.Update(ctx, c.Param("id"), item)
	if err != nil {
		ginx.ResErrorCode(c, -1050, err)
		return
	}

	ginx.ResSuccess(c, result)
}

// Delete 删除数据
func (a *Dict) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.DictSrv.Delete(ctx, c.Param("id"))
	if err != nil {
		ginx.ResErrorCode(c, -1060, err)
		return
	}
	ginx.ResOK(c, "成功删除数据")
}

// Enable 启用数据
func (a *Dict) Enable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.DictSrv.UpdateActive(ctx, c.Param("id"), true)
	if err != nil {
		ginx.ResErrorCode(c, -1070, err)
		return
	}
	ginx.ResOK(c, "启用成功")
}

// Disable 禁用数据
func (a *Dict) Disable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.DictSrv.UpdateActive(ctx, c.Param("id"), false)
	if err != nil {
		ginx.ResErrorCode(c, -1080, err)
		return
	}

	ginx.ResOK(c, "启用成功")
}
