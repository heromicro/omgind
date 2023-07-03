package service

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/google/wire"
	"github.com/gotidy/ptr"
	"github.com/hibiken/asynq"
	"github.com/jossef/format"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/entscheme"
	"github.com/heromicro/omgind/internal/gen/entscheme/sysmenu"
	"github.com/heromicro/omgind/internal/gen/entscheme/sysmenuactionresource"
	"github.com/heromicro/omgind/internal/scheme/repo"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/helper/yaml"
	"github.com/heromicro/omgind/pkg/mw/asyncq/worker"
	"github.com/heromicro/omgind/pkg/mw/queue"
	"github.com/heromicro/omgind/pkg/types"
)

// MenuSet 注入Menu
// var SysMenuSet = wire.NewSet(wire.Struct(new(Menu), "*"))
var SysMenuSet = wire.NewSet(NewMenuSrv)

// Menu 菜单管理
type Menu struct {
	EntCli *entscheme.Client

	MenuRepo               *repo.Menu
	MenuActionRepo         *repo.MenuAction
	MenuActionResourceRepo *repo.MenuActionResource

	Queue    queue.Queuer
	consumer *worker.Consumer
}

func NewMenuSrv(entCli *entscheme.Client, menuRepo *repo.Menu, menuActionRepo *repo.MenuAction, menuActionResourceRepo *repo.MenuActionResource, q queue.Queuer, c *worker.Consumer) *Menu {

	menuSrv := &Menu{
		EntCli:                 entCli,
		MenuRepo:               menuRepo,
		MenuActionRepo:         menuActionRepo,
		MenuActionResourceRepo: menuActionResourceRepo,
		Queue:                  q,
		consumer:               c,
	}

	// TODO: register tasks
	menuSrv.consumer.RegisterHandlers(types.TaskName_REPAIR_MENU_TREE_PATH, menuSrv)

	return menuSrv
}

