package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// OrgOrganSet 注入OrgOrgan
var OrgOrganSet = wire.NewSet(wire.Struct(new(OrgOrgan), "*"))

// OrgOrgan 组织管理
type OrgOrgan struct {
}

// Query 查询数据
//	@Tags		组织管理
//	@Summary	查询数据
//	@Security	ApiKeyAuth
//	@Param		Authorization	header		string						false	"Bearer 用户令牌"
//	@Param		""				query		schema.OrgOrganQueryParam	false	"查询参数"	default{}
//	@Success	200				{array}		schema.ListResult{list=schema.OrgOrgans,pagination=schema.PaginationResult}
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/org-organs [get]
func (a *OrgOrgan) Query(c *gin.Context) {
}

// Get 查询指定数据
//	@Tags		组织管理
//	@Summary	查询指定数据
//	@Security	ApiKeyAuth
//	@Param		Authorization	header		string	false	"Bearer 用户令牌"
//	@Param		id				path		string	true	"唯一标识"
//	@Success	200				{object}	schema.OrgOrgan
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	404				{object}	schema.ErrorResult	"{error:{code:0,message:资源不存在}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/org-organs/{id} [get]
func (a *OrgOrgan) Get(c *gin.Context) {
}

// View 查询指定数据
//	@Tags		组织管理
//	@Summary	查询指定数据
//	@Security	ApiKeyAuth
//	@Param		Authorization	header		string	false	"Bearer 用户令牌"
//	@Param		id				path		string	true	"唯一标识"
//	@Success	200				{object}	schema.OrgOrgan
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	404				{object}	schema.ErrorResult	"{error:{code:0,message:资源不存在}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/org-organs/{id}/view [get]
func (a *OrgOrgan) View(c *gin.Context) {
}

// Create 创建数据
//	@Tags		组织管理
//	@Summary	创建数据
//	@Security	ApiKeyAuth
//	@Param		Authorization	header		string			false	"Bearer 用户令牌"
//	@Param		body			body		schema.OrgOrgan	true	"创建数据"
//	@Success	200				{object}	schema.IDResult
//	@Failure	400				{object}	schema.ErrorResult	"{error:{code:0,message:无效的请求参数}}"
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/org-organs [post]
func (a *OrgOrgan) Create(c *gin.Context) {
}

// Update 更新数据
//	@Tags		组织管理
//	@Summary	更新数据
//	@Security	ApiKeyAuth
//	@Param		Authorization	header		string				false	"Bearer 用户令牌"
//	@Param		id				path		string				true	"唯一标识"
//	@Param		body			body		schema.OrgOrgan		true	"更新数据"
//	@Success	200				{object}	schema.StatusResult	"{status:OK}"
//	@Failure	400				{object}	schema.ErrorResult	"{error:{code:0,message:无效的请求参数}}"
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/org-organs/{id} [put]
func (a *OrgOrgan) Update(c *gin.Context) {
}

// Delete 删除数据
//	@Tags		组织管理
//	@Summary	删除数据
//	@Security	ApiKeyAuth
//	@Param		Authorization	header		string				false	"Bearer 用户令牌"
//	@Param		id				path		string				true	"唯一标识"
//	@Success	200				{object}	schema.StatusResult	"{status:OK}"
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/org-organs/{id} [delete]
func (a *OrgOrgan) Delete(c *gin.Context) {
}

// Enable 启用数据
//	@Tags		组织管理
//	@Summary	启用数据
//	@Security	ApiKeyAuth
//	@Param		id	path		string				true	"唯一标识"
//	@Success	200	{object}	schema.StatusResult	"{status:OK}"
//	@Failure	401	{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500	{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/org-organs/{id}/enable [patch]
func (a *OrgOrgan) Enable(c *gin.Context) {
}

// Disable 禁用数据
//	@Tags		组织管理
//	@Summary	禁用数据
//	@Security	ApiKeyAuth
//	@Param		id	path		string				true	"唯一标识"
//	@Success	200	{object}	schema.StatusResult	"{status:OK}"
//	@Failure	401	{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500	{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/org-organs/{id}/disable [patch]
func (a *OrgOrgan) Disable(c *gin.Context) {
}
