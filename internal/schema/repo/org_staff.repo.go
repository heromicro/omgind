package repo

import (
	"context"
	"time"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/ent"
	"github.com/heromicro/omgind/internal/gen/ent/orgorgan"
	"github.com/heromicro/omgind/internal/gen/ent/orgstaff"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/helper/structure"

	"github.com/google/wire"
)

// OrgStaffSet 注入OrgStaff
var OrgStaffSet = wire.NewSet(wire.Struct(new(OrgStaff), "*"))

// OrgStaff 员工存储
type OrgStaff struct {
	EntCli *ent.Client
}

// ToSchemaOrgStaff 转换为
func ToSchemaOrgStaff(et *ent.OrgStaff) *schema.OrgStaff {
	item := new(schema.OrgStaff)
	structure.Copy(et, item)
	if et.Edges.Organ != nil {
		item.Org = ToSchemaOrgOrganShow(et.Edges.Organ)
	}

	if et.Edges.IdenAddr != nil {
		item.IdenAddr = ToSchemaSysAddress(et.Edges.IdenAddr)
	}
	if et.Edges.ResiAddr != nil {
		item.ResiAddr = ToSchemaSysAddress(et.Edges.ResiAddr)
	}

	return item
}

func ToSchemaOrgStaffs(ets ent.OrgStaffs) []*schema.OrgStaff {
	list := make([]*schema.OrgStaff, len(ets))
	for i, item := range ets {
		list[i] = ToSchemaOrgStaff(item)
	}
	return list
}

func (a *OrgStaff) ToEntCreateOrgStaffInput(sch *schema.OrgStaff) *ent.CreateOrgStaffInput {
	createinput := new(ent.CreateOrgStaffInput)
	structure.Copy(sch, &createinput)

	return createinput
}

func (a *OrgStaff) ToEntUpdateOrgStaffInput(sch *schema.OrgStaff) *ent.UpdateOrgStaffInput {
	updateinput := new(ent.UpdateOrgStaffInput)
	structure.Copy(sch, &updateinput)

	return updateinput
}

