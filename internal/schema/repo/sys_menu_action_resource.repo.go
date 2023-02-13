package repo

import (
	"context"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/wire"
	"github.com/heromicro/omgind/internal/gen/ent/sysmenuaction"
	"github.com/heromicro/omgind/internal/gen/ent/sysmenuactionresource"
	"github.com/heromicro/omgind/pkg/helper/structure"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/ent"
	"github.com/heromicro/omgind/pkg/errors"
)

// MenuActionResourceSet 注入MenuActionResource
var MenuActionResourceSet = wire.NewSet(wire.Struct(new(MenuActionResource), "*"))

// MenuActionResource 菜单动作关联资源存储
type MenuActionResource struct {
	EntCli *ent.Client
	//TxCli *ent.Tx
}


func (a *MenuActionResource) toSchemaSysMenuActionResource(ma *ent.SysMenuActionResource) *schema.MenuActionResource {
	item := new(schema.MenuActionResource)
	structure.Copy(ma, item)
	return item
}

func (a *MenuActionResource) toSchemaSysMenuActionResources(mas ent.SysMenuActionResources) []*schema.
	MenuActionResource {
	list := make([]*schema.MenuActionResource, len(mas))
	for i, item := range mas {
		list[i] = a.toSchemaSysMenuActionResource(item)
	}
	return list
}


func (a *MenuActionResource) ToEntCreateSysMenuActionResourceInput(ma *schema.MenuActionResource) *ent.CreateSysMenuActionResourceInput {
	createinput := new(ent.CreateSysMenuActionResourceInput)
	structure.Copy(ma, &createinput)

	return createinput
}

func (a *MenuActionResource) ToEntUpdateSysMenuActionResourceInput(ma *schema.MenuActionResource) *ent.UpdateSysMenuActionResourceInput {
	updateinput := new(ent.UpdateSysMenuActionResourceInput)
	structure.Copy(ma, &updateinput)

	return updateinput
}


func (a *MenuActionResource) getQueryOption(opts ...schema.MenuActionResourceQueryOptions) schema.MenuActionResourceQueryOptions {
	var opt schema.MenuActionResourceQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *MenuActionResource) Query(ctx context.Context, params schema.MenuActionResourceQueryParam, opts ...schema.MenuActionResourceQueryOptions) (*schema.MenuActionResourceQueryResult, error) {
	opt := a.getQueryOption(opts...)

	query := a.EntCli.SysMenuActionResource.Query().Where(sysmenuactionresource.DeletedAtIsNil())

	if v := params.MenuID; v != "" {
		// TODO::
		query = query.Where(func(s *sql.Selector) {
			sma_t := sql.Table(sysmenuaction.Table)
			s.Where(sql.In(
				sysmenuactionresource.FieldActionID,
				sql.Select(sysmenuaction.FieldID).From(sma_t).Where(sql.EQ(sysmenuaction.FieldMenuID, v)),
				))
		})
		//subQuery := entity.GetMenuActionDB(ctx, a.DB).
		//	Where("menu_id=?", v).
		//	Select("id").SubQuery()
		//db = db.Where("action_id IN ?", subQuery)
	}
	if v := params.MenuIDs; len(v) > 0 {
		// TODO::
		query = query.Where(func(s *sql.Selector) {
			sma_t := sql.Table(sysmenuaction.Table)
			s.Where(sql.In(
				sysmenuactionresource.FieldActionID,
				sql.Select(sysmenuaction.FieldID).From(sma_t).Where(sql.In(sysmenuaction.FieldMenuID, strings.Join(v,
					","))),
				))
		})
		//subQuery := entity.GetMenuActionDB(ctx, a.DB).Where("menu_id IN (?)", v).Select("id").SubQuery()
		//db = db.Where("action_id IN ?", subQuery)
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// get total
	pr := &schema.PaginationResult{Total: count}
	if params.PaginationParam.OnlyCount {
		return &schema.MenuActionResourceQueryResult{PageResult: pr}, nil
	}

	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("id", schema.OrderByASC))
	query = query.Order(ParseOrder(opt.OrderFields)...)

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()
	if params.Offset() > count {
		return &schema.MenuActionResourceQueryResult{PageResult: pr}, nil
	}
	query = query.Limit(params.Limit()).Offset(params.Offset())

	list, err := query.All(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	rlist := ent.SysMenuActionResources(list)

	qr := &schema.MenuActionResourceQueryResult{
		PageResult: pr,
		Data:       a.toSchemaSysMenuActionResources(rlist),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *MenuActionResource) Get(ctx context.Context, id string, opts ...schema.MenuActionResourceQueryOptions) (*schema.MenuActionResource, error) {

	sys_mar, err := a.EntCli.SysMenuActionResource.Query().Where(sysmenuactionresource.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return a.toSchemaSysMenuActionResource(sys_mar), nil
}

// Create 创建数据
func (a *MenuActionResource) Create(ctx context.Context, item schema.MenuActionResource) (*schema.MenuActionResource, error) {

	iteminput := a.ToEntCreateSysMenuActionResourceInput(&item)
	sys_mar, err := a.EntCli.SysMenuActionResource.Create().SetInput(*iteminput).Save(ctx)
	if err != nil {
		return nil, err
	}
	sch_mar := a.toSchemaSysMenuActionResource(sys_mar)
	return sch_mar, nil
}

// Update 更新数据
func (a *MenuActionResource) Update(ctx context.Context, id string, item schema.MenuActionResource) (*schema.MenuActionResource, error) {
	oitem, err := a.EntCli.SysMenuActionResource.Query().Where(sysmenuactionresource.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	iteminput := a.ToEntUpdateSysMenuActionResourceInput(&item)
	sys_mar, err := oitem.Update().SetInput(*iteminput).Save(ctx)
	if err != nil {
		return nil, err
	}
	sch_mar := a.toSchemaSysMenuActionResource(sys_mar)

	return sch_mar, nil
}

// Delete 删除数据
func (a *MenuActionResource) Delete(ctx context.Context, id string) error {
	_, err := a.EntCli.SysMenuActionResource.UpdateOneID(id).SetDeletedAt(time.Now()).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

// DeleteByActionID 根据动作ID删除数据
func (a *MenuActionResource) DeleteByActionID(ctx context.Context, actionID string) error {
	_, err := a.EntCli.SysMenuActionResource.Update().Where(sysmenuactionresource.ActionIDEQ(actionID)).SetDeletedAt(time.Now()).Save(ctx)
	return errors.WithStack(err)
}

// DeleteByMenuID 根据菜单ID删除数据
func (a *MenuActionResource) DeleteByMenuID(ctx context.Context, menuID string) error {
	_, err := a.EntCli.SysMenuActionResource.Update().Where(func(s *sql.Selector) {
		sma_t := sql.Table(sysmenuaction.Table)
		s.Where(
			sql.In(
				sysmenuactionresource.FieldActionID,
				sql.Select(sysmenuaction.FieldID).
				From(sma_t).
				Where(sql.EQ(sysmenuaction.FieldMenuID, menuID)),
			),
		)
	}).SetIsDel(true).SetDeletedAt(time.Now()).Save(ctx)

	return errors.WithStack(err)
}
