package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// SysTeamSet 注入SysTeam
var SysTeamSet = wire.NewSet(wire.Struct(new(SysTeam), "*"))

// SysTeam 地址管理
type SysTeam struct {
}

// Query 查询数据
// @Tags 地址管理
// @Summary 查询数据
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param "" query schema.SysTeamQueryParam false "查询参数" default{}
// @Param current query int true "分页索引" default(1)
// @Param pageSize query int true "分页大小" default(10)
// @Param queryValue query string false "查询值"
// @Success 200 {array} schema.ListResult{list=schema.Systeams,pagination=schema.PaginationResult}
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v2/systeams [get]
func (a *SysTeam) Query(c *gin.Context) {
}

// Get 查询指定数据
// @Tags 地址管理
// @Summary 查询指定数据
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "唯一标识"
// @Success 200 {object} schema.SysTeam
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 404 {object} schema.ErrorResult "{error:{code:0,message:资源不存在}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v2/systeams/{id} [get]
func (a *SysTeam) Get(c *gin.Context) {
}

// View 查询指定数据
// @Tags 地址管理
// @Summary 查询指定数据
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "唯一标识"
// @Success 200 {object} schema.SysTeam
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 404 {object} schema.ErrorResult "{error:{code:0,message:资源不存在}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v2/systeams/{id}/view [get]
func (a *SysTeam) View(c *gin.Context) {
}

// Create 创建数据
// @Tags 地址管理
// @Summary 创建数据
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param body body schema.SysTeam true "创建数据"
// @Success 200 {object} schema.IDResult
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v2/systeams [post]
func (a *SysTeam) Create(c *gin.Context) {
}

// Update 更新数据
// @Tags 地址管理
// @Summary 更新数据
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "唯一标识"
// @Param body body schema.SysTeam true "更新数据"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v2/systeams/{id} [put]
func (a *SysTeam) Update(c *gin.Context) {
}

// Delete 删除数据
// @Tags 地址管理
// @Summary 删除数据
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "唯一标识"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v2/systeams/{id} [delete]
func (a *SysTeam) Delete(c *gin.Context) {
}

// Enable 启用数据
// @Tags 地址管理
// @Summary 启用数据
// @Security ApiKeyAuth
// @Param id path string true "唯一标识"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v2/systeams/{id}/enable [patch]
func (a *SysTeam) Enable(c *gin.Context) {
}

// Disable 禁用数据
// @Tags 地址管理
// @Summary 禁用数据
// @Security ApiKeyAuth
// @Param id path string true "唯一标识"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v2/systeams/{id}/disable [patch]
func (a *SysTeam) Disable(c *gin.Context) {
}
