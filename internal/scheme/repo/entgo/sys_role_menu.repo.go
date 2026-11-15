package repo

import (
	"context"
	"time"

	"github.com/google/wire"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/mainent"
	"github.com/heromicro/omgind/internal/gen/mainent/sysrolemenu"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/helper/structure"
)

// RoleMenuSet 注入RoleMenu
var RoleMenuSet = wire.NewSet(wire.Struct(new(RoleMenu), "*"))

// RoleMenu 角色菜单存储
type RoleMenu struct {
	EntCli *mainent.Client
}

func ToSchemaSysRoleMenu(role *mainent.SysRoleMenu) *schema.RoleMenu {
	item := new(schema.RoleMenu)
	structure.Copy(role, item)
	return item
}

func ToSchemaSysRoleMenus(roles mainent.SysRoleMenus) []*schema.RoleMenu {
	list := make([]*schema.RoleMenu, len(roles))
	for i, item := range roles {
		list[i] = ToSchemaSysRoleMenu(item)
	}
	return list
}

func ToEntCreateSysRoleMenuInput(schrole *schema.RoleMenu) *mainent.CreateSysRoleMenuInput {
	createinput := new(mainent.CreateSysRoleMenuInput)
	structure.Copy(schrole, &createinput)

	return createinput
}

func ToEntUpdateSysRoleMenuInput(schrole *schema.RoleMenu) *mainent.UpdateSysRoleMenuInput {
	updateinput := new(mainent.UpdateSysRoleMenuInput)
	structure.Copy(schrole, &updateinput)

	return updateinput
}

func (a *RoleMenu) getQueryOption(opts ...schema.RoleMenuQueryOptions) schema.RoleMenuQueryOptions {
	var opt schema.RoleMenuQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *RoleMenu) Query(ctx context.Context, params schema.RoleMenuQueryParam, opts ...schema.RoleMenuQueryOptions) (*schema.RoleMenuQueryResult, error) {
	// opt := a.getQueryOption(opts...)

	query := a.EntCli.SysRoleMenu.Query().Where(sysrolemenu.DeletedAtIsNil())

	if v := params.RoleID; v != "" {
		query = query.Where(sysrolemenu.RoleIDEQ(v))
	}
	if v := params.RoleIDs; len(v) > 0 {
		query = query.Where(sysrolemenu.RoleIDIn(v...))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// get total
	pr := &schema.PaginationResult{Total: count}

	if params.PaginationParam.OnlyCount {
		return &schema.RoleMenuQueryResult{PageResult: pr}, nil
	}

	query = query.Order(sysrolemenu.ByID(OrderDirection("desc")))

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()

	if params.Offset() > count {
		return &schema.RoleMenuQueryResult{PageResult: pr}, nil
	}

	query = query.Limit(params.Limit()).Offset(params.Offset())

	list, err := query.All(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	rlist := mainent.SysRoleMenus(list)

	qr := &schema.RoleMenuQueryResult{
		PageResult: pr,
		Data:       ToSchemaSysRoleMenus(rlist),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *RoleMenu) Get(ctx context.Context, id string, opts ...schema.RoleMenuQueryOptions) (*schema.RoleMenu, error) {

	rolemenu, err := a.EntCli.SysRoleMenu.Query().Where(sysrolemenu.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return ToSchemaSysRoleMenu(rolemenu), nil
}

// Create 创建数据
func (a *RoleMenu) Create(ctx context.Context, item schema.RoleMenu) (*schema.RoleMenu, error) {

	iteminput := ToEntCreateSysRoleMenuInput(&item)
	iteminput.CreatedAt = nil
	iteminput.UpdatedAt = nil

	menu, err := a.EntCli.SysRoleMenu.Create().SetInput(*iteminput).Save(ctx)
	if err != nil {
		return nil, err
	}
	sch_rm := ToSchemaSysRoleMenu(menu)
	return sch_rm, nil
}

// Update 更新数据
func (a *RoleMenu) Update(ctx context.Context, id string, item schema.RoleMenu) (*schema.RoleMenu, error) {
	oitem, err := a.EntCli.SysRoleMenu.Query().Where(sysrolemenu.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, errors.ErrNotFound
	}
	iteminput := ToEntUpdateSysRoleMenuInput(&item)
	iteminput.UpdatedAt = nil
	rolemenu, err := oitem.Update().SetInput(*iteminput).Save(ctx)
	if err != nil {
		return nil, err
	}
	sch_rm := ToSchemaSysRoleMenu(rolemenu)
	return sch_rm, nil
}

// Delete 删除数据
func (a *RoleMenu) Delete(ctx context.Context, id string) error {
	rolemenu, err := a.EntCli.SysRoleMenu.Query().Where(sysrolemenu.IDEQ(id)).Only(ctx)
	if err != nil {
		return errors.ErrNotFound
	}
	_, err1 := rolemenu.Update().SetDeletedAt(time.Now()).SetIsDel(true).Save(ctx)
	return errors.WithStack(err1)
}

// DeleteByRoleID 根据角色ID删除数据
func (a *RoleMenu) DeleteByRoleID(ctx context.Context, roleID string) error {
	_, err := a.EntCli.SysRoleMenu.Update().SetDeletedAt(time.Now()).SetIsDel(true).Where(sysrolemenu.RoleIDEQ(roleID)).Save(ctx)
	return errors.WithStack(err)
}
