package repo

import (
	"context"
	"time"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/ent"
	"github.com/heromicro/omgind/internal/gen/ent/orgdepartment"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/helper/structure"

	"github.com/google/wire"
)

// OrgDepartmentSet 注入OrgDepartment
var OrgDepartmentSet = wire.NewSet(wire.Struct(new(OrgDepartment), "*"))

// OrgDepartment 部门管理存储
type OrgDepartment struct {
	EntCli *ent.Client
}

// ToSchemaOrgDepartment 转换为
func (a *OrgDepartment) ToSchemaOrgDepartment(et *ent.OrgDepartment) *schema.OrgDepartment {
	item := new(schema.OrgDepartment)
	structure.Copy(et, item)
	return item
}

func (a *OrgDepartment) ToSchemaOrgDepartments(ets ent.OrgDepartments) []*schema.OrgDepartment {
	list := make([]*schema.OrgDepartment, len(ets))
	for i, item := range ets {
		list[i] = a.ToSchemaOrgDepartment(item)
	}
	return list
}

func (a *OrgDepartment) ToEntCreateOrgDepartmentInput(sch *schema.OrgDepartment) *ent.CreateOrgDepartmentInput {
	createinput := new(ent.CreateOrgDepartmentInput)
	structure.Copy(sch, &createinput)

	return createinput
}

func (a *OrgDepartment) ToEntUpdateOrgDepartmentInput(sch *schema.OrgDepartment) *ent.UpdateOrgDepartmentInput {
	updateinput := new(ent.UpdateOrgDepartmentInput)
	structure.Copy(sch, &updateinput)

	return updateinput
}

func (a *OrgDepartment) getQueryOption(opts ...schema.OrgDepartmentQueryOptions) schema.OrgDepartmentQueryOptions {
	var opt schema.OrgDepartmentQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *OrgDepartment) Query(ctx context.Context, params schema.OrgDepartmentQueryParam, opts ...schema.OrgDepartmentQueryOptions) (*schema.OrgDepartmentQueryResult, error) {
	opt := a.getQueryOption(opts...)

	query := a.EntCli.OrgDepartment.Query()

	query = query.Where(orgdepartment.DeletedAtIsNil())
	// TODO: 查询条件

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// get total
	pr := &schema.PaginationResult{Total: count}
	if params.PaginationParam.OnlyCount {
		return &schema.OrgDepartmentQueryResult{PageResult: pr}, nil
	}

	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("id", schema.OrderByDESC))
	query = query.Order(ParseOrder(opt.OrderFields)...)

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()
	if params.Offset() > count {
		return &schema.OrgDepartmentQueryResult{PageResult: pr}, nil
	}
	query = query.Limit(params.Limit()).Offset(params.Offset())

	list, err1 := query.All(ctx)
	if err1 != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.OrgDepartmentQueryResult{
		PageResult: pr,
		Data:       a.ToSchemaOrgDepartments(list),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *OrgDepartment) Get(ctx context.Context, id string, opts ...schema.OrgDepartmentQueryOptions) (*schema.OrgDepartment, error) {

	r_orgdepartment, err := a.EntCli.OrgDepartment.Query().Where(orgdepartment.IDEQ(id)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	return a.ToSchemaOrgDepartment(r_orgdepartment), nil
}

// View 查询指定数据
func (a *OrgDepartment) View(ctx context.Context, id string, opts ...schema.OrgDepartmentQueryOptions) (*schema.OrgDepartment, error) {

	r_orgdepartment, err := a.EntCli.OrgDepartment.Query().Where(orgdepartment.IDEQ(id)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	return a.ToSchemaOrgDepartment(r_orgdepartment), nil
}

// Create 创建数据
func (a *OrgDepartment) Create(ctx context.Context, item schema.OrgDepartment) (*schema.OrgDepartment, error) {

	iteminput := a.ToEntCreateOrgDepartmentInput(&item)
	r_orgdepartment, err := a.EntCli.OrgDepartment.Create().SetInput(*iteminput).Save(ctx)

	if err != nil {
		return nil, err
	}
	sch_orgdepartment := a.ToSchemaOrgDepartment(r_orgdepartment)
	return sch_orgdepartment, nil
}

// Update 更新数据
func (a *OrgDepartment) Update(ctx context.Context, id string, item schema.OrgDepartment) (*schema.OrgDepartment, error) {

	oitem, err := a.EntCli.OrgDepartment.Query().Where(orgdepartment.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	iteminput := a.ToEntUpdateOrgDepartmentInput(&item)

	r_orgdepartment, err := oitem.Update().SetInput(*iteminput).Save(ctx)
	sch_orgdepartment := a.ToSchemaOrgDepartment(r_orgdepartment)

	return sch_orgdepartment, nil
}

// Delete 删除数据
func (a *OrgDepartment) Delete(ctx context.Context, id string) error {

	r_orgdepartment, err := a.EntCli.OrgDepartment.Query().Where(orgdepartment.IDEQ(id)).Only(ctx)

	if err != nil {
		return err
	}

	_, err1 := r_orgdepartment.Update().SetDeletedAt(time.Now()).SetIsDel(true).Save(ctx)

	return errors.WithStack(err1)
}

// UpdateActive 更新状态
func (a *OrgDepartment) UpdateActive(ctx context.Context, id string, active bool) error {

	_, err1 := a.EntCli.OrgDepartment.Update().Where(orgdepartment.IDEQ(id)).SetIsActive(active).Save(ctx)

	return errors.WithStack(err1)
}
