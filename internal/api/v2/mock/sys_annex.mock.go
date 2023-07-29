package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// SysAnnexSet 注入SysAnnex
var SysAnnexSet = wire.NewSet(wire.Struct(new(SysAnnex), "*"))

// SysAnnex
type SysAnnex struct {
}

// Query 查询数据
// @Tags
// @Summary 查询数据
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param "" query schema.SysAnnexQueryParam false "查询参数" default{}
// @Param current query int true "分页索引" default(1)
// @Param pageSize query int true "分页大小" default(10)
// @Param q query string false "查询值"
// @Success 200 {array} schema.ListResult{list=schema.Sysannexes,pagination=schema.PaginationResult}
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v2/sysannexes [get]
func (a *SysAnnex) Query(c *gin.Context) {
}

// Get 查询指定数据
// @Tags
// @Summary 查询指定数据
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "唯一标识"
// @Success 200 {object} schema.SysAnnex
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 404 {object} schema.ErrorResult "{error:{code:0,message:资源不存在}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v2/sysannexes/{id} [get]
func (a *SysAnnex) Get(c *gin.Context) {
}

// View 查询指定数据
// @Tags
// @Summary 查询指定数据
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "唯一标识"
// @Success 200 {object} schema.SysAnnex
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 404 {object} schema.ErrorResult "{error:{code:0,message:资源不存在}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v2/sysannexes/{id}/view [get]
func (a *SysAnnex) View(c *gin.Context) {
}

// Create 创建数据
// @Tags
// @Summary 创建数据
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param body body schema.SysAnnex true "创建数据"
// @Success 200 {object} schema.ResSuccess
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v2/sysannexes [post]
func (a *SysAnnex) Create(c *gin.Context) {
}

// Update 更新数据
// @Tags
// @Summary 更新数据
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "唯一标识"
// @Param body body schema.SysAnnex true "更新数据"
// @Success 200 {object} schema.ResSuccess
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v2/sysannexes/{id} [put]
func (a *SysAnnex) Update(c *gin.Context) {
}

// Delete 删除数据
// @Tags
// @Summary 删除数据
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "唯一标识"
// @Success 200 {object} schema.ResOK
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v2/sysannexes/{id} [delete]
func (a *SysAnnex) Delete(c *gin.Context) {
}

// Enable 启用数据
// @Tags
// @Summary 启用数据
// @Security ApiKeyAuth
// @Param id path string true "唯一标识"
// @Success 200 {object} schema.ResOK
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v2/sysannexes/{id}/enable [patch]
func (a *SysAnnex) Enable(c *gin.Context) {
}

// Disable 禁用数据
// @Tags
// @Summary 禁用数据
// @Security ApiKeyAuth
// @Param id path string true "唯一标识"
// @Success 200 {object} schema.ResOK
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v2/sysannexes/{id}/disable [patch]
func (a *SysAnnex) Disable(c *gin.Context) {
}
