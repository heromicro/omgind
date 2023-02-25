package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// SysDistrictSet 注入SysDistrict
var SysDistrictSet = wire.NewSet(wire.Struct(new(SysDistrict), "*"))

// SysDistrict 行政区域
type SysDistrict struct {
}

// Query 查询数据
//	@Tags		行政区域
//	@Summary	查询数据
//	@Security	ApiKeyAuth
//	@Param		Authorization	header		string	false	"Bearer 用户令牌"
//	@Param		current			query		int		true	"分页索引"	default(1)
//	@Param		pageSize		query		int		true	"分页大小"	default(10)
//	@Param		queryValue		query		string	false	"查询值"
//	@Success	200				{array}		schema.ListResult{list=schema.SysDistricts,pagination=schema.PaginationResult}
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-districts [get]
func (a *SysDistrict) Query(c *gin.Context) {
}

// Get 查询指定数据
//	@Tags		行政区域
//	@Summary	查询指定数据
//	@Security	ApiKeyAuth
//	@Param		Authorization	header		string	false	"Bearer 用户令牌"
//	@Param		id				path		string	true	"唯一标识"
//	@Success	200				{object}	schema.SysDistrict
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	404				{object}	schema.ErrorResult	"{error:{code:0,message:资源不存在}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-districts/{id} [get]
func (a *SysDistrict) Get(c *gin.Context) {
}

// Create 创建数据
//	@Tags		行政区域
//	@Summary	创建数据
//	@Security	ApiKeyAuth
//	@Param		Authorization	header		string				false	"Bearer 用户令牌"
//	@Param		body			body		schema.SysDistrict	true	"创建数据"
//	@Success	200				{object}	schema.IDResult
//	@Failure	400				{object}	schema.ErrorResult	"{error:{code:0,message:无效的请求参数}}"
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-districts [post]
func (a *SysDistrict) Create(c *gin.Context) {
}

// Update 更新数据
//	@Tags		行政区域
//	@Summary	更新数据
//	@Security	ApiKeyAuth
//	@Param		Authorization	header		string				false	"Bearer 用户令牌"
//	@Param		id				path		string				true	"唯一标识"
//	@Param		body			body		schema.SysDistrict	true	"更新数据"
//	@Success	200				{object}	schema.StatusResult	"{status:OK}"
//	@Failure	400				{object}	schema.ErrorResult	"{error:{code:0,message:无效的请求参数}}"
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-districts/{id} [put]
func (a *SysDistrict) Update(c *gin.Context) {
}

// Delete 删除数据
//	@Tags		行政区域
//	@Summary	删除数据
//	@Security	ApiKeyAuth
//	@Param		Authorization	header		string				false	"Bearer 用户令牌"
//	@Param		id				path		string				true	"唯一标识"
//	@Success	200				{object}	schema.StatusResult	"{status:OK}"
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-districts/{id} [delete]
func (a *SysDistrict) Delete(c *gin.Context) {
}

// Enable 启用数据
//	@Tags		行政区域
//	@Summary	启用数据
//	@Security	ApiKeyAuth
//	@Param		id	path		string				true	"唯一标识"
//	@Success	200	{object}	schema.StatusResult	"{status:OK}"
//	@Failure	401	{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500	{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-districts/{id}/enable [patch]
func (a *SysDistrict) Enable(c *gin.Context) {
}

// Disable 禁用数据
//	@Tags		行政区域
//	@Summary	禁用数据
//	@Security	ApiKeyAuth
//	@Param		id	path		string				true	"唯一标识"
//	@Success	200	{object}	schema.StatusResult	"{status:OK}"
//	@Failure	401	{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500	{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-districts/{id}/disable [patch]
func (a *SysDistrict) Disable(c *gin.Context) {
}
