package repo

import (
	"context"
	"time"

	"github.com/google/wire"

	"github.com/heromicro/omgind/internal/app/schema"
	"github.com/heromicro/omgind/internal/gen/ent"
	"github.com/heromicro/omgind/internal/gen/ent/sysaddress"
	"github.com/heromicro/omgind/pkg/errors"
	"github.com/heromicro/omgind/pkg/helper/structure"
)

// SysAddressSet 注入SysAddress
var SysAddressSet = wire.NewSet(wire.Struct(new(SysAddress), "*"))

// SysAddress 地址管理存储
type SysAddress struct {
	EntCli *ent.Client
}

// ToSchemaSysAddress 转换为
func (a *SysAddress) ToSchemaSysAddress(et *ent.SysAddress) *schema.SysAddress {
	item := new(schema.SysAddress)
	structure.Copy(et, item)
	return item
}

func (a *SysAddress) ToSchemaSysAddresses(ets ent.SysAddresses) []*schema.SysAddress {
	list := make([]*schema.SysAddress, len(ets))
	for i, item := range ets {
		list[i] = a.ToSchemaSysAddress(item)
	}
	return list
}

func (a *SysAddress) ToEntCreateSysAddressInput(sch *schema.SysAddress) *ent.CreateSysAddressInput {
	createinput := new(ent.CreateSysAddressInput)
	structure.Copy(sch, &createinput)

	return createinput
}

func (a *SysAddress) ToEntUpdateSysAddressInput(sch *schema.SysAddress) *ent.UpdateSysAddressInput {
	updateinput := new(ent.UpdateSysAddressInput)
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
	opt := a.getQueryOption(opts...)

	query := a.EntCli.SysAddress.Query()

	query = query.Where(sysaddress.DeletedAtIsNil())
	// TODO: 查询条件

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// get total
	pr := &schema.PaginationResult{Total: count}
	if params.PaginationParam.OnlyCount {
		return &schema.SysAddressQueryResult{PageResult: pr}, nil
	}

	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("id", schema.OrderByDESC))
	query = query.Order(ParseOrder(opt.OrderFields)...)

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
		Data:       a.ToSchemaSysAddresses(list),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *SysAddress) Get(ctx context.Context, id string, opts ...schema.SysAddressQueryOptions) (*schema.SysAddress, error) {

	r_sysaddress, err := a.EntCli.SysAddress.Query().Where(sysaddress.IDEQ(id)).Only(ctx)
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	return a.ToSchemaSysAddress(r_sysaddress), nil
}

// Create 创建数据
func (a *SysAddress) Create(ctx context.Context, item schema.SysAddress) (*schema.SysAddress, error) {

	iteminput := a.ToEntCreateSysAddressInput(&item)
	r_sysaddress, err := a.EntCli.SysAddress.Create().SetInput(*iteminput).Save(ctx)

	if err != nil {
		return nil, err
	}
	sch_sysaddress := a.ToSchemaSysAddress(r_sysaddress)
	return sch_sysaddress, nil
}

// Update 更新数据
func (a *SysAddress) Update(ctx context.Context, id string, item schema.SysAddress) (*schema.SysAddress, error) {

	oitem, err := a.EntCli.SysAddress.Query().Where(sysaddress.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	iteminput := a.ToEntUpdateSysAddressInput(&item)

	r_sysaddress, err := oitem.Update().SetInput(*iteminput).Save(ctx)
	sch_sysaddress := a.ToSchemaSysAddress(r_sysaddress)

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
