package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// DictSet 注入Dict
var DictSet = wire.NewSet(wire.Struct(new(Dict), "*"))

// Dict 字典
type Dict struct {
}

// Query 查询数据
//
//	@Tags		字典
//	@Summary	查询数据
//	@Param		Authorization	header		string					false	"Bearer 用户令牌"
//	@Param		object			query		schema.DictQueryParam	false	"查询参数"	default{}
//	@Success	200				{object}	schema.ListResult{list=schema.Dicts,pagination=schema.PaginationResult}"
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-dicts [get]
func (a *Dict) Query(c *gin.Context) {
}

// Get 查询指定数据
//
//	@Tags		字典
//	@Summary	查询指定数据
//	@Param		Authorization	header		string	false	"Bearer 用户令牌"
//	@Param		id				path		string	true	"唯一标识"
//	@Success	200				{object}	schema.Dict
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	404				{object}	schema.ErrorResult	"{error:{code:0,message:资源不存在}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-dicts/{id} [get]
func (a *Dict) Get(c *gin.Context) {
}

// Create 创建数据
//
//	@Tags		字典
//	@Summary	创建数据
//	@Param		Authorization	header		string		false	"Bearer 用户令牌"
//	@Param		body			body		schema.Dict	true	"创建数据"
//	@Success	200				{object}	schema.IDResult
//	@Failure	400				{object}	schema.ErrorResult	"{error:{code:0,message:无效的请求参数}}"
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-dicts [post]
func (a *Dict) Create(c *gin.Context) {
}

// Update 更新数据
//
//	@Tags		字典
//	@Summary	更新数据
//	@Param		Authorization	header		string				false	"Bearer 用户令牌"
//	@Param		id				path		string				true	"唯一标识"
//	@Param		body			body		schema.Dict			true	"更新数据"
//	@Success	200				{object}	schema.StatusResult	"{status:OK}"
//	@Failure	400				{object}	schema.ErrorResult	"{error:{code:0,message:无效的请求参数}}"
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-dicts/{id} [put]
func (a *Dict) Update(c *gin.Context) {
}

// Delete 删除数据
//
//	@Tags		字典
//	@Summary	删除数据
//	@Param		Authorization	header		string				false	"Bearer 用户令牌"
//	@Param		id				path		string				true	"唯一标识"
//	@Success	200				{object}	schema.StatusResult	"{status:OK}"
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-dicts/{id} [delete]
func (a *Dict) Delete(c *gin.Context) {
}

// Enable 启用数据
//
//	@Tags		字典
//	@Summary	启用数据
//	@Security	ApiKeyAuth
//	@Param		id	path		string				true	"唯一标识"
//	@Success	200	{object}	schema.StatusResult	"{status:OK}"
//	@Failure	401	{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500	{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-dicts/{id}/enable [patch]
func (a *Dict) Enable(c *gin.Context) {
}

// Disable 禁用数据
//
//	@Tags		字典
//	@Summary	禁用数据
//	@Security	ApiKeyAuth
//	@Param		id	path		string				true	"唯一标识"
//	@Success	200	{object}	schema.StatusResult	"{status:OK}"
//	@Failure	401	{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500	{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-dicts/{id}/disable [patch]
func (a *Dict) Disable(c *gin.Context) {
}
