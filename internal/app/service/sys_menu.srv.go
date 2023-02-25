package service

import (
	"context"
	"os"
	"time"

	"github.com/google/wire"
	"github.com/gotidy/ptr"
	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/ent"
	"github.com/heromicro/omgind/internal/gen/ent/sysmenu"
	"github.com/heromicro/omgind/internal/gen/ent/sysmenuactionresource"
	"github.com/heromicro/omgind/internal/schema/repo"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/helper/yaml"
)

// MenuSet 注入Menu
var SysMenuSet = wire.NewSet(wire.Struct(new(Menu), "*"))

// Menu 菜单管理
type Menu struct {
	//TransModel              *repo.Trans

	MenuRepo               *repo.Menu
	MenuActionRepo         *repo.MenuAction
	MenuActionResourceRepo *repo.MenuActionResource
}

// InitData 初始化菜单数据
func (a *Menu) InitData(ctx context.Context, dataFile string) error {

	result, err := a.MenuRepo.Query(ctx, schema.MenuQueryParam{
		PaginationParam: schema.PaginationParam{OnlyCount: true},
	})

	if err != nil {
		return err
	} else if result.PageResult.Total > 0 {
		// 如果存在则不进行初始化
		return nil
	}

	data, err := a.readData(dataFile)
	// log.Printf(" ------- ===== data: %+v \n", data)
	// log.Println(" ------- ===== err: \n", err)

	if err != nil {
		return err
	}

	err = a.createMenus(ctx, "", data)

	return err
}

func (a *Menu) readData(name string) (schema.MenuTrees, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data schema.MenuTrees
	d := yaml.NewDecoder(file)
	d.SetStrict(true)
	err = d.Decode(&data)
	return data, err
}

func (a *Menu) createMenus(ctx context.Context, parentID string, list schema.MenuTrees) error {

	for _, tritem := range list {
		sitem := schema.Menu{
			Name:     tritem.Name,
			Sort:     tritem.Sort,
			Level:    tritem.Level,
			Icon:     tritem.Icon,
			Router:   tritem.Router,
			ParentID: parentID,
			IsActive: ptr.Bool(true),
			IsShow:   tritem.IsShow,
			Actions:  tritem.Actions,
		}

		if err := a.checkName(ctx, sitem); err != nil {
			return err
		}

		parentPath, err := a.getParentPath(ctx, sitem.ParentID)
		if err != nil {
			return err
		}
		sitem.ParentPath = parentPath

		menuinput := a.MenuRepo.ToEntCreateSysMenuInput(&sitem)

		amenu, err := a.MenuRepo.EntCli.SysMenu.Create().SetInput(*menuinput).Save(ctx)
		if err != nil {
			return err
		}
		// 保存actions
		err = a.createActions(ctx, amenu.ID, sitem.Actions)
		if err != nil {
			return err
		}

		if tritem.Children != nil && len(*tritem.Children) > 0 {
			err := a.createMenus(ctx, amenu.ID, *tritem.Children)

			if err != nil {
				return err
			}
		}

	}

	//return nil
	//})

	return nil
}

