package schema

import "C"
import (
	"context"
	"time"

	"github.com/heromicro/omgind/pkg/global"
	"github.com/heromicro/omgind/pkg/helper/hash"
	"github.com/heromicro/omgind/pkg/helper/json"
	"github.com/heromicro/omgind/pkg/helper/structure"
)

// GetRootUser 获取root用户
func GetRootUser() *User {
	user := global.CFG.Root
	return &User{
		ID:        user.UserName,
		UserName:  user.UserName,
		RealName:  user.RealName,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  hash.MD5String(user.Password),
	}
}

// CheckIsRootUser 检查是否是root用户
func CheckIsRootUser(ctx context.Context, userID string) bool {
	return GetRootUser().ID == userID
}

// User 用户对象
type User struct {
	TimeMixin `yaml:"-"`

	ID        string `json:"id" `                           // 唯一标识
	UserName  string `json:"user_name" binding:"required"`  // 用户名
	RealName  string `json:"real_name"`                     // 真实姓名
	FirstName string `json:"first_name" binding:"required"` // 真实姓名
	LastName  string `json:"last_name" binding:"required"`  // 真实姓名
	Password  string `json:"password"`                      // 密码
	Mobile    string `json:"mobile"`                        // 手机号
	Email     string `json:"email"`                         // 邮箱
	IsActive  *bool  `json:"is_active" binding:"required"`  // 状态
	Gender    int    `json:"gender" binding:"max=3,min=1"`  // 性别(1:男,2:女)

	UserRoles UserRoles `json:"user_roles" binding:"required,gt=0"` // 角色授权
}

func (a *User) String() string {
	return json.MarshalToString(a)
}

// CleanSecure 清理安全数据
func (a *User) CleanSecure() *User {
	a.Password = ""
	return a
}

// UserQueryParam 查询条件
type UserQueryParam struct {
	PaginationParam
	UserName   string   `form:"userName"`  // 用户名
	QueryValue string   `form:"q"`         // 模糊查询
	IsActive   *bool    `form:"is_active"` //
	RoleIDs    []string `form:"role_ids"`  // 角色ID列表

	BasicOrderParam
	TimeOrderParam
}

// UserQueryOptions 查询可选参数项
type UserQueryOptions struct {
	OrderFields []*OrderField // 排序字段
	UserRoles   *bool         // query user roles
}

// UserQueryResult 查询结果
type UserQueryResult struct {
	Data       Users
	PageResult *PaginationResult
}

// ToShowResult 转换为显示结果
func (a UserQueryResult) ToShowResult(mUserRoles map[string]UserRoles, mRoles map[string]*Role) *UserShowQueryResult {
	return &UserShowQueryResult{
		PageResult: a.PageResult,
		Data:       a.Data.ToUserShows(mUserRoles, mRoles),
	}
}

// Users 用户对象列表
type Users []*User

// ToIDs 转换为唯一标识列表
func (a Users) ToIDs() []string {
	idList := make([]string, len(a))
	for i, item := range a {
		idList[i] = item.ID
	}
	return idList
}

// ToUserShows 转换为用户显示列表
func (a Users) ToUserShows(mUserRoles map[string]UserRoles, mRoles map[string]*Role) UserShows {
	list := make(UserShows, len(a))
	for i, item := range a {
		showItem := new(UserShow)
		structure.Copy(item, showItem)
		for _, roleID := range mUserRoles[item.ID].ToRoleIDs() {
			if v, ok := mRoles[roleID]; ok {
				showItem.Roles = append(showItem.Roles, v)
			}
		}
		list[i] = showItem
	}

	return list
}

// ----------------------------------------UserRole--------------------------------------

// UserRole 用户角色
type UserRole struct {
	ID     string `json:"id"`      // 唯一标识
	UserID string `json:"user_id"` // 用户ID
	RoleID string `json:"role_id"` // 角色ID
}

// UserRoleQueryParam 查询条件
type UserRoleQueryParam struct {
	PaginationParam
	UserID  string   // 用户ID
	UserIDs []string // 用户ID列表

	BasicOrderParam
	TimeOrderParam
}

// UserRoleQueryOptions 查询可选参数项
type UserRoleQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

// UserRoleQueryResult 查询结果
type UserRoleQueryResult struct {
	Data       UserRoles
	PageResult *PaginationResult
}

// UserRoles 角色菜单列表
type UserRoles []*UserRole

// ToMap 转换为map
func (a UserRoles) ToMap() map[string]*UserRole {
	m := make(map[string]*UserRole)
	for _, item := range a {
		m[item.RoleID] = item
	}
	return m
}

// ToRoleIDs 转换为角色ID列表
func (a UserRoles) ToRoleIDs() []string {
	list := make([]string, len(a))
	for i, item := range a {
		list[i] = item.RoleID
	}
	return list
}

// ToUserIDMap 转换为用户ID映射
func (a UserRoles) ToUserIDMap() map[string]UserRoles {
	m := make(map[string]UserRoles)
	for _, item := range a {
		m[item.UserID] = append(m[item.UserID], item)
	}
	return m
}

// ----------------------------------------UserShow--------------------------------------

// UserShow 用户显示项
type UserShow struct {
	ID        string    `json:"id"`         // 唯一标识
	UserName  string    `json:"user_name"`  // 用户名
	RealName  string    `json:"real_name"`  // 真实姓名
	FirstName string    `json:"first_name"` // 真实名
	LastName  string    `json:"last_name"`  // 真实姓
	Mobile    string    `json:"mobile"`     // 手机号
	Email     string    `json:"email"`      // 邮箱
	IsActive  *bool     `json:"is_active"`  // 状态
	CreatedAt time.Time `json:"created_at"` // 创建时间
	Roles     []*Role   `json:"roles"`      // 授权角色列表
}

// UserShows 用户显示项列表
type UserShows []*UserShow

// UserShowQueryResult 用户显示项查询结果
type UserShowQueryResult struct {
	Data       UserShows
	PageResult *PaginationResult
}
