package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// MenuSet 注入Menu
var MenuSet = wire.NewSet(wire.Struct(new(Menu), "*"))

// Menu 菜单管理
type Menu struct{}

// Query 查询数据
//	@Tags		菜单管理
//	@Summary	查询数据
//	@Security	ApiKeyAuth
//	@Param		current		query		int		true	"分页索引"	default(1)
//	@Param		pageSize	query		int		true	"分页大小"	default(10)
//	@Param		q	query		string	false	"查询值"
//	@Param		is_active	query		bool	false	"是否有效"
//	@Param		isShow		query		bool	false	"显示状态(true:显示 false:隐藏)"
//	@Param		parentID	query		string	false	"父级ID"
//	@Success	200			{object}	schema.ListResult{list=schema.Menus,pagination=schema.PaginationResult}"
//	@Failure	401			{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500			{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-menus [get]
func (a *Menu) Query(c *gin.Context) {
}

// QueryTree 查询菜单树
//	@Tags		菜单管理
//	@Summary	查询菜单树
//	@Security	ApiKeyAuth
//	@Param		is_active	query		bool										false	"是否有效"
//	@Param		parentID	query		string										false	"父级ID"
//	@Success	200			{object}	schema.ListResult{list=[]schema.MenuTree}	"查询结果"
//	@Failure	401			{object}	schema.ErrorResult							"{error:{code:0,message:未授权}}"
//	@Failure	500			{object}	schema.ErrorResult							"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-menus.tree [get]
func (a *Menu) QueryTree(c *gin.Context) {
}

// Get 查询指定数据
//	@Tags		菜单管理
//	@Summary	查询指定数据
//	@Security	ApiKeyAuth
//	@Param		id	path		string	true	"唯一标识"
//	@Success	200	{object}	schema.Menu
//	@Failure	401	{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	404	{object}	schema.ErrorResult	"{error:{code:0,message:资源不存在}}"
//	@Failure	500	{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-menus/{id} [get]
func (a *Menu) Get(c *gin.Context) {
}

// Create 创建数据
//	@Tags		菜单管理
//	@Summary	创建数据
//	@Security	ApiKeyAuth
//	@Param		body	body		schema.Menu	true	"创建数据"
//	@Success	200		{object}	schema.IDResult
//	@Failure	400		{object}	schema.ErrorResult	"{error:{code:0,message:无效的请求参数}}"
//	@Failure	401		{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500		{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-menus [post]
func (a *Menu) Create(c *gin.Context) {
}

// Update 更新数据
//	@Tags		菜单管理
//	@Summary	更新数据
//	@Security	ApiKeyAuth
//	@Param		id		path		string				true	"唯一标识"
//	@Param		body	body		schema.Menu			true	"更新数据"
//	@Success	200		{object}	schema.StatusResult	"{status:OK}"
//	@Failure	400		{object}	schema.ErrorResult	"{error:{code:0,message:无效的请求参数}}"
//	@Failure	401		{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500		{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-menus/{id} [put]
func (a *Menu) Update(c *gin.Context) {
}

// Delete 删除数据
//	@Tags		菜单管理
//	@Summary	删除数据
//	@Security	ApiKeyAuth
//	@Param		id	path		string				true	"唯一标识"
//	@Success	200	{object}	schema.StatusResult	"{status:OK}"
//	@Failure	401	{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500	{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-menus/{id} [delete]
func (a *Menu) Delete(c *gin.Context) {
}

// Enable 启用数据
//	@Tags		菜单管理
//	@Summary	启用数据
//	@Security	ApiKeyAuth
//	@Param		id	path		string				true	"唯一标识"
//	@Success	200	{object}	schema.StatusResult	"{status:OK}"
//	@Failure	401	{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500	{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-menus/{id}/enable [patch]
func (a *Menu) Enable(c *gin.Context) {
}

// Disable 禁用数据
//	@Tags		菜单管理
//	@Summary	禁用数据
//	@Security	ApiKeyAuth
//	@Param		id	path		string				true	"唯一标识"
//	@Success	200	{object}	schema.StatusResult	"{status:OK}"
//	@Failure	401	{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500	{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-menus/{id}/disable [patch]
func (a *Menu) Disable(c *gin.Context) {
}
