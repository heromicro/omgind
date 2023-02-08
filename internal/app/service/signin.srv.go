package service

import (
	"context"
	"net/http"
	"sort"

	"github.com/google/wire"
	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/schema/repo"
	"github.com/heromicro/omgind/pkg/auth"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/helper/hash"
	"github.com/heromicro/omgind/pkg/vcode"
)

// SignInSet 注入SignIn
var SignInSet = wire.NewSet(wire.Struct(new(SignIn), "*"))

// SignIn 登录管理
type SignIn struct {
	Auth            auth.Auther
	UserModel       *repo.User
	UserRoleModel   *repo.UserRole
	RoleModel       *repo.Role
	RoleMenuModel   *repo.RoleMenu
	MenuModel       *repo.Menu
	MenuActionModel *repo.MenuAction
	Vcode           *vcode.Vcode
}

// GetCaptcha 获取图形验证码信息
func (a *SignIn) GetCaptcha(ctx context.Context, length int) (*schema.SignInCaptcha, error) {

	captchaID := a.Vcode.NewLen(length)
	item := &schema.SignInCaptcha{
		CaptchaID: captchaID,
	}
	return item, nil
}

// ResCaptcha 生成并响应图形验证码
func (a *SignIn) ResCaptcha(ctx context.Context, w http.ResponseWriter, captchaID string, width, height int) error {

	err := a.Vcode.GenerateImage(w, captchaID)
	if err != nil {
		if err == vcode.ErrNotFound {
			return errors.ErrNotFound
		}
		return errors.WithStack(err)
	}

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Type", "image/png")

	return nil
}

// Verify 登录验证
func (a *SignIn) Verify(ctx context.Context, userName, password string) (*schema.User, error) {
	// 检查是否是超级用户
	root := schema.GetRootUser()
	if userName == root.UserName && root.Password == password {
		return root, nil
	}

	result, err := a.UserModel.Query(ctx, schema.UserQueryParam{
		UserName: userName,
	})
	if err != nil {
		return nil, err
	} else if len(result.Data) == 0 {
		return nil, errors.ErrInvalidUserName
	}

	item := result.Data[0]
	ok, err := hash.CheckPassword(password, item.Password)
	if err != nil || !ok {
		return nil, errors.ErrInvalidPassword
	} else if item.Status != 1 {
		return nil, errors.ErrUserDisable
	}

	return item, nil
}

