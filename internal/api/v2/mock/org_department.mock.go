package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// OrgDepartmentSet 注入OrgDepartment
var OrgDepartmentSet = wire.NewSet(wire.Struct(new(OrgDepartment), "*"))

// OrgDepartment 部门管理
type OrgDepartment struct {
}

// Query 查询数据
// @Tags 部门管理
// @Summary 查询数据
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param current query int true "分页索引" default(1)
// @Param pageSize query int true "分页大小" default(10)
// @Param queryValue query string false "查询值"
// @Success 200 {array} schema.ListResult{list=schema.OrgDepartments,pagination=schema.PaginationResult}
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v2/org-departments [get]
func (a *OrgDepartment) Query(c *gin.Context) {
}

// Get 查询指定数据
// @Tags 部门管理
// @Summary 查询指定数据
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "唯一标识"
// @Success 200 {object} schema.OrgDepartment
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 404 {object} schema.ErrorResult "{error:{code:0,message:资源不存在}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v2/org-departments/{id} [get]
func (a *OrgDepartment) Get(c *gin.Context) {
}

// View 查询指定数据
// @Tags 部门管理
// @Summary 查询指定数据
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "唯一标识"
// @Success 200 {object} schema.OrgDepartment
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 404 {object} schema.ErrorResult "{error:{code:0,message:资源不存在}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v2/org-departments/{id}/view [get]
func (a *OrgDepartment) View(c *gin.Context) {
}

// Create 创建数据
// @Tags 部门管理
// @Summary 创建数据
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param body body schema.OrgDepartment true "创建数据"
// @Success 200 {object} schema.IDResult
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v2/org-departments [post]
func (a *OrgDepartment) Create(c *gin.Context) {
}

// Update 更新数据
// @Tags 部门管理
// @Summary 更新数据
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "唯一标识"
// @Param body body schema.OrgDepartment true "更新数据"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v2/org-departments/{id} [put]
func (a *OrgDepartment) Update(c *gin.Context) {
}

// Delete 删除数据
// @Tags 部门管理
// @Summary 删除数据
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "唯一标识"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v2/org-departments/{id} [delete]
func (a *OrgDepartment) Delete(c *gin.Context) {
}

// Enable 启用数据
// @Tags 部门管理
// @Summary 启用数据
// @Security ApiKeyAuth
// @Param id path string true "唯一标识"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v2/org-departments/{id}/enable [patch]
func (a *OrgDepartment) Enable(c *gin.Context) {
}

// Disable 禁用数据
// @Tags 部门管理
// @Summary 禁用数据
// @Security ApiKeyAuth
// @Param id path string true "唯一标识"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v2/org-departments/{id}/disable [patch]
func (a *OrgDepartment) Disable(c *gin.Context) {
}
