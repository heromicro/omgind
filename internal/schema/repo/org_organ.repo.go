package repo

import (
	"context"
	"time"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/ent"
	"github.com/heromicro/omgind/internal/gen/ent/orgorgan"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/helper/structure"

	"github.com/google/wire"
)

// OrgOrganSet 注入OrgOrgan
var OrgOrganSet = wire.NewSet(wire.Struct(new(OrgOrgan), "*"))

// OrgOrgan 组织管理存储
type OrgOrgan struct {
	EntCli *ent.Client
}

// ToSchemaOrgOrgan 转换为
func (a *OrgOrgan) ToSchemaOrgOrgan(et *ent.OrgOrgan) *schema.OrgOrgan {
	item := new(schema.OrgOrgan)
	structure.Copy(et, item)
	return item
}

func (a *OrgOrgan) ToSchemaOrgOrgans(ets ent.OrgOrgans) []*schema.OrgOrgan {
	list := make([]*schema.OrgOrgan, len(ets))
	for i, item := range ets {
		list[i] = a.ToSchemaOrgOrgan(item)
	}
	return list
}

func (a *OrgOrgan) ToEntCreateOrgOrganInput(sch *schema.OrgOrgan) *ent.CreateOrgOrganInput {
	createinput := new(ent.CreateOrgOrganInput)
	structure.Copy(sch, &createinput)

	return createinput
}

func (a *OrgOrgan) ToEntUpdateOrgOrganInput(sch *schema.OrgOrgan) *ent.UpdateOrgOrganInput {
	updateinput := new(ent.UpdateOrgOrganInput)
	structure.Copy(sch, &updateinput)

	return updateinput
}

func (a *OrgOrgan) getQueryOption(opts ...schema.OrgOrganQueryOptions) schema.OrgOrganQueryOptions {
	var opt schema.OrgOrganQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *OrgOrgan) Query(ctx context.Context, params schema.OrgOrganQueryParam, opts ...schema.OrgOrganQueryOptions) (*schema.OrgOrganQueryResult, error) {
	opt := a.getQueryOption(opts...)

	query := a.EntCli.OrgOrgan.Query()

	query = query.Where(orgorgan.DeletedAtIsNil())
	// TODO: 查询条件

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// get total
	pr := &schema.PaginationResult{Total: count}
	if params.PaginationParam.OnlyCount {
		return &schema.OrgOrganQueryResult{PageResult: pr}, nil
	}

	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("id", schema.OrderByDESC))
	query = query.Order(ParseOrder(opt.OrderFields)...)

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()
	if params.Offset() > count {
		return &schema.OrgOrganQueryResult{PageResult: pr}, nil
	}
	query = query.Limit(params.Limit()).Offset(params.Offset())

	list, err1 := query.All(ctx)
	if err1 != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.OrgOrganQueryResult{
		PageResult: pr,
		Data:       a.ToSchemaOrgOrgans(list),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *OrgOrgan) Get(ctx context.Context, id string, opts ...schema.OrgOrganQueryOptions) (*schema.OrgOrgan, error) {

	r_orgorgan, err := a.EntCli.OrgOrgan.Query().Where(orgorgan.IDEQ(id)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	return a.ToSchemaOrgOrgan(r_orgorgan), nil
}

// View 查询指定数据
func (a *OrgOrgan) View(ctx context.Context, id string, opts ...schema.OrgOrganQueryOptions) (*schema.OrgOrgan, error) {

	r_orgorgan, err := a.EntCli.OrgOrgan.Query().Where(orgorgan.IDEQ(id)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	return a.ToSchemaOrgOrgan(r_orgorgan), nil
}

// Create 创建数据
func (a *OrgOrgan) Create(ctx context.Context, item schema.OrgOrgan) (*schema.OrgOrgan, error) {

	iteminput := a.ToEntCreateOrgOrganInput(&item)
	r_orgorgan, err := a.EntCli.OrgOrgan.Create().SetInput(*iteminput).Save(ctx)

	if err != nil {
		return nil, err
	}
	sch_orgorgan := a.ToSchemaOrgOrgan(r_orgorgan)
	return sch_orgorgan, nil
}

// Update 更新数据
func (a *OrgOrgan) Update(ctx context.Context, id string, item schema.OrgOrgan) (*schema.OrgOrgan, error) {

	oitem, err := a.EntCli.OrgOrgan.Query().Where(orgorgan.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	iteminput := a.ToEntUpdateOrgOrganInput(&item)

	r_orgorgan, err := oitem.Update().SetInput(*iteminput).Save(ctx)
	sch_orgorgan := a.ToSchemaOrgOrgan(r_orgorgan)

	return sch_orgorgan, nil
}

// Delete 删除数据
func (a *OrgOrgan) Delete(ctx context.Context, id string) error {

	r_orgorgan, err := a.EntCli.OrgOrgan.Query().Where(orgorgan.IDEQ(id)).Only(ctx)

	if err != nil {
		return err
	}

	_, err1 := r_orgorgan.Update().SetDeletedAt(time.Now()).SetIsDel(true).Save(ctx)

	return errors.WithStack(err1)
}

// UpdateActive 更新状态
func (a *OrgOrgan) UpdateActive(ctx context.Context, id string, active bool) error {

	_, err1 := a.EntCli.OrgOrgan.Update().Where(orgorgan.IDEQ(id)).SetIsActive(active).Save(ctx)

	return errors.WithStack(err1)
}
