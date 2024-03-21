package service

import (
	"context"
	"log"
	"time"

	"github.com/google/wire"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/mainent"
	"github.com/heromicro/omgind/internal/gen/mainent/sysrolemenu"
	"github.com/heromicro/omgind/internal/scheme/repo"
	"github.com/heromicro/omgind/pkg/errors"
)

// RoleSet 注入Role
var SysRoleSet = wire.NewSet(wire.Struct(new(Role), "*"))

// Role 角色管理
type Role struct {
	// Enforcer *casbin.SyncedEnforcer

	EntCli *mainent.Client

	RoleRepo               *repo.Role
	RoleMenuRepo           *repo.RoleMenu
	UserRepo               *repo.User
	MenuActionResourceRepo *repo.MenuActionResource
}

// Query 查询数据
func (a *Role) Query(ctx context.Context, params schema.RoleQueryParam, opts ...schema.RoleQueryOptions) (*schema.RoleQueryResult, error) {
	return a.RoleRepo.Query(ctx, params, opts...)
}

// Query 查询数据
func (a *Role) QuerySelectPage(ctx context.Context, params schema.RoleQueryParam, opts ...schema.RoleQueryOptions) (*schema.RoleQueryResult, error) {
	return a.RoleRepo.QuerySelectPage(ctx, params, opts...)
}

// Get 查询指定数据
func (a *Role) Get(ctx context.Context, id string, opts ...schema.RoleQueryOptions) (*schema.Role, error) {
	item, err := a.RoleRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	roleMenus, err := a.QueryRoleMenus(ctx, id)
	if err != nil {
		return nil, err
	}
	item.RoleMenus = roleMenus

	return item, nil
}

// QueryRoleMenus 查询角色菜单列表
func (a *Role) QueryRoleMenus(ctx context.Context, roleID string) (schema.RoleMenus, error) {
	result, err := a.RoleMenuRepo.Query(ctx, schema.RoleMenuQueryParam{
		RoleID: roleID,
	})
	if err != nil {
		return nil, err
	}
	return result.Data, nil
}

// Create 创建数据
func (a *Role) Create(ctx context.Context, item schema.Role) (*schema.Role, error) {
	err := a.checkName(ctx, item)
	if err != nil {
		return nil, err
	}

	err = repo.WithTx(ctx, a.RoleRepo.EntCli, func(tx *mainent.Tx) error {
		role_input := repo.ToEntCreateSysRoleInput(&item)

		arole, err := tx.SysRole.Create().SetInput(*role_input).Save(ctx)
		if err != nil {
			return err
		}
		for _, rmitem := range item.RoleMenus {
			rminput := repo.ToEntCreateSysRoleMenuInput(rmitem)
			rminput.CreatedAt = nil
			rminput.UpdatedAt = nil
			rminput.RoleID = arole.ID
			_, err := tx.SysRoleMenu.Create().SetInput(*rminput).Save(ctx)
			if err != nil {
				return err
			}
		}
		item.ID = arole.ID
		return nil
	})
	if err != nil {
		return nil, err
	}

	// resources, err := a.MenuActionResourceRepo.Query(ctx, schema.MenuActionResourceQueryParam{
	// 	MenuIDs: item.RoleMenus.ToMenuIDs(),
	// })
	// if err != nil {
	// 	return nil, err
	// }
	// for _, ritem := range resources.Data.ToMap() {
	// 	a.Enforcer.AddPermissionForUser(item.ID, ritem.Path, ritem.Method)
	// }

	nitem, err := a.Get(ctx, item.ID)
	if err != nil {
		return nil, err
	}
	return nitem, nil
}

func (a *Role) checkName(ctx context.Context, item schema.Role) error {
	result, err := a.RoleRepo.Query(ctx, schema.RoleQueryParam{
		PaginationParam: schema.PaginationParam{OnlyCount: true},
		Name:            item.Name,
	})
	if err != nil {
		return err
	} else if result.PageResult.Total > 0 {
		return errors.New400Response("角色名称已经存在")
	}
	return nil
}

