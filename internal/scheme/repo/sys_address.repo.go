package repo

import (
	"context"
	"time"

	"github.com/google/wire"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/mainent"
	"github.com/heromicro/omgind/internal/gen/mainent/sysaddress"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/helper/structure"
)

// SysAddressSet 注入SysAddress
var SysAddressSet = wire.NewSet(wire.Struct(new(SysAddress), "*"))

// SysAddress 地址管理存储
type SysAddress struct {
	EntCli *mainent.Client
}

// ToSchemaSysAddress 转换为
func ToSchemaSysAddress(et *mainent.SysAddress) *schema.SysAddress {
	item := new(schema.SysAddress)
	structure.Copy(et, item)
	return item
}

func ToSchemaSysAddresses(ets mainent.SysAddresses) []*schema.SysAddress {
	list := make([]*schema.SysAddress, len(ets))
	for i, item := range ets {
		list[i] = ToSchemaSysAddress(item)
	}
	return list
}

func ToEntCreateSysAddressInput(sch *schema.SysAddress) *mainent.CreateSysAddressInput {
	createinput := new(mainent.CreateSysAddressInput)
	structure.Copy(sch, &createinput)

	return createinput
}

func ToEntUpdateSysAddressInput(sch *schema.SysAddress) *mainent.UpdateSysAddressInput {
	updateinput := new(mainent.UpdateSysAddressInput)
	structure.Copy(sch, &updateinput)

	return updateinput
}

func (a *SysAddress) getQueryOption(opts ...schema.SysAddressQueryOptions) schema.SysAddressQueryOptions {
	var opt schema.SysAddressQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *SysAddress) Query(ctx context.Context, params schema.SysAddressQueryParam, opts ...schema.SysAddressQueryOptions) (*schema.SysAddressQueryResult, error) {
	// opt := a.getQueryOption(opts...)

	query := a.EntCli.SysAddress.Query()

	query = query.Where(sysaddress.DeletedAtIsNil())

	// TODO: 查询条件
	if v := params.QueryValue; v != "" {
		query = query.Where(sysaddress.Or(sysaddress.FirstNameContains(v), sysaddress.LastNameContains(v), sysaddress.MobileContains(v)))
	}

	if v := params.CountryID; v != "" {
		query = query.Where(sysaddress.CountryIDEQ(v))
	}

	if v := params.ProvinceID; v != "" {
		query = query.Where(sysaddress.ProvinceIDEQ(v))
	}

	if v := params.CityID; v != "" {
		query = query.Where(sysaddress.CityIDEQ(v))
	}

	if v := params.CountyID; v != "" {
		query = query.Where(sysaddress.CountyIDEQ(v))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// get total
	pr := &schema.PaginationResult{Total: count}
	if params.PaginationParam.OnlyCount {
		return &schema.SysAddressQueryResult{PageResult: pr}, nil
	}

	if v := params.CreatedAt_Order; v != nil {
		query = query.Order(sysaddress.ByCreatedAt(OrderDirection(*v)))
	}
	if v := params.CountryID_Order; v != "" {
		query = query.Order(sysaddress.ByCountryID(OrderDirection(v)))
	}
	if v := params.ProvinceID_Order; v != "" {
		query = query.Order(sysaddress.ByProvinceID(OrderDirection(v)))

	}
	if v := params.CityID_Order; v != "" {
		query = query.Order(sysaddress.ByCityID(OrderDirection(v)))
	}
	if v := params.CountyID_Order; v != "" {
		query = query.Order(sysaddress.ByCountyID(OrderDirection(v)))
	}

	query = query.Order(sysaddress.ByID(OrderDirection("desc")))

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()
	if params.Offset() > count {
		return &schema.SysAddressQueryResult{PageResult: pr}, nil
	}
	query = query.Limit(params.Limit()).Offset(params.Offset())

	list, err1 := query.All(ctx)
	if err1 != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.SysAddressQueryResult{
		PageResult: pr,
		Data:       ToSchemaSysAddresses(list),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *SysAddress) Get(ctx context.Context, id string, opts ...schema.SysAddressQueryOptions) (*schema.SysAddress, error) {

	r_sysaddress, err := a.EntCli.SysAddress.Query().Where(sysaddress.IDEQ(id)).Only(ctx)
	if err != nil {
		if _, ok := err.(*mainent.NotFoundError); ok {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	return ToSchemaSysAddress(r_sysaddress), nil
}

// Get 查询指定数据
func (a *SysAddress) View(ctx context.Context, id string, opts ...schema.SysAddressQueryOptions) (*schema.SysAddress, error) {

	r_sysaddress, err := a.EntCli.SysAddress.Query().Where(sysaddress.IDEQ(id)).Only(ctx)
	if err != nil {
		if _, ok := err.(*mainent.NotFoundError); ok {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	return ToSchemaSysAddress(r_sysaddress), nil
}

// Create 创建数据
func (a *SysAddress) Create(ctx context.Context, item schema.SysAddress) (*schema.SysAddress, error) {

	iteminput := ToEntCreateSysAddressInput(&item)
	r_sysaddress, err := a.EntCli.SysAddress.Create().SetInput(*iteminput).Save(ctx)

	if err != nil {
		return nil, err
	}
	sch_sysaddress := ToSchemaSysAddress(r_sysaddress)
	return sch_sysaddress, nil
}

// Update 更新数据
func (a *SysAddress) Update(ctx context.Context, id string, item schema.SysAddress) (*schema.SysAddress, error) {

	oitem, err := a.EntCli.SysAddress.Query().Where(sysaddress.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	iteminput := ToEntUpdateSysAddressInput(&item)

	r_sysaddress, err := oitem.Update().SetInput(*iteminput).Save(ctx)
	if err != nil {
		return nil, err
	}
	sch_sysaddress := ToSchemaSysAddress(r_sysaddress)

	return sch_sysaddress, nil
}

// Delete 删除数据
func (a *SysAddress) Delete(ctx context.Context, id string) error {

	r_sysaddress, err := a.EntCli.SysAddress.Query().Where(sysaddress.IDEQ(id)).Only(ctx)

	if err != nil {
		return err
	}

	_, err1 := r_sysaddress.Update().SetDeletedAt(time.Now()).SetIsDel(true).Save(ctx)

	return errors.WithStack(err1)
}

// UpdateActive 更新状态
func (a *SysAddress) UpdateActive(ctx context.Context, id string, active bool) error {

	_, err1 := a.EntCli.SysAddress.Update().Where(sysaddress.IDEQ(id)).SetIsActive(active).Save(ctx)

	return errors.WithStack(err1)
}