// 创建动作数据
func (a *Menu) createActions(ctx context.Context, menuID string, items schema.MenuActions) error {

	for _, actitem := range items {
		actitem.MenuID = menuID
		mainput := a.MenuActionRepo.ToEntCreateSysMenuActionInput(actitem)
		mainput.MenuID = menuID

		anaction, err := a.MenuActionRepo.EntCli.SysMenuAction.Create().SetInput(*mainput).Save(ctx)
		if err != nil {
			return err
		}
		// 保存 resources
		for _, ritem := range actitem.Resources {
			ritem.ActionID = actitem.ID
			marinput := a.MenuActionResourceRepo.ToEntCreateSysMenuActionResourceInput(ritem)
			marinput.ActionID = anaction.ID

			_, err := a.MenuActionResourceRepo.EntCli.SysMenuActionResource.Create().SetInput(*marinput).Save(ctx)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Query 查询数据
func (a *Menu) Query(ctx context.Context, params schema.MenuQueryParam, opts ...schema.MenuQueryOptions) (*schema.MenuQueryResult, error) {
	menuActionResult, err := a.MenuActionRepo.Query(ctx, schema.MenuActionQueryParam{})
	if err != nil {
		return nil, err
	}

	result, err := a.MenuRepo.Query(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	result.Data.FillMenuAction(menuActionResult.Data.ToMenuIDMap())
	return result, nil
}

// Get 查询指定数据
func (a *Menu) Get(ctx context.Context, id string, opts ...schema.MenuQueryOptions) (*schema.Menu, error) {
	item, err := a.MenuRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	actions, err := a.QueryActions(ctx, id)
	if err != nil {
		return nil, err
	}
	item.Actions = actions

	return item, nil
}

// QueryActions 查询动作数据
func (a *Menu) QueryActions(ctx context.Context, id string) (schema.MenuActions, error) {
	result, err := a.MenuActionRepo.Query(ctx, schema.MenuActionQueryParam{
		MenuID: id,
	})
	if err != nil {
		return nil, err
	} else if len(result.Data) == 0 {
		return nil, nil
	}

	resourceResult, err := a.MenuActionResourceRepo.Query(ctx, schema.MenuActionResourceQueryParam{
		MenuID: id,
	})
	if err != nil {
		return nil, err
	}

	result.Data.FillResources(resourceResult.Data.ToActionIDMap())

	return result.Data, nil
}

func (a *Menu) checkName(ctx context.Context, item schema.Menu) error {
	result, err := a.MenuRepo.Query(ctx, schema.MenuQueryParam{
		PaginationParam: schema.PaginationParam{
			OnlyCount: true,
		},
		ParentID: &item.ParentID,
		Name:     item.Name,
	})
	if err != nil {
		return err
	} else if result.PageResult.Total > 0 {
		return errors.New400Response("菜单名称已经存在")
	}
	return nil
}

// Create 创建数据
func (a *Menu) Create(ctx context.Context, item schema.Menu) (*schema.IDResult, error) {

	if err := a.checkName(ctx, item); err != nil {
		return nil, err
	}

	pitem, err := a.getParent(ctx, item.ParentID)
	if err != nil {
		return nil, err
	}
	item.ParentPath = a.getParentPathNet(pitem, item.ParentID)
	item.Level = 1
	if pitem != nil {
		item.Level = pitem.Level + 1
	}
	item.IsLeaf = true

	err = repo.WithTx(ctx, a.MenuRepo.EntCli, func(tx *ent.Tx) error {

		menuinput := a.MenuRepo.ToEntCreateSysMenuInput(&item)
		amenu, err := tx.SysMenu.Create().SetInput(*menuinput).Save(ctx)
		if err != nil {
			return err
		}
		item.ID = amenu.ID
		if pitem != nil {
			tx.SysMenu.Update().SetIsLeaf(false).Where(sysmenu.IDEQ(pitem.ID)).Save(ctx)
		}

		// 保存actions
		err = a.createActionsTx(ctx, tx, amenu.ID, item.Actions)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return schema.NewIDResult(item.ID), nil
}

// 创建动作数据
func (a *Menu) createActionsTx(ctx context.Context, tx *ent.Tx, menuID string, items schema.MenuActions) error {

	for _, actitem := range items {
		actitem.MenuID = menuID
		mainput := a.MenuActionRepo.ToEntCreateSysMenuActionInput(actitem)
		mainput.MenuID = menuID

		anaction, err := tx.SysMenuAction.Create().SetInput(*mainput).Save(ctx)
		if err != nil {
			return err
		}
		// 保存 resources
		for _, ritem := range actitem.Resources {
			ritem.ActionID = actitem.ID
			marinput := a.MenuActionResourceRepo.ToEntCreateSysMenuActionResourceInput(ritem)
			marinput.ActionID = anaction.ID

			_, err := tx.SysMenuActionResource.Create().SetInput(*marinput).Save(ctx)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// 获取父级
func (a *Menu) getParent(ctx context.Context, parentID string) (*schema.Menu, error) {
	if parentID == "" {
		return nil, nil
	}

	pitem, err := a.MenuRepo.Get(ctx, parentID)
	if err != nil {
		return nil, err
	} else if pitem == nil {
		return nil, errors.ErrInvalidParent
	}

	return pitem, nil
}

func (a *Menu) getParentPathNet(pitem *schema.Menu, parentID string) string {
	if parentID == "" {
		return ""
	}
	if pitem == nil {
		return ""
	}
	return a.joinParentPath(pitem.ParentPath, pitem.ID)
}

// 获取父级路径
func (a *Menu) getParentPath(ctx context.Context, parentID string) (string, error) {
	if parentID == "" {
		return "", nil
	}

	pitem, err := a.getParent(ctx, parentID)
	if err != nil {
		return "", err
	} else if pitem == nil {
		return "", errors.ErrInvalidParent
	}

	return a.joinParentPath(pitem.ParentPath, pitem.ID), nil
}

func (a *Menu) joinParentPath(parent, id string) string {
	if parent != "" {
		return parent + "/" + id
	}
	return id
}

// Update 更新数据
func (a *Menu) Update(ctx context.Context, id string, item schema.Menu) error {
	if id == item.ParentID {
		return errors.ErrInvalidParent
	}

	oldItem, err := a.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	} else if oldItem.Name != item.Name {
		if err := a.checkName(ctx, item); err != nil {
			return err
		}
	}

	item.ID = oldItem.ID
	item.Creator = oldItem.Creator
	item.CreatedAt = oldItem.CreatedAt

	pitem, err := a.getParent(ctx, item.ParentID)

	if oldItem.ParentID != item.ParentID {
		// parentPath, err := a.getParentPath(ctx, item.ParentID)
		if err != nil {
			return err
		}
		item.ParentPath = a.getParentPathNet(pitem, item.ParentID)
		item.Level = 1
		if pitem == nil {
			item.Level = pitem.Level + 1
		}
	} else {
		item.ParentPath = oldItem.ParentPath
		item.Level = oldItem.Level
	}
	item.IsLeaf = oldItem.IsLeaf

	err = repo.WithTx(ctx, a.MenuRepo.EntCli, func(tx *ent.Tx) error {

		menuinput := a.MenuRepo.ToEntUpdateSysMenuInput(&item)
		_, err = tx.SysMenu.UpdateOneID(id).SetInput(*menuinput).Save(ctx)
		if err != nil {
			return err
		}

		if pitem != nil {
			tx.SysMenu.UpdateOneID(pitem.ID).SetIsLeaf(false).Save(ctx)
		}

		err := a.updateActions(ctx, tx, id, oldItem.Actions, item.Actions)
		if err != nil {
			return err
		}
		err = a.updateChildParentPath(ctx, tx, *oldItem, item)
		if err != nil {
			return err
		}

		return nil
	})
	return err
}

// 更新动作数据
func (a *Menu) updateActions(ctx context.Context, tx *ent.Tx, menuID string, oldItems,
	newItems schema.MenuActions) error {
	addActions, delActions, updateActions := a.compareActions(ctx, oldItems, newItems)

	err := a.createActions(ctx, menuID, addActions)
	if err != nil {
		return err
	}

	for _, item := range delActions {
		_, err := tx.SysMenuAction.UpdateOneID(item.ID).SetDeletedAt(time.Now()).SetIsDel(true).Save(ctx)

		if err != nil {
			return err
		}

		_, err1 := tx.SysMenuActionResource.Update().Where(sysmenuactionresource.ActionIDEQ(item.ID)).SetIsDel(true).SetDeletedAt(time.Now()).SetIsDel(true).Save(ctx)
		if err1 != nil {
			return err1
		}
	}

	mOldItems := oldItems.ToMap()
	for _, item := range updateActions {
		oitem := mOldItems[item.Code]
		// 只更新动作名称
		if item.Name != oitem.Name {
			oitem.Name = item.Name
			_, err := tx.SysMenuAction.UpdateOneID(item.ID).SetName(item.Name).Save(ctx)
			if err != nil {
				return err
			}
		}

		// 计算需要更新的资源配置（只包括新增和删除的，更新的不关心）
		addResources, delResources := a.compareResources(ctx, oitem.Resources, item.Resources)
		for _, aritem := range addResources {
			//aritem.ID = uid.MustString()
			aritem.ActionID = oitem.ID

			marinput := a.MenuActionResourceRepo.ToEntCreateSysMenuActionResourceInput(aritem)
			marinput.ActionID = oitem.ID
			_, err := tx.SysMenuActionResource.Create().SetInput(*marinput).Save(ctx)

			if err != nil {
				return err
			}
		}

		for _, ditem := range delResources {
			err := tx.SysMenuActionResource.DeleteOneID(ditem.ID).Exec(ctx)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// 对比动作列表
func (a *Menu) compareActions(ctx context.Context, oldActions, newActions schema.MenuActions) (addList, delList, updateList schema.MenuActions) {
	mOldActions := oldActions.ToMap()
	mNewActions := newActions.ToMap()

	for k, item := range mNewActions {
		if _, ok := mOldActions[k]; ok {
			updateList = append(updateList, item)
			delete(mOldActions, k)
			continue
		}
		addList = append(addList, item)
	}

	for _, item := range mOldActions {
		delList = append(delList, item)
	}
	return
}

// 对比资源列表
func (a *Menu) compareResources(ctx context.Context, oldResources, newResources schema.MenuActionResources) (addList, delList schema.MenuActionResources) {
	mOldResources := oldResources.ToMap()
	mNewResources := newResources.ToMap()

	for k, item := range mNewResources {
		if _, ok := mOldResources[k]; ok {
			delete(mOldResources, k)
			continue
		}
		addList = append(addList, item)
	}

	for _, item := range mOldResources {
		delList = append(delList, item)
	}
	return
}

// 检查并更新下级节点的父级路径
func (a *Menu) updateChildParentPath(ctx context.Context, tx *ent.Tx, oldItem, newItem schema.Menu) error {
	if oldItem.ParentID == newItem.ParentID {
		return nil
	}

	opath := a.joinParentPath(oldItem.ParentPath, oldItem.ID)
	result, err := a.MenuRepo.Query(ctx, schema.MenuQueryParam{
		PrefixParentPath: opath,
	})
	if err != nil {
		return err
	}

	npath := a.joinParentPath(newItem.ParentPath, newItem.ID)
	for _, menu := range result.Data {
		_, err := tx.SysMenu.UpdateOneID(menu.ID).SetParentPath(npath + menu.ParentPath[len(opath):]).Save(ctx)

		if err != nil {
			return err
		}
	}
	return nil
}

// Delete 删除数据
func (a *Menu) Delete(ctx context.Context, id string) error {

	oldItem, err := a.MenuRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	result, err := a.MenuRepo.Query(ctx, schema.MenuQueryParam{
		PaginationParam: schema.PaginationParam{OnlyCount: true},
		ParentID:        &id,
	})
	if err != nil {
		return err
	} else if result.PageResult.Total > 0 {
		return errors.ErrNotAllowDeleteWithChild
	}
	err = repo.WithTx(ctx, a.MenuRepo.EntCli, func(tx *ent.Tx) error {

		ma_result, err := a.MenuActionRepo.Query(ctx, schema.MenuActionQueryParam{
			MenuID: id,
		})
		if err != nil {
			return err
		}
		ma_ids := make([]string, len(ma_result.Data))
		for i, nit := range ma_result.Data {
			ma_ids[i] = nit.ID
		}

		_, err = tx.SysMenuActionResource.Update().
			SetIsDel(true).
			SetDeletedAt(time.Now()).
			Where(sysmenuactionresource.ActionIDIn(ma_ids...)).Save(ctx)

		if err != nil {
			return err
		}
		err = tx.SysMenu.DeleteOneID(id).Exec(ctx)
		if err != nil {
			return err
		}
		return nil
	})
	return err

	/*
		return a.TransModel.Exec(ctx, func(ctx context.Context) error {
			err = a.MenuActionResourceRepo.DeleteByMenuID(ctx, id)
			if err != nil {
				return err
			}

			err := a.MenuActionRepo.DeleteByMenuID(ctx, id)
			if err != nil {
				return err
			}

			return a.MenuRepo.Delete(ctx, id)
		})*/
}

// UpdateStatus 更新状态
func (a *Menu) UpdateStatus(ctx context.Context, id string, isActive bool) error {
	oldItem, err := a.MenuRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.MenuRepo.UpdateStatus(ctx, id, isActive)
}
