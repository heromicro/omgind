package service

import (
	"context"

	"github.com/casbin/casbin/v2"
	"github.com/google/wire"
	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/ent"
	"github.com/heromicro/omgind/internal/gen/ent/sysuserrole"
	"github.com/heromicro/omgind/internal/schema/repo"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/helper/hash"
)

// UserSet 注入User
var UserSet = wire.NewSet(wire.Struct(new(User), "*"))

// User 用户管理
type User struct {
	Enforcer *casbin.SyncedEnforcer

	UserRepo     *repo.User
	UserRoleRepo *repo.UserRole
	RoleRepo     *repo.Role
}

// Query 查询数据
func (a *User) Query(ctx context.Context, params schema.UserQueryParam, opts ...schema.UserQueryOptions) (*schema.UserQueryResult, error) {
	return a.UserRepo.Query(ctx, params, opts...)
}

// QueryShow 查询显示项数据
func (a *User) QueryShow(ctx context.Context, params schema.UserQueryParam, opts ...schema.UserQueryOptions) (*schema.UserShowQueryResult, error) {
	result, err := a.UserRepo.Query(ctx, params, opts...)
	if err != nil {
		return nil, err
	} else if result == nil {
		return nil, nil
	}

	userRoleResult, err := a.UserRoleRepo.Query(ctx, schema.UserRoleQueryParam{
		UserIDs: result.Data.ToIDs(),
	})
	if err != nil {
		return nil, err
	}

	roleResult, err := a.RoleRepo.Query(ctx, schema.RoleQueryParam{
		IDs: userRoleResult.Data.ToRoleIDs(),
	})
	if err != nil {
		return nil, err
	}

	return result.ToShowResult(userRoleResult.Data.ToUserIDMap(), roleResult.Data.ToMap()), nil
}

// Get 查询指定数据
func (a *User) Get(ctx context.Context, id string, opts ...schema.UserQueryOptions) (*schema.User, error) {
	item, err := a.UserRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	userRoleResult, err := a.UserRoleRepo.Query(ctx, schema.UserRoleQueryParam{
		UserID: id,
	})
	if err != nil {
		return nil, err
	}
	item.UserRoles = userRoleResult.Data

	return item, nil
}

// Create 创建数据
func (a *User) Create(ctx context.Context, item schema.User) (*schema.IDResult, error) {

	err := a.checkUserName(ctx, item)

	if err != nil {
		return nil, err
	}

	pword, _ := hash.MakePassword(item.Password)
	item.Password = pword

	err = repo.WithTx(ctx, a.UserRepo.EntCli, func(tx *ent.Tx) error {

		userInput := a.UserRepo.ToEntCreateSysUserInput(&item)
		userInput.CreatedAt = nil
		userInput.UpdatedAt = nil

		auser, err := tx.SysUser.Create().SetInput(*userInput).Save(ctx)
		if err != nil {
			return err
		}
		item.ID = auser.ID
		for _, urItem := range item.UserRoles {
			urItem.UserID = auser.ID

			_, err := tx.SysUserRole.Create().SetUserID(urItem.UserID).SetRoleID(urItem.RoleID).Save(ctx)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	LoadCasbinPolicy(ctx, a.Enforcer)
	return schema.NewIDResult(item.ID), nil
}

func (a *User) checkUserName(ctx context.Context, item schema.User) error {

	if item.UserName == schema.GetRootUser().UserName {
		return errors.New400Response("用户名不合法")
	}
	if len(item.UserName) < 6 {
		return errors.New400Response("用户名不合法")
	}
	result, err := a.UserRepo.Query(ctx, schema.UserQueryParam{
		PaginationParam: schema.PaginationParam{OnlyCount: true},
		UserName:        item.UserName,
	})
	if err != nil {
		return err
	} else if result.PageResult.Total > 0 {
		return errors.New400Response("用户名已经存在")
	}
	return nil
}

// Update 更新数据
func (a *User) Update(ctx context.Context, id string, item schema.User) error {
	oldItem, err := a.Get(ctx, id)

	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	} else if oldItem.UserName != item.UserName {
		err := a.checkUserName(ctx, item)
		if err != nil {
			return err
		}
	}

	if item.Password != "" {
		pword, _ := hash.MakePassword(item.Password)
		item.Password = pword
	} else {
		item.Password = oldItem.Password
	}

	item.ID = oldItem.ID
	item.Creator = oldItem.Creator
	item.CreatedAt = oldItem.CreatedAt
	err = repo.WithTx(ctx, a.UserRepo.EntCli, func(tx *ent.Tx) error {
		addUserRoles, delUserRoles := a.compareUserRoles(ctx, oldItem.UserRoles, item.UserRoles)

		// 添加的
		for _, uritem := range addUserRoles {
			_, err := tx.SysUserRole.Create().SetUserID(id).SetRoleID(uritem.RoleID).Save(ctx)
			if err != nil {
				return err
			}
		}
		// 删除的
		for _, uritem := range delUserRoles {
			_, err := tx.SysUserRole.Delete().Where(sysuserrole.UserIDEQ(id)).Where(sysuserrole.RoleIDEQ(uritem.
				RoleID)).Exec(ctx)
			if err != nil {
				return err
			}
		}

		user_input := a.UserRepo.ToEntUpdateSysUserInput(&item)
		user_input.UpdatedAt = nil
		_, err := tx.SysUser.UpdateOneID(id).SetInput(*user_input).Save(ctx)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	LoadCasbinPolicy(ctx, a.Enforcer)
	return nil
}

func (a *User) compareUserRoles(ctx context.Context, oldUserRoles, newUserRoles schema.UserRoles) (addList, delList schema.UserRoles) {
	mOldUserRoles := oldUserRoles.ToMap()
	mNewUserRoles := newUserRoles.ToMap()

	for k, item := range mNewUserRoles {
		if _, ok := mOldUserRoles[k]; ok {
			delete(mOldUserRoles, k)
			continue
		}
		addList = append(addList, item)
	}

	for _, item := range mOldUserRoles {
		delList = append(delList, item)
	}
	return
}

// Delete 删除数据
func (a *User) Delete(ctx context.Context, id string) error {

	oldItem, err := a.UserRepo.Get(ctx, id)

	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	err = a.UserRepo.Delete(ctx, id)

	if err != nil {
		return err
	}

	LoadCasbinPolicy(ctx, a.Enforcer)
	return nil
}

// UpdateStatus 更新状态
func (a *User) UpdateStatus(ctx context.Context, id string, status int16) error {
	oldItem, err := a.UserRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}
	oldItem.Status = status

	err = a.UserRepo.UpdateStatus(ctx, id, status)
	if err != nil {
		return err
	}

	LoadCasbinPolicy(ctx, a.Enforcer)
	return nil
}
