package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// SignInSet 注入SignIn
var SignInSet = wire.NewSet(wire.Struct(new(SignIn), "*"))

// SignIn 登录管理
type SignIn struct {
}

// GetCaptcha 获取验证码信息
// @Tags 登录管理
// @Summary 获取验证码信息
// @Success 200 {object} schema.SignInCaptcha
// @Router /api/v2/pub/signin/captchaid [get]
func (a *SignIn) GetCaptcha(c *gin.Context) {
}

// ResCaptcha 响应图形验证码
// @Tags 登录管理
// @Summary 响应图形验证码
// @Param id query string true "验证码ID"
// @Param reload query string false "重新加载"
// @Produce image/png
// @Success 200 "图形验证码"
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:无效的请求参数}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v2/pub/signin/captcha [get]
func (a *SignIn) ResCaptcha(c *gin.Context) {
}

// SignIn 用户登录
// @Tags 登录管理
// @Summary 用户登录
// @Param body body schema.SignInParam true "请求参数"
// @Success 200 {object} schema.SignInTokenInfo
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:无效的请求参数}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v2/pub/signin [post]
func (a *SignIn) SignIn(c *gin.Context) {
}

// SignOut 用户登出
// @Tags 登录管理
// @Summary 用户登出
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Router /api/v2/pub/signin/exit [post]
func (a *SignIn) SignOut(c *gin.Context) {
}

// RefreshToken 刷新令牌
// @Tags 登录管理
// @Summary 刷新令牌
// @Security ApiKeyAuth
// @Success 200 {object} schema.SignInTokenInfo
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v2/pub/refresh-token [post]
func (a *SignIn) RefreshToken(c *gin.Context) {
}

// GetUserInfo 获取当前用户信息
// @Tags 登录管理
// @Summary 获取当前用户信息
// @Security ApiKeyAuth
// @Success 200 {object} schema.UserSignInInfo
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v2/pub/current/user [get]
func (a *SignIn) GetUserInfo(c *gin.Context) {
}

// QueryUserMenuTree 查询当前用户菜单树
// @Tags 登录管理
// @Summary 查询当前用户菜单树
// @Security ApiKeyAuth
// @Success 200 {object} schema.ListResult{list=[]schema.MenuTree} "查询结果"
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v2/pub/current/menutree [get]
func (a *SignIn) QueryUserMenuTree(c *gin.Context) {
}

// UpdatePassword 更新个人密码
// @Tags 登录管理
// @Summary 更新个人密码
// @Security ApiKeyAuth
// @Param body body schema.UpdatePasswordParam true "请求参数"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/v2/pub/current/password [put]
func (a *SignIn) UpdatePassword(c *gin.Context) {
}
