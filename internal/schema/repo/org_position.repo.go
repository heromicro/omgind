package repo

import (
	"context"
	"time"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/ent"
	"github.com/heromicro/omgind/internal/gen/ent/orgposition"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/helper/structure"

	"github.com/google/wire"
)

// OrgPositionSet 注入OrgPosition
var OrgPositionSet = wire.NewSet(wire.Struct(new(OrgPosition), "*"))

// OrgPosition 职位管理存储
type OrgPosition struct {
	EntCli *ent.Client
}

// ToSchemaOrgPosition 转换为
func (a *OrgPosition) ToSchemaOrgPosition(et *ent.OrgPosition) *schema.OrgPosition {
	item := new(schema.OrgPosition)
	structure.Copy(et, item)
	return item
}

func (a *OrgPosition) ToSchemaOrgPositions(ets ent.OrgPositions) []*schema.OrgPosition {
	list := make([]*schema.OrgPosition, len(ets))
	for i, item := range ets {
		list[i] = a.ToSchemaOrgPosition(item)
	}
	return list
}

func (a *OrgPosition) ToEntCreateOrgPositionInput(sch *schema.OrgPosition) *ent.CreateOrgPositionInput {
	createinput := new(ent.CreateOrgPositionInput)
	structure.Copy(sch, &createinput)

	return createinput
}

func (a *OrgPosition) ToEntUpdateOrgPositionInput(sch *schema.OrgPosition) *ent.UpdateOrgPositionInput {
	updateinput := new(ent.UpdateOrgPositionInput)
	structure.Copy(sch, &updateinput)

	return updateinput
}

func (a *OrgPosition) getQueryOption(opts ...schema.OrgPositionQueryOptions) schema.OrgPositionQueryOptions {
	var opt schema.OrgPositionQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *OrgPosition) Query(ctx context.Context, params schema.OrgPositionQueryParam, opts ...schema.OrgPositionQueryOptions) (*schema.OrgPositionQueryResult, error) {
	opt := a.getQueryOption(opts...)

	query := a.EntCli.OrgPosition.Query()

	query = query.Where(orgposition.DeletedAtIsNil())
	// TODO: 查询条件

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// get total
	pr := &schema.PaginationResult{Total: count}
	if params.PaginationParam.OnlyCount {
		return &schema.OrgPositionQueryResult{PageResult: pr}, nil
	}

	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("id", schema.OrderByDESC))
	query = query.Order(ParseOrder(opt.OrderFields)...)

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()
	if params.Offset() > count {
		return &schema.OrgPositionQueryResult{PageResult: pr}, nil
	}
	query = query.Limit(params.Limit()).Offset(params.Offset())

	list, err1 := query.All(ctx)
	if err1 != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.OrgPositionQueryResult{
		PageResult: pr,
		Data:       a.ToSchemaOrgPositions(list),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *OrgPosition) Get(ctx context.Context, id string, opts ...schema.OrgPositionQueryOptions) (*schema.OrgPosition, error) {

	r_orgposition, err := a.EntCli.OrgPosition.Query().Where(orgposition.IDEQ(id)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	return a.ToSchemaOrgPosition(r_orgposition), nil
}

// View 查询指定数据
func (a *OrgPosition) View(ctx context.Context, id string, opts ...schema.OrgPositionQueryOptions) (*schema.OrgPosition, error) {

	r_orgposition, err := a.EntCli.OrgPosition.Query().Where(orgposition.IDEQ(id)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	return a.ToSchemaOrgPosition(r_orgposition), nil
}

// Create 创建数据
func (a *OrgPosition) Create(ctx context.Context, item schema.OrgPosition) (*schema.OrgPosition, error) {

	iteminput := a.ToEntCreateOrgPositionInput(&item)
	r_orgposition, err := a.EntCli.OrgPosition.Create().SetInput(*iteminput).Save(ctx)

	if err != nil {
		return nil, err
	}
	sch_orgposition := a.ToSchemaOrgPosition(r_orgposition)
	return sch_orgposition, nil
}

// Update 更新数据
func (a *OrgPosition) Update(ctx context.Context, id string, item schema.OrgPosition) (*schema.OrgPosition, error) {

	oitem, err := a.EntCli.OrgPosition.Query().Where(orgposition.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	iteminput := a.ToEntUpdateOrgPositionInput(&item)

	r_orgposition, err := oitem.Update().SetInput(*iteminput).Save(ctx)
	sch_orgposition := a.ToSchemaOrgPosition(r_orgposition)

	return sch_orgposition, nil
}

// Delete 删除数据
func (a *OrgPosition) Delete(ctx context.Context, id string) error {

	r_orgposition, err := a.EntCli.OrgPosition.Query().Where(orgposition.IDEQ(id)).Only(ctx)

	if err != nil {
		return err
	}

	_, err1 := r_orgposition.Update().SetDeletedAt(time.Now()).SetIsDel(true).Save(ctx)

	return errors.WithStack(err1)
}

// UpdateActive 更新状态
func (a *OrgPosition) UpdateActive(ctx context.Context, id string, active bool) error {

	_, err1 := a.EntCli.OrgPosition.Update().Where(orgposition.IDEQ(id)).SetIsActive(active).Save(ctx)

	return errors.WithStack(err1)
}