// Update 更新数据
func (a *Role) Update(ctx context.Context, id string, item schema.Role) (*schema.Role, error) {
	oldItem, err := a.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if oldItem == nil {
		return nil, errors.ErrNotFound
	} else if oldItem.Name != item.Name {
		err := a.checkName(ctx, item)
		if err != nil {
			return nil, err
		}
	}

	err = repo.WithTx(ctx, a.RoleRepo.EntCli, func(tx *mainent.Tx) error {

		addRoleMenus, delRoleMenus := a.compareRoleMenus(ctx, oldItem.RoleMenus, item.RoleMenus)
		for _, rmitem := range addRoleMenus {
			rmitem.RoleID = id
			rolemenu := repo.ToEntCreateSysRoleMenuInput(rmitem)

			_, err := tx.SysRoleMenu.Create().SetInput(*rolemenu).Save(ctx)
			if err != nil {
				return err
			}
		}

		for _, rmitem := range delRoleMenus {
			_, err := tx.SysRoleMenu.UpdateOneID(rmitem.ID).SetDeletedAt(time.Now()).SetIsDel(true).Save(ctx)

			if err != nil {
				return err
			}
		}

		rminput := repo.ToEntUpdateSysRoleInput(&item)
		_, err := tx.SysRole.UpdateOneID(id).SetInput(*rminput).Save(ctx)
		if err != nil {
			return err
		}

		return nil
	})

	log.Println(" ------- ====== ", err)

	if err != nil {
		return nil, err
	}

	// roleMenus, err := a.RoleMenuRepo.Query(ctx, schema.RoleMenuQueryParam{
	// 	RoleID: id,
	// })
	// if err != nil {
	// 	return nil, err
	// }

	// resources, err := a.MenuActionResourceRepo.Query(ctx, schema.MenuActionResourceQueryParam{
	// 	MenuIDs: roleMenus.Data.ToMenuIDs(),
	// })
	// if err != nil {
	// 	return nil, err
	// }

	// a.Enforcer.DeletePermissionsForUser(item.ID)
	// for _, ritem := range resources.Data.ToMap() {
	// 	a.Enforcer.AddPermissionForUser(item.ID, ritem.Path, ritem.Method)
	// }

	nitem, err := a.Get(ctx, item.ID)
	if err != nil {
		return nil, err
	}
	return nitem, nil
}

func (a *Role) compareRoleMenus(ctx context.Context, oldRoleMenus, newRoleMenus schema.RoleMenus) (addList, delList schema.RoleMenus) {
	mOldRoleMenus := oldRoleMenus.ToMap()
	mNewRoleMenus := newRoleMenus.ToMap()

	for k, item := range mNewRoleMenus {
		if _, ok := mOldRoleMenus[k]; ok {
			delete(mOldRoleMenus, k)
			continue
		}
		addList = append(addList, item)
	}

	for _, item := range mOldRoleMenus {
		delList = append(delList, item)
	}
	return
}

// Delete 删除数据
func (a *Role) Delete(ctx context.Context, id string) error {
	oldItem, err := a.RoleRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	userResult, err := a.UserRepo.Query(ctx, schema.UserQueryParam{
		PaginationParam: schema.PaginationParam{OnlyCount: true},
		RoleIDs:         []string{id},
	})
	if err != nil {
		return err
	} else if userResult.PageResult.Total > 0 {
		return errors.New400Response("该角色已被赋予用户，不允许删除")
	}

	err = repo.WithTx(ctx, a.RoleRepo.EntCli, func(tx *mainent.Tx) error {
		_, err := tx.SysRoleMenu.Update().Where(sysrolemenu.RoleIDEQ(id)).SetDeletedAt(time.Now()).SetIsDel(true).Save(ctx)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	// a.Enforcer.DeleteRole(id)

	return nil
}

// UpdateActive 更新状态
func (a *Role) UpdateActive(ctx context.Context, id string, isActive bool) error {
	oldItem, err := a.RoleRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	err = a.RoleRepo.UpdateActive(ctx, id, isActive)
	if err != nil {
		return err
	}

	if isActive {
		// roleMenus, err := a.RoleMenuRepo.Query(ctx, schema.RoleMenuQueryParam{
		// 	RoleID: id,
		// })
		// if err != nil {
		// 	return err
		// }
		// resources, err := a.MenuActionResourceRepo.Query(ctx, schema.MenuActionResourceQueryParam{
		// 	MenuIDs: roleMenus.Data.ToMenuIDs(),
		// })
		// if err != nil {
		// 	return err
		// }

		// for _, ritem := range resources.Data.ToMap() {
		// 	a.Enforcer.AddPermissionForUser(id, ritem.Path, ritem.Method)
		// }
	} else {
		// a.Enforcer.DeleteRole(id)
	}

	return nil
}
