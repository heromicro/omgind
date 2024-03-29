package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// SysAddressSet 注入SysAddress
var SysAddressSet = wire.NewSet(wire.Struct(new(SysAddress), "*"))

// SysAddress 地址管理
type SysAddress struct {
}

// Query 查询数据
//
//	@Tags		地址管理
//	@Summary	查询数据
//	@Security	ApiKeyAuth
//	@Param		Authorization	header		string						false	"Bearer 用户令牌"
//	@Param		object			query		schema.SysAddressQueryParam	false	"查询参数"	default{}
//	@Success	200				{array}		schema.ListResult{list=schema.SysAddresses,pagination=schema.PaginationResult}
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-addresses [get]
func (a *SysAddress) Query(c *gin.Context) {
}

// Get 查询指定数据
//
//	@Tags		地址管理
//	@Summary	查询指定数据
//	@Security	ApiKeyAuth
//	@Param		Authorization	header		string	false	"Bearer 用户令牌"
//	@Param		id				path		string	true	"唯一标识"
//	@Success	200				{object}	schema.SysAddress
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	404				{object}	schema.ErrorResult	"{error:{code:0,message:资源不存在}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-addresses/{id} [get]
func (a *SysAddress) Get(c *gin.Context) {
}

// Create 创建数据
//
//	@Tags		地址管理
//	@Summary	创建数据
//	@Security	ApiKeyAuth
//	@Param		Authorization	header		string				false	"Bearer 用户令牌"
//	@Param		body			body		schema.SysAddress	true	"创建数据"
//	@Success	200				{object}	schema.IDResult
//	@Failure	400				{object}	schema.ErrorResult	"{error:{code:0,message:无效的请求参数}}"
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-addresses [post]
func (a *SysAddress) Create(c *gin.Context) {
}

// Update 更新数据
//
//	@Tags		地址管理
//	@Summary	更新数据
//	@Security	ApiKeyAuth
//	@Param		Authorization	header		string				false	"Bearer 用户令牌"
//	@Param		id				path		string				true	"唯一标识"
//	@Param		body			body		schema.SysAddress	true	"更新数据"
//	@Success	200				{object}	schema.StatusResult	"{status:OK}"
//	@Failure	400				{object}	schema.ErrorResult	"{error:{code:0,message:无效的请求参数}}"
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-addresses/{id} [put]
func (a *SysAddress) Update(c *gin.Context) {
}

// Delete 删除数据
//
//	@Tags		地址管理
//	@Summary	删除数据
//	@Security	ApiKeyAuth
//	@Param		Authorization	header		string				false	"Bearer 用户令牌"
//	@Param		id				path		string				true	"唯一标识"
//	@Success	200				{object}	schema.StatusResult	"{status:OK}"
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-addresses/{id} [delete]
func (a *SysAddress) Delete(c *gin.Context) {
}

// Enable 启用数据
//
//	@Tags		地址管理
//	@Summary	启用数据
//	@Security	ApiKeyAuth
//	@Param		id	path		string				true	"唯一标识"
//	@Success	200	{object}	schema.StatusResult	"{status:OK}"
//	@Failure	401	{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500	{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-addresses/{id}/enable [patch]
func (a *SysAddress) Enable(c *gin.Context) {
}

// Disable 禁用数据
//
//	@Tags		地址管理
//	@Summary	禁用数据
//	@Security	ApiKeyAuth
//	@Param		id	path		string				true	"唯一标识"
//	@Success	200	{object}	schema.StatusResult	"{status:OK}"
//	@Failure	401	{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500	{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-addresses/{id}/disable [patch]
func (a *SysAddress) Disable(c *gin.Context) {
}