func (a *OrgStaff) getQueryOption(opts ...schema.OrgStaffQueryOptions) schema.OrgStaffQueryOptions {
	var opt schema.OrgStaffQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *OrgStaff) Query(ctx context.Context, params schema.OrgStaffQueryParam, opts ...schema.OrgStaffQueryOptions) (*schema.OrgStaffQueryResult, error) {
	opt := a.getQueryOption(opts...)

	query := a.EntCli.OrgStaff.Query().WithOrgan(func(ooq *ent.OrgOrganQuery) {
		ooq.Select(orgorgan.FieldID, orgorgan.FieldName, orgorgan.FieldSname)
	})

	query = query.Where(orgstaff.DeletedAtIsNil())
	// TODO: 查询条件

	if v := params.FirstName; v != "" {
		query = query.Where(orgstaff.FirstNameContains(v))
	}

	if v := params.IsActive; v != nil {
		query = query.Where(orgstaff.IsActiveEQ(*v))
	}

	if v := params.Gender; v != "" {
		query = query.Where(orgstaff.GenderEQ(orgstaff.Gender(v)))
	}

	if v := params.OrgID; v != "" {
		query = query.Where(orgstaff.OrgIDEQ(v))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// get total
	pr := &schema.PaginationResult{Total: count}
	if params.PaginationParam.OnlyCount {
		return &schema.OrgStaffQueryResult{PageResult: pr}, nil
	}

	if v := params.IsActive_Order; v != "" {
		of := MakeUpOrderField(orgstaff.FieldIsActive, v)
		opt.OrderFields = append(opt.OrderFields, of)
	}

	if v := params.Sort_Order; v != "" {
		of := MakeUpOrderField(orgstaff.FieldSort, v)
		opt.OrderFields = append(opt.OrderFields, of)
	}

	if v := params.CreatedAt_Order; v != "" {
		of := MakeUpOrderField(orgstaff.FieldBirthDate, v)
		opt.OrderFields = append(opt.OrderFields, of)
	}

	if v := params.EntryDate_Order; v != "" {
		of := MakeUpOrderField(orgstaff.FieldEntryDate, v)
		opt.OrderFields = append(opt.OrderFields, of)
	}

	if v := params.RegularDate_Order; v != "" {
		of := MakeUpOrderField(orgstaff.FieldRegularDate, v)
		opt.OrderFields = append(opt.OrderFields, of)
	}

	if v := params.ResignDate_Order; v != "" {
		of := MakeUpOrderField(orgstaff.FieldRegularDate, v)
		opt.OrderFields = append(opt.OrderFields, of)
	}

	if len(opt.OrderFields) == 0 {
		opt.OrderFields = append(opt.OrderFields, schema.NewOrderField(orgstaff.FieldID, schema.OrderByDESC))
	}

	query = query.Order(ParseOrder(opt.OrderFields)...)

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()
	if params.Offset() > count {
		return &schema.OrgStaffQueryResult{PageResult: pr}, nil
	}
	query = query.Limit(params.Limit()).Offset(params.Offset())

	list, err1 := query.All(ctx)
	if err1 != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.OrgStaffQueryResult{
		PageResult: pr,
		Data:       ToSchemaOrgStaffs(list),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *OrgStaff) Get(ctx context.Context, id string, opts ...schema.OrgStaffQueryOptions) (*schema.OrgStaff, error) {

	query := a.EntCli.OrgStaff.Query().WithOrgan(func(ooq *ent.OrgOrganQuery) {
		ooq.Select(orgorgan.FieldID, orgorgan.FieldName, orgorgan.FieldSname)
	})
	r_orgstaff, err := query.Where(orgstaff.IDEQ(id)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	return ToSchemaOrgStaff(r_orgstaff), nil
}

// View 查询指定数据
func (a *OrgStaff) View(ctx context.Context, id string, opts ...schema.OrgStaffQueryOptions) (*schema.OrgStaff, error) {

	query := a.EntCli.OrgStaff.Query().WithOrgan(func(ooq *ent.OrgOrganQuery) {
		ooq.Select(orgorgan.FieldID, orgorgan.FieldName, orgorgan.FieldSname, orgorgan.FieldIdenNo).WithHaddr()
	})
	r_orgstaff, err := query.Where(orgstaff.IDEQ(id)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	return ToSchemaOrgStaff(r_orgstaff), nil
}

// Create 创建数据
func (a *OrgStaff) Create(ctx context.Context, item schema.OrgStaff) (*schema.OrgStaff, error) {

	// TODO: check org_id

	iteminput := a.ToEntCreateOrgStaffInput(&item)
	r_orgstaff, err := a.EntCli.OrgStaff.Create().SetInput(*iteminput).Save(ctx)

	if err != nil {
		return nil, err
	}
	sch_orgstaff := ToSchemaOrgStaff(r_orgstaff)
	return sch_orgstaff, nil
}

// Update 更新数据
func (a *OrgStaff) Update(ctx context.Context, id string, item schema.OrgStaff) (*schema.OrgStaff, error) {

	// TODO: check org_id

	oitem, err := a.EntCli.OrgStaff.Query().Where(orgstaff.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	iteminput := a.ToEntUpdateOrgStaffInput(&item)

	r_orgstaff, err := oitem.Update().SetInput(*iteminput).Save(ctx)
	sch_orgstaff := ToSchemaOrgStaff(r_orgstaff)

	return sch_orgstaff, nil
}

// Delete 删除数据
func (a *OrgStaff) Delete(ctx context.Context, id string) error {

	r_orgstaff, err := a.EntCli.OrgStaff.Query().Where(orgstaff.IDEQ(id)).Only(ctx)

	if err != nil {
		return err
	}

	_, err1 := r_orgstaff.Update().SetDeletedAt(time.Now()).SetIsDel(true).Save(ctx)

	return errors.WithStack(err1)
}

// UpdateActive 更新状态
func (a *OrgStaff) UpdateActive(ctx context.Context, id string, active bool) error {

	_, err1 := a.EntCli.OrgStaff.Update().Where(orgstaff.IDEQ(id)).SetIsActive(active).Save(ctx)

	return errors.WithStack(err1)
}
