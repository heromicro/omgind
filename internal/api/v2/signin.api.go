package api_v2

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/heromicro/omgind/internal/app/ginx"
	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/app/service"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/global"
	"github.com/heromicro/omgind/pkg/logger"
	"github.com/heromicro/omgind/pkg/vcode"
)

// SignInSet 注入SignIn
var SignInSet = wire.NewSet(wire.Struct(new(SignIn), "*"))

// SignIn 登录管理
type SignIn struct {
	SigninSrv *service.SignIn
	Vcode     *vcode.Vcode
}

// GetCaptcha 获取验证码信息
func (a *SignIn) GetCaptcha(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.SigninSrv.GetCaptcha(ctx, global.CFG.Captcha.Length)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess2(c, item)
}

// ResCaptcha 响应图形验证码
func (a *SignIn) ResCaptcha(c *gin.Context) {
	ctx := c.Request.Context()
	captchaID := c.Query("id")
	if captchaID == "" {
		ginx.ResError(c, errors.New400Response("请提供验证码ID"))
		return
	}

	if c.Query("reload") != "" {
		if !a.Vcode.Reload(captchaID) {
			ginx.ResError(c, errors.New400Response("未找到验证码ID"))
			return
		}
	}

	cfg := global.CFG.Captcha
	err := a.SigninSrv.ResCaptcha(ctx, c.Writer, captchaID, cfg.Width, cfg.Height)
	if err != nil {
		ginx.ResError(c, err)
	}
}

// SignIn 用户登录
func (a *SignIn) SignIn(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.SignInParam

	if err := ginx.ParseJSON(c, &item); err != nil {
		//log.Println(" ----err- ", err)
		ginx.ResError(c, err)
		return
	}

	if !a.Vcode.Verify(item.CaptchaID, item.CaptchaCode, true) {
		ginx.ResError(c, errors.New400Response("无效的验证码"))
		return
	}

	user, err := a.SigninSrv.Verify(ctx, item.UserName, item.Password)

	if err != nil {
		ginx.ResError(c, err)
		return
	}

	userID := user.ID
	// 将用户ID放入上下文
	ginx.SetUserID(c, userID)

	tokenInfo, err := a.SigninSrv.GenerateToken(ctx, userID)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	ctx = logger.NewUserIDContext(ctx, userID)
	ctx = logger.NewTagContext(ctx, "__signin__")
	logger.WithContext(ctx).Infof("登入系统")
	ginx.ResSuccess2(c, tokenInfo)
}

// SignOut 用户登出
func (a *SignIn) SignOut(c *gin.Context) {
	ctx := c.Request.Context()

	// 检查用户是否处于登录状态，如果是则执行销毁
	userID := ginx.GetUserID(c)
	if userID != "" {
		ctx = logger.NewTagContext(ctx, "__signOut__")
		err := a.SigninSrv.DestroyToken(ctx, ginx.GetToken(c))
		if err != nil {
			logger.WithContext(ctx).Errorf(err.Error())
		}
		logger.WithContext(ctx).Infof("登出系统")
	}
	ginx.ResOK2(c, "登出系统")
}

// RefreshToken 刷新令牌
func (a *SignIn) RefreshToken(c *gin.Context) {
	ctx := c.Request.Context()
	tokenInfo, err := a.SigninSrv.GenerateToken(ctx, ginx.GetUserID(c))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess2(c, tokenInfo)
}

// GetUserInfo 获取当前用户信息
func (a *SignIn) GetUserInfo(c *gin.Context) {
	ctx := c.Request.Context()
	info, err := a.SigninSrv.GetSignInInfo(ctx, ginx.GetUserID(c))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess2(c, info)
}

// QueryUserMenuTree 查询当前用户菜单树
func (a *SignIn) QueryUserMenuTree(c *gin.Context) {

	ctx := c.Request.Context()
	menus, err := a.SigninSrv.QueryUserMenuTree(ctx, ginx.GetUserID(c))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResList2(c, menus)
}

// UpdatePassword 更新个人密码
func (a *SignIn) UpdatePassword(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.UpdatePasswordParam
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	err := a.SigninSrv.UpdatePassword(ctx, ginx.GetUserID(c), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK2(c, "更新密码成功")
}
