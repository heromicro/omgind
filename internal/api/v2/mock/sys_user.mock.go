package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// UserSet 注入User
var UserSet = wire.NewSet(wire.Struct(new(User), "*"))

// User 用户管理
type User struct {
}

// Query 查询数据
//
//	@Tags		用户管理
//	@Summary	查询数据
//	@Security	ApiKeyAuth
//	@Param		""			query		schema.UserQueryParam	false	"查询参数"	default{}
//	@Success	200			{object}	schema.ListResult{list=schema.UserShows,pagination=schema.PaginationResult}"
//	@Failure	401			{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500			{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-users [get]
func (a *User) Query(c *gin.Context) {
}

// Get 查询指定数据
// Get 查询指定数据
//
//	@Tags		用户管理
//	@Summary	查询指定数据
//	@Security	ApiKeyAuth
//	@Param		id	path		string	true	"唯一标识"
//	@Success	200	{object}	schema.User
//	@Failure	401	{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	404	{object}	schema.ErrorResult	"{error:{code:0,message:资源不存在}}"
//	@Failure	500	{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-users/{id} [get]
func (a *User) Get(c *gin.Context) {
}

// Create 创建数据
//
//	@Tags		用户管理
//	@Summary	创建数据
//	@Security	ApiKeyAuth
//	@Param		body	body		schema.User	true	"创建数据"
//	@Success	200		{object}	schema.IDResult
//	@Failure	400		{object}	schema.ErrorResult	"{error:{code:0,message:无效的请求参数}}"
//	@Failure	401		{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500		{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-users [post]
func (a *User) Create(c *gin.Context) {
}

// Update 更新数据
//
//	@Tags		用户管理
//	@Summary	更新数据
//	@Security	ApiKeyAuth
//	@Param		id		path		string		true	"唯一标识"
//	@Param		body	body		schema.User	true	"更新数据"
//	@Success	200		{object}	schema.User
//	@Failure	400		{object}	schema.ErrorResult	"{error:{code:0,message:无效的请求参数}}"
//	@Failure	401		{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500		{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-users/{id} [put]
func (a *User) Update(c *gin.Context) {
}

// Delete 删除数据
//
//	@Tags		用户管理
//	@Summary	删除数据
//	@Security	ApiKeyAuth
//	@Param		id	path		string				true	"唯一标识"
//	@Success	200	{object}	schema.StatusResult	"{status:OK}"
//	@Failure	401	{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500	{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-users/{id} [delete]
func (a *User) Delete(c *gin.Context) {
}

// Enable 启用数据
//
//	@Tags		用户管理
//	@Summary	启用数据
//	@Security	ApiKeyAuth
//	@Param		id	path		string				true	"唯一标识"
//	@Success	200	{object}	schema.StatusResult	"{status:OK}"
//	@Failure	401	{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500	{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-users/{id}/enable [patch]
func (a *User) Enable(c *gin.Context) {
}

// Disable 禁用数据
//
//	@Tags		用户管理
//	@Summary	禁用数据
//	@Security	ApiKeyAuth
//	@Param		id	path		string				true	"唯一标识"
//	@Success	200	{object}	schema.StatusResult	"{status:OK}"
//	@Failure	401	{object}	schema.ErrorResult	"{error:{code:0,message:未授权}}"
//	@Failure	500	{object}	schema.ErrorResult	"{error:{code:0,message:服务器错误}}"
//	@Router		/api/v2/sys-users/{id}/disable [patch]
func (a *User) Disable(c *gin.Context) {
}