// GenerateToken 生成令牌
func (a *SignIn) GenerateToken(ctx context.Context, userID string) (*schema.SignInTokenInfo, error) {
	tokenInfo, err := a.Auth.GenerateToken(ctx, userID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	item := &schema.SignInTokenInfo{
		AccessToken: tokenInfo.GetAccessToken(),
		TokenType:   tokenInfo.GetTokenType(),
		ExpiresAt:   tokenInfo.GetExpiresAt(),
	}
	return item, nil
}

// DestroyToken 销毁令牌
func (a *SignIn) DestroyToken(ctx context.Context, tokenString string) error {
	err := a.Auth.DestroyToken(ctx, tokenString)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (a *SignIn) checkAndGetUser(ctx context.Context, userID string) (*schema.User, error) {
	user, err := a.UserModel.Get(ctx, userID)
	if err != nil {
		return nil, err
	} else if user == nil {
		return nil, errors.ErrInvalidUser
	} else if user.Status != 1 {
		return nil, errors.ErrUserDisable
	}
	return user, nil
}

// GetSignInInfo 获取当前用户登录信息
func (a *SignIn) GetSignInInfo(ctx context.Context, userID string) (*schema.UserSignInInfo, error) {
	if isRoot := schema.CheckIsRootUser(ctx, userID); isRoot {
		root := schema.GetRootUser()
		signinInfo := &schema.UserSignInInfo{
			UserName: root.UserName,
			RealName: root.RealName,
			FistName: root.FirstName,
			LastName: root.LastName,
		}
		return signinInfo, nil
	}

	user, err := a.checkAndGetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	info := &schema.UserSignInInfo{
		UserID:   user.ID,
		UserName: user.UserName,
		RealName: user.RealName,
	}

	userRoleResult, err := a.UserRoleModel.Query(ctx, schema.UserRoleQueryParam{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}

	if roleIDs := userRoleResult.Data.ToRoleIDs(); len(roleIDs) > 0 {
		roleResult, err := a.RoleModel.Query(ctx, schema.RoleQueryParam{
			IDs:    roleIDs,
			Status: 1,
		})
		if err != nil {
			return nil, err
		}
		info.Roles = roleResult.Data
	}

	return info, nil
}

// QueryUserMenuTree 查询当前用户的权限菜单树
func (a *SignIn) QueryUserMenuTree(ctx context.Context, userID string) (schema.MenuTrees, error) {

	isRoot := schema.CheckIsRootUser(ctx, userID)
	// 如果是root用户，则查询所有显示的菜单树
	if isRoot {
		result, err := a.MenuModel.Query(ctx, schema.MenuQueryParam{
			Status: 1,
		}, schema.MenuQueryOptions{
			OrderFields: schema.NewOrderFields(schema.NewOrderField("sort", schema.OrderByASC)),
		})

		if err != nil {
			return nil, err
		}

		menuActionResult, err := a.MenuActionModel.Query(ctx, schema.MenuActionQueryParam{})

		if err != nil {
			return nil, err
		}
		return result.Data.FillMenuAction(menuActionResult.Data.ToMenuIDMap()).ToTree(), nil
	}

	userRoleResult, err := a.UserRoleModel.Query(ctx, schema.UserRoleQueryParam{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	} else if len(userRoleResult.Data) == 0 {
		return nil, errors.ErrNoPerm
	}

	roleMenuResult, err := a.RoleMenuModel.Query(ctx, schema.RoleMenuQueryParam{
		RoleIDs: userRoleResult.Data.ToRoleIDs(),
	})
	if err != nil {
		return nil, err
	} else if len(roleMenuResult.Data) == 0 {
		return nil, errors.ErrNoPerm
	}

	menuResult, err := a.MenuModel.Query(ctx, schema.MenuQueryParam{
		IDs:    roleMenuResult.Data.ToMenuIDs(),
		Status: 1,
	})
	if err != nil {
		return nil, err
	} else if len(menuResult.Data) == 0 {
		return nil, errors.ErrNoPerm
	}

	mData := menuResult.Data.ToMap()
	// 获取授权菜单的父级菜单，判断哪些父级菜单不在之前的授权菜单中，存放于qIDs切片
	var qIDs []string
	for _, pid := range menuResult.Data.SplitParentIDs() {
		if _, ok := mData[pid]; !ok {
			qIDs = append(qIDs, pid)
		}
	}
	// 获取这些差异的父级菜单的信息，补充到menuResult.Data中
	if len(qIDs) > 0 {
		pmenuResult, err := a.MenuModel.Query(ctx, schema.MenuQueryParam{
			IDs: qIDs,
		})
		if err != nil {
			return nil, err
		}
		menuResult.Data = append(menuResult.Data, pmenuResult.Data...)
	}

	sort.Sort(menuResult.Data)
	menuActionResult, err := a.MenuActionModel.Query(ctx, schema.MenuActionQueryParam{
		IDs: roleMenuResult.Data.ToActionIDs(),
	})
	if err != nil {
		return nil, err
	}
	return menuResult.Data.FillMenuAction(menuActionResult.Data.ToMenuIDMap()).ToTree(), nil
}

// UpdatePassword 更新当前用户登录密码
func (a *SignIn) UpdatePassword(ctx context.Context, userID string, params schema.UpdatePasswordParam) error {
	if schema.CheckIsRootUser(ctx, userID) {
		return errors.New400Response("root用户不允许更新密码")
	}

	user, err := a.checkAndGetUser(ctx, userID)
	if err != nil {
		return err
	} else {
		ok, err := hash.CheckPassword(params.OldPassword, user.Password)
		if err != nil || !ok {
			return errors.New400Response("旧密码不正确")
		}
	}

	params.NewPassword = hash.SHA1String(params.NewPassword)
	return a.UserModel.UpdatePassword(ctx, userID, params.NewPassword)
}