// repair parent_path, is_leaf, level
func (s *Menu) ProcessTask(ctx context.Context, t *asynq.Task) error {

	log.Println(" --- ---- === = ", t.Type())
	id := string(t.Payload())
	if id == "" {
		return nil
	}

	log.Println(" --- ---- === = id: ", id)

	zmenu, err := s.EntCli.SysMenu.Query().Where(sysmenu.IDEQ(id)).First(ctx)
	if err != nil {
		if !entscheme.IsNotFound(err) {
			return err
		}
	}

	var parent *entscheme.SysMenu = nil
	if zmenu.ParentID != nil && *zmenu.ParentID != "" {
		parent, err = s.EntCli.SysMenu.Query().Where(sysmenu.IDEQ(*zmenu.ParentID)).First(ctx)
		if err != nil {

			if !entscheme.IsNotFound(err) {

				return err
			}
		}
	}

	if zmenu.IsDel {

		s.EntCli.SysMenu.Update().Where(sysmenu.ParentIDEQ(id)).SetIsDel(true).Save(ctx)
		// if err != nil {
		// 	return err
		// }
		if zmenu.ParentID != nil && *zmenu.ParentID != "" {
			mcount, err := s.EntCli.SysMenu.Query().Where(sysmenu.ParentIDEQ(*zmenu.ParentID), sysmenu.IsDelEQ(false)).Count(ctx)

			if err != nil {
				return err
			}

			if mcount == 0 {
				s.EntCli.SysMenu.Update().Where(sysmenu.IDEQ(*zmenu.ParentID)).SetIsLeaf(true).Save(ctx)
			}

		}

	} else {

		submenus, err := s.EntCli.SysMenu.Query().Where(sysmenu.ParentIDEQ(id)).Where(sysmenu.IsDelEQ(false)).All(ctx)
		if err != nil {
			log.Println(" ----- ==== --- err - ", err)
			if !entscheme.IsNotFound(err) {
				return err
			}
		}

		if len(submenus) > 0 {

			ppath := s.joinParentPath(*zmenu.ParentPath, zmenu.ID)
			err := repo.WithTx(ctx, s.EntCli, func(tx *entscheme.Tx) error {

				_, err := tx.SysMenu.Update().Where(sysmenu.ParentIDEQ(zmenu.ID)).SetParentPath(ppath).SetLevel(zmenu.Level + 1).Save(ctx)
				if err != nil {
					if !entscheme.IsNotFound(err) {
						return err
					}
				}
				_, err = tx.SysMenu.UpdateOneID(zmenu.ID).SetIsLeaf(false).Save(ctx)
				if err != nil {
					return err
				}

				return err
			})

			if err != nil {
				return err
			}

			for _, submenu := range submenus {
				jobid := format.String(`{id}-{ml}`, format.Items{"id": submenu.ID, "ml": time.Now().UnixMilli()})

				job := &queue.Job{
					ID:      jobid,
					Payload: json.RawMessage(submenu.ID),
					Delay:   100 * time.Millisecond,
				}

				s.Queue.Write(types.TaskName_REPAIR_MENU_TREE_PATH, types.RepaireTreeQueue, job)
			}

			return nil
		} else {

			if parent != nil {
				_, err = s.EntCli.SysMenu.UpdateOneID(zmenu.ID).SetLevel(parent.Level + 1).SetIsLeaf(true).Save(ctx)
			} else {
				_, err = s.EntCli.SysMenu.UpdateOneID(zmenu.ID).SetLevel(1).SetIsLeaf(true).Save(ctx)
			}

			if err != nil {
				return err
			}

		}

	}

	return nil
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
			ParentID: &parentID,
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

		menuinput := repo.ToEntCreateSysMenuInput(&sitem)

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
		mainput := repo.ToEntCreateSysMenuActionInput(actitem)
		mainput.MenuID = menuID

		anaction, err := a.MenuActionRepo.EntCli.SysMenuAction.Create().SetInput(*mainput).Save(ctx)
		if err != nil {
			return err
		}
		// 保存 resources
		for _, ritem := range actitem.Resources {
			ritem.ActionID = actitem.ID
			marinput := repo.ToEntCreateSysMenuActionResourceInput(ritem)
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

// Get 查询指定数据
func (a *Menu) View(ctx context.Context, id string, params schema.MenuQueryParam, opts ...schema.MenuQueryOptions) (*schema.Menu, error) {
	item, err := a.MenuRepo.View(ctx, id, params, opts...)
	if err != nil {
		// log.Println(" ----- =1111= err == ", err)
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	actions, err := a.QueryActions(ctx, id)
	if err != nil {
		// log.Println(" ----- =2222= err == ", err)
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
		ParentID: item.ParentID,
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
	if item.ParentID != nil && *item.ParentID == "" {
		item.ParentID = nil
	}

	err = repo.WithTx(ctx, a.MenuRepo.EntCli, func(tx *entscheme.Tx) error {

		menuinput := repo.ToEntCreateSysMenuInput(&item)
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
func (a *Menu) createActionsTx(ctx context.Context, tx *entscheme.Tx, menuID string, items schema.MenuActions) error {

	for _, actitem := range items {
		actitem.MenuID = menuID
		mainput := repo.ToEntCreateSysMenuActionInput(actitem)
		mainput.MenuID = menuID

		anaction, err := tx.SysMenuAction.Create().SetInput(*mainput).Save(ctx)
		if err != nil {
			return err
		}
		// 保存 resources
		for _, ritem := range actitem.Resources {
			ritem.ActionID = actitem.ID
			marinput := repo.ToEntCreateSysMenuActionResourceInput(ritem)
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
func (a *Menu) getParent(ctx context.Context, parentID *string) (*schema.Menu, error) {
	if parentID == nil || *parentID == "" {
		return nil, nil
	}

	pitem, err := a.MenuRepo.Get(ctx, *parentID)
	if err != nil {
		return nil, err
	} else if pitem == nil {
		return nil, errors.ErrInvalidParent
	}

	return pitem, nil
}

func (a *Menu) getParentPathNet(pitem *schema.Menu, parentID *string) string {
	if parentID == nil || *parentID == "" {
		return ""
	}
	if pitem == nil {
		return ""
	}
	return a.joinParentPath(pitem.ParentPath, pitem.ID)
}

// 获取父级路径
func (a *Menu) getParentPath(ctx context.Context, parentID *string) (string, error) {

	if parentID == nil || *parentID == "" {
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
func (a *Menu) Update(ctx context.Context, id string, item schema.Menu) (*schema.Menu, error) {

	if item.ParentID != nil && id == *item.ParentID {
		return nil, errors.ErrInvalidParent
	}

	oldItem, err := a.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if oldItem == nil {
		return nil, errors.ErrNotFound
	} else if oldItem.Name != item.Name {
		if err := a.checkName(ctx, item); err != nil {
			return nil, err
		}
	}

	item.ID = oldItem.ID

	item.CreatedAt = oldItem.CreatedAt

	if item.ParentID != nil && *item.ParentID != "" {
		nparent, err := a.getParent(ctx, item.ParentID)
		if err != nil {
			return nil, err
		}
		if nparent == nil {
			return nil, errors.ErrInvalidParent
		}

		item.Level = nparent.Level + 1
		item.ParentPath = a.getParentPathNet(nparent, item.ParentID)
		item.IsLeaf = oldItem.IsLeaf

		menuinput := repo.ToEntUpdateSysMenuInput(&item)
		err = repo.WithTx(ctx, a.MenuRepo.EntCli, func(tx *entscheme.Tx) error {

			_, err = tx.SysMenu.UpdateOneID(id).SetInput(*menuinput).Save(ctx)
			if err != nil {
				return err
			}

			tx.SysMenu.UpdateOneID(*item.ParentID).SetIsLeaf(false).Save(ctx)

			if oldItem.ParentID != nil {
				// update old item's parent leaf
				cchildren, err := tx.SysMenu.Query().Where(sysmenu.IsDelEQ(false), sysmenu.ParentIDEQ(*oldItem.ParentID)).Count(ctx)
				if err != nil {
					return err
				}

				if cchildren == 0 {
					_, err = tx.SysMenu.UpdateOneID(*oldItem.ParentID).SetIsLeaf(true).Save(ctx)
					if err != nil {
						return err
					}
				}
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
		if err != nil {
			return nil, err
		}

		jobid := format.String(`{id}-{ml}`, format.Items{"id": id, "ml": time.Now().UnixMilli()})

		job := &queue.Job{
			ID:      jobid,
			Payload: json.RawMessage(id),
			Delay:   100 * time.Millisecond,
		}

		log.Println(" ------ === jobid ==== ", jobid)
		log.Println(" ------ === jobid ==== ", jobid)

		a.Queue.Write(types.TaskName_REPAIR_MENU_TREE_PATH, types.RepaireTreeQueue, job)

	} else {

		item.Level = 1
		item.ParentID = nil
		item.ParentPath = ""
		menuinput := repo.ToEntUpdateSysMenuInput(&item)
		menuinput.ClearParentID = true

		err = repo.WithTx(ctx, a.MenuRepo.EntCli, func(tx *entscheme.Tx) error {

			_, err = tx.SysMenu.UpdateOneID(id).SetInput(*menuinput).Save(ctx)
			if err != nil {
				return err
			}
			log.Println(" ------ ===== oldItem.ParentID ", oldItem.ParentID)

			if oldItem.ParentID != nil {
				// update old item's parent leaf
				cchildren, err := tx.SysMenu.Query().Where(sysmenu.IsDelEQ(false), sysmenu.ParentIDEQ(*oldItem.ParentID)).Count(ctx)

				if err != nil {
					return err
				}
				if cchildren == 0 {
					_, err = tx.SysMenu.UpdateOneID(*oldItem.ParentID).SetIsLeaf(true).Save(ctx)
					if err != nil {
						return err
					}
				}
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
		if err != nil {
			return nil, err
		}

		jobid := format.String(`{id}-{ml}`, format.Items{"id": id, "ml": time.Now().UnixMilli()})

		job := &queue.Job{
			ID:      jobid,
			Payload: json.RawMessage(id),
			Delay:   100 * time.Millisecond,
		}

		a.Queue.Write(types.TaskName_REPAIR_MENU_TREE_PATH, types.RepaireTreeQueue, job)

	}

	nitem, err := a.Get(ctx, id)
	return nitem, err
}

// 更新动作数据
func (a *Menu) updateActions(ctx context.Context, tx *entscheme.Tx, menuID string, oldItems,
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

			marinput := repo.ToEntCreateSysMenuActionResourceInput(aritem)
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
func (a *Menu) updateChildParentPath(ctx context.Context, tx *entscheme.Tx, oldItem, newItem schema.Menu) error {
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

	ma_result, err := a.MenuActionRepo.Query(ctx, schema.MenuActionQueryParam{
		MenuID: id,
	})
	if err != nil {
		if !entscheme.IsNotFound(err) {
			return err
		}
	}

	err = repo.WithTx(ctx, a.MenuRepo.EntCli, func(tx *entscheme.Tx) error {

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
		// err = tx.SysMenu.DeleteOneID(id).Exec(ctx)
		_, err = tx.SysMenu.Update().Where(sysmenu.IDEQ(id)).SetDeletedAt(time.Now()).SetIsDel(true).Save(ctx)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	jobid := format.String(`{id}-{ml}`, format.Items{"id": id, "ml": time.Now().UnixMilli()})

	job := &queue.Job{
		ID:      jobid,
		Payload: json.RawMessage(id),
		Delay:   100 * time.Millisecond,
	}

	a.Queue.Write(types.TaskName_REPAIR_MENU_TREE_PATH, types.RepaireTreeQueue, job)

	return err

}

// UpdateActive 更新状态
func (a *Menu) UpdateActive(ctx context.Context, id string, isActive bool) error {
	oldItem, err := a.MenuRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.MenuRepo.UpdateActive(ctx, id, isActive)
}
