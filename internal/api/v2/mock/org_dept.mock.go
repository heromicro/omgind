package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// OrgDeptSet 注入OrgDept
var OrgDeptSet = wire.NewSet(wire.Struct(new(OrgDept), "*"))

// OrgDept 部门管理
type OrgDept struct {
}

// Query 查询数据
//
//	@Tags		部门管理
//	@Summary	查询数据
//	@Security	ApiKeyAuth
//	@Param		Authorization	header		string							false	"Bearer 用户令牌"
//	@Param		object			query		schema.OrgDeptQueryParam	false	"查询参数"	default{}
//	@Success	200				{array}		schema.ListResult{list=schema.OrgDepts,pagination=schema.PaginationResult}
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/org-departments [get]
func (a *OrgDept) Query(c *gin.Context) {
}

// Get 查询指定数据
//
//	@Tags		部门管理
//	@Summary	查询指定数据
//	@Security	ApiKeyAuth
//	@Param		Authorization	header		string	false	"Bearer 用户令牌"
//	@Param		id				path		string	true	"唯一标识"
//	@Success	200				{object}	schema.OrgDept
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	404				{object}	schema.ErrorResult	"{error:{code:0,message:资源不存在}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/org-departments/{id} [get]
func (a *OrgDept) Get(c *gin.Context) {
}

// View 查询指定数据
//
//	@Tags		部门管理
//	@Summary	查询指定数据
//	@Security	ApiKeyAuth
//	@Param		Authorization	header		string	false	"Bearer 用户令牌"
//	@Param		id				path		string	true	"唯一标识"
//	@Success	200				{object}	schema.OrgDept
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	404				{object}	schema.ErrorResult	"{error:{code:0,message:资源不存在}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/org-departments/{id}/view [get]
func (a *OrgDept) View(c *gin.Context) {
}

// Create 创建数据
//
//	@Tags		部门管理
//	@Summary	创建数据
//	@Security	ApiKeyAuth
//	@Param		Authorization	header		string					false	"Bearer 用户令牌"
//	@Param		body			body		schema.OrgDept	true	"创建数据"
//	@Success	200				{object}	schema.IDResult
//	@Failure	400				{object}	schema.ErrorResult	"{error:{code:0,message:无效的请求参数}}"
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/org-departments [post]
func (a *OrgDept) Create(c *gin.Context) {
}

// Update 更新数据
//
//	@Tags		部门管理
//	@Summary	更新数据
//	@Security	ApiKeyAuth
//	@Param		Authorization	header		string					false	"Bearer 用户令牌"
//	@Param		id				path		string					true	"唯一标识"
//	@Param		body			body		schema.OrgDept	true	"更新数据"
//	@Success	200				{object}	schema.StatusResult		"{status:OK}"
//	@Failure	400				{object}	schema.ErrorResult		"{error:{code:0,message:无效的请求参数}}"
//	@Failure	401				{object}	schema.ErrorResult		"{error:{code:0,message:未授权}}"
//	@Failure	500				{object}	schema.ErrorResult		"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/org-departments/{id} [put]
func (a *OrgDept) Update(c *gin.Context) {
}

// Delete 删除数据
//
//	@Tags		部门管理
//	@Summary	删除数据
//	@Security	ApiKeyAuth
//	@Param		Authorization	header		string				false	"Bearer 用户令牌"
//	@Param		id				path		string				true	"唯一标识"
//	@Success	200				{object}	schema.StatusResult	"{status:OK}"
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/org-departments/{id} [delete]
func (a *OrgDept) Delete(c *gin.Context) {
}

// Enable 启用数据
//
//	@Tags		部门管理
//	@Summary	启用数据
//	@Security	ApiKeyAuth
//	@Param		id	path		string				true	"唯一标识"
//	@Success	200	{object}	schema.StatusResult	"{status:OK}"
//	@Failure	401	{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500	{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/org-departments/{id}/enable [patch]
func (a *OrgDept) Enable(c *gin.Context) {
}

// Disable 禁用数据
//
//	@Tags		部门管理
//	@Summary	禁用数据
//	@Security	ApiKeyAuth
//	@Param		id	path		string				true	"唯一标识"
//	@Success	200	{object}	schema.StatusResult	"{status:OK}"
//	@Failure	401	{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500	{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/org-departments/{id}/disable [patch]
func (a *OrgDept) Disable(c *gin.Context) {
}
