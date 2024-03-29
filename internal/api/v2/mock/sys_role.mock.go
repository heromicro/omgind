package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// RoleSet 注入Role
var RoleSet = wire.NewSet(wire.Struct(new(Role), "*"))

// Role 角色管理
type Role struct {
}

// Query 查询数据
//	@Tags		角色管理
//	@Summary	查询数据
//	@Security	ApiKeyAuth
//	@Param		current		query		int		true	"分页索引"	default(1)
//	@Param		pageSize	query		int		true	"分页大小"	default(10)
//	@Param		q	query		string	false	"查询值"
//	@Param		is_active	query		bool	false	"是否有效"
//	@Success	200			{object}	schema.ListResult{list=schema.Roles,pagination=schema.PaginationResult}"
//	@Failure	401			{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500			{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-roles [get]
func (a *Role) Query(c *gin.Context) {
}

// QuerySelect 查询选择数据
//	@Tags		角色管理
//	@Summary	查询选择数据
//	@Security	ApiKeyAuth
//	@Param		q	query		string									false	"查询值"
//	@Param		is_active	query		bool									false	"是否有效"
//	@Success	200			{object}	schema.ListResult{list=[]schema.Role}	"查询结果"
//	@Failure	400			{object}	schema.ErrorResult						"{error:{code:0,message:未知的查询类型}}"
//	@Failure	401			{object}	schema.ErrorResult						"{error:{code:0,message:未授权}}"
//	@Failure	500			{object}	schema.ErrorResult						"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-roles.select [get]
func (a *Role) QuerySelect(c *gin.Context) {
}

// Get 查询指定数据
//	@Tags		角色管理
//	@Summary	查询指定数据
//	@Security	ApiKeyAuth
//	@Param		id	path		string	true	"唯一标识"
//	@Success	200	{object}	schema.Role
//	@Failure	401	{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	404	{object}	schema.ErrorResult	"{error:{code:0,message:资源不存在}}"
//	@Failure	500	{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-roles/{id} [get]
func (a *Role) Get(c *gin.Context) {
}

// Create 创建数据
//	@Tags		角色管理
//	@Summary	创建数据
//	@Security	ApiKeyAuth
//	@Param		body	body		schema.Role	true	"创建数据"
//	@Success	200		{object}	schema.IDResult
//	@Failure	400		{object}	schema.ErrorResult	"{error:{code:0,message:无效的请求参数}}"
//	@Failure	401		{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500		{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-roles [post]
func (a *Role) Create(c *gin.Context) {
}

// Update 更新数据
//	@Tags		角色管理
//	@Summary	更新数据
//	@Security	ApiKeyAuth
//	@Param		id		path		string		true	"唯一标识"
//	@Param		body	body		schema.Role	true	"更新数据"
//	@Success	200		{object}	schema.Role
//	@Failure	400		{object}	schema.ErrorResult	"{error:{code:0,message:无效的请求参数}}"
//	@Failure	401		{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500		{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-roles/{id} [put]
func (a *Role) Update(c *gin.Context) {
}

// Delete 删除数据
//	@Tags		角色管理
//	@Summary	删除数据
//	@Security	ApiKeyAuth
//	@Param		id	path		string				true	"唯一标识"
//	@Success	200	{object}	schema.StatusResult	"{status:OK}"
//	@Failure	401	{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500	{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-roles/{id} [delete]
func (a *Role) Delete(c *gin.Context) {
}

// Enable 启用数据
//	@Tags		角色管理
//	@Summary	启用数据
//	@Security	ApiKeyAuth
//	@Param		id	path		string				true	"唯一标识"
//	@Success	200	{object}	schema.StatusResult	"{status:OK}"
//	@Failure	401	{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500	{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-roles/{id}/enable [patch]
func (a *Role) Enable(c *gin.Context) {
}

// Disable 禁用数据
//	@Tags		角色管理
//	@Summary	禁用数据
//	@Security	ApiKeyAuth
//	@Param		id	path		string				true	"唯一标识"
//	@Success	200	{object}	schema.StatusResult	"{status:OK}"
//	@Failure	401	{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500	{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-roles/{id}/disable [patch]
func (a *Role) Disable(c *gin.Context) {
}
