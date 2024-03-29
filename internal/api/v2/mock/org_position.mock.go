package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// OrgPositionSet 注入OrgPosition
var OrgPositionSet = wire.NewSet(wire.Struct(new(OrgPosition), "*"))

// OrgPosition 职位管理
type OrgPosition struct {
}

// Query 查询数据
//
//	@Tags		职位管理
//	@Summary	查询数据
//	@Security	ApiKeyAuth
//	@Param		Authorization	header		string							false	"Bearer 用户令牌"
//	@Param		object			query		schema.OrgPositionQueryParam	false	"查询参数"	default{}
//	@Success	200				{array}		schema.ListResult{list=schema.OrgPositions,pagination=schema.PaginationResult}
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/org-positions [get]
func (a *OrgPosition) Query(c *gin.Context) {
}

// Get 查询指定数据
//
//	@Tags		职位管理
//	@Summary	查询指定数据
//	@Security	ApiKeyAuth
//	@Param		Authorization	header		string	false	"Bearer 用户令牌"
//	@Param		id				path		string	true	"唯一标识"
//	@Success	200				{object}	schema.OrgPosition
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	404				{object}	schema.ErrorResult	"{error:{code:0,message:资源不存在}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/org-positions/{id} [get]
func (a *OrgPosition) Get(c *gin.Context) {
}

// View 查询指定数据
//
//	@Tags		职位管理
//	@Summary	查询指定数据
//	@Security	ApiKeyAuth
//	@Param		Authorization	header		string	false	"Bearer 用户令牌"
//	@Param		id				path		string	true	"唯一标识"
//	@Success	200				{object}	schema.OrgPosition
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	404				{object}	schema.ErrorResult	"{error:{code:0,message:资源不存在}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/org-positions/{id}/view [get]
func (a *OrgPosition) View(c *gin.Context) {
}

// Create 创建数据
//
//	@Tags		职位管理
//	@Summary	创建数据
//	@Security	ApiKeyAuth
//	@Param		Authorization	header		string				false	"Bearer 用户令牌"
//	@Param		body			body		schema.OrgPosition	true	"创建数据"
//	@Success	200				{object}	schema.IDResult
//	@Failure	400				{object}	schema.ErrorResult	"{error:{code:0,message:无效的请求参数}}"
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/org-positions [post]
func (a *OrgPosition) Create(c *gin.Context) {
}

// Update 更新数据
//
//	@Tags		职位管理
//	@Summary	更新数据
//	@Security	ApiKeyAuth
//	@Param		Authorization	header		string				false	"Bearer 用户令牌"
//	@Param		id				path		string				true	"唯一标识"
//	@Param		body			body		schema.OrgPosition	true	"更新数据"
//	@Success	200				{object}	schema.StatusResult	"{status:OK}"
//	@Failure	400				{object}	schema.ErrorResult	"{error:{code:0,message:无效的请求参数}}"
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/org-positions/{id} [put]
func (a *OrgPosition) Update(c *gin.Context) {
}

// Delete 删除数据
//
//	@Tags		职位管理
//	@Summary	删除数据
//	@Security	ApiKeyAuth
//	@Param		Authorization	header		string				false	"Bearer 用户令牌"
//	@Param		id				path		string				true	"唯一标识"
//	@Success	200				{object}	schema.StatusResult	"{status:OK}"
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/org-positions/{id} [delete]
func (a *OrgPosition) Delete(c *gin.Context) {
}

// Enable 启用数据
//
//	@Tags		职位管理
//	@Summary	启用数据
//	@Security	ApiKeyAuth
//	@Param		id	path		string				true	"唯一标识"
//	@Success	200	{object}	schema.StatusResult	"{status:OK}"
//	@Failure	401	{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500	{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/org-positions/{id}/enable [patch]
func (a *OrgPosition) Enable(c *gin.Context) {
}

// Disable 禁用数据
//
//	@Tags		职位管理
//	@Summary	禁用数据
//	@Security	ApiKeyAuth
//	@Param		id	path		string				true	"唯一标识"
//	@Success	200	{object}	schema.StatusResult	"{status:OK}"
//	@Failure	401	{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500	{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/org-positions/{id}/disable [patch]
func (a *OrgPosition) Disable(c *gin.Context) {
}
