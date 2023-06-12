package repo

import (
	"context"
	"time"

	"github.com/google/wire"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/ent"
	"github.com/heromicro/omgind/internal/gen/ent/sysmenu"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/helper/structure"
)

// MenuSet 注入Menu
var MenuSet = wire.NewSet(wire.Struct(new(Menu), "*"))

// Menu 菜单存储
type Menu struct {
	EntCli *ent.Client
}

func ToSchemaSysMenu(dit *ent.SysMenu) *schema.Menu {
	item := new(schema.Menu)
	structure.Copy(dit, item)
	return item
}

func ToSchemaSysMenus(dits ent.SysMenus) []*schema.Menu {
	list := make([]*schema.Menu, len(dits))
	for i, item := range dits {
		list[i] = ToSchemaSysMenu(item)
	}
	return list
}

func ToEntCreateSysMenuInput(m *schema.Menu) *ent.CreateSysMenuInput {
	createinput := new(ent.CreateSysMenuInput)
	structure.Copy(m, &createinput)

	return createinput
}

func ToEntUpdateSysMenuInput(m *schema.Menu) *ent.UpdateSysMenuInput {
	updateinput := new(ent.UpdateSysMenuInput)
	structure.Copy(m, &updateinput)

	return updateinput
}

func (a *Menu) getQueryOption(opts ...schema.MenuQueryOptions) schema.MenuQueryOptions {
	var opt schema.MenuQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *Menu) Query(ctx context.Context, params schema.MenuQueryParam, opts ...schema.MenuQueryOptions) (*schema.MenuQueryResult, error) {
	// opt := a.getQueryOption(opts...)

	query := a.EntCli.SysMenu.Query().Where(sysmenu.DeletedAtIsNil())

	if v := params.IDs; len(v) > 0 {
		query = query.Where(sysmenu.IDIn(v...))
	}
	if v := params.Name; v != "" {
		query = query.Where(sysmenu.NameEQ(v))
	}
	if v := params.ParentID; v != nil {
		query = query.Where(sysmenu.ParentIDEQ(*v))
	}

	if v := params.PrefixParentPath; v != "" {
		query = query.Where(sysmenu.ParentPathContains(v))
	}
	if v := params.IsActive; v != nil {
		query = query.Where(sysmenu.IsActiveEQ(*v))
	}
	if v := params.IsShow; v != nil {
		query = query.Where(sysmenu.IsShowEQ(*v))
	}

	if v := params.QueryValue; v != "" {
		query = query.Where(sysmenu.Or(
			sysmenu.NameContains(v),
			sysmenu.MemoContains(v),
		))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// get total
	pr := &schema.PaginationResult{Total: count}
	if params.PaginationParam.OnlyCount {
		return &schema.MenuQueryResult{PageResult: pr}, nil
	}

	if v := params.Sort_Order; v != "" {
		query = query.Order(sysmenu.ByIsActive(OrderDirection(v)))
	}

	if v := params.IsActive_Order; v != "" {
		query = query.Order(sysmenu.ByIsActive(OrderDirection(v)))
	}

	query = query.Order(sysmenu.ByID(OrderDirection("desc")))

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()
	if params.Offset() > count {
		return &schema.MenuQueryResult{PageResult: pr}, nil
	}

	query = query.Limit(params.Limit()).Offset(params.Offset())

	list, err := query.All(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	rlist := ent.SysMenus(list)

	qr := &schema.MenuQueryResult{
		PageResult: pr,
		Data:       ToSchemaSysMenus(rlist),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *Menu) Get(ctx context.Context, id string, opts ...schema.MenuQueryOptions) (*schema.Menu, error) {

	menu, err := a.EntCli.SysMenu.Query().Where(sysmenu.IDEQ(id)).Only(ctx)

	if err != nil {
		return nil, err
	}
	return ToSchemaSysMenu(menu), nil
}

// Create 创建数据
func (a *Menu) Create(ctx context.Context, item schema.Menu) (*schema.Menu, error) {

	iteminput := ToEntCreateSysMenuInput(&item)
	iteminput.CreatedAt = nil
	iteminput.UpdatedAt = nil
	sys_menu, err := a.EntCli.SysMenu.Create().SetInput(*iteminput).Save(ctx)
	if err != nil {
		return nil, err
	}
	sch_menu := ToSchemaSysMenu(sys_menu)
	return sch_menu, nil
}

// Update 更新数据
func (a *Menu) Update(ctx context.Context, id string, item schema.Menu) (*schema.Menu, error) {

	oitem, err := a.EntCli.SysMenu.Query().Where(sysmenu.IDEQ(id)).Only(ctx)

	if err != nil {
		return nil, err
	}

	iteminput := ToEntUpdateSysMenuInput(&item)
	sys_menu, err := oitem.Update().SetInput(*iteminput).Save(ctx)
	if err != nil {
		return nil, err
	}
	sch_menu := ToSchemaSysMenu(sys_menu)
	return sch_menu, nil
}

// UpdateParentPath 更新父级路径
func (a *Menu) UpdateParentPath(ctx context.Context, id, parentPath string) error {
	_, err := a.EntCli.SysMenu.UpdateOneID(id).SetParentPath(parentPath).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Delete 删除数据
func (a *Menu) Delete(ctx context.Context, id string) error {
	sys_menu, err := a.EntCli.SysMenu.Query().Where(sysmenu.IDEQ(id)).Only(ctx)
	if err != nil {
		return err
	}
	_, err1 := sys_menu.Update().SetDeletedAt(time.Now()).SetIsDel(true).Save(ctx)
	return errors.WithStack(err1)
}

// UpdateStatus 更新状态
func (a *Menu) UpdateStatus(ctx context.Context, id string, isActive bool) error {
	_, err := a.EntCli.SysMenu.UpdateOneID(id).SetIsActive(isActive).Save(ctx)
	return errors.WithStack(err)
}
