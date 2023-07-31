package repo

import (
	"context"
	"time"

	"github.com/google/wire"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/mainent"
	"github.com/heromicro/omgind/internal/gen/mainent/orgorgan"
	"github.com/heromicro/omgind/internal/gen/mainent/sysaddress"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/helper/structure"
)

// OrgOrganSet 注入OrgOrgan
var OrgOrganSet = wire.NewSet(wire.Struct(new(OrgOrgan), "*"))

// OrgOrgan 组织管理存储
type OrgOrgan struct {
	EntCli *mainent.Client
}

// ToSchemaOrgOrgan 转换为
func ToSchemaOrgOrganShow(et *mainent.OrgOrgan) *schema.OrgOrganShow {
	item := new(schema.OrgOrganShow)

	structure.Copy(et, item)
	if et.Edges.Haddr != nil {
		item.Haddr = ToSchemaSysAddress(et.Edges.Haddr)
	}
	return item
}

func ToSchemaOrgOrgan(et *mainent.OrgOrgan) *schema.OrgOrgan {
	item := new(schema.OrgOrgan)

	structure.Copy(et, item)
	if et.Edges.Haddr != nil {
		item.Haddr = ToSchemaSysAddress(et.Edges.Haddr)
	}
	return item
}

func ToSchemaOrgOrgans(ets mainent.OrgOrgans) []*schema.OrgOrgan {
	list := make([]*schema.OrgOrgan, len(ets))
	for i, item := range ets {
		list[i] = ToSchemaOrgOrgan(item)
	}
	return list
}

func ToEntCreateOrgOrganInput(sch *schema.OrgOrgan) *mainent.CreateOrgOrganInput {
	createinput := new(mainent.CreateOrgOrganInput)
	structure.Copy(sch, &createinput)

	return createinput
}

func ToEntUpdateOrgOrganInput(sch *schema.OrgOrgan) *mainent.UpdateOrgOrganInput {
	updateinput := new(mainent.UpdateOrgOrganInput)
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

	// if opt.FieldsAll || (slices.Contains(opt.FieldsIncludes, orgorgan.EdgeHaddr) && !slices.Contains(opt.FieldsExcludes, orgorgan.EdgeHaddr)) {
	query = query.WithHaddr(func(saq *mainent.SysAddressQuery) {
		saq.Select(sysaddress.FieldID, sysaddress.FieldCountry, sysaddress.FieldCountryID, sysaddress.FieldProvince, sysaddress.FieldProvinceID, sysaddress.FieldCity, sysaddress.FieldCityID, sysaddress.FieldCounty, sysaddress.FieldCountyID, sysaddress.FieldAreaCode, sysaddress.FieldMobile, sysaddress.FieldFirstName, sysaddress.FieldLastName, sysaddress.FieldDaddr, sysaddress.FieldZipCode)
	})
	// }

	query = query.Where(orgorgan.DeletedAtIsNil())

	// TODO: 查询条件
	if v := params.Name; v != "" {
		query = query.Where(orgorgan.NameContains(v))
	}

	if v := params.Code; v != nil && *v != "" {
		query = query.Where(orgorgan.CodeContains(*v))
	}

	if v := params.IsActive; v != nil {
		query = query.Where(orgorgan.IsActiveEQ(*v))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// get total
	pr := &schema.PaginationResult{Total: count}
	if params.PaginationParam.OnlyCount {
		return &schema.OrgOrganQueryResult{PageResult: pr}, nil
	}

	if v := params.IsActive_Order; v != "" {
		query = query.Order(orgorgan.ByIsActive(OrderDirection(v)))
	}

	if v := params.Sort_Order; v != "" {
		query = query.Order(orgorgan.BySort(OrderDirection(v)))
	}
	query = query.Order(orgorgan.ByID(OrderDirection("desc")))

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()
	if params.Offset() > count {
		return &schema.OrgOrganQueryResult{PageResult: pr}, nil
	}
	query = query.Limit(params.Limit()).Offset(params.Offset())

	var query_select *mainent.OrgOrganSelect
	if len(opt.FieldsIncludes) > 0 {
		query_select = query.Select(opt.FieldsIncludes...)
	}

	var list []*mainent.OrgOrgan
	if query_select != nil {
		list, err = query_select.All(ctx)
	} else {
		list, err = query.All(ctx)
	}

	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.OrgOrganQueryResult{
		PageResult: pr,
		Data:       ToSchemaOrgOrgans(list),
	}

	return qr, nil
}

// QuerySelectPage 查询数据
func (a *OrgOrgan) QuerySelectPage(ctx context.Context, params schema.OrgOrganQueryParam, opts ...schema.OrgOrganQueryOptions) (*schema.OrgOrganQueryResult, error) {

	opt := a.getQueryOption(opts...)

	query := a.EntCli.OrgOrgan.Query().Where(orgorgan.IsDelEQ(false))

	query = query.Where(orgorgan.DeletedAtIsNil())

	if v := params.QueryValue; v != "" {
		query = query.Where(orgorgan.Or(orgorgan.NameContains(v), orgorgan.CodeContains(v)))
	}

	if v := params.Name; v != "" {
		query = query.Where(orgorgan.NameContains(v))
	}

	if v := params.Code; v != nil && *v != "" {
		query = query.Where(orgorgan.CodeContains(*v))
	}

	if v := params.IsActive; v != nil {
		query = query.Where(orgorgan.IsActiveEQ(*v))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// get total
	pr := &schema.PaginationResult{Total: count}
	if params.PaginationParam.OnlyCount {
		return &schema.OrgOrganQueryResult{PageResult: pr}, nil
	}

	if v := params.IsActive_Order; v != "" {
		query = query.Order(orgorgan.ByIsActive(OrderDirection(v)))
	}

	if v := params.Sort_Order; v != "" {
		query = query.Order(orgorgan.BySort(OrderDirection(v)))
	}
	query = query.Order(orgorgan.ByID(OrderDirection(OrderByDESC.String())))

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()
	if params.Offset() > count {
		return &schema.OrgOrganQueryResult{PageResult: pr}, nil
	}
	query = query.Limit(params.Limit()).Offset(params.Offset())

	var query_select *mainent.OrgOrganSelect
	if opt.FieldsIncludes != nil {
		query_select = query.Select(opt.FieldsIncludes...)
	}

	list, err1 := query_select.All(ctx)
	if err1 != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.OrgOrganQueryResult{
		PageResult: pr,
		Data:       ToSchemaOrgOrgans(list),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *OrgOrgan) Get(ctx context.Context, id string, opts ...schema.OrgOrganQueryOptions) (*schema.OrgOrgan, error) {

	r_orgorgan, err := a.EntCli.OrgOrgan.Query().Where(orgorgan.IDEQ(id)).WithHaddr().Only(ctx)
	if err != nil {
		if mainent.IsNotFound(err) {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	return ToSchemaOrgOrgan(r_orgorgan), nil
}

// View 查询指定数据
func (a *OrgOrgan) View(ctx context.Context, id string, opts ...schema.OrgOrganQueryOptions) (*schema.OrgOrgan, error) {

	r_orgorgan, err := a.EntCli.OrgOrgan.Query().Where(orgorgan.IDEQ(id)).WithHaddr().Only(ctx)
	if err != nil {
		if mainent.IsNotFound(err) {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	return ToSchemaOrgOrgan(r_orgorgan), nil
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
