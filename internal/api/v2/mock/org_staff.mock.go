package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// OrgStaffSet 注入OrgStaff
var OrgStaffSet = wire.NewSet(wire.Struct(new(OrgStaff), "*"))

// OrgStaff 员工管理
type OrgStaff struct {
}

// Query 查询数据
//
//	@Tags		员工管理
//	@Summary	查询数据
//	@Security	ApiKeyAuth
//	@Param		Authorization	header		string						false	"Bearer 用户令牌"
//	@Param		object			query		schema.OrgStaffQueryParam	false	"查询参数"	default{}
//	@Success	200				{array}		schema.ListResult{list=schema.OrgStaffs,pagination=schema.PaginationResult}
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/org-staffs [get]
func (a *OrgStaff) Query(c *gin.Context) {
}

// Get 查询指定数据
//
//	@Tags		员工管理
//	@Summary	查询指定数据
//	@Security	ApiKeyAuth
//	@Param		Authorization	header		string	false	"Bearer 用户令牌"
//	@Param		id				path		string	true	"唯一标识"
//	@Success	200				{object}	schema.OrgStaff
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	404				{object}	schema.ErrorResult	"{error:{code:0,message:资源不存在}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/org-staffs/{id} [get]
func (a *OrgStaff) Get(c *gin.Context) {
}

// View 查询指定数据
//
//	@Tags		员工管理
//	@Summary	查询指定数据
//	@Security	ApiKeyAuth
//	@Param		Authorization	header		string	false	"Bearer 用户令牌"
//	@Param		id				path		string	true	"唯一标识"
//	@Success	200				{object}	schema.OrgStaff
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	404				{object}	schema.ErrorResult	"{error:{code:0,message:资源不存在}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/org-staffs/{id}/view [get]
func (a *OrgStaff) View(c *gin.Context) {
}

// Create 创建数据
//
//	@Tags		员工管理
//	@Summary	创建数据
//	@Security	ApiKeyAuth
//	@Param		Authorization	header		string			false	"Bearer 用户令牌"
//	@Param		body			body		schema.OrgStaff	true	"创建数据"
//	@Success	200				{object}	schema.IDResult
//	@Failure	400				{object}	schema.ErrorResult	"{error:{code:0,message:无效的请求参数}}"
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/org-staffs [post]
func (a *OrgStaff) Create(c *gin.Context) {
}

// Update 更新数据
//
//	@Tags		员工管理
//	@Summary	更新数据
//	@Security	ApiKeyAuth
//	@Param		Authorization	header		string				false	"Bearer 用户令牌"
//	@Param		id				path		string				true	"唯一标识"
//	@Param		body			body		schema.OrgStaff		true	"更新数据"
//	@Success	200				{object}	schema.StatusResult	"{status:OK}"
//	@Failure	400				{object}	schema.ErrorResult	"{error:{code:0,message:无效的请求参数}}"
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/org-staffs/{id} [put]
func (a *OrgStaff) Update(c *gin.Context) {
}

// Delete 删除数据
//
//	@Tags		员工管理
//	@Summary	删除数据
//	@Security	ApiKeyAuth
//	@Param		Authorization	header		string				false	"Bearer 用户令牌"
//	@Param		id				path		string				true	"唯一标识"
//	@Success	200				{object}	schema.StatusResult	"{status:OK}"
//	@Failure	401				{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500				{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/org-staffs/{id} [delete]
func (a *OrgStaff) Delete(c *gin.Context) {
}

// Enable 启用数据
//
//	@Tags		员工管理
//	@Summary	启用数据
//	@Security	ApiKeyAuth
//	@Param		id	path		string				true	"唯一标识"
//	@Success	200	{object}	schema.StatusResult	"{status:OK}"
//	@Failure	401	{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500	{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/org-staffs/{id}/enable [patch]
func (a *OrgStaff) Enable(c *gin.Context) {
}

// Disable 禁用数据
//
//	@Tags		员工管理
//	@Summary	禁用数据
//	@Security	ApiKeyAuth
//	@Param		id	path		string				true	"唯一标识"
//	@Success	200	{object}	schema.StatusResult	"{status:OK}"
//	@Failure	401	{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500	{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/org-staffs/{id}/disable [patch]
func (a *OrgStaff) Disable(c *gin.Context) {
}
