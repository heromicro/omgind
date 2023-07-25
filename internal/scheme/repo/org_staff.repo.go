package repo

import (
	"context"
	"time"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/mainent"
	"github.com/heromicro/omgind/internal/gen/mainent/orgdept"
	"github.com/heromicro/omgind/internal/gen/mainent/orgorgan"
	"github.com/heromicro/omgind/internal/gen/mainent/orgposition"
	"github.com/heromicro/omgind/internal/gen/mainent/orgstaff"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/helper/structure"

	"github.com/google/wire"
)

// OrgStaffSet 注入OrgStaff
var OrgStaffSet = wire.NewSet(wire.Struct(new(OrgStaff), "*"))

// OrgStaff 员工存储
type OrgStaff struct {
	EntCli *mainent.Client
}

// ToSchemaOrgStaff 转换为
func ToSchemaOrgStaff(et *mainent.OrgStaff) *schema.OrgStaff {
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

	if et.Edges.Dept != nil {
		item.Dept = ToSchemaOrgDept(et.Edges.Dept)
	}

	if et.Edges.Posi != nil {
		item.Posi = ToSchemaOrgPosition(et.Edges.Posi)
	}

	return item
}

func ToSchemaOrgStaffs(ets mainent.OrgStaffs) []*schema.OrgStaff {
	list := make([]*schema.OrgStaff, len(ets))
	for i, item := range ets {
		list[i] = ToSchemaOrgStaff(item)
	}
	return list
}

func ToEntCreateOrgStaffInput(sch *schema.OrgStaff) *mainent.CreateOrgStaffInput {
	createinput := new(mainent.CreateOrgStaffInput)
	structure.Copy(sch, &createinput)

	return createinput
}

func ToEntUpdateOrgStaffInput(sch *schema.OrgStaff) *mainent.UpdateOrgStaffInput {
	updateinput := new(mainent.UpdateOrgStaffInput)
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
	// opt := a.getQueryOption(opts...)

	query := a.EntCli.OrgStaff.Query().WithOrgan(func(ooq *mainent.OrgOrganQuery) {
		ooq.Select(orgorgan.FieldID, orgorgan.FieldName, orgorgan.FieldSname)
	}).WithIdenAddr().WithResiAddr()

	query = query.WithDept(func(odq *mainent.OrgDeptQuery) {
		odq.Select(orgdept.FieldID, orgdept.FieldName, orgdept.FieldMergeName, orgdept.FieldTreePath)
	})

	query = query.WithPosi(func(opq *mainent.OrgPositionQuery) {
		opq.Select(orgposition.FieldID, orgposition.FieldName)
	})

	query = query.Where(orgstaff.DeletedAtIsNil())
	// TODO: 查询条件

	if v := params.FirstName; v != "" {
		query = query.Where(orgstaff.FirstNameContains(v))
	}

	if v := params.IsActive; v != nil {
		query = query.Where(orgstaff.IsActiveEQ(*v))
	}

	if v := params.Gender; v != nil {
		query = query.Where(orgstaff.GenderEQ(*v))
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
		query = query.Order(orgstaff.ByIsActive(OrderDirection(v)))
	}

	if v := params.Sort_Order; v != "" {
		query = query.Order(orgstaff.BySort(OrderDirection(v)))
	}

	if v := params.CreatedAt_Order; v != "" {
		query = query.Order(orgstaff.ByCreatedAt(OrderDirection(v)))
	}

	if v := params.BirthDate_Order; v != "" {
		query = query.Order(orgstaff.ByBirthDate(OrderDirection(v)))
	}

	if v := params.EntryDate_Order; v != "" {
		query = query.Order(orgstaff.ByEntryDate(OrderDirection(v)))
	}

	if v := params.RegularDate_Order; v != "" {
		query = query.Order(orgstaff.ByRegularDate(OrderDirection(v)))
	}

	if v := params.ResignDate_Order; v != "" {
		query = query.Order(orgstaff.ByResignDate(OrderDirection(v)))
	}

	query = query.Order(orgstaff.ByID(OrderDirection("desc")))

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

	query := a.EntCli.OrgStaff.Query().WithOrgan(func(ooq *mainent.OrgOrganQuery) {
		ooq.Select(orgorgan.FieldID, orgorgan.FieldName, orgorgan.FieldSname)
	}).WithIdenAddr().WithResiAddr()

	query = query.WithDept(func(odq *mainent.OrgDeptQuery) {
		odq.Select(orgdept.FieldID, orgdept.FieldName, orgdept.FieldMergeName, orgdept.FieldTreePath)
	})

	query = query.WithPosi(func(opq *mainent.OrgPositionQuery) {
		opq.Select(orgposition.FieldID, orgposition.FieldName)
	})

	r_orgstaff, err := query.Where(orgstaff.IDEQ(id)).Only(ctx)
	if err != nil {
		if mainent.IsNotFound(err) {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	return ToSchemaOrgStaff(r_orgstaff), nil
}

// View 查询指定数据
func (a *OrgStaff) View(ctx context.Context, id string, opts ...schema.OrgStaffQueryOptions) (*schema.OrgStaff, error) {

	query := a.EntCli.OrgStaff.Query().WithOrgan(func(ooq *mainent.OrgOrganQuery) {
		ooq.Select(orgorgan.FieldID, orgorgan.FieldName, orgorgan.FieldSname, orgorgan.FieldIdenNo).WithHaddr()
	}).WithIdenAddr().WithResiAddr()

	query = query.WithDept(func(odq *mainent.OrgDeptQuery) {
		odq.Select(orgdept.FieldID, orgdept.FieldName, orgdept.FieldMergeName, orgdept.FieldTreePath)
	})

	query = query.WithPosi(func(opq *mainent.OrgPositionQuery) {
		opq.Select(orgposition.FieldID, orgposition.FieldName)
	})

	r_orgstaff, err := query.Where(orgstaff.IDEQ(id)).Only(ctx)
	if err != nil {
		if mainent.IsNotFound(err) {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	return ToSchemaOrgStaff(r_orgstaff), nil
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
